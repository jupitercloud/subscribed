package api

type CreateSubscriptionRequest struct {
    // Account ID owning the subscription.
    AccountId string `json:"accountId"`
    // Subscription ID assigned.
    SubscriptionId string `json:"subscriptionId"`
    // SKU to subscribe.
    Sku int64 `json:"sku"`
}

type CreateSubscriptionResponse struct {
    // Optional vendor-defined data. You may use this field to return ID's or tokens mapping this
    // subscription to your internal application identifiers.
    VendorData map[string]interface{} `json:"vendorData"`
}

type TerminateSubscriptionRequest struct {
    // Account ID owning the subscription.
    AccountId string `json:"accountId"`
    // Subscription ID to terminate.
    SubscriptionId string `json:"subscriptionId"`
    // Vendor-defined data associated with this subscription.
    VendorData map[string]interface{} `json:"vendorData"`
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
    // SKU linked to the subscription.
    Sku int64 `json:"sku"`
    // Vendor-defined configuration for this SKU.
    Configuration map[string]interface{} `json:"configuration"`
}

type ResourceInstructions struct {
    // Format of this instruction content.
    // Valid values: "plain", "markdown"
    Format string `json:"format"`
    // Encoded content
    Content string `json:"content"`
}

type CreateResourceResponse struct {
    // Vendor-defined data associated with this resource.
    VendorData map[string]interface{} `json:"vendorData"`
    // Human readable instructions to access this resource.
    Instructions ResourceInstructions `json:"instructions"`
}

type TerminateResourceRequest struct {
    // Account ID owning the resource.
    AccountId string `json:"accountId"`
    // Subscription ID assocated with the resource.
    SubscriptionId string `json:"subscriptionId"`
    // Resource ID to terminate.
    ResourceId string `json:"resourceId"`
    // Vendor-defined data associated with this resource.
    VendorData map[string]interface{} `json:"vendorData"`
}

type TerminateResourceResponse struct {
    // Empty
}
