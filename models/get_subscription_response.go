package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetSubscriptionResponse represents a GetSubscriptionResponse struct.
type GetSubscriptionResponse struct {
    Id                   *string                                 `json:"id"`
    Code                 *string                                 `json:"code"`
    StartAt              *time.Time                              `json:"start_at"`
    Interval             *string                                 `json:"interval"`
    IntervalCount        *int                                    `json:"interval_count"`
    BillingType          *string                                 `json:"billing_type"`
    CurrentCycle         Optional[GetPeriodResponse]             `json:"current_cycle"`
    PaymentMethod        *string                                 `json:"payment_method"`
    Currency             *string                                 `json:"currency"`
    Installments         *int                                    `json:"installments"`
    Status               *string                                 `json:"status"`
    CreatedAt            *time.Time                              `json:"created_at"`
    UpdatedAt            *time.Time                              `json:"updated_at"`
    Customer             Optional[GetCustomerResponse]           `json:"customer"`
    Card                 *GetCardResponse                        `json:"card"`
    Items                []GetSubscriptionItemResponse           `json:"items"`
    StatementDescriptor  *string                                 `json:"statement_descriptor"`
    Metadata             map[string]string                       `json:"metadata"`
    Setup                *GetSetupResponse                       `json:"setup"`
    // Affiliation Code
    GatewayAffiliationId *string                                 `json:"gateway_affiliation_id"`
    NextBillingAt        Optional[time.Time]                     `json:"next_billing_at"`
    BillingDay           Optional[int]                           `json:"billing_day"`
    MinimumPrice         Optional[int]                           `json:"minimum_price"`
    CanceledAt           Optional[time.Time]                     `json:"canceled_at"`
    // Subscription discounts
    Discounts            Optional[[]GetDiscountResponse]         `json:"discounts"`
    // Subscription increments
    Increments           []GetIncrementResponse                  `json:"increments"`
    // Days until boleto expires
    BoletoDueDays        Optional[int]                           `json:"boleto_due_days"`
    // Subscription's split response
    Split                *GetSubscriptionSplitResponse           `json:"split"`
    Boleto               Optional[GetSubscriptionBoletoResponse] `json:"boleto"`
}

