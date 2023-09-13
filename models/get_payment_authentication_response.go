package models

import (
    "encoding/json"
)

// GetPaymentAuthenticationResponse represents a GetPaymentAuthenticationResponse struct.
// Payment Authentication response
type GetPaymentAuthenticationResponse struct {
    Type         *string                  `json:"type"`
    // 3D-S payment authentication response
    ThreedSecure *GetThreeDSecureResponse `json:"threed_secure"`
}

// MarshalJSON implements the json.Marshaler interface for GetPaymentAuthenticationResponse.
// It customizes the JSON marshaling process for GetPaymentAuthenticationResponse objects.
func (g *GetPaymentAuthenticationResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetPaymentAuthenticationResponse object to a map representation for JSON marshaling.
func (g *GetPaymentAuthenticationResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["type"] = g.Type
    structMap["threed_secure"] = g.ThreedSecure
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPaymentAuthenticationResponse.
// It customizes the JSON unmarshaling process for GetPaymentAuthenticationResponse objects.
func (g *GetPaymentAuthenticationResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Type         *string                  `json:"type"`
        ThreedSecure *GetThreeDSecureResponse `json:"threed_secure"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Type = temp.Type
    g.ThreedSecure = temp.ThreedSecure
    return nil
}
