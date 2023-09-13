package models

import (
    "encoding/json"
)

// GetCheckoutCreditCardPaymentResponse represents a GetCheckoutCreditCardPaymentResponse struct.
type GetCheckoutCreditCardPaymentResponse struct {
    // Descrição na fatura
    StatementDescriptor *string                                     `json:"statementDescriptor"`
    // Parcelas
    Installments        []GetCheckoutCardInstallmentOptionsResponse `json:"installments"`
    // Payment Authentication response
    Authentication      *GetPaymentAuthenticationResponse           `json:"authentication"`
}

// MarshalJSON implements the json.Marshaler interface for GetCheckoutCreditCardPaymentResponse.
// It customizes the JSON marshaling process for GetCheckoutCreditCardPaymentResponse objects.
func (g *GetCheckoutCreditCardPaymentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetCheckoutCreditCardPaymentResponse object to a map representation for JSON marshaling.
func (g *GetCheckoutCreditCardPaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["statementDescriptor"] = g.StatementDescriptor
    structMap["installments"] = g.Installments
    structMap["authentication"] = g.Authentication
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCheckoutCreditCardPaymentResponse.
// It customizes the JSON unmarshaling process for GetCheckoutCreditCardPaymentResponse objects.
func (g *GetCheckoutCreditCardPaymentResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        StatementDescriptor *string                                     `json:"statementDescriptor"`
        Installments        []GetCheckoutCardInstallmentOptionsResponse `json:"installments"`
        Authentication      *GetPaymentAuthenticationResponse           `json:"authentication"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.StatementDescriptor = temp.StatementDescriptor
    g.Installments = temp.Installments
    g.Authentication = temp.Authentication
    return nil
}
