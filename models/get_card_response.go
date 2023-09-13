package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetCardResponse represents a GetCardResponse struct.
// Response object for getting a credit card
type GetCardResponse struct {
    Id             *string                       `json:"id"`
    LastFourDigits *string                       `json:"last_four_digits"`
    Brand          *string                       `json:"brand"`
    HolderName     *string                       `json:"holder_name"`
    ExpMonth       *int                          `json:"exp_month"`
    ExpYear        *int                          `json:"exp_year"`
    Status         *string                       `json:"status"`
    CreatedAt      *time.Time                    `json:"created_at"`
    UpdatedAt      *time.Time                    `json:"updated_at"`
    BillingAddress *GetBillingAddressResponse    `json:"billing_address"`
    Customer       Optional[GetCustomerResponse] `json:"customer"`
    Metadata       map[string]string             `json:"metadata"`
    // Card type
    Type           *string                       `json:"type"`
    // Document number for the card's holder
    HolderDocument *string                       `json:"holder_document"`
    DeletedAt      Optional[time.Time]           `json:"deleted_at"`
    // First six digits
    FirstSixDigits *string                       `json:"first_six_digits"`
    Label          *string                       `json:"label"`
}

// MarshalJSON implements the json.Marshaler interface for GetCardResponse.
// It customizes the JSON marshaling process for GetCardResponse objects.
func (g *GetCardResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetCardResponse object to a map representation for JSON marshaling.
func (g *GetCardResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["last_four_digits"] = g.LastFourDigits
    structMap["brand"] = g.Brand
    structMap["holder_name"] = g.HolderName
    structMap["exp_month"] = g.ExpMonth
    structMap["exp_year"] = g.ExpYear
    structMap["status"] = g.Status
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    structMap["billing_address"] = g.BillingAddress
    if g.Customer.IsValueSet() {
        structMap["customer"] = g.Customer.Value()
    }
    structMap["metadata"] = g.Metadata
    structMap["type"] = g.Type
    structMap["holder_document"] = g.HolderDocument
    if g.DeletedAt.IsValueSet() {
        var DeletedAtVal *string = nil
        if g.DeletedAt.Value() != nil {
            val := g.DeletedAt.Value().Format(time.RFC3339)
            DeletedAtVal = &val
        }
        structMap["deleted_at"] = DeletedAtVal
    }
    structMap["first_six_digits"] = g.FirstSixDigits
    structMap["label"] = g.Label
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCardResponse.
// It customizes the JSON unmarshaling process for GetCardResponse objects.
func (g *GetCardResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id             *string                       `json:"id"`
        LastFourDigits *string                       `json:"last_four_digits"`
        Brand          *string                       `json:"brand"`
        HolderName     *string                       `json:"holder_name"`
        ExpMonth       *int                          `json:"exp_month"`
        ExpYear        *int                          `json:"exp_year"`
        Status         *string                       `json:"status"`
        CreatedAt      *string                       `json:"created_at"`
        UpdatedAt      *string                       `json:"updated_at"`
        BillingAddress *GetBillingAddressResponse    `json:"billing_address"`
        Customer       Optional[GetCustomerResponse] `json:"customer"`
        Metadata       map[string]string             `json:"metadata"`
        Type           *string                       `json:"type"`
        HolderDocument *string                       `json:"holder_document"`
        DeletedAt      Optional[string]              `json:"deleted_at"`
        FirstSixDigits *string                       `json:"first_six_digits"`
        Label          *string                       `json:"label"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.LastFourDigits = temp.LastFourDigits
    g.Brand = temp.Brand
    g.HolderName = temp.HolderName
    g.ExpMonth = temp.ExpMonth
    g.ExpYear = temp.ExpYear
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
    g.BillingAddress = temp.BillingAddress
    g.Customer = temp.Customer
    g.Metadata = temp.Metadata
    g.Type = temp.Type
    g.HolderDocument = temp.HolderDocument
    g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
    if temp.DeletedAt.Value() != nil {
        DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
        }
        g.DeletedAt.SetValue(&DeletedAtVal)
    }
    g.FirstSixDigits = temp.FirstSixDigits
    g.Label = temp.Label
    return nil
}
