package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetWithdrawResponse represents a GetWithdrawResponse struct.
type GetWithdrawResponse struct {
    Id                   *string                    `json:"id"`
    GatewayId            *string                    `json:"gateway_id"`
    Amount               *int                       `json:"amount"`
    Status               *string                    `json:"status"`
    CreatedAt            *time.Time                 `json:"created_at"`
    UpdatedAt            *time.Time                 `json:"updated_at"`
    Metadata             Optional[[]string]         `json:"metadata"`
    Fee                  Optional[int]              `json:"fee"`
    FundingDate          Optional[time.Time]        `json:"funding_date"`
    FundingEstimatedDate Optional[time.Time]        `json:"funding_estimated_date"`
    Type                 *string                    `json:"type"`
    Source               *GetWithdrawSourceResponse `json:"source"`
    Target               *GetWithdrawTargetResponse `json:"target"`
}

// MarshalJSON implements the json.Marshaler interface for GetWithdrawResponse.
// It customizes the JSON marshaling process for GetWithdrawResponse objects.
func (g *GetWithdrawResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetWithdrawResponse object to a map representation for JSON marshaling.
func (g *GetWithdrawResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["gateway_id"] = g.GatewayId
    structMap["amount"] = g.Amount
    structMap["status"] = g.Status
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    if g.Metadata.IsValueSet() {
        structMap["metadata"] = g.Metadata.Value()
    }
    if g.Fee.IsValueSet() {
        structMap["fee"] = g.Fee.Value()
    }
    if g.FundingDate.IsValueSet() {
        var FundingDateVal *string = nil
        if g.FundingDate.Value() != nil {
            val := g.FundingDate.Value().Format(time.RFC3339)
            FundingDateVal = &val
        }
        structMap["funding_date"] = FundingDateVal
    }
    if g.FundingEstimatedDate.IsValueSet() {
        var FundingEstimatedDateVal *string = nil
        if g.FundingEstimatedDate.Value() != nil {
            val := g.FundingEstimatedDate.Value().Format(time.RFC3339)
            FundingEstimatedDateVal = &val
        }
        structMap["funding_estimated_date"] = FundingEstimatedDateVal
    }
    structMap["type"] = g.Type
    structMap["source"] = g.Source
    structMap["target"] = g.Target
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetWithdrawResponse.
// It customizes the JSON unmarshaling process for GetWithdrawResponse objects.
func (g *GetWithdrawResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                   *string                    `json:"id"`
        GatewayId            *string                    `json:"gateway_id"`
        Amount               *int                       `json:"amount"`
        Status               *string                    `json:"status"`
        CreatedAt            *string                    `json:"created_at"`
        UpdatedAt            *string                    `json:"updated_at"`
        Metadata             Optional[[]string]         `json:"metadata"`
        Fee                  Optional[int]              `json:"fee"`
        FundingDate          Optional[string]           `json:"funding_date"`
        FundingEstimatedDate Optional[string]           `json:"funding_estimated_date"`
        Type                 *string                    `json:"type"`
        Source               *GetWithdrawSourceResponse `json:"source"`
        Target               *GetWithdrawTargetResponse `json:"target"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.GatewayId = temp.GatewayId
    g.Amount = temp.Amount
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
    g.Metadata = temp.Metadata
    g.Fee = temp.Fee
    g.FundingDate.ShouldSetValue(temp.FundingDate.IsValueSet())
    if temp.FundingDate.Value() != nil {
        FundingDateVal, err := time.Parse(time.RFC3339, (*temp.FundingDate.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse funding_date as % s format.", time.RFC3339)
        }
        g.FundingDate.SetValue(&FundingDateVal)
    }
    g.FundingEstimatedDate.ShouldSetValue(temp.FundingEstimatedDate.IsValueSet())
    if temp.FundingEstimatedDate.Value() != nil {
        FundingEstimatedDateVal, err := time.Parse(time.RFC3339, (*temp.FundingEstimatedDate.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse funding_estimated_date as % s format.", time.RFC3339)
        }
        g.FundingEstimatedDate.SetValue(&FundingEstimatedDateVal)
    }
    g.Type = temp.Type
    g.Source = temp.Source
    g.Target = temp.Target
    return nil
}
