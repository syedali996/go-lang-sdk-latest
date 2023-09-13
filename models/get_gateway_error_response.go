package models

import (
    "encoding/json"
)

// GetGatewayErrorResponse represents a GetGatewayErrorResponse struct.
// Gateway Response
type GetGatewayErrorResponse struct {
    // The message error
    Message *string `json:"message"`
}

// MarshalJSON implements the json.Marshaler interface for GetGatewayErrorResponse.
// It customizes the JSON marshaling process for GetGatewayErrorResponse objects.
func (g *GetGatewayErrorResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetGatewayErrorResponse object to a map representation for JSON marshaling.
func (g *GetGatewayErrorResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["message"] = g.Message
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetGatewayErrorResponse.
// It customizes the JSON unmarshaling process for GetGatewayErrorResponse objects.
func (g *GetGatewayErrorResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Message *string `json:"message"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Message = temp.Message
    return nil
}
