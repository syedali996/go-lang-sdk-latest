package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetDiscountResponse represents a GetDiscountResponse struct.
// Response object for getting a discount
type GetDiscountResponse struct {
    Id               *string                               `json:"id"`
    Value            *float64                              `json:"value"`
    DiscountType     *string                               `json:"discount_type"`
    Status           *string                               `json:"status"`
    CreatedAt        *time.Time                            `json:"created_at"`
    Cycles           Optional[int]                         `json:"cycles"`
    DeletedAt        Optional[time.Time]                   `json:"deleted_at"`
    Description      Optional[string]                      `json:"description"`
    Subscription     Optional[GetSubscriptionResponse]     `json:"subscription"`
    // The subscription item
    SubscriptionItem Optional[GetSubscriptionItemResponse] `json:"subscription_item"`
}

// MarshalJSON implements the json.Marshaler interface for GetDiscountResponse.
// It customizes the JSON marshaling process for GetDiscountResponse objects.
func (g *GetDiscountResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetDiscountResponse object to a map representation for JSON marshaling.
func (g *GetDiscountResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["value"] = g.Value
    structMap["discount_type"] = g.DiscountType
    structMap["status"] = g.Status
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    if g.Cycles.IsValueSet() {
        structMap["cycles"] = g.Cycles.Value()
    }
    if g.DeletedAt.IsValueSet() {
        var DeletedAtVal *string = nil
        if g.DeletedAt.Value() != nil {
            val := g.DeletedAt.Value().Format(time.RFC3339)
            DeletedAtVal = &val
        }
        structMap["deleted_at"] = DeletedAtVal
    }
    if g.Description.IsValueSet() {
        structMap["description"] = g.Description.Value()
    }
    if g.Subscription.IsValueSet() {
        structMap["subscription"] = g.Subscription.Value()
    }
    if g.SubscriptionItem.IsValueSet() {
        structMap["subscription_item"] = g.SubscriptionItem.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetDiscountResponse.
// It customizes the JSON unmarshaling process for GetDiscountResponse objects.
func (g *GetDiscountResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id               *string                               `json:"id"`
        Value            *float64                              `json:"value"`
        DiscountType     *string                               `json:"discount_type"`
        Status           *string                               `json:"status"`
        CreatedAt        *string                               `json:"created_at"`
        Cycles           Optional[int]                         `json:"cycles"`
        DeletedAt        Optional[string]                      `json:"deleted_at"`
        Description      Optional[string]                      `json:"description"`
        Subscription     Optional[GetSubscriptionResponse]     `json:"subscription"`
        SubscriptionItem Optional[GetSubscriptionItemResponse] `json:"subscription_item"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Value = temp.Value
    g.DiscountType = temp.DiscountType
    g.Status = temp.Status
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        g.CreatedAt = &CreatedAtVal
    }
    g.Cycles = temp.Cycles
    g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
    if temp.DeletedAt.Value() != nil {
        DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
        }
        g.DeletedAt.SetValue(&DeletedAtVal)
    }
    g.Description = temp.Description
    g.Subscription = temp.Subscription
    g.SubscriptionItem = temp.SubscriptionItem
    return nil
}
