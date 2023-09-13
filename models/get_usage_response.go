package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetUsageResponse represents a GetUsageResponse struct.
// Response object for getting a usage
type GetUsageResponse struct {
    // Id
    Id               *string                      `json:"id"`
    // Quantity
    Quantity         *int                         `json:"quantity"`
    // Description
    Description      *string                      `json:"description"`
    // Used at
    UsedAt           *time.Time                   `json:"used_at"`
    // Creation date
    CreatedAt        *time.Time                   `json:"created_at"`
    // Status
    Status           *string                      `json:"status"`
    DeletedAt        Optional[time.Time]          `json:"deleted_at"`
    // Subscription item
    SubscriptionItem *GetSubscriptionItemResponse `json:"subscription_item"`
    // Identification code in the client system
    Code             Optional[string]             `json:"code"`
    // Identification group in the client system
    Group            Optional[string]             `json:"group"`
    // Field used in item scheme type 'Percent'
    Amount           Optional[int]                `json:"amount"`
}

// MarshalJSON implements the json.Marshaler interface for GetUsageResponse.
// It customizes the JSON marshaling process for GetUsageResponse objects.
func (g *GetUsageResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetUsageResponse object to a map representation for JSON marshaling.
func (g *GetUsageResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["quantity"] = g.Quantity
    structMap["description"] = g.Description
    structMap["used_at"] = g.UsedAt.Format(time.RFC3339)
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["status"] = g.Status
    if g.DeletedAt.IsValueSet() {
        var DeletedAtVal *string = nil
        if g.DeletedAt.Value() != nil {
            val := g.DeletedAt.Value().Format(time.RFC3339)
            DeletedAtVal = &val
        }
        structMap["deleted_at"] = DeletedAtVal
    }
    structMap["subscription_item"] = g.SubscriptionItem
    if g.Code.IsValueSet() {
        structMap["code"] = g.Code.Value()
    }
    if g.Group.IsValueSet() {
        structMap["group"] = g.Group.Value()
    }
    if g.Amount.IsValueSet() {
        structMap["amount"] = g.Amount.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetUsageResponse.
// It customizes the JSON unmarshaling process for GetUsageResponse objects.
func (g *GetUsageResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id               *string                      `json:"id"`
        Quantity         *int                         `json:"quantity"`
        Description      *string                      `json:"description"`
        UsedAt           *string                      `json:"used_at"`
        CreatedAt        *string                      `json:"created_at"`
        Status           *string                      `json:"status"`
        DeletedAt        Optional[string]             `json:"deleted_at"`
        SubscriptionItem *GetSubscriptionItemResponse `json:"subscription_item"`
        Code             Optional[string]             `json:"code"`
        Group            Optional[string]             `json:"group"`
        Amount           Optional[int]                `json:"amount"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Quantity = temp.Quantity
    g.Description = temp.Description
    if temp.UsedAt != nil {
        UsedAtVal, err := time.Parse(time.RFC3339, *temp.UsedAt)
        if err != nil {
            log.Fatalf("Cannot Parse used_at as % s format.", time.RFC3339)
        }
        g.UsedAt = &UsedAtVal
    }
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        g.CreatedAt = &CreatedAtVal
    }
    g.Status = temp.Status
    g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
    if temp.DeletedAt.Value() != nil {
        DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
        }
        g.DeletedAt.SetValue(&DeletedAtVal)
    }
    g.SubscriptionItem = temp.SubscriptionItem
    g.Code = temp.Code
    g.Group = temp.Group
    g.Amount = temp.Amount
    return nil
}
