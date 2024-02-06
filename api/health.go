package api

type HealthCheckRequest struct {
}

type HealthCheckResponse struct {
    Ok bool `json:"ok"`
}
