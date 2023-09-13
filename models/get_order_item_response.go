package models

import (
    "encoding/json"
)

// GetOrderItemResponse represents a GetOrderItemResponse struct.
// Response object for getting an order item
type GetOrderItemResponse struct {
    // Id
    Id          *string `json:"id"`
    Amount      *int    `json:"amount"`
    Description *string `json:"description"`
    Quantity    *int    `json:"quantity"`
    // Category
    Category    *string `json:"category"`
    // Code
    Code        *string `json:"code"`
}

// MarshalJSON implements the json.Marshaler interface for GetOrderItemResponse.
// It customizes the JSON marshaling process for GetOrderItemResponse objects.
func (g *GetOrderItemResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetOrderItemResponse object to a map representation for JSON marshaling.
func (g *GetOrderItemResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["amount"] = g.Amount
    structMap["description"] = g.Description
    structMap["quantity"] = g.Quantity
    structMap["category"] = g.Category
    structMap["code"] = g.Code
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetOrderItemResponse.
// It customizes the JSON unmarshaling process for GetOrderItemResponse objects.
func (g *GetOrderItemResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id          *string `json:"id"`
        Amount      *int    `json:"amount"`
        Description *string `json:"description"`
        Quantity    *int    `json:"quantity"`
        Category    *string `json:"category"`
        Code        *string `json:"code"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Amount = temp.Amount
    g.Description = temp.Description
    g.Quantity = temp.Quantity
    g.Category = temp.Category
    g.Code = temp.Code
    return nil
}
