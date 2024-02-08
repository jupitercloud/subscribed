// This program provides a JSON-RPC server using gorilla/rpc
// See https://dev.to/iamelesq/more-go-rpc-using-gorilla-rpc-json-5hb for an introduction.

package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"jupitercloud.com/subscribed/auth"
	"jupitercloud.com/subscribed/service"
)

var tracer = otel.Tracer("server")

func CorsHandler(response http.ResponseWriter, request *http.Request) {
    log.Debug("OPTIONS /rpc")
    response.WriteHeader(http.StatusOK)
}

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
        header := response.Header()
        header.Add("Access-Control-Allow-Origin", "*")
        header.Add("Access-Control-Allow-Headers", "*")
        header.Add("Access-Control-Allow-Methods", "*")
        header.Add("Access-Control-Allow-Credentials", "true")
        next.ServeHTTP(response, request)
    })
}

func httpTraceMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
        span := trace.SpanFromContext(request.Context())
        if span.IsRecording() {
            span.SetAttributes(
                attribute.String("http.request.header.content-type", request.Header.Get("Content-Type")),
            )
        }
        next.ServeHTTP(response, request)
    })
}

func rpcHookBefore(info *rpc.RequestInfo) *http.Request {
    ctx1, span := tracer.Start(info.Request.Context(), "RPC " + info.Method)
    ctx2 := context.WithValue(ctx1, "rpcSpan", span)
    span.SetAttributes(
        attribute.String("rpc.system", "json_rpc"),
        attribute.String("rpc.method", info.Method),
    )
    if info.Error == nil {
        span.SetStatus(codes.Ok, codes.Ok.String())
    } else {
        span.SetStatus(codes.Error, info.Error.Error())
    }
    req := info.Request.WithContext(ctx2)
    return req
}

func rpcHookAfter(info *rpc.RequestInfo) {
    span := info.Request.Context().Value("rpcSpan").(trace.Span)
    span.End()
}

func (cmd *ServerCmd) Run() error {
    log.Info("Launching SubscribeD", "address", cmd.Address)

    // Create a new RPC server
    s := rpc.NewServer()
    // Register the type of data requested as JSON
    s.RegisterCodec(json.NewCodec(), "application/json")
    // Register the service by creating a new JSON server
    s.RegisterService(new(service.SubscriptionService), "")
    s.RegisterInterceptFunc(rpcHookBefore)
    s.RegisterAfterFunc(rpcHookAfter)

    auth := auth.NewAuthService(cmd.Issuer)
    err := auth.Initialize(context.Background())
    if (err != nil) {
        log.Error("Failed to initialize authorization")
        return err
    }

    defer auth.Shutdown(context.Background())

    r := mux.NewRouter()
    r.Use(otelmux.Middleware("subscribed"))
    r.Use(httpTraceMiddleware)
    r.Use(corsMiddleware)
    r.Use(auth.Middleware)
    r.HandleFunc("/rpc", CorsHandler).Methods("OPTIONS")
    r.Handle("/rpc", s)
    err = http.ListenAndServe(cmd.Address, r)
    if (err != nil) {
      log.Error("Failed to launch server", "error", err)
      return err
    }

   	return nil
}
