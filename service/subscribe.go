package service

import (
	"net/http"

	"jupitercloud.com/subscribed/api"
	"jupitercloud.com/subscribed/auth"
	"jupitercloud.com/subscribed/logger"
)

var log = logger.Named("SubscriptionService");

type SubscriptionService struct{}


func verifyAuthorization (request *http.Request) (*auth.Claims, error) {
    var claims *auth.Claims = request.Context().Value("claims").(*auth.Claims)
    if claims == nil {
        return nil, auth.Unauthenticated()
    }
    if claims.Error != nil {
        return nil, claims.Error
    }
    return claims, nil
}

func (t *SubscriptionService) HealthCheck(request *http.Request, args *api.HealthCheckRequest, reply *api.HealthCheckResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC HealthCheck")
    reply.Ok = true
    return nil
}

func (t *SubscriptionService) CreateSubscription(request *http.Request, args *api.CreateSubscriptionRequest, reply *api.CreateSubscriptionResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }
    log.Debug("RPC CreateSubscription")

    return nil
}

func (t *SubscriptionService) TerminateSubscription(request *http.Request, args *api.TerminateSubscriptionRequest, reply *api.TerminateSubscriptionResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }
    log.Debug("RPC TerminateSubscription")

    return nil
}

func (t *SubscriptionService) CreateResource(request *http.Request, args *api.CreateResourceRequest, reply *api.CreateResourceResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }
    log.Debug("RPC CreateResource")

    return nil
}

func (t *SubscriptionService) TerminateResource(request *http.Request, args *api.TerminateResourceRequest, reply *api.TerminateResourceResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }
    log.Debug("RPC TerminateResource")

    return nil
}
