package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetCheckoutBoletoPaymentResponse represents a GetCheckoutBoletoPaymentResponse struct.
type GetCheckoutBoletoPaymentResponse struct {
    // Data de vencimento do boleto
    DueAt        *time.Time `json:"due_at"`
    // Instruções do boleto
    Instructions *string    `json:"instructions"`
}

// MarshalJSON implements the json.Marshaler interface for GetCheckoutBoletoPaymentResponse.
// It customizes the JSON marshaling process for GetCheckoutBoletoPaymentResponse objects.
func (g *GetCheckoutBoletoPaymentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetCheckoutBoletoPaymentResponse object to a map representation for JSON marshaling.
func (g *GetCheckoutBoletoPaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["due_at"] = g.DueAt.Format(time.RFC3339)
    structMap["instructions"] = g.Instructions
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCheckoutBoletoPaymentResponse.
// It customizes the JSON unmarshaling process for GetCheckoutBoletoPaymentResponse objects.
func (g *GetCheckoutBoletoPaymentResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        DueAt        *string `json:"due_at"`
        Instructions *string `json:"instructions"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    if temp.DueAt != nil {
        DueAtVal, err := time.Parse(time.RFC3339, *temp.DueAt)
        if err != nil {
            log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
        }
        g.DueAt = &DueAtVal
    }
    g.Instructions = temp.Instructions
    return nil
}
