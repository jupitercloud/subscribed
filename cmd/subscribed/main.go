// This program provides a JSON-RPC server using gorilla/rpc
// See https://dev.to/iamelesq/more-go-rpc-using-gorilla-rpc-json-5hb for an introduction.

package main

import (
    "net/http"
    "os"

    "github.com/alecthomas/kong"
    "github.com/gorilla/mux"
    "github.com/gorilla/rpc"
    "github.com/gorilla/rpc/json"
    "jupitercloud.com/subscribed/logger"
    "jupitercloud.com/subscribed/service"
)

var log = logger.Named("main");

type Globals struct {
    LogLevel string `enum:"debug,info,warn,error" default:"info"`
}

type ServerCmd struct {
    Address string `default:":8081"`
}

type CLI struct {
    Globals
    Server ServerCmd `cmd:"" help:"Run a server"`
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

func (cmd *ServerCmd) Run() error {
    log.Info("Launching SubscribeD", "address", cmd.Address)

    // Create a new RPC server
    s := rpc.NewServer()
    // Register the type of data requested as JSON
    s.RegisterCodec(json.NewCodec(), "application/json")
    // Register the service by creating a new JSON server
    s.RegisterService(new(service.SubscriptionService), "")

    r := mux.NewRouter()
    r.Use(corsMiddleware)
    r.HandleFunc("/rpc", CorsHandler).Methods("OPTIONS")
    r.Handle("/rpc", s)
    err := http.ListenAndServe(cmd.Address, r)
    log.Info("Online")
    if (err != nil) {
      log.Error("Failed to launch server", "error", err)
    }

   	return nil
}

func main() {
    // This program uses Kong to parse the CLI
    // See https://danielms.site/zet/2023/kong-is-an-amazing-cli-for-go-apps/

    cli := CLI{
        Globals: Globals{
            LogLevel: "INFO",
        },
    }
    ctx := kong.Parse(&cli)
    logger.Initialize(cli.Globals.LogLevel)
    err := ctx.Run()
    if (err != nil) {
        log.Error("Fatal error", "error", err)
        os.Exit(1)
    }
}
