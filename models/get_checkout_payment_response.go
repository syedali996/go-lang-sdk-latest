package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetCheckoutPaymentResponse represents a GetCheckoutPaymentResponse struct.
// Resposta das configurações de pagamento do checkout
type GetCheckoutPaymentResponse struct {
    Id                      *string                                          `json:"id"`
    // Valor em centavos
    Amount                  Optional[int]                                    `json:"amount"`
    // Meio de pagamento padrão no checkout
    DefaultPaymentMethod    *string                                          `json:"default_payment_method"`
    // Url de redirecionamento de sucesso após o checkou
    SuccessUrl              *string                                          `json:"success_url"`
    // Url para pagamento usando o checkout
    PaymentUrl              *string                                          `json:"payment_url"`
    // Código da afiliação onde o pagamento será processado no gateway
    GatewayAffiliationId    *string                                          `json:"gateway_affiliation_id"`
    // Meios de pagamento aceitos no checkout
    AcceptedPaymentMethods  []string                                         `json:"accepted_payment_methods"`
    // Status do checkout
    Status                  *string                                          `json:"status"`
    // Pular tela de sucesso pós-pagamento?
    SkipCheckoutSuccessPage *bool                                            `json:"skip_checkout_success_page"`
    // Data de criação
    CreatedAt               *time.Time                                       `json:"created_at"`
    // Data de atualização
    UpdatedAt               *time.Time                                       `json:"updated_at"`
    // Data de cancelamento
    CanceledAt              Optional[time.Time]                              `json:"canceled_at"`
    // Torna o objeto customer editável
    CustomerEditable        *bool                                            `json:"customer_editable"`
    // Dados do comprador
    Customer                Optional[GetCustomerResponse]                    `json:"customer"`
    // Dados do endereço de cobrança
    Billingaddress          *GetAddressResponse                              `json:"billingaddress"`
    // Configurações de cartão de crédito
    CreditCard              *GetCheckoutCreditCardPaymentResponse            `json:"credit_card"`
    // Configurações de boleto
    Boleto                  *GetCheckoutBoletoPaymentResponse                `json:"boleto"`
    // Indica se o billing address poderá ser editado
    BillingAddressEditable  *bool                                            `json:"billing_address_editable"`
    // Configurações  de entrega
    Shipping                *GetShippingResponse                             `json:"shipping"`
    // Indica se possui entrega
    Shippable               *bool                                            `json:"shippable"`
    // Data de fechamento
    ClosedAt                Optional[time.Time]                              `json:"closed_at"`
    // Data de expiração
    ExpiresAt               Optional[time.Time]                              `json:"expires_at"`
    // Moeda
    Currency                *string                                          `json:"currency"`
    // Configurações de cartão de débito
    DebitCard               Optional[GetCheckoutDebitCardPaymentResponse]    `json:"debit_card"`
    // Bank transfer payment response
    BankTransfer            Optional[GetCheckoutBankTransferPaymentResponse] `json:"bank_transfer"`
    // Accepted Brands
    AcceptedBrands          []string                                         `json:"accepted_brands"`
    // Pix payment response
    Pix                     Optional[GetCheckoutPixPaymentResponse]          `json:"pix"`
}

