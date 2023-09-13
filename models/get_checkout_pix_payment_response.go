package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetCheckoutPixPaymentResponse represents a GetCheckoutPixPaymentResponse struct.
// Checkout pix payment response
type GetCheckoutPixPaymentResponse struct {
    // Expires at
    ExpiresAt             *time.Time                 `json:"expires_at"`
    // Additional information
    AdditionalInformation []PixAdditionalInformation `json:"additional_information"`
}

// MarshalJSON implements the json.Marshaler interface for GetCheckoutPixPaymentResponse.
// It customizes the JSON marshaling process for GetCheckoutPixPaymentResponse objects.
func (g *GetCheckoutPixPaymentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetCheckoutPixPaymentResponse object to a map representation for JSON marshaling.
func (g *GetCheckoutPixPaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["expires_at"] = g.ExpiresAt.Format(time.RFC3339)
    structMap["additional_information"] = g.AdditionalInformation
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCheckoutPixPaymentResponse.
// It customizes the JSON unmarshaling process for GetCheckoutPixPaymentResponse objects.
func (g *GetCheckoutPixPaymentResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        ExpiresAt             *string                    `json:"expires_at"`
        AdditionalInformation []PixAdditionalInformation `json:"additional_information"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    if temp.ExpiresAt != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        g.ExpiresAt = &ExpiresAtVal
    }
    g.AdditionalInformation = temp.AdditionalInformation
    return nil
}
