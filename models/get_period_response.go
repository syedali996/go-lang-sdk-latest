package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetPeriodResponse represents a GetPeriodResponse struct.
// Response object for getting a period
type GetPeriodResponse struct {
    StartAt      *time.Time               `json:"start_at"`
    EndAt        *time.Time               `json:"end_at"`
    Id           *string                  `json:"id"`
    BillingAt    *time.Time               `json:"billing_at"`
    Subscription *GetSubscriptionResponse `json:"subscription"`
    Status       *string                  `json:"status"`
    Duration     *int                     `json:"duration"`
    CreatedAt    *string                  `json:"created_at"`
    UpdatedAt    *string                  `json:"updated_at"`
    Cycle        *int                     `json:"cycle"`
}

// MarshalJSON implements the json.Marshaler interface for GetPeriodResponse.
// It customizes the JSON marshaling process for GetPeriodResponse objects.
func (g *GetPeriodResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetPeriodResponse object to a map representation for JSON marshaling.
func (g *GetPeriodResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["start_at"] = g.StartAt.Format(time.RFC3339)
    structMap["end_at"] = g.EndAt.Format(time.RFC3339)
    structMap["id"] = g.Id
    structMap["billing_at"] = g.BillingAt.Format(time.RFC3339)
    structMap["subscription"] = g.Subscription
    structMap["status"] = g.Status
    structMap["duration"] = g.Duration
    structMap["created_at"] = g.CreatedAt
    structMap["updated_at"] = g.UpdatedAt
    structMap["cycle"] = g.Cycle
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPeriodResponse.
// It customizes the JSON unmarshaling process for GetPeriodResponse objects.
func (g *GetPeriodResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        StartAt      *string                  `json:"start_at"`
        EndAt        *string                  `json:"end_at"`
        Id           *string                  `json:"id"`
        BillingAt    *string                  `json:"billing_at"`
        Subscription *GetSubscriptionResponse `json:"subscription"`
        Status       *string                  `json:"status"`
        Duration     *int                     `json:"duration"`
        CreatedAt    *string                  `json:"created_at"`
        UpdatedAt    *string                  `json:"updated_at"`
        Cycle        *int                     `json:"cycle"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    if temp.StartAt != nil {
        StartAtVal, err := time.Parse(time.RFC3339, *temp.StartAt)
        if err != nil {
            log.Fatalf("Cannot Parse start_at as % s format.", time.RFC3339)
        }
        g.StartAt = &StartAtVal
    }
    if temp.EndAt != nil {
        EndAtVal, err := time.Parse(time.RFC3339, *temp.EndAt)
        if err != nil {
            log.Fatalf("Cannot Parse end_at as % s format.", time.RFC3339)
        }
        g.EndAt = &EndAtVal
    }
    g.Id = temp.Id
    if temp.BillingAt != nil {
        BillingAtVal, err := time.Parse(time.RFC3339, *temp.BillingAt)
        if err != nil {
            log.Fatalf("Cannot Parse billing_at as % s format.", time.RFC3339)
        }
        g.BillingAt = &BillingAtVal
    }
    g.Subscription = temp.Subscription
    g.Status = temp.Status
    g.Duration = temp.Duration
    g.CreatedAt = temp.CreatedAt
    g.UpdatedAt = temp.UpdatedAt
    g.Cycle = temp.Cycle
    return nil
}
