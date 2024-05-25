package service

import (
    "net/http"

    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/trace"
    "github.com/jupitercloud/subscribed/api"
    "github.com/jupitercloud/subscribed/auth"
    "github.com/jupitercloud/subscribed/errors"
    "github.com/jupitercloud/subscribed/logger"
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

func (self *SubscriptionService) OpenAccount(request *http.Request, args *api.OpenAccountRequest, reply *api.OpenAccountResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC OpenAccount")

    span := trace.SpanFromContext(request.Context())
    span.SetAttributes(
        attribute.String("account.account_id", args.AccountId),
    )

    return self.impl.OpenAccount(request, args, reply)
}

func (self *SubscriptionService) CloseAccount(request *http.Request, args *api.CloseAccountRequest, reply *api.CloseAccountResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC CloseAccount")

    span := trace.SpanFromContext(request.Context())
    span.SetAttributes(
        attribute.String("account.account_id", args.AccountId),
    )

    return self.impl.CloseAccount(request, args, reply)
}

func (self *SubscriptionService) CreateSubscription(request *http.Request, args *api.CreateSubscriptionRequest, reply *api.CreateSubscriptionResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC CreateSubscription")

    span := trace.SpanFromContext(request.Context())
    span.SetAttributes(
        attribute.String("subscription.account_id", args.AccountId),
        attribute.String("subscription.subscription_id", args.SubscriptionId),
        attribute.Int64("subscription.sku", args.Sku),
    )

    return self.impl.CreateSubscription(request, args, reply)
}

func (self *SubscriptionService) TerminateSubscription(request *http.Request, args *api.TerminateSubscriptionRequest, reply *api.TerminateSubscriptionResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC TerminateSubscription")

    span := trace.SpanFromContext(request.Context())
    span.SetAttributes(
        attribute.String("subscription.account_id", args.AccountId),
        attribute.String("subscription.subscription_id", args.SubscriptionId),
        attribute.Int64("subscription.sku", args.Sku),
    )

    return self.impl.TerminateSubscription(request, args, reply)
}

func (self *SubscriptionService) CreateResource(request *http.Request, args *api.CreateResourceRequest, reply *api.CreateResourceResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC CreateResource")

    span := trace.SpanFromContext(request.Context())
    span.SetAttributes(
        attribute.String("resource.account_id", args.AccountId),
        attribute.String("resource.subscription_id", args.SubscriptionId),
        attribute.String("resource.resource_id", args.ResourceId),
        attribute.Int64("resource.sku", args.Sku),
    )

    return self.impl.CreateResource(request, args, reply)
}

func (self *SubscriptionService) TerminateResource(request *http.Request, args *api.TerminateResourceRequest, reply *api.TerminateResourceResponse) error {
    _, err := verifyAuthorization(request)
    if err != nil {
        return err
    }

    log.Debug("RPC TerminateResource")

    span := trace.SpanFromContext(request.Context())
    span.SetAttributes(
        attribute.String("resource.account_id", args.AccountId),
        attribute.String("resource.subscription_id", args.SubscriptionId),
        attribute.String("resource.resource_id", args.ResourceId),
        attribute.Int64("resource.sku", args.Sku),
    )

    return self.impl.TerminateResource(request, args, reply)
}

func createSubscriptionService(impl api.SubscriptionServiceInterface) *SubscriptionService {
    return &SubscriptionService{
      impl: impl,
    }
}
