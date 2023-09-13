package models

import (
    "encoding/json"
)

// GetFineResponse represents a GetFineResponse struct.
// Fine Response
type GetFineResponse struct {
    // Days
    Days   *int    `json:"days"`
    // Type
    Type   *string `json:"type"`
    // Amount
    Amount *int    `json:"amount"`
}

// MarshalJSON implements the json.Marshaler interface for GetFineResponse.
// It customizes the JSON marshaling process for GetFineResponse objects.
func (g *GetFineResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetFineResponse object to a map representation for JSON marshaling.
func (g *GetFineResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["days"] = g.Days
    structMap["type"] = g.Type
    structMap["amount"] = g.Amount
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetFineResponse.
// It customizes the JSON unmarshaling process for GetFineResponse objects.
func (g *GetFineResponse) UnmarshalJSON(input []byte) error {
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
