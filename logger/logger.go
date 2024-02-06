package logger

import (
    "github.com/hashicorp/go-hclog"
)

var appLogger = hclog.New(&hclog.LoggerOptions{
    Name: "subscribed",
    Level: hclog.LevelFromString("DEBUG"),
});

func Named(name string) hclog.Logger {
    return appLogger.Named(name)
}
