// This program provides a JSON-RPC server using gorilla/rpc
// See https://dev.to/iamelesq/more-go-rpc-using-gorilla-rpc-json-5hb for an introduction.

package main

import (
    "github.com/hashicorp/go-hclog"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/rpc"
    "github.com/gorilla/rpc/json"
)

var appLogger = hclog.New(&hclog.LoggerOptions{
  Name: "subscribed",
  Level: hclog.LevelFromString("DEBUG"),
});
var log = appLogger.Named("main");

type HealthCheckRequest struct {
}

type HealthCheckResponse struct {
    Ok bool `json:"ok"`
}

type SubscriptionService struct{}

func (t *SubscriptionService) HealthCheck(r *http.Request, args *HealthCheckRequest, reply *HealthCheckResponse) error {
    log.Debug("RPC HealthCheck")
    reply.Ok = true
    return nil
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

func main() {
    port := ":8081"

    log.Info("Launching SubscribeD", "port", port)

    // Create a new RPC server
    s := rpc.NewServer()
    // Register the type of data requested as JSON
    s.RegisterCodec(json.NewCodec(), "application/json")
    // Register the service by creating a new JSON server
    s.RegisterService(new(SubscriptionService), "")

    r := mux.NewRouter()
    r.Use(corsMiddleware)
    r.HandleFunc("/rpc", CorsHandler).Methods("OPTIONS")
    r.Handle("/rpc", s)
    err := http.ListenAndServe(port, r)
    log.Info("Online")
    if (err != nil) {
      log.Error("Failed to launch server", "error", err)
    }
}
