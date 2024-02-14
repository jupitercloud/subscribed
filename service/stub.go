package service

import (
	"net/http"

	"jupitercloud.com/subscribed/api"
)

type SubscriptionServiceStub struct{}

func (t *SubscriptionServiceStub) HealthCheck(request *http.Request, args *api.HealthCheckRequest, reply *api.HealthCheckResponse) error {
    reply.Ok = true
    return nil
}

func (t *SubscriptionServiceStub) CreateAccount(request *http.Request, args *api.CreateAccountRequest, reply *api.CreateAccountResponse) error {
    return nil
}

func (t *SubscriptionServiceStub) TerminateAccount(request *http.Request, args *api.TerminateAccountRequest, reply *api.TerminateAccountResponse) error {
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
