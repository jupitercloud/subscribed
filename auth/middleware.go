package auth

import (
    "context"
	"net/http"
  	"github.com/coreos/go-oidc/v3/oidc"
    "jupitercloud.com/subscribed/logger"
)

var log = logger.Named("auth");

type authService struct {
    issuer string
    provider *oidc.Provider
}

func (auth *authService) Initialize(ctx context.Context) error {
    log.Info("Initializing authorization service", "issuer", auth.issuer)
    provider, err := oidc.NewProvider(ctx, auth.issuer)

    if err != nil {
        return err
    }

    auth.provider = provider
    return nil
}


func (auth *authService) Shutdown(ctx context.Context) error {
    log.Debug("Shutting down authorization service")
    return nil
}

func (amw *authService) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
        next.ServeHTTP(response, request)
    })
}

func NewAuthService(issuer string) *authService {
    return &authService{issuer: issuer}
}
