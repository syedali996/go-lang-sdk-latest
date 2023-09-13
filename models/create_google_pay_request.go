package models

import (
    "encoding/json"
)

// CreateGooglePayRequest represents a CreateGooglePayRequest struct.
// The GooglePay Token Payment Request
type CreateGooglePayRequest struct {
    // The token version
    Version            string                       `json:"version"`
    // The cryptography data
    Data               string                       `json:"data"`
    // The GooglePay header request
    Header             CreateGooglePayHeaderRequest `json:"header"`
    // Detached PKCS #7 signature, Base64 encoded as string
    Signature          string                       `json:"signature"`
    // GooglePay Merchant identifier
    MerchantIdentifier string                       `json:"merchant_identifier"`
}

// MarshalJSON implements the json.Marshaler interface for CreateGooglePayRequest.
// It customizes the JSON marshaling process for CreateGooglePayRequest objects.
func (c *CreateGooglePayRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateGooglePayRequest object to a map representation for JSON marshaling.
func (c *CreateGooglePayRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["version"] = c.Version
    structMap["data"] = c.Data
    structMap["header"] = c.Header
    structMap["signature"] = c.Signature
    structMap["merchant_identifier"] = c.MerchantIdentifier
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateGooglePayRequest.
// It customizes the JSON unmarshaling process for CreateGooglePayRequest objects.
func (c *CreateGooglePayRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Version            string                       `json:"version"`
        Data               string                       `json:"data"`
        Header             CreateGooglePayHeaderRequest `json:"header"`
        Signature          string                       `json:"signature"`
        MerchantIdentifier string                       `json:"merchant_identifier"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Version = temp.Version
    c.Data = temp.Data
    c.Header = temp.Header
    c.Signature = temp.Signature
    c.MerchantIdentifier = temp.MerchantIdentifier
    return nil
}
