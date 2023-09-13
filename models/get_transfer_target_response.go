package models

import (
    "encoding/json"
)

// GetTransferTargetResponse represents a GetTransferTargetResponse struct.
type GetTransferTargetResponse struct {
    TargetId *string `json:"target_id"`
    Type     *string `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for GetTransferTargetResponse.
// It customizes the JSON marshaling process for GetTransferTargetResponse objects.
func (g *GetTransferTargetResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetTransferTargetResponse object to a map representation for JSON marshaling.
func (g *GetTransferTargetResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["target_id"] = g.TargetId
    structMap["type"] = g.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetTransferTargetResponse.
// It customizes the JSON unmarshaling process for GetTransferTargetResponse objects.
func (g *GetTransferTargetResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        TargetId *string `json:"target_id"`
        Type     *string `json:"type"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.TargetId = temp.TargetId
    g.Type = temp.Type
    return nil
}
