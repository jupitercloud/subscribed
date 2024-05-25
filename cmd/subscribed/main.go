package main

import (
    "context"
    "os"
    "os/signal"
    "syscall"

    "github.com/alecthomas/kong"
    "github.com/jupitercloud/subscribed/logger"
    "github.com/jupitercloud/subscribed/service"
    "github.com/jupitercloud/subscribed/telemetry"
)

var log = logger.Named("main");

type Globals struct {
    LogLevel string `enum:"debug,info,warn,error" default:"info"`
    Telemetry string `enum:"console,grpc,none" default:"none"`
}

type ServerCmd struct {
    Address string `default:":8081" help:"Server bind address"`
    Issuer string `default:"https://jupitercloud.com" help:"OIDC compatible token issuer URL"`
    VendorId string `required:"" help:"Vendor ID operated by this server"`
    Dev bool `default:"false" help:"Development mode. Authorization is disabled"`
}

type CLI struct {
    Globals
    Server ServerCmd `cmd:"" help:"Run a server"`
}

func (cmd *ServerCmd) Run (quit chan os.Signal) error {
    config := service.ServerConfig{
        Address: cmd.Address,
        Issuer: cmd.Issuer,
        VendorId: cmd.VendorId,
        Dev: cmd.Dev,
    }
    impl := service.CreateSubscriptionServiceStub()
    return service.RunServer(config, impl, quit)
}

func main() {
    // This program uses Kong to parse the CLI
    // See https://danielms.site/zet/2023/kong-is-an-amazing-cli-for-go-apps/
    cli := CLI{
        Globals: Globals{
            LogLevel: "INFO",
            Telemetry: "none",
        },
    }
    ctx := kong.Parse(&cli)

    // Initialize logging
    logger.Initialize(cli.Globals.LogLevel)

   	// Set up OpenTelemetry.
  	shutdownTelemetry, err := telemetry.Initialize(context.Background(), telemetry.ExportModeFromString(cli.Globals.Telemetry))

    if err == nil {
        quit := make(chan os.Signal)
        signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
        // Run the Server
        err = ctx.Run(quit)
    }

    if shutdownTelemetry != nil {
        shutdownTelemetry(context.Background())
    }

    if (err == nil) {
        ctx.Kong.Exit(0)
    } else {
        log.Error("Fatal error", "error", err)
        ctx.Kong.Exit(1)
    }
}
