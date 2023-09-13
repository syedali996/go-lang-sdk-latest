package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetOrderResponse represents a GetOrderResponse struct.
// Response object for getting an Order
type GetOrderResponse struct {
    Id         *string                                `json:"id"`
    Code       *string                                `json:"code"`
    Currency   *string                                `json:"currency"`
    Items      []GetOrderItemResponse                 `json:"items"`
    Customer   Optional[GetCustomerResponse]          `json:"customer"`
    Status     *string                                `json:"status"`
    CreatedAt  *time.Time                             `json:"created_at"`
    UpdatedAt  *time.Time                             `json:"updated_at"`
    Charges    []GetChargeResponse                    `json:"charges"`
    InvoiceUrl *string                                `json:"invoice_url"`
    Shipping   *GetShippingResponse                   `json:"shipping"`
    Metadata   map[string]string                      `json:"metadata"`
    // Checkout Payment Settings Response
    Checkouts  Optional[[]GetCheckoutPaymentResponse] `json:"checkouts"`
    // Ip address
    Ip         Optional[string]                       `json:"ip"`
    // Session id
    SessionId  Optional[string]                       `json:"session_id"`
    // Location
    Location   Optional[GetLocationResponse]          `json:"location"`
    // Device's informations
    Device     Optional[GetDeviceResponse]            `json:"device"`
    // Indicates whether the order is closed
    Closed     *bool                                  `json:"closed"`
}

// MarshalJSON implements the json.Marshaler interface for GetOrderResponse.
// It customizes the JSON marshaling process for GetOrderResponse objects.
func (g *GetOrderResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetOrderResponse object to a map representation for JSON marshaling.
func (g *GetOrderResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["code"] = g.Code
    structMap["currency"] = g.Currency
    structMap["items"] = g.Items
    if g.Customer.IsValueSet() {
        structMap["customer"] = g.Customer.Value()
    }
    structMap["status"] = g.Status
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    structMap["charges"] = g.Charges
    structMap["invoice_url"] = g.InvoiceUrl
    structMap["shipping"] = g.Shipping
    structMap["metadata"] = g.Metadata
    if g.Checkouts.IsValueSet() {
        structMap["checkouts"] = g.Checkouts.Value()
    }
    if g.Ip.IsValueSet() {
        structMap["ip"] = g.Ip.Value()
    }
    if g.SessionId.IsValueSet() {
        structMap["session_id"] = g.SessionId.Value()
    }
    if g.Location.IsValueSet() {
        structMap["location"] = g.Location.Value()
    }
    if g.Device.IsValueSet() {
        structMap["device"] = g.Device.Value()
    }
    structMap["closed"] = g.Closed
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetOrderResponse.
// It customizes the JSON unmarshaling process for GetOrderResponse objects.
func (g *GetOrderResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id         *string                                `json:"id"`
        Code       *string                                `json:"code"`
        Currency   *string                                `json:"currency"`
        Items      []GetOrderItemResponse                 `json:"items"`
        Customer   Optional[GetCustomerResponse]          `json:"customer"`
        Status     *string                                `json:"status"`
        CreatedAt  *string                                `json:"created_at"`
        UpdatedAt  *string                                `json:"updated_at"`
        Charges    []GetChargeResponse                    `json:"charges"`
        InvoiceUrl *string                                `json:"invoice_url"`
        Shipping   *GetShippingResponse                   `json:"shipping"`
        Metadata   map[string]string                      `json:"metadata"`
        Checkouts  Optional[[]GetCheckoutPaymentResponse] `json:"checkouts"`
        Ip         Optional[string]                       `json:"ip"`
        SessionId  Optional[string]                       `json:"session_id"`
        Location   Optional[GetLocationResponse]          `json:"location"`
        Device     Optional[GetDeviceResponse]            `json:"device"`
        Closed     *bool                                  `json:"closed"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Code = temp.Code
    g.Currency = temp.Currency
    g.Items = temp.Items
    g.Customer = temp.Customer
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
    g.Charges = temp.Charges
    g.InvoiceUrl = temp.InvoiceUrl
    g.Shipping = temp.Shipping
    g.Metadata = temp.Metadata
    g.Checkouts = temp.Checkouts
    g.Ip = temp.Ip
    g.SessionId = temp.SessionId
    g.Location = temp.Location
    g.Device = temp.Device
    g.Closed = temp.Closed
    return nil
}
