package auth

import (
	"context"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"jupitercloud.com/subscribed/errors"
	"jupitercloud.com/subscribed/logger"
)

var log = logger.Named("auth");

type Claims struct {
    // Failure to parse claims
    Error error
    // The JWT *MUST* have a vendorId claim matching our vendor ID.
    VendorId string `json:"https://jupitercloud.com/vendorId"`
    // Account ID scope authorized for this token.
    AccountId string `json:"https://jupitercloud.com/accountId"`
}

type authService struct {
    issuer string
    vendorId string
    provider *oidc.Provider
    verifier *oidc.IDTokenVerifier
}


func (auth *authService) readToken(ctx context.Context, tokenString string) *Claims {
    var claims Claims
    if (tokenString == "") {
        claims.Error = errors.Unauthenticated()
        return &claims
    }
    token, err := auth.verifier.Verify(ctx, tokenString)
    if err != nil {
        claims.Error = errors.JwtError(err)
        return &claims
    }
    err = token.Claims(&claims)
    if err != nil {
        claims.Error = errors.JwtError(err)
        return &claims
    }
    if claims.VendorId != auth.vendorId {
        claims.Error = errors.InvalidVendorIdClaim()
        return &claims
    }
    return &claims
}

func (auth *authService) Initialize(ctx context.Context) error {
    log.Info("Initializing authorization service", "issuer", auth.issuer, "vendor-id", auth.vendorId)
    provider, err := oidc.NewProvider(ctx, auth.issuer)

    if err != nil {
        return err
    }

    auth.provider = provider
    // No client ID, as we are not executing a full OIDC handshake.
    auth.verifier = provider.Verifier(&oidc.Config{SkipClientIDCheck: true})
    return nil
}


func (auth *authService) Shutdown(ctx context.Context) error {
    log.Debug("Shutting down authorization service")
    return nil
}

func (auth *authService) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
        ctx := request.Context()
        claims := auth.readToken(ctx, request.Header.Get("Authorization"))
        ctx2 := context.WithValue(ctx, "claims", claims)
        request2 := request.WithContext(ctx2)
        next.ServeHTTP(response, request2)
    })
}

func NewAuthService(issuer string, vendorId string) *authService {
    return &authService{issuer: issuer, vendorId: vendorId}
}
