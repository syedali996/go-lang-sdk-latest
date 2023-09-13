package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetAnticipationResponse represents a GetAnticipationResponse struct.
// Anticipation
type GetAnticipationResponse struct {
    // Id
    Id              *string                        `json:"id"`
    // Requested amount
    RequestedAmount *int                           `json:"requested_amount"`
    // Approved amount
    ApprovedAmount  *int                           `json:"approved_amount"`
    // Recipient
    Recipient       Optional[GetRecipientResponse] `json:"recipient"`
    // Anticipation id on the gateway
    Pgid            *string                        `json:"pgid"`
    // Creation date
    CreatedAt       *time.Time                     `json:"created_at"`
    // Last update date
    UpdatedAt       *time.Time                     `json:"updated_at"`
    // Payment date
    PaymentDate     *time.Time                     `json:"payment_date"`
    // Status
    Status          *string                        `json:"status"`
    // Timeframe
    Timeframe       *string                        `json:"timeframe"`
}

// MarshalJSON implements the json.Marshaler interface for GetAnticipationResponse.
// It customizes the JSON marshaling process for GetAnticipationResponse objects.
func (g *GetAnticipationResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetAnticipationResponse object to a map representation for JSON marshaling.
func (g *GetAnticipationResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["requested_amount"] = g.RequestedAmount
    structMap["approved_amount"] = g.ApprovedAmount
    if g.Recipient.IsValueSet() {
        structMap["recipient"] = g.Recipient.Value()
    }
    structMap["pgid"] = g.Pgid
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    structMap["payment_date"] = g.PaymentDate.Format(time.RFC3339)
    structMap["status"] = g.Status
    structMap["timeframe"] = g.Timeframe
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetAnticipationResponse.
// It customizes the JSON unmarshaling process for GetAnticipationResponse objects.
func (g *GetAnticipationResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id              *string                        `json:"id"`
        RequestedAmount *int                           `json:"requested_amount"`
        ApprovedAmount  *int                           `json:"approved_amount"`
        Recipient       Optional[GetRecipientResponse] `json:"recipient"`
        Pgid            *string                        `json:"pgid"`
        CreatedAt       *string                        `json:"created_at"`
        UpdatedAt       *string                        `json:"updated_at"`
        PaymentDate     *string                        `json:"payment_date"`
        Status          *string                        `json:"status"`
        Timeframe       *string                        `json:"timeframe"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.RequestedAmount = temp.RequestedAmount
    g.ApprovedAmount = temp.ApprovedAmount
    g.Recipient = temp.Recipient
    g.Pgid = temp.Pgid
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
    if temp.PaymentDate != nil {
        PaymentDateVal, err := time.Parse(time.RFC3339, *temp.PaymentDate)
        if err != nil {
            log.Fatalf("Cannot Parse payment_date as % s format.", time.RFC3339)
        }
        g.PaymentDate = &PaymentDateVal
    }
    g.Status = temp.Status
    g.Timeframe = temp.Timeframe
    return nil
}
