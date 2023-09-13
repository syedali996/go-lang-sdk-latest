package models

import (
    "encoding/json"
)

// GetGatewayRecipientResponse represents a GetGatewayRecipientResponse struct.
// Information about the recipient on the gateway
type GetGatewayRecipientResponse struct {
    // Gateway name
    Gateway   *string `json:"gateway"`
    // Status of the recipient on the gateway
    Status    *string `json:"status"`
    // Recipient id on the gateway
    Pgid      *string `json:"pgid"`
    // Creation date
    CreatedAt *string `json:"created_at"`
    // Last update date
    UpdatedAt *string `json:"updated_at"`
}

// MarshalJSON implements the json.Marshaler interface for GetGatewayRecipientResponse.
// It customizes the JSON marshaling process for GetGatewayRecipientResponse objects.
func (g *GetGatewayRecipientResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetGatewayRecipientResponse object to a map representation for JSON marshaling.
func (g *GetGatewayRecipientResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["gateway"] = g.Gateway
    structMap["status"] = g.Status
    structMap["pgid"] = g.Pgid
    structMap["created_at"] = g.CreatedAt
    structMap["updated_at"] = g.UpdatedAt
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetGatewayRecipientResponse.
// It customizes the JSON unmarshaling process for GetGatewayRecipientResponse objects.
func (g *GetGatewayRecipientResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Gateway   *string `json:"gateway"`
        Status    *string `json:"status"`
        Pgid      *string `json:"pgid"`
        CreatedAt *string `json:"created_at"`
        UpdatedAt *string `json:"updated_at"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Gateway = temp.Gateway
    g.Status = temp.Status
    g.Pgid = temp.Pgid
    g.CreatedAt = temp.CreatedAt
    g.UpdatedAt = temp.UpdatedAt
    return nil
}
