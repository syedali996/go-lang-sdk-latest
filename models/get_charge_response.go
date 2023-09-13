package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetChargeResponse represents a GetChargeResponse struct.
// Response object for getting a charge
type GetChargeResponse struct {
    Id                  *string                         `json:"id"`
    Code                *string                         `json:"code"`
    GatewayId           *string                         `json:"gateway_id"`
    Amount              *int                            `json:"amount"`
    Status              *string                         `json:"status"`
    Currency            *string                         `json:"currency"`
    PaymentMethod       *string                         `json:"payment_method"`
    DueAt               *time.Time                      `json:"due_at"`
    CreatedAt           *time.Time                      `json:"created_at"`
    UpdatedAt           *time.Time                      `json:"updated_at"`
    LastTransaction     GetTransactionResponseInterface `json:"last_transaction"`
    Invoice             Optional[GetInvoiceResponse]    `json:"invoice"`
    Order               Optional[GetOrderResponse]      `json:"order"`
    Customer            Optional[GetCustomerResponse]   `json:"customer"`
    Metadata            map[string]string               `json:"metadata"`
    PaidAt              Optional[time.Time]             `json:"paid_at"`
    CanceledAt          Optional[time.Time]             `json:"canceled_at"`
    // Canceled Amount
    CanceledAmount      *int                            `json:"canceled_amount"`
    // Paid amount
    PaidAmount          *int                            `json:"paid_amount"`
    // interest and fine paid
    InterestAndFinePaid Optional[int]                   `json:"interest_and_fine_paid"`
    // Defines whether the card has been used one or more times.
    RecurrencyCycle     Optional[string]                `json:"recurrency_cycle"`
}

// MarshalJSON implements the json.Marshaler interface for GetChargeResponse.
// It customizes the JSON marshaling process for GetChargeResponse objects.
func (g *GetChargeResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetChargeResponse object to a map representation for JSON marshaling.
func (g *GetChargeResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["code"] = g.Code
    structMap["gateway_id"] = g.GatewayId
    structMap["amount"] = g.Amount
    structMap["status"] = g.Status
    structMap["currency"] = g.Currency
    structMap["payment_method"] = g.PaymentMethod
    structMap["due_at"] = g.DueAt.Format(time.RFC3339)
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    structMap["last_transaction"] = g.LastTransaction
    if g.Invoice.IsValueSet() {
        structMap["invoice"] = g.Invoice.Value()
    }
    if g.Order.IsValueSet() {
        structMap["order"] = g.Order.Value()
    }
    if g.Customer.IsValueSet() {
        structMap["customer"] = g.Customer.Value()
    }
    structMap["metadata"] = g.Metadata
    if g.PaidAt.IsValueSet() {
        var PaidAtVal *string = nil
        if g.PaidAt.Value() != nil {
            val := g.PaidAt.Value().Format(time.RFC3339)
            PaidAtVal = &val
        }
        structMap["paid_at"] = PaidAtVal
    }
    if g.CanceledAt.IsValueSet() {
        var CanceledAtVal *string = nil
        if g.CanceledAt.Value() != nil {
            val := g.CanceledAt.Value().Format(time.RFC3339)
            CanceledAtVal = &val
        }
        structMap["canceled_at"] = CanceledAtVal
    }
    structMap["canceled_amount"] = g.CanceledAmount
    structMap["paid_amount"] = g.PaidAmount
    if g.InterestAndFinePaid.IsValueSet() {
        structMap["interest_and_fine_paid"] = g.InterestAndFinePaid.Value()
    }
    if g.RecurrencyCycle.IsValueSet() {
        structMap["recurrency_cycle"] = g.RecurrencyCycle.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetChargeResponse.
// It customizes the JSON unmarshaling process for GetChargeResponse objects.
func (g *GetChargeResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                  *string                         `json:"id"`
        Code                *string                         `json:"code"`
        GatewayId           *string                         `json:"gateway_id"`
        Amount              *int                            `json:"amount"`
        Status              *string                         `json:"status"`
        Currency            *string                         `json:"currency"`
        PaymentMethod       *string                         `json:"payment_method"`
        DueAt               *string                         `json:"due_at"`
        CreatedAt           *string                         `json:"created_at"`
        UpdatedAt           *string                         `json:"updated_at"`
        LastTransaction     GetTransactionResponseInterface `json:"last_transaction"`
        Invoice             Optional[GetInvoiceResponse]    `json:"invoice"`
        Order               Optional[GetOrderResponse]      `json:"order"`
        Customer            Optional[GetCustomerResponse]   `json:"customer"`
        Metadata            map[string]string               `json:"metadata"`
        PaidAt              Optional[string]                `json:"paid_at"`
        CanceledAt          Optional[string]                `json:"canceled_at"`
        CanceledAmount      *int                            `json:"canceled_amount"`
        PaidAmount          *int                            `json:"paid_amount"`
        InterestAndFinePaid Optional[int]                   `json:"interest_and_fine_paid"`
        RecurrencyCycle     Optional[string]                `json:"recurrency_cycle"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Code = temp.Code
    g.GatewayId = temp.GatewayId
    g.Amount = temp.Amount
    g.Status = temp.Status
    g.Currency = temp.Currency
    g.PaymentMethod = temp.PaymentMethod
    if temp.DueAt != nil {
        DueAtVal, err := time.Parse(time.RFC3339, *temp.DueAt)
        if err != nil {
            log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
        }
        g.DueAt = &DueAtVal
    }
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        g.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        g.UpdatedAt = &UpdatedAtVal
    }
    g.LastTransaction = temp.LastTransaction
    g.Invoice = temp.Invoice
    g.Order = temp.Order
    g.Customer = temp.Customer
    g.Metadata = temp.Metadata
    g.PaidAt.ShouldSetValue(temp.PaidAt.IsValueSet())
    if temp.PaidAt.Value() != nil {
        PaidAtVal, err := time.Parse(time.RFC3339, (*temp.PaidAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse paid_at as % s format.", time.RFC3339)
        }
        g.PaidAt.SetValue(&PaidAtVal)
    }
    g.CanceledAt.ShouldSetValue(temp.CanceledAt.IsValueSet())
    if temp.CanceledAt.Value() != nil {
        CanceledAtVal, err := time.Parse(time.RFC3339, (*temp.CanceledAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
        }
        g.CanceledAt.SetValue(&CanceledAtVal)
    }
    g.CanceledAmount = temp.CanceledAmount
    g.PaidAmount = temp.PaidAmount
    g.InterestAndFinePaid = temp.InterestAndFinePaid
    g.RecurrencyCycle = temp.RecurrencyCycle
    return nil
}
