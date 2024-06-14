# SubscribeD

SubscribeD provides a JSON-RPC server which links your services to the Jupiter Cloud software distribution channel.

## Usage
Import the subscribed library into your go project, and build a [main.go](cmd/subscribed).

```go
package main

import (
    "github.com/jupitercloud/subscribed/service"
)

func main() {
    // Set up config, interface implementation, and shutdown hook
    // ...
    
    // Launch the JSON-RPC server
    service.RunServer(config, impl, quit)
}

```


## Build Instructions
This builds the stub `subscribed` server for an example application.

### Native

    go build -o subscribed ./cmd/subscribed

### Docker

    docker build . -t jupitercloud/subscribed:(date +%s)