// MarshalJSON implements the json.Marshaler interface for GetCheckoutPaymentResponse.
// It customizes the JSON marshaling process for GetCheckoutPaymentResponse objects.
func (g *GetCheckoutPaymentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetCheckoutPaymentResponse object to a map representation for JSON marshaling.
func (g *GetCheckoutPaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    if g.Amount.IsValueSet() {
        structMap["amount"] = g.Amount.Value()
    }
    structMap["default_payment_method"] = g.DefaultPaymentMethod
    structMap["success_url"] = g.SuccessUrl
    structMap["payment_url"] = g.PaymentUrl
    structMap["gateway_affiliation_id"] = g.GatewayAffiliationId
    structMap["accepted_payment_methods"] = g.AcceptedPaymentMethods
    structMap["status"] = g.Status
    structMap["skip_checkout_success_page"] = g.SkipCheckoutSuccessPage
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    if g.CanceledAt.IsValueSet() {
        var CanceledAtVal *string = nil
        if g.CanceledAt.Value() != nil {
            val := g.CanceledAt.Value().Format(time.RFC3339)
            CanceledAtVal = &val
        }
        structMap["canceled_at"] = CanceledAtVal
    }
    structMap["customer_editable"] = g.CustomerEditable
    if g.Customer.IsValueSet() {
        structMap["customer"] = g.Customer.Value()
    }
    structMap["billingaddress"] = g.Billingaddress
    structMap["credit_card"] = g.CreditCard
    structMap["boleto"] = g.Boleto
    structMap["billing_address_editable"] = g.BillingAddressEditable
    structMap["shipping"] = g.Shipping
    structMap["shippable"] = g.Shippable
    if g.ClosedAt.IsValueSet() {
        var ClosedAtVal *string = nil
        if g.ClosedAt.Value() != nil {
            val := g.ClosedAt.Value().Format(time.RFC3339)
            ClosedAtVal = &val
        }
        structMap["closed_at"] = ClosedAtVal
    }
    if g.ExpiresAt.IsValueSet() {
        var ExpiresAtVal *string = nil
        if g.ExpiresAt.Value() != nil {
            val := g.ExpiresAt.Value().Format(time.RFC3339)
            ExpiresAtVal = &val
        }
        structMap["expires_at"] = ExpiresAtVal
    }
    structMap["currency"] = g.Currency
    if g.DebitCard.IsValueSet() {
        structMap["debit_card"] = g.DebitCard.Value()
    }
    if g.BankTransfer.IsValueSet() {
        structMap["bank_transfer"] = g.BankTransfer.Value()
    }
    structMap["accepted_brands"] = g.AcceptedBrands
    if g.Pix.IsValueSet() {
        structMap["pix"] = g.Pix.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCheckoutPaymentResponse.
// It customizes the JSON unmarshaling process for GetCheckoutPaymentResponse objects.
func (g *GetCheckoutPaymentResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                      *string                                          `json:"id"`
        Amount                  Optional[int]                                    `json:"amount"`
        DefaultPaymentMethod    *string                                          `json:"default_payment_method"`
        SuccessUrl              *string                                          `json:"success_url"`
        PaymentUrl              *string                                          `json:"payment_url"`
        GatewayAffiliationId    *string                                          `json:"gateway_affiliation_id"`
        AcceptedPaymentMethods  []string                                         `json:"accepted_payment_methods"`
        Status                  *string                                          `json:"status"`
        SkipCheckoutSuccessPage *bool                                            `json:"skip_checkout_success_page"`
        CreatedAt               *string                                          `json:"created_at"`
        UpdatedAt               *string                                          `json:"updated_at"`
        CanceledAt              Optional[string]                                 `json:"canceled_at"`
        CustomerEditable        *bool                                            `json:"customer_editable"`
        Customer                Optional[GetCustomerResponse]                    `json:"customer"`
        Billingaddress          *GetAddressResponse                              `json:"billingaddress"`
        CreditCard              *GetCheckoutCreditCardPaymentResponse            `json:"credit_card"`
        Boleto                  *GetCheckoutBoletoPaymentResponse                `json:"boleto"`
        BillingAddressEditable  *bool                                            `json:"billing_address_editable"`
        Shipping                *GetShippingResponse                             `json:"shipping"`
        Shippable               *bool                                            `json:"shippable"`
        ClosedAt                Optional[string]                                 `json:"closed_at"`
        ExpiresAt               Optional[string]                                 `json:"expires_at"`
        Currency                *string                                          `json:"currency"`
        DebitCard               Optional[GetCheckoutDebitCardPaymentResponse]    `json:"debit_card"`
        BankTransfer            Optional[GetCheckoutBankTransferPaymentResponse] `json:"bank_transfer"`
        AcceptedBrands          []string                                         `json:"accepted_brands"`
        Pix                     Optional[GetCheckoutPixPaymentResponse]          `json:"pix"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Amount = temp.Amount
    g.DefaultPaymentMethod = temp.DefaultPaymentMethod
    g.SuccessUrl = temp.SuccessUrl
    g.PaymentUrl = temp.PaymentUrl
    g.GatewayAffiliationId = temp.GatewayAffiliationId
    g.AcceptedPaymentMethods = temp.AcceptedPaymentMethods
    g.Status = temp.Status
    g.SkipCheckoutSuccessPage = temp.SkipCheckoutSuccessPage
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
    g.CanceledAt.ShouldSetValue(temp.CanceledAt.IsValueSet())
    if temp.CanceledAt.Value() != nil {
        CanceledAtVal, err := time.Parse(time.RFC3339, (*temp.CanceledAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
        }
        g.CanceledAt.SetValue(&CanceledAtVal)
    }
    g.CustomerEditable = temp.CustomerEditable
    g.Customer = temp.Customer
    g.Billingaddress = temp.Billingaddress
    g.CreditCard = temp.CreditCard
    g.Boleto = temp.Boleto
    g.BillingAddressEditable = temp.BillingAddressEditable
    g.Shipping = temp.Shipping
    g.Shippable = temp.Shippable
    g.ClosedAt.ShouldSetValue(temp.ClosedAt.IsValueSet())
    if temp.ClosedAt.Value() != nil {
        ClosedAtVal, err := time.Parse(time.RFC3339, (*temp.ClosedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse closed_at as % s format.", time.RFC3339)
        }
        g.ClosedAt.SetValue(&ClosedAtVal)
    }
    g.ExpiresAt.ShouldSetValue(temp.ExpiresAt.IsValueSet())
    if temp.ExpiresAt.Value() != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, (*temp.ExpiresAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        g.ExpiresAt.SetValue(&ExpiresAtVal)
    }
    g.Currency = temp.Currency
    g.DebitCard = temp.DebitCard
    g.BankTransfer = temp.BankTransfer
    g.AcceptedBrands = temp.AcceptedBrands
    g.Pix = temp.Pix
    return nil
}
