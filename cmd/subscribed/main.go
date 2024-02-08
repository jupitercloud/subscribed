package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"
	"jupitercloud.com/subscribed/logger"
	"jupitercloud.com/subscribed/telemetry"
)

var log = logger.Named("main");

type Globals struct {
    LogLevel string `enum:"debug,info,warn,error" default:"info"`
    Telemetry string `enum:"console,grpc,none" default:"none"`
}

type ServerCmd struct {
    Address string `default:":8081" help:"Server bind address"`
    Issuer string `default:"https://login.poseidon.cloud" help:"OIDC compatible token issuer URL"`
    VendorId string `required:"" help:"Vendor ID for this instance"`
    Dev bool `default:"false" help:"Development mode. Authorization is disabled"`
}

type CLI struct {
    Globals
    Server ServerCmd `cmd:"" help:"Run a server"`
}

func exit(err error) {
    log.Error("Fatal error", "error", err)
    os.Exit(1)
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
    logger.Initialize(cli.Globals.LogLevel)
   	// Set up OpenTelemetry.
	shutdownTelemetry, err := telemetry.Initialize(context.Background(), telemetry.ExportModeFromString(cli.Globals.Telemetry))
	if err != nil {
		exit(err)
	}
    // Shutdown telemetry on exit.
	defer func() {
        shutdownTelemetry(context.Background())
	}()

    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    err = ctx.Run(quit)
    if (err != nil) {
        exit(err)
    }
}
