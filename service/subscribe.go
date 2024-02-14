package service

import (
	"net/http"

	"jupitercloud.com/subscribed/api"
	"jupitercloud.com/subscribed/auth"
	"jupitercloud.com/subscribed/errors"
	"jupitercloud.com/subscribed/logger"
)

var log = logger.Named("SubscriptionService");

// This SubscriptionService wrapper wraps an implementation with token verification and logging.
type SubscriptionService struct{
    impl api.SubscriptionServiceInterface
}

func verifyAuthorization (request *http.Request) (*auth.Claims, error) {
    var claims *auth.Claims = request.Context().Value("claims").(*auth.Claims)
    if claims == nil {
        return nil, errors.Unauthenticated()
    }
    if claims.Error != nil {
        return nil, claims.Error
    }
    return claims, nil
}

func (self *SubscriptionService) HealthCheck(request *http.Request, args *api.HealthCheckRequest, reply *api.HealthCheckResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC HealthCheck")
    return self.impl.HealthCheck(request, args, reply)
}

func (self *SubscriptionService) CreateAccount(request *http.Request, args *api.CreateAccountRequest, reply *api.CreateAccountResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC CreateAccount")
    return self.impl.CreateAccount(request, args, reply)
}

func (self *SubscriptionService) TerminateAccount(request *http.Request, args *api.TerminateAccountRequest, reply *api.TerminateAccountResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC TerminateAccount")
    return self.impl.TerminateAccount(request, args, reply)
}

func (self *SubscriptionService) CreateSubscription(request *http.Request, args *api.CreateSubscriptionRequest, reply *api.CreateSubscriptionResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC CreateSubscription")
    return self.impl.CreateSubscription(request, args, reply)
}

func (self *SubscriptionService) TerminateSubscription(request *http.Request, args *api.TerminateSubscriptionRequest, reply *api.TerminateSubscriptionResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC TerminateSubscription")
    return self.impl.TerminateSubscription(request, args, reply)
}

func (self *SubscriptionService) CreateResource(request *http.Request, args *api.CreateResourceRequest, reply *api.CreateResourceResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC CreateResource")
    return self.impl.CreateResource(request, args, reply)
}

func (self *SubscriptionService) TerminateResource(request *http.Request, args *api.TerminateResourceRequest, reply *api.TerminateResourceResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC TerminateResource")
    return self.impl.TerminateResource(request, args, reply)
}

func CreateSubscriptionService(impl api.SubscriptionServiceInterface) *SubscriptionService {
    return &SubscriptionService{
      impl: impl,
    }
}
