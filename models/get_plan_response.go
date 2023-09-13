package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetPlanResponse represents a GetPlanResponse struct.
// Response object for getting a plan
type GetPlanResponse struct {
    Id                  *string               `json:"id"`
    Name                *string               `json:"name"`
    Description         *string               `json:"description"`
    Url                 *string               `json:"url"`
    StatementDescriptor *string               `json:"statement_descriptor"`
    Interval            *string               `json:"interval"`
    IntervalCount       *int                  `json:"interval_count"`
    BillingType         *string               `json:"billing_type"`
    PaymentMethods      []string              `json:"payment_methods"`
    Installments        []int                 `json:"installments"`
    Status              *string               `json:"status"`
    Currency            *string               `json:"currency"`
    CreatedAt           *time.Time            `json:"created_at"`
    UpdatedAt           *time.Time            `json:"updated_at"`
    Items               []GetPlanItemResponse `json:"items"`
    BillingDays         []int                 `json:"billing_days"`
    Shippable           *bool                 `json:"shippable"`
    Metadata            map[string]string     `json:"metadata"`
    TrialPeriodDays     Optional[int]         `json:"trial_period_days"`
    MinimumPrice        Optional[int]         `json:"minimum_price"`
    DeletedAt           Optional[time.Time]   `json:"deleted_at"`
}

// MarshalJSON implements the json.Marshaler interface for GetPlanResponse.
// It customizes the JSON marshaling process for GetPlanResponse objects.
func (g *GetPlanResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetPlanResponse object to a map representation for JSON marshaling.
func (g *GetPlanResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["name"] = g.Name
    structMap["description"] = g.Description
    structMap["url"] = g.Url
    structMap["statement_descriptor"] = g.StatementDescriptor
    structMap["interval"] = g.Interval
    structMap["interval_count"] = g.IntervalCount
    structMap["billing_type"] = g.BillingType
    structMap["payment_methods"] = g.PaymentMethods
    structMap["installments"] = g.Installments
    structMap["status"] = g.Status
    structMap["currency"] = g.Currency
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    structMap["items"] = g.Items
    structMap["billing_days"] = g.BillingDays
    structMap["shippable"] = g.Shippable
    structMap["metadata"] = g.Metadata
    if g.TrialPeriodDays.IsValueSet() {
        structMap["trial_period_days"] = g.TrialPeriodDays.Value()
    }
    if g.MinimumPrice.IsValueSet() {
        structMap["minimum_price"] = g.MinimumPrice.Value()
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

// UnmarshalJSON implements the json.Unmarshaler interface for GetPlanResponse.
// It customizes the JSON unmarshaling process for GetPlanResponse objects.
func (g *GetPlanResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                  *string               `json:"id"`
        Name                *string               `json:"name"`
        Description         *string               `json:"description"`
        Url                 *string               `json:"url"`
        StatementDescriptor *string               `json:"statement_descriptor"`
        Interval            *string               `json:"interval"`
        IntervalCount       *int                  `json:"interval_count"`
        BillingType         *string               `json:"billing_type"`
        PaymentMethods      []string              `json:"payment_methods"`
        Installments        []int                 `json:"installments"`
        Status              *string               `json:"status"`
        Currency            *string               `json:"currency"`
        CreatedAt           *string               `json:"created_at"`
        UpdatedAt           *string               `json:"updated_at"`
        Items               []GetPlanItemResponse `json:"items"`
        BillingDays         []int                 `json:"billing_days"`
        Shippable           *bool                 `json:"shippable"`
        Metadata            map[string]string     `json:"metadata"`
        TrialPeriodDays     Optional[int]         `json:"trial_period_days"`
        MinimumPrice        Optional[int]         `json:"minimum_price"`
        DeletedAt           Optional[string]      `json:"deleted_at"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Name = temp.Name
    g.Description = temp.Description
    g.Url = temp.Url
    g.StatementDescriptor = temp.StatementDescriptor
    g.Interval = temp.Interval
    g.IntervalCount = temp.IntervalCount
    g.BillingType = temp.BillingType
    g.PaymentMethods = temp.PaymentMethods
    g.Installments = temp.Installments
    g.Status = temp.Status
    g.Currency = temp.Currency
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
    g.Items = temp.Items
    g.BillingDays = temp.BillingDays
    g.Shippable = temp.Shippable
    g.Metadata = temp.Metadata
    g.TrialPeriodDays = temp.TrialPeriodDays
    g.MinimumPrice = temp.MinimumPrice
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
