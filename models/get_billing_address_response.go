package models

import (
    "encoding/json"
)

// GetBillingAddressResponse represents a GetBillingAddressResponse struct.
// Response object for getting a billing address
type GetBillingAddressResponse struct {
    Street       *string `json:"street"`
    Number       *string `json:"number"`
    ZipCode      *string `json:"zip_code"`
    Neighborhood *string `json:"neighborhood"`
    City         *string `json:"city"`
    State        *string `json:"state"`
    Country      *string `json:"country"`
    Complement   *string `json:"complement"`
    // Line 1 for address
    Line1        *string `json:"line_1"`
    // Line 2 for address
    Line2        *string `json:"line_2"`
}

// MarshalJSON implements the json.Marshaler interface for GetBillingAddressResponse.
// It customizes the JSON marshaling process for GetBillingAddressResponse objects.
func (g *GetBillingAddressResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetBillingAddressResponse object to a map representation for JSON marshaling.
func (g *GetBillingAddressResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["street"] = g.Street
    structMap["number"] = g.Number
    structMap["zip_code"] = g.ZipCode
    structMap["neighborhood"] = g.Neighborhood
    structMap["city"] = g.City
    structMap["state"] = g.State
    structMap["country"] = g.Country
    structMap["complement"] = g.Complement
    structMap["line_1"] = g.Line1
    structMap["line_2"] = g.Line2
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetBillingAddressResponse.
// It customizes the JSON unmarshaling process for GetBillingAddressResponse objects.
func (g *GetBillingAddressResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Street       *string `json:"street"`
        Number       *string `json:"number"`
        ZipCode      *string `json:"zip_code"`
        Neighborhood *string `json:"neighborhood"`
        City         *string `json:"city"`
        State        *string `json:"state"`
        Country      *string `json:"country"`
        Complement   *string `json:"complement"`
        Line1        *string `json:"line_1"`
        Line2        *string `json:"line_2"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Street = temp.Street
    g.Number = temp.Number
    g.ZipCode = temp.ZipCode
    g.Neighborhood = temp.Neighborhood
    g.City = temp.City
    g.State = temp.State
    g.Country = temp.Country
    g.Complement = temp.Complement
    g.Line1 = temp.Line1
    g.Line2 = temp.Line2
    return nil
}
