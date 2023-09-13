package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetInvoiceResponse represents a GetInvoiceResponse struct.
// Response object for getting an invoice
type GetInvoiceResponse struct {
    Id             *string                       `json:"id"`
    Code           *string                       `json:"code"`
    Url            *string                       `json:"url"`
    Amount         *int                          `json:"amount"`
    Status         *string                       `json:"status"`
    PaymentMethod  *string                       `json:"payment_method"`
    CreatedAt      *time.Time                    `json:"created_at"`
    Items          []GetInvoiceItemResponse      `json:"items"`
    Customer       Optional[GetCustomerResponse] `json:"customer"`
    Charge         *GetChargeResponse            `json:"charge"`
    Installments   *int                          `json:"installments"`
    BillingAddress *GetBillingAddressResponse    `json:"billing_address"`
    Subscription   *GetSubscriptionResponse      `json:"subscription"`
    Cycle          Optional[GetPeriodResponse]   `json:"cycle"`
    Shipping       *GetShippingResponse          `json:"shipping"`
    Metadata       map[string]string             `json:"metadata"`
    DueAt          Optional[time.Time]           `json:"due_at"`
    CanceledAt     Optional[time.Time]           `json:"canceled_at"`
    BillingAt      Optional[time.Time]           `json:"billing_at"`
    SeenAt         Optional[time.Time]           `json:"seen_at"`
    // Total discounted value
    TotalDiscount  Optional[int]                 `json:"total_discount"`
    // Total discounted value
    TotalIncrement Optional[int]                 `json:"total_increment"`
    // Subscription Id
    SubscriptionId *string                       `json:"subscription_id"`
}

// MarshalJSON implements the json.Marshaler interface for GetInvoiceResponse.
// It customizes the JSON marshaling process for GetInvoiceResponse objects.
func (g *GetInvoiceResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetInvoiceResponse object to a map representation for JSON marshaling.
func (g *GetInvoiceResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["code"] = g.Code
    structMap["url"] = g.Url
    structMap["amount"] = g.Amount
    structMap["status"] = g.Status
    structMap["payment_method"] = g.PaymentMethod
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["items"] = g.Items
    if g.Customer.IsValueSet() {
        structMap["customer"] = g.Customer.Value()
    }
    structMap["charge"] = g.Charge
    structMap["installments"] = g.Installments
    structMap["billing_address"] = g.BillingAddress
    structMap["subscription"] = g.Subscription
    if g.Cycle.IsValueSet() {
        structMap["cycle"] = g.Cycle.Value()
    }
    structMap["shipping"] = g.Shipping
    structMap["metadata"] = g.Metadata
    if g.DueAt.IsValueSet() {
        var DueAtVal *string = nil
        if g.DueAt.Value() != nil {
            val := g.DueAt.Value().Format(time.RFC3339)
            DueAtVal = &val
        }
        structMap["due_at"] = DueAtVal
    }
    if g.CanceledAt.IsValueSet() {
        var CanceledAtVal *string = nil
        if g.CanceledAt.Value() != nil {
            val := g.CanceledAt.Value().Format(time.RFC3339)
            CanceledAtVal = &val
        }
        structMap["canceled_at"] = CanceledAtVal
    }
    if g.BillingAt.IsValueSet() {
        var BillingAtVal *string = nil
        if g.BillingAt.Value() != nil {
            val := g.BillingAt.Value().Format(time.RFC3339)
            BillingAtVal = &val
        }
        structMap["billing_at"] = BillingAtVal
    }
    if g.SeenAt.IsValueSet() {
        var SeenAtVal *string = nil
        if g.SeenAt.Value() != nil {
            val := g.SeenAt.Value().Format(time.RFC3339)
            SeenAtVal = &val
        }
        structMap["seen_at"] = SeenAtVal
    }
    if g.TotalDiscount.IsValueSet() {
        structMap["total_discount"] = g.TotalDiscount.Value()
    }
    if g.TotalIncrement.IsValueSet() {
        structMap["total_increment"] = g.TotalIncrement.Value()
    }
    structMap["subscription_id"] = g.SubscriptionId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetInvoiceResponse.
// It customizes the JSON unmarshaling process for GetInvoiceResponse objects.
func (g *GetInvoiceResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id             *string                       `json:"id"`
        Code           *string                       `json:"code"`
        Url            *string                       `json:"url"`
        Amount         *int                          `json:"amount"`
        Status         *string                       `json:"status"`
        PaymentMethod  *string                       `json:"payment_method"`
        CreatedAt      *string                       `json:"created_at"`
        Items          []GetInvoiceItemResponse      `json:"items"`
        Customer       Optional[GetCustomerResponse] `json:"customer"`
        Charge         *GetChargeResponse            `json:"charge"`
        Installments   *int                          `json:"installments"`
        BillingAddress *GetBillingAddressResponse    `json:"billing_address"`
        Subscription   *GetSubscriptionResponse      `json:"subscription"`
        Cycle          Optional[GetPeriodResponse]   `json:"cycle"`
        Shipping       *GetShippingResponse          `json:"shipping"`
        Metadata       map[string]string             `json:"metadata"`
        DueAt          Optional[string]              `json:"due_at"`
        CanceledAt     Optional[string]              `json:"canceled_at"`
        BillingAt      Optional[string]              `json:"billing_at"`
        SeenAt         Optional[string]              `json:"seen_at"`
        TotalDiscount  Optional[int]                 `json:"total_discount"`
        TotalIncrement Optional[int]                 `json:"total_increment"`
        SubscriptionId *string                       `json:"subscription_id"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Code = temp.Code
    g.Url = temp.Url
    g.Amount = temp.Amount
    g.Status = temp.Status
    g.PaymentMethod = temp.PaymentMethod
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        g.CreatedAt = &CreatedAtVal
    }
    g.Items = temp.Items
    g.Customer = temp.Customer
    g.Charge = temp.Charge
    g.Installments = temp.Installments
    g.BillingAddress = temp.BillingAddress
    g.Subscription = temp.Subscription
    g.Cycle = temp.Cycle
    g.Shipping = temp.Shipping
    g.Metadata = temp.Metadata
    g.DueAt.ShouldSetValue(temp.DueAt.IsValueSet())
    if temp.DueAt.Value() != nil {
        DueAtVal, err := time.Parse(time.RFC3339, (*temp.DueAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
        }
        g.DueAt.SetValue(&DueAtVal)
    }
    g.CanceledAt.ShouldSetValue(temp.CanceledAt.IsValueSet())
    if temp.CanceledAt.Value() != nil {
        CanceledAtVal, err := time.Parse(time.RFC3339, (*temp.CanceledAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
        }
        g.CanceledAt.SetValue(&CanceledAtVal)
    }
    g.BillingAt.ShouldSetValue(temp.BillingAt.IsValueSet())
    if temp.BillingAt.Value() != nil {
        BillingAtVal, err := time.Parse(time.RFC3339, (*temp.BillingAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse billing_at as % s format.", time.RFC3339)
        }
        g.BillingAt.SetValue(&BillingAtVal)
    }
    g.SeenAt.ShouldSetValue(temp.SeenAt.IsValueSet())
    if temp.SeenAt.Value() != nil {
        SeenAtVal, err := time.Parse(time.RFC3339, (*temp.SeenAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse seen_at as % s format.", time.RFC3339)
        }
        g.SeenAt.SetValue(&SeenAtVal)
    }
    g.TotalDiscount = temp.TotalDiscount
    g.TotalIncrement = temp.TotalIncrement
    g.SubscriptionId = temp.SubscriptionId
    return nil
}
