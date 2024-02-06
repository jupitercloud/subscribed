package service

import (
    "net/http"
    "jupitercloud.com/subscribed/api"
    "jupitercloud.com/subscribed/logger"
)

var log = logger.Named("SubscriptionService");

type SubscriptionService struct{}

func (t *SubscriptionService) HealthCheck(r *http.Request, args *api.HealthCheckRequest, reply *api.HealthCheckResponse) error {
    log.Debug("RPC HealthCheck")
    reply.Ok = true
    return nil
}

func (t *SubscriptionService) CreateSubscription(r *http.Request, args *api.CreateSubscriptionRequest, reply *api.CreateSubscriptionResponse) error {
  log.Debug("RPC CreateSubscription")

  return nil
}

func (t *SubscriptionService) TerminateSubscription(r *http.Request, args *api.TerminateSubscriptionRequest, reply *api.TerminateSubscriptionResponse) error {
  log.Debug("RPC TerminateSubscription")

  return nil
}

func (t *SubscriptionService) CreateResource(r *http.Request, args *api.CreateResourceRequest, reply *api.CreateResourceResponse) error {
  log.Debug("RPC CreateResource")

  return nil
}

func (t *SubscriptionService) TerminateResource(r *http.Request, args *api.TerminateResourceRequest, reply *api.TerminateResourceResponse) error {
  log.Debug("RPC TerminateResource")

  return nil
}
