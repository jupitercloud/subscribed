// This program provides a JSON-RPC server using gorilla/rpc
// See https://dev.to/iamelesq/more-go-rpc-using-gorilla-rpc-json-5hb for an introduction.

package service

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	rpc "github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/jupitercloud/subscribed/api"
	"github.com/jupitercloud/subscribed/auth"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("server")

type ServerConfig struct {
    // Server bind address
    Address string
    // OIDC compatible token issuer URL. Should be "jupitercloud.com"
    Issuer string
    // Vendor ID operated by this server
    VendorId string
    // Development mode - authorization is disabled.
    Dev bool
}

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

// Run a server, exiting on the quit signal. This function returns an error
// on failure to launch the server, otherwise blocks until the server exits.
func RunServer(config ServerConfig, impl api.SubscriptionServiceInterface, quit chan os.Signal) error {
    svc := createSubscriptionService(impl)
    // Create a new RPC server
    s := rpc.NewServer()
    // Register the type of data requested as JSON
    s.RegisterCodec(json2.NewCodec(), "application/json")
    // Register the service by creating a new JSON server
    s.RegisterService(svc, "")
    s.RegisterInterceptFunc(rpcHookBefore)
    s.RegisterAfterFunc(rpcHookAfter)

    auth := auth.NewAuthService(config.Issuer, config.VendorId, config.Dev)
    err := auth.Initialize(context.Background())
    if (err != nil) {
        log.Error("Failed to initialize authorization")
        return err
    }

    defer auth.Shutdown(context.Background())

    err = impl.Initialize(context.Background())
    if (err != nil) {
        log.Error("Failed to initialize service")
        return err
    }

    defer impl.Shutdown(context.Background())

    r := mux.NewRouter()
    r.Use(otelmux.Middleware("subscribed"))
    r.Use(httpTraceMiddleware)
    r.Use(corsMiddleware)
    r.Use(auth.Middleware)
    r.HandleFunc("/rpc", CorsHandler).Methods("OPTIONS")
    r.Handle("/rpc", s)

    server := &http.Server{Addr: config.Address, Handler: r}

    go func() {
        <-quit
        server.Shutdown(context.Background())
    }()

    log.Info("Launching SubscribeD", "address", config.Address)
    err = server.ListenAndServe()
    if (err != nil && err != http.ErrServerClosed ) {
        log.Error("Failed to launch server", "error", err)
        return err
    }

   	return nil
}
