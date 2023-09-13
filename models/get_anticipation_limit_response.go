package models

import (
    "encoding/json"
)

// GetAnticipationLimitResponse represents a GetAnticipationLimitResponse struct.
// Anticipation limit
type GetAnticipationLimitResponse struct {
    // Amount
    Amount          *int `json:"amount"`
    // Anticipation fee
    AnticipationFee *int `json:"anticipation_fee"`
}

// MarshalJSON implements the json.Marshaler interface for GetAnticipationLimitResponse.
// It customizes the JSON marshaling process for GetAnticipationLimitResponse objects.
func (g *GetAnticipationLimitResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetAnticipationLimitResponse object to a map representation for JSON marshaling.
func (g *GetAnticipationLimitResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["amount"] = g.Amount
    structMap["anticipation_fee"] = g.AnticipationFee
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetAnticipationLimitResponse.
// It customizes the JSON unmarshaling process for GetAnticipationLimitResponse objects.
func (g *GetAnticipationLimitResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Amount          *int `json:"amount"`
        AnticipationFee *int `json:"anticipation_fee"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Amount = temp.Amount
    g.AnticipationFee = temp.AnticipationFee
    return nil
}
