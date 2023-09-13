package models

import (
    "encoding/json"
)

// GetInvoiceItemResponse represents a GetInvoiceItemResponse struct.
// Response object for getting an invoice item
type GetInvoiceItemResponse struct {
    Amount             *int                      `json:"amount"`
    Description        *string                   `json:"description"`
    PricingScheme      *GetPricingSchemeResponse `json:"pricing_scheme"`
    PriceBracket       *GetPriceBracketResponse  `json:"price_bracket"`
    Quantity           Optional[int]             `json:"quantity"`
    Name               Optional[string]          `json:"name"`
    // Subscription Item Id
    SubscriptionItemId *string                   `json:"subscription_item_id"`
}

// MarshalJSON implements the json.Marshaler interface for GetInvoiceItemResponse.
// It customizes the JSON marshaling process for GetInvoiceItemResponse objects.
func (g *GetInvoiceItemResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetInvoiceItemResponse object to a map representation for JSON marshaling.
func (g *GetInvoiceItemResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["amount"] = g.Amount
    structMap["description"] = g.Description
    structMap["pricing_scheme"] = g.PricingScheme
    structMap["price_bracket"] = g.PriceBracket
    if g.Quantity.IsValueSet() {
        structMap["quantity"] = g.Quantity.Value()
    }
    if g.Name.IsValueSet() {
        structMap["name"] = g.Name.Value()
    }
    structMap["subscription_item_id"] = g.SubscriptionItemId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetInvoiceItemResponse.
// It customizes the JSON unmarshaling process for GetInvoiceItemResponse objects.
func (g *GetInvoiceItemResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Amount             *int                      `json:"amount"`
        Description        *string                   `json:"description"`
        PricingScheme      *GetPricingSchemeResponse `json:"pricing_scheme"`
        PriceBracket       *GetPriceBracketResponse  `json:"price_bracket"`
        Quantity           Optional[int]             `json:"quantity"`
        Name               Optional[string]          `json:"name"`
        SubscriptionItemId *string                   `json:"subscription_item_id"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Amount = temp.Amount
    g.Description = temp.Description
    g.PricingScheme = temp.PricingScheme
    g.PriceBracket = temp.PriceBracket
    g.Quantity = temp.Quantity
    g.Name = temp.Name
    g.SubscriptionItemId = temp.SubscriptionItemId
    return nil
}
