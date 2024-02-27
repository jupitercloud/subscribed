package api

import (
	"context"
	"encoding/json"
	"net/http"
)

type Metadata map[string]interface{}

type RichText struct {
    // Format of this instruction content.
    // Valid values: "plain", "markdown"
    Format string `json:"format"`
    // Encoded content
    Content string `json:"content"`
}

type Address struct {
    // Unique ID for this address
    AddressId string `json:"addressId"`
    // Address type. Valid values: 'PRIMARY', 'BILLING', 'SHIPPING'
    AddressType string `json:"addressType"`
    // Street address line 1
    Line1 string `json:"line1"`
    // Street address line 2
    Line2 string `json:"line2"`
    // City or local jurisdiction
    City string `json:"city"`
    // State or province code
    State string `json:"state"`
    // Two-letter ISO country code
    Country string `json:"country"`
    // Postal code
    PostalCode string `json:"postalCode"`
}

type OpenAccountRequest struct {
    // Account ID to create
    AccountId string `json:"accountId"`
    // Account name
    Name string `json:"name"`
    // Address info
    Addresses []Address `json:"addresses"`
}

type OpenAccountResponse struct {
    // Optional vendor-defined data to attach to the account.
    AccountData Metadata `json:"accountData"`
}

type CloseAccountRequest struct {
    // Account ID to terminate.
    AccountId string `json:"accountId"`
    // Vendor-defined data associated with this account.
    AccountData Metadata `json:"accountData"`
}

type CloseAccountResponse struct {
    // Empty
}

type CreateSubscriptionRequest struct {
    // Account ID owning the subscription.
    AccountId string `json:"accountId"`
    // Subscription ID assigned.
    SubscriptionId string `json:"subscriptionId"`
    // SKU to subscribe.
    Sku int64 `json:"sku"`
    // Vendor-defined data for the account.
    AccountData Metadata `json:"accountData"`
}

type CreateSubscriptionResponse struct {
    // Vendor-defined data. You may use this field to return ID's or tokens mapping this
    // subscription to your internal application identifiers.
    SubscriptionData Metadata `json:"subscriptionData"`
    // URL to access this subscription.
    Url string `json:"url"`
    // Human readable instructions to access this subscription.
    Instructions RichText `json:"instructions"`
}

type TerminateSubscriptionRequest struct {
    // Account ID owning the subscription.
    AccountId string `json:"accountId"`
    // Subscription ID to terminate.
    SubscriptionId string `json:"subscriptionId"`
    // SKU for the subscription.
    Sku int64 `json:"sku"`
    // Vendor-defined data for the account.
    AccountData Metadata `json:"accountData"`
    // Vendor-defined data for the subscription.
    SubscriptionData Metadata `json:"subscriptionData"`
}

type TerminateSubscriptionResponse struct {
    // Empty
}

type CreateResourceRequest struct {
    // Account ID owning the resource.
    AccountId string `json:"accountId"`
    // Subscription ID associated with the resource.
    SubscriptionId string `json:"subscriptionId"`
    // Resource ID assigned.
    ResourceId string `json:"resourceId"`
    // SKU for the subscription.
    Sku int64 `json:"sku"`
    // Resource name assigned by the user.
    ResourceName string `json:"resourceName"`
    // Vendor-defined configuration for this SKU.
    Configuration json.RawMessage `json:"configuration"`
    // Vendor-defined data for the account.
    AccountData Metadata `json:"accountData"`
    // Vendor-defined data for the subscription.
    SubscriptionData Metadata `json:"subscriptionData"`
}

type CreateResourceResponse struct {
    // URL to access this resource
    Url string `json:"url"`
    // Vendor-defined data associated with this resource.
    ResourceData Metadata `json:"resourceData"`
    // Human readable instructions to access this resource.
    Instructions RichText `json:"instructions"`
}

type TerminateResourceRequest struct {
    // Account ID owning the resource.
    AccountId string `json:"accountId"`
    // Subscription ID assocated with the resource.
    SubscriptionId string `json:"subscriptionId"`
    // Resource ID to terminate.
    ResourceId string `json:"resourceId"`
    // SKU for the subscription resource.
    Sku int64 `json:"sku"`
    // Resource name assigned by the user.
    ResourceName string `json:"resourceName"`
    // Vendor-defined data for the account.
    AccountData Metadata `json:"accountData"`
    // Vendor-defined data for the subscription.
    SubscriptionData Metadata `json:"subscriptionData"`
    // Vendor-defined data for the resource.
    ResourceData Metadata `json:"resourceData"`
}

type TerminateResourceResponse struct {
    // Empty
}

type Initializable interface {
    // Perform resource initialization, such as establishing a database connection.
    Initialize(ctx context.Context) error
    // Close resources opened during initialization.
    Shutdown(ctx context.Context) error
}

type SubscriptionServiceInterface interface {
    Initializable

    // Probe the service for liveness.
    HealthCheck(request *http.Request, args *HealthCheckRequest, reply *HealthCheckResponse) error

    // Create (or reopen) a customer account.
    OpenAccount(request *http.Request, args *OpenAccountRequest, reply *OpenAccountResponse) error

    // Close a customer account.
    CloseAccount(request *http.Request, args *CloseAccountRequest, reply *CloseAccountResponse) error

    // Create a new subscription.
    CreateSubscription(request *http.Request, args *CreateSubscriptionRequest, reply *CreateSubscriptionResponse) error

    // Terminate an existing subscription.
    TerminateSubscription(request *http.Request, args *TerminateSubscriptionRequest, reply *TerminateSubscriptionResponse) error

    // Create a new resource in a subscription.
    CreateResource(request *http.Request, args *CreateResourceRequest, reply *CreateResourceResponse) error

    // Terminate a resource in a subscription.
    TerminateResource(request *http.Request, args *TerminateResourceRequest, reply *TerminateResourceResponse) error
}
