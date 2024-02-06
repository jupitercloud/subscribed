package service

import (
    "net/http"
    "jupitercloud.com/subscribed/logger"
)

var log = logger.Named("SubscriptionService");

type HealthCheckRequest struct {
}

type HealthCheckResponse struct {
    Ok bool `json:"ok"`
}

type SubscriptionService struct{}

func (t *SubscriptionService) HealthCheck(r *http.Request, args *HealthCheckRequest, reply *HealthCheckResponse) error {
    log.Debug("RPC HealthCheck")
    reply.Ok = true
    return nil
}
