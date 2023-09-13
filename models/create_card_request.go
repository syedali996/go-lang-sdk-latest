package models

import (
    "encoding/json"
)

// CreateCardRequest represents a CreateCardRequest struct.
// Card data
type CreateCardRequest struct {
    // Credit card number
    Number           string                   `json:"number"`
    // Holder name, as written on the card
    HolderName       string                   `json:"holder_name"`
    // The expiration month
    ExpMonth         int                      `json:"exp_month"`
    // The expiration year, that can be informed with 2 or 4 digits
    ExpYear          int                      `json:"exp_year"`
    // The card's security code
    Cvv              string                   `json:"cvv"`
    // Card's billing address
    BillingAddress   CreateAddressRequest     `json:"billing_address"`
    // Card brand
    Brand            string                   `json:"brand"`
    // The address id for the billing address
    BillingAddressId string                   `json:"billing_address_id"`
    // Metadata
    Metadata         map[string]string        `json:"metadata"`
    // Card type
    Type             string                   `json:"type"`
    // Options for creating the card
    Options          CreateCardOptionsRequest `json:"options"`
    // Document number for the card's holder
    HolderDocument   *string                  `json:"holder_document,omitempty"`
    // Indicates whether it is a private label card
    PrivateLabel     bool                     `json:"private_label"`
    Label            string                   `json:"label"`
    // Identifier
    Id               *string                  `json:"id,omitempty"`
    // token identifier
    Token            *string                  `json:"token,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCardRequest.
// It customizes the JSON marshaling process for CreateCardRequest objects.
func (c *CreateCardRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateCardRequest object to a map representation for JSON marshaling.
func (c *CreateCardRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["number"] = c.Number
    structMap["holder_name"] = c.HolderName
    structMap["exp_month"] = c.ExpMonth
    structMap["exp_year"] = c.ExpYear
    structMap["cvv"] = c.Cvv
    structMap["billing_address"] = c.BillingAddress
    structMap["brand"] = c.Brand
    structMap["billing_address_id"] = c.BillingAddressId
    structMap["metadata"] = c.Metadata
    structMap["type"] = c.Type
    structMap["options"] = c.Options
    if c.HolderDocument != nil {
        structMap["holder_document"] = c.HolderDocument
    }
    structMap["private_label"] = c.PrivateLabel
    structMap["label"] = c.Label
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.Token != nil {
        structMap["token"] = c.Token
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCardRequest.
// It customizes the JSON unmarshaling process for CreateCardRequest objects.
func (c *CreateCardRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Number           string                   `json:"number"`
        HolderName       string                   `json:"holder_name"`
        ExpMonth         int                      `json:"exp_month"`
        ExpYear          int                      `json:"exp_year"`
        Cvv              string                   `json:"cvv"`
        BillingAddress   CreateAddressRequest     `json:"billing_address"`
        Brand            string                   `json:"brand"`
        BillingAddressId string                   `json:"billing_address_id"`
        Metadata         map[string]string        `json:"metadata"`
        Type             string                   `json:"type"`
        Options          CreateCardOptionsRequest `json:"options"`
        HolderDocument   *string                  `json:"holder_document,omitempty"`
        PrivateLabel     bool                     `json:"private_label"`
        Label            string                   `json:"label"`
        Id               *string                  `json:"id,omitempty"`
        Token            *string                  `json:"token,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Number = temp.Number
    c.HolderName = temp.HolderName
    c.ExpMonth = temp.ExpMonth
    c.ExpYear = temp.ExpYear
    c.Cvv = temp.Cvv
    c.BillingAddress = temp.BillingAddress
    c.Brand = temp.Brand
    c.BillingAddressId = temp.BillingAddressId
    c.Metadata = temp.Metadata
    c.Type = temp.Type
    c.Options = temp.Options
    c.HolderDocument = temp.HolderDocument
    c.PrivateLabel = temp.PrivateLabel
    c.Label = temp.Label
    c.Id = temp.Id
    c.Token = temp.Token
    return nil
}
