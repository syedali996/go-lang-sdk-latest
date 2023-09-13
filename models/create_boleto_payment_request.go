package models

import (
    "encoding/json"
    "log"
    "time"
)

// CreateBoletoPaymentRequest represents a CreateBoletoPaymentRequest struct.
// Contains the settings for creating a boleto payment
type CreateBoletoPaymentRequest struct {
    // Number of retries
    Retries             int                    `json:"retries"`
    // The bank code, containing three characters. The available codes are on the API specification
    Bank                string                 `json:"bank"`
    // The instructions field that will be printed on the boleto.
    Instructions        string                 `json:"instructions"`
    // Boleto due date
    DueAt               *time.Time             `json:"due_at,omitempty"`
    // Card's billing address
    BillingAddress      CreateAddressRequest   `json:"billing_address"`
    // The address id for the billing address
    BillingAddressId    string                 `json:"billing_address_id"`
    // Customer identification number with the bank
    NossoNumero         *string                `json:"nosso_numero,omitempty"`
    // Boleto identification
    DocumentNumber      string                 `json:"document_number"`
    // Soft Descriptor
    StatementDescriptor string                 `json:"statement_descriptor"`
    Interest            *CreateInterestRequest `json:"interest,omitempty"`
    Fine                *CreateFineRequest     `json:"fine,omitempty"`
    MaxDaysToPayPastDue *int                   `json:"max_days_to_pay_past_due,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateBoletoPaymentRequest.
// It customizes the JSON marshaling process for CreateBoletoPaymentRequest objects.
func (c *CreateBoletoPaymentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateBoletoPaymentRequest object to a map representation for JSON marshaling.
func (c *CreateBoletoPaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["retries"] = c.Retries
    structMap["bank"] = c.Bank
    structMap["instructions"] = c.Instructions
    if c.DueAt != nil {
        structMap["due_at"] = c.DueAt.Format(time.RFC3339)
    }
    structMap["billing_address"] = c.BillingAddress
    structMap["billing_address_id"] = c.BillingAddressId
    if c.NossoNumero != nil {
        structMap["nosso_numero"] = c.NossoNumero
    }
    structMap["document_number"] = c.DocumentNumber
    structMap["statement_descriptor"] = c.StatementDescriptor
    if c.Interest != nil {
        structMap["interest"] = c.Interest
    }
    if c.Fine != nil {
        structMap["fine"] = c.Fine
    }
    if c.MaxDaysToPayPastDue != nil {
        structMap["max_days_to_pay_past_due"] = c.MaxDaysToPayPastDue
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateBoletoPaymentRequest.
// It customizes the JSON unmarshaling process for CreateBoletoPaymentRequest objects.
func (c *CreateBoletoPaymentRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Retries             int                    `json:"retries"`
        Bank                string                 `json:"bank"`
        Instructions        string                 `json:"instructions"`
        DueAt               *string                `json:"due_at,omitempty"`
        BillingAddress      CreateAddressRequest   `json:"billing_address"`
        BillingAddressId    string                 `json:"billing_address_id"`
        NossoNumero         *string                `json:"nosso_numero,omitempty"`
        DocumentNumber      string                 `json:"document_number"`
        StatementDescriptor string                 `json:"statement_descriptor"`
        Interest            *CreateInterestRequest `json:"interest,omitempty"`
        Fine                *CreateFineRequest     `json:"fine,omitempty"`
        MaxDaysToPayPastDue *int                   `json:"max_days_to_pay_past_due,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Retries = temp.Retries
    c.Bank = temp.Bank
    c.Instructions = temp.Instructions
    if temp.DueAt != nil {
        DueAtVal, err := time.Parse(time.RFC3339, *temp.DueAt)
        if err != nil {
            log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
        }
        c.DueAt = &DueAtVal
    }
    c.BillingAddress = temp.BillingAddress
    c.BillingAddressId = temp.BillingAddressId
    c.NossoNumero = temp.NossoNumero
    c.DocumentNumber = temp.DocumentNumber
    c.StatementDescriptor = temp.StatementDescriptor
    c.Interest = temp.Interest
    c.Fine = temp.Fine
    c.MaxDaysToPayPastDue = temp.MaxDaysToPayPastDue
    return nil
}
