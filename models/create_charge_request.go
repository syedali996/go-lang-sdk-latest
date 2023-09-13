package models

import (
    "encoding/json"
    "log"
    "time"
)

// CreateChargeRequest represents a CreateChargeRequest struct.
// Request for creating a new charge
type CreateChargeRequest struct {
    // Code
    Code       string                 `json:"code"`
    // The amount of the charge, in cents
    Amount     int                    `json:"amount"`
    // The customer's id
    CustomerId string                 `json:"customer_id"`
    // Customer data
    Customer   CreateCustomerRequest  `json:"customer"`
    // Payment data
    Payment    CreatePaymentRequest   `json:"payment"`
    // Metadata
    Metadata   map[string]string      `json:"metadata"`
    // The charge due date
    DueAt      *time.Time             `json:"due_at,omitempty"`
    Antifraud  CreateAntifraudRequest `json:"antifraud"`
    // Order Id
    OrderId    string                 `json:"order_id"`
}

// MarshalJSON implements the json.Marshaler interface for CreateChargeRequest.
// It customizes the JSON marshaling process for CreateChargeRequest objects.
func (c *CreateChargeRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateChargeRequest object to a map representation for JSON marshaling.
func (c *CreateChargeRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["code"] = c.Code
    structMap["amount"] = c.Amount
    structMap["customer_id"] = c.CustomerId
    structMap["customer"] = c.Customer
    structMap["payment"] = c.Payment
    structMap["metadata"] = c.Metadata
    if c.DueAt != nil {
        structMap["due_at"] = c.DueAt.Format(time.RFC3339)
    }
    structMap["antifraud"] = c.Antifraud
    structMap["order_id"] = c.OrderId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateChargeRequest.
// It customizes the JSON unmarshaling process for CreateChargeRequest objects.
func (c *CreateChargeRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Code       string                 `json:"code"`
        Amount     int                    `json:"amount"`
        CustomerId string                 `json:"customer_id"`
        Customer   CreateCustomerRequest  `json:"customer"`
        Payment    CreatePaymentRequest   `json:"payment"`
        Metadata   map[string]string      `json:"metadata"`
        DueAt      *string                `json:"due_at,omitempty"`
        Antifraud  CreateAntifraudRequest `json:"antifraud"`
        OrderId    string                 `json:"order_id"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Code = temp.Code
    c.Amount = temp.Amount
    c.CustomerId = temp.CustomerId
    c.Customer = temp.Customer
    c.Payment = temp.Payment
    c.Metadata = temp.Metadata
    if temp.DueAt != nil {
        DueAtVal, err := time.Parse(time.RFC3339, *temp.DueAt)
        if err != nil {
            log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
        }
        c.DueAt = &DueAtVal
    }
    c.Antifraud = temp.Antifraud
    c.OrderId = temp.OrderId
    return nil
}
