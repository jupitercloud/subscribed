package service

import (
    "context"
    "net/http"

    "github.com/jupitercloud/subscribed/api"
)

type SubscriptionServiceStub struct{}

func (t *SubscriptionServiceStub) Initialize(ctx context.Context) error {
    return nil
}

func (t *SubscriptionServiceStub) Shutdown(ctx context.Context) error {
    return nil
}

func (t *SubscriptionServiceStub) HealthCheck(request *http.Request, args *api.HealthCheckRequest, reply *api.HealthCheckResponse) error {
    reply.Ok = true
    return nil
}

func (t *SubscriptionServiceStub) OpenAccount(request *http.Request, args *api.OpenAccountRequest, reply *api.OpenAccountResponse) error {
    return nil
}

func (t *SubscriptionServiceStub) CloseAccount(request *http.Request, args *api.CloseAccountRequest, reply *api.CloseAccountResponse) error {
    return nil
}

func (t *SubscriptionServiceStub) CreateSubscription(request *http.Request, args *api.CreateSubscriptionRequest, reply *api.CreateSubscriptionResponse) error {
    return nil
}

func (t *SubscriptionServiceStub) TerminateSubscription(request *http.Request, args *api.TerminateSubscriptionRequest, reply *api.TerminateSubscriptionResponse) error {
    return nil
}

func (t *SubscriptionServiceStub) CreateResource(request *http.Request, args *api.CreateResourceRequest, reply *api.CreateResourceResponse) error {
    return nil
}

func (t *SubscriptionServiceStub) TerminateResource(request *http.Request, args *api.TerminateResourceRequest, reply *api.TerminateResourceResponse) error {
    return nil
}

func CreateSubscriptionServiceStub() api.SubscriptionServiceInterface {
  return &SubscriptionServiceStub{}
}