// MarshalJSON implements the json.Marshaler interface for GetSubscriptionResponse.
// It customizes the JSON marshaling process for GetSubscriptionResponse objects.
func (g *GetSubscriptionResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetSubscriptionResponse object to a map representation for JSON marshaling.
func (g *GetSubscriptionResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["code"] = g.Code
    structMap["start_at"] = g.StartAt.Format(time.RFC3339)
    structMap["interval"] = g.Interval
    structMap["interval_count"] = g.IntervalCount
    structMap["billing_type"] = g.BillingType
    if g.CurrentCycle.IsValueSet() {
        structMap["current_cycle"] = g.CurrentCycle.Value()
    }
    structMap["payment_method"] = g.PaymentMethod
    structMap["currency"] = g.Currency
    structMap["installments"] = g.Installments
    structMap["status"] = g.Status
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    if g.Customer.IsValueSet() {
        structMap["customer"] = g.Customer.Value()
    }
    structMap["card"] = g.Card
    structMap["items"] = g.Items
    structMap["statement_descriptor"] = g.StatementDescriptor
    structMap["metadata"] = g.Metadata
    structMap["setup"] = g.Setup
    structMap["gateway_affiliation_id"] = g.GatewayAffiliationId
    if g.NextBillingAt.IsValueSet() {
        var NextBillingAtVal *string = nil
        if g.NextBillingAt.Value() != nil {
            val := g.NextBillingAt.Value().Format(time.RFC3339)
            NextBillingAtVal = &val
        }
        structMap["next_billing_at"] = NextBillingAtVal
    }
    if g.BillingDay.IsValueSet() {
        structMap["billing_day"] = g.BillingDay.Value()
    }
    if g.MinimumPrice.IsValueSet() {
        structMap["minimum_price"] = g.MinimumPrice.Value()
    }
    if g.CanceledAt.IsValueSet() {
        var CanceledAtVal *string = nil
        if g.CanceledAt.Value() != nil {
            val := g.CanceledAt.Value().Format(time.RFC3339)
            CanceledAtVal = &val
        }
        structMap["canceled_at"] = CanceledAtVal
    }
    if g.Discounts.IsValueSet() {
        structMap["discounts"] = g.Discounts.Value()
    }
    structMap["increments"] = g.Increments
    if g.BoletoDueDays.IsValueSet() {
        structMap["boleto_due_days"] = g.BoletoDueDays.Value()
    }
    structMap["split"] = g.Split
    if g.Boleto.IsValueSet() {
        structMap["boleto"] = g.Boleto.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSubscriptionResponse.
// It customizes the JSON unmarshaling process for GetSubscriptionResponse objects.
func (g *GetSubscriptionResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                   *string                                 `json:"id"`
        Code                 *string                                 `json:"code"`
        StartAt              *string                                 `json:"start_at"`
        Interval             *string                                 `json:"interval"`
        IntervalCount        *int                                    `json:"interval_count"`
        BillingType          *string                                 `json:"billing_type"`
        CurrentCycle         Optional[GetPeriodResponse]             `json:"current_cycle"`
        PaymentMethod        *string                                 `json:"payment_method"`
        Currency             *string                                 `json:"currency"`
        Installments         *int                                    `json:"installments"`
        Status               *string                                 `json:"status"`
        CreatedAt            *string                                 `json:"created_at"`
        UpdatedAt            *string                                 `json:"updated_at"`
        Customer             Optional[GetCustomerResponse]           `json:"customer"`
        Card                 *GetCardResponse                        `json:"card"`
        Items                []GetSubscriptionItemResponse           `json:"items"`
        StatementDescriptor  *string                                 `json:"statement_descriptor"`
        Metadata             map[string]string                       `json:"metadata"`
        Setup                *GetSetupResponse                       `json:"setup"`
        GatewayAffiliationId *string                                 `json:"gateway_affiliation_id"`
        NextBillingAt        Optional[string]                        `json:"next_billing_at"`
        BillingDay           Optional[int]                           `json:"billing_day"`
        MinimumPrice         Optional[int]                           `json:"minimum_price"`
        CanceledAt           Optional[string]                        `json:"canceled_at"`
        Discounts            Optional[[]GetDiscountResponse]         `json:"discounts"`
        Increments           []GetIncrementResponse                  `json:"increments"`
        BoletoDueDays        Optional[int]                           `json:"boleto_due_days"`
        Split                *GetSubscriptionSplitResponse           `json:"split"`
        Boleto               Optional[GetSubscriptionBoletoResponse] `json:"boleto"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Code = temp.Code
    if temp.StartAt != nil {
        StartAtVal, err := time.Parse(time.RFC3339, *temp.StartAt)
        if err != nil {
            log.Fatalf("Cannot Parse start_at as % s format.", time.RFC3339)
        }
        g.StartAt = &StartAtVal
    }
    g.Interval = temp.Interval
    g.IntervalCount = temp.IntervalCount
    g.BillingType = temp.BillingType
    g.CurrentCycle = temp.CurrentCycle
    g.PaymentMethod = temp.PaymentMethod
    g.Currency = temp.Currency
    g.Installments = temp.Installments
    g.Status = temp.Status
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
    g.Customer = temp.Customer
    g.Card = temp.Card
    g.Items = temp.Items
    g.StatementDescriptor = temp.StatementDescriptor
    g.Metadata = temp.Metadata
    g.Setup = temp.Setup
    g.GatewayAffiliationId = temp.GatewayAffiliationId
    g.NextBillingAt.ShouldSetValue(temp.NextBillingAt.IsValueSet())
    if temp.NextBillingAt.Value() != nil {
        NextBillingAtVal, err := time.Parse(time.RFC3339, (*temp.NextBillingAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse next_billing_at as % s format.", time.RFC3339)
        }
        g.NextBillingAt.SetValue(&NextBillingAtVal)
    }
    g.BillingDay = temp.BillingDay
    g.MinimumPrice = temp.MinimumPrice
    g.CanceledAt.ShouldSetValue(temp.CanceledAt.IsValueSet())
    if temp.CanceledAt.Value() != nil {
        CanceledAtVal, err := time.Parse(time.RFC3339, (*temp.CanceledAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
        }
        g.CanceledAt.SetValue(&CanceledAtVal)
    }
    g.Discounts = temp.Discounts
    g.Increments = temp.Increments
    g.BoletoDueDays = temp.BoletoDueDays
    g.Split = temp.Split
    g.Boleto = temp.Boleto
    return nil
}
