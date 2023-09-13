package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetAddressResponse represents a GetAddressResponse struct.
// Response object for getting an Address
type GetAddressResponse struct {
    Id           *string                       `json:"id"`
    Street       *string                       `json:"street"`
    Number       *string                       `json:"number"`
    Complement   *string                       `json:"complement"`
    ZipCode      *string                       `json:"zip_code"`
    Neighborhood *string                       `json:"neighborhood"`
    City         *string                       `json:"city"`
    State        *string                       `json:"state"`
    Country      *string                       `json:"country"`
    Status       *string                       `json:"status"`
    CreatedAt    *time.Time                    `json:"created_at"`
    UpdatedAt    *time.Time                    `json:"updated_at"`
    Customer     Optional[GetCustomerResponse] `json:"customer"`
    Metadata     map[string]string             `json:"metadata"`
    // Line 1 for address
    Line1        *string                       `json:"line_1"`
    // Line 2 for address
    Line2        *string                       `json:"line_2"`
    DeletedAt    Optional[time.Time]           `json:"deleted_at"`
}

// MarshalJSON implements the json.Marshaler interface for GetAddressResponse.
// It customizes the JSON marshaling process for GetAddressResponse objects.
func (g *GetAddressResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetAddressResponse object to a map representation for JSON marshaling.
func (g *GetAddressResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["street"] = g.Street
    structMap["number"] = g.Number
    structMap["complement"] = g.Complement
    structMap["zip_code"] = g.ZipCode
    structMap["neighborhood"] = g.Neighborhood
    structMap["city"] = g.City
    structMap["state"] = g.State
    structMap["country"] = g.Country
    structMap["status"] = g.Status
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    if g.Customer.IsValueSet() {
        structMap["customer"] = g.Customer.Value()
    }
    structMap["metadata"] = g.Metadata
    structMap["line_1"] = g.Line1
    structMap["line_2"] = g.Line2
    if g.DeletedAt.IsValueSet() {
        var DeletedAtVal *string = nil
        if g.DeletedAt.Value() != nil {
            val := g.DeletedAt.Value().Format(time.RFC3339)
            DeletedAtVal = &val
        }
        structMap["deleted_at"] = DeletedAtVal
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetAddressResponse.
// It customizes the JSON unmarshaling process for GetAddressResponse objects.
func (g *GetAddressResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id           *string                       `json:"id"`
        Street       *string                       `json:"street"`
        Number       *string                       `json:"number"`
        Complement   *string                       `json:"complement"`
        ZipCode      *string                       `json:"zip_code"`
        Neighborhood *string                       `json:"neighborhood"`
        City         *string                       `json:"city"`
        State        *string                       `json:"state"`
        Country      *string                       `json:"country"`
        Status       *string                       `json:"status"`
        CreatedAt    *string                       `json:"created_at"`
        UpdatedAt    *string                       `json:"updated_at"`
        Customer     Optional[GetCustomerResponse] `json:"customer"`
        Metadata     map[string]string             `json:"metadata"`
        Line1        *string                       `json:"line_1"`
        Line2        *string                       `json:"line_2"`
        DeletedAt    Optional[string]              `json:"deleted_at"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Street = temp.Street
    g.Number = temp.Number
    g.Complement = temp.Complement
    g.ZipCode = temp.ZipCode
    g.Neighborhood = temp.Neighborhood
    g.City = temp.City
    g.State = temp.State
    g.Country = temp.Country
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
    g.Metadata = temp.Metadata
    g.Line1 = temp.Line1
    g.Line2 = temp.Line2
    g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
    if temp.DeletedAt.Value() != nil {
        DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
        }
        g.DeletedAt.SetValue(&DeletedAtVal)
    }
    return nil
}
