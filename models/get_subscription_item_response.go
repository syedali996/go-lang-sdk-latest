package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetSubscriptionItemResponse represents a GetSubscriptionItemResponse struct.
type GetSubscriptionItemResponse struct {
    Id            *string                   `json:"id"`
    Description   *string                   `json:"description"`
    Status        *string                   `json:"status"`
    CreatedAt     *time.Time                `json:"created_at"`
    UpdatedAt     *time.Time                `json:"updated_at"`
    PricingScheme *GetPricingSchemeResponse `json:"pricing_scheme"`
    Discounts     []GetDiscountResponse     `json:"discounts"`
    Increments    []GetIncrementResponse    `json:"increments"`
    Subscription  *GetSubscriptionResponse  `json:"subscription"`
    // Item name
    Name          *string                   `json:"name"`
    Quantity      Optional[int]             `json:"quantity"`
    Cycles        Optional[int]             `json:"cycles"`
    DeletedAt     Optional[time.Time]       `json:"deleted_at"`
}

// MarshalJSON implements the json.Marshaler interface for GetSubscriptionItemResponse.
// It customizes the JSON marshaling process for GetSubscriptionItemResponse objects.
func (g *GetSubscriptionItemResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetSubscriptionItemResponse object to a map representation for JSON marshaling.
func (g *GetSubscriptionItemResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["description"] = g.Description
    structMap["status"] = g.Status
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    structMap["pricing_scheme"] = g.PricingScheme
    structMap["discounts"] = g.Discounts
    structMap["increments"] = g.Increments
    structMap["subscription"] = g.Subscription
    structMap["name"] = g.Name
    if g.Quantity.IsValueSet() {
        structMap["quantity"] = g.Quantity.Value()
    }
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
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSubscriptionItemResponse.
// It customizes the JSON unmarshaling process for GetSubscriptionItemResponse objects.
func (g *GetSubscriptionItemResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id            *string                   `json:"id"`
        Description   *string                   `json:"description"`
        Status        *string                   `json:"status"`
        CreatedAt     *string                   `json:"created_at"`
        UpdatedAt     *string                   `json:"updated_at"`
        PricingScheme *GetPricingSchemeResponse `json:"pricing_scheme"`
        Discounts     []GetDiscountResponse     `json:"discounts"`
        Increments    []GetIncrementResponse    `json:"increments"`
        Subscription  *GetSubscriptionResponse  `json:"subscription"`
        Name          *string                   `json:"name"`
        Quantity      Optional[int]             `json:"quantity"`
        Cycles        Optional[int]             `json:"cycles"`
        DeletedAt     Optional[string]          `json:"deleted_at"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Description = temp.Description
    g.Status = temp.Status
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        g.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        g.UpdatedAt = &UpdatedAtVal
    }
    g.PricingScheme = temp.PricingScheme
    g.Discounts = temp.Discounts
    g.Increments = temp.Increments
    g.Subscription = temp.Subscription
    g.Name = temp.Name
    g.Quantity = temp.Quantity
    g.Cycles = temp.Cycles
    g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
    if temp.DeletedAt.Value() != nil {
        DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
        }
        g.DeletedAt.SetValue(&DeletedAtVal)
    }
    return nil
}
