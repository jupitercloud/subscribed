package api

type GetSubscriptionUsageRequest struct {
    // Account ID owning the subscription.
    AccountId string `json:"accountId"`
    // Subscription ID to query billable usage from.
    SubscriptionId string `json:"subscriptionId"`
    // SKU for the subscription.
    Sku int64 `json:"sku"`
    // Vendor-defined data for the account.
    AccountData Metadata `json:"accountData"`
    // Vendor-defined data for the subscription.
    SubscriptionData Metadata `json:"subscriptionData"`
    // Query time period start date-time in RFC 3339 format. Inclusive.
    StartTime string `json:"startTime"`
    // Query time period end date-time in RFC 3339 format. Exclusive.
    EndTime string `json:"endTime"`
}

type CurrencyValue struct {
    // Currency units billed. Supported values: USD
    Currency string `json:"currency"`
    // Discrete currency units billable. For USD, this value is cents.
    // Fractions of cents are acceptable here and will be accumulated
    // over the entire billing period.
    Value float64 `json:"value"`
}

// A line item of billable usage.
type SubscriptionUsage struct {
    // Billable amount
    Amount CurrencyValue `json:"amount"`
    // Line item description
    Description string `json:"description"`
    // Numeric value of units consumed, when applicable.
    Volume float32 `json:"volume"`
    // Label for a billable unit of volume, e.g. "hours", "users", etc.
    Unit string `json:"unit"`
    // Resource ID, when tied to a specific resource.
    ResourceId string `json:"resourceId"`
    // Resource Name, when tied to a specific resource.
    // Prefer resource ID when available.
    ResourceName string `json:"resourceName"`
}

type GetSubscriptionUsageResponse struct {
    // Array of billable usage data
    Usage []SubscriptionUsage `json:"usage"`
}
