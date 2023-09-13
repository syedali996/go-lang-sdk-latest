package models

import (
    "encoding/json"
)

// GetInterestResponse represents a GetInterestResponse struct.
// Interest Response
type GetInterestResponse struct {
    // Days
    Days   *int    `json:"days"`
    // Type
    Type   *string `json:"type"`
    // Amount
    Amount *int    `json:"amount"`
}

// MarshalJSON implements the json.Marshaler interface for GetInterestResponse.
// It customizes the JSON marshaling process for GetInterestResponse objects.
func (g *GetInterestResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetInterestResponse object to a map representation for JSON marshaling.
func (g *GetInterestResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["days"] = g.Days
    structMap["type"] = g.Type
    structMap["amount"] = g.Amount
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetInterestResponse.
// It customizes the JSON unmarshaling process for GetInterestResponse objects.
func (g *GetInterestResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Days   *int    `json:"days"`
        Type   *string `json:"type"`
        Amount *int    `json:"amount"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Days = temp.Days
    g.Type = temp.Type
    g.Amount = temp.Amount
    return nil
}
