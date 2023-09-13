package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetAccessTokenResponse represents a GetAccessTokenResponse struct.
// Response object for getting a access token
type GetAccessTokenResponse struct {
    Id        *string                       `json:"id"`
    Code      *string                       `json:"code"`
    Status    *string                       `json:"status"`
    CreatedAt *time.Time                    `json:"created_at"`
    Customer  Optional[GetCustomerResponse] `json:"customer"`
}

// MarshalJSON implements the json.Marshaler interface for GetAccessTokenResponse.
// It customizes the JSON marshaling process for GetAccessTokenResponse objects.
func (g *GetAccessTokenResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetAccessTokenResponse object to a map representation for JSON marshaling.
func (g *GetAccessTokenResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["code"] = g.Code
    structMap["status"] = g.Status
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    if g.Customer.IsValueSet() {
        structMap["customer"] = g.Customer.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetAccessTokenResponse.
// It customizes the JSON unmarshaling process for GetAccessTokenResponse objects.
func (g *GetAccessTokenResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id        *string                       `json:"id"`
        Code      *string                       `json:"code"`
        Status    *string                       `json:"status"`
        CreatedAt *string                       `json:"created_at"`
        Customer  Optional[GetCustomerResponse] `json:"customer"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Code = temp.Code
    g.Status = temp.Status
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        g.CreatedAt = &CreatedAtVal
    }
    g.Customer = temp.Customer
    return nil
}
