package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetShippingResponse represents a GetShippingResponse struct.
// Response object for getting the shipping data
type GetShippingResponse struct {
    Amount                *int                `json:"amount"`
    Description           *string             `json:"description"`
    RecipientName         *string             `json:"recipient_name"`
    RecipientPhone        *string             `json:"recipient_phone"`
    Address               *GetAddressResponse `json:"address"`
    // Data máxima de entrega
    MaxDeliveryDate       Optional[time.Time] `json:"max_delivery_date"`
    // Prazo estimado de entrega
    EstimatedDeliveryDate Optional[time.Time] `json:"estimated_delivery_date"`
    // Shipping Type
    Type                  *string             `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for GetShippingResponse.
// It customizes the JSON marshaling process for GetShippingResponse objects.
func (g *GetShippingResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetShippingResponse object to a map representation for JSON marshaling.
func (g *GetShippingResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["amount"] = g.Amount
    structMap["description"] = g.Description
    structMap["recipient_name"] = g.RecipientName
    structMap["recipient_phone"] = g.RecipientPhone
    structMap["address"] = g.Address
    if g.MaxDeliveryDate.IsValueSet() {
        var MaxDeliveryDateVal *string = nil
        if g.MaxDeliveryDate.Value() != nil {
            val := g.MaxDeliveryDate.Value().Format(time.RFC3339)
            MaxDeliveryDateVal = &val
        }
        structMap["max_delivery_date"] = MaxDeliveryDateVal
    }
    if g.EstimatedDeliveryDate.IsValueSet() {
        var EstimatedDeliveryDateVal *string = nil
        if g.EstimatedDeliveryDate.Value() != nil {
            val := g.EstimatedDeliveryDate.Value().Format(time.RFC3339)
            EstimatedDeliveryDateVal = &val
        }
        structMap["estimated_delivery_date"] = EstimatedDeliveryDateVal
    }
    structMap["type"] = g.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetShippingResponse.
// It customizes the JSON unmarshaling process for GetShippingResponse objects.
func (g *GetShippingResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Amount                *int                `json:"amount"`
        Description           *string             `json:"description"`
        RecipientName         *string             `json:"recipient_name"`
        RecipientPhone        *string             `json:"recipient_phone"`
        Address               *GetAddressResponse `json:"address"`
        MaxDeliveryDate       Optional[string]    `json:"max_delivery_date"`
        EstimatedDeliveryDate Optional[string]    `json:"estimated_delivery_date"`
        Type                  *string             `json:"type"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Amount = temp.Amount
    g.Description = temp.Description
    g.RecipientName = temp.RecipientName
    g.RecipientPhone = temp.RecipientPhone
    g.Address = temp.Address
    g.MaxDeliveryDate.ShouldSetValue(temp.MaxDeliveryDate.IsValueSet())
    if temp.MaxDeliveryDate.Value() != nil {
        MaxDeliveryDateVal, err := time.Parse(time.RFC3339, (*temp.MaxDeliveryDate.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse max_delivery_date as % s format.", time.RFC3339)
        }
        g.MaxDeliveryDate.SetValue(&MaxDeliveryDateVal)
    }
    g.EstimatedDeliveryDate.ShouldSetValue(temp.EstimatedDeliveryDate.IsValueSet())
    if temp.EstimatedDeliveryDate.Value() != nil {
        EstimatedDeliveryDateVal, err := time.Parse(time.RFC3339, (*temp.EstimatedDeliveryDate.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse estimated_delivery_date as % s format.", time.RFC3339)
        }
        g.EstimatedDeliveryDate.SetValue(&EstimatedDeliveryDateVal)
    }
    g.Type = temp.Type
    return nil
}
