package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetCustomerResponse represents a GetCustomerResponse struct.
// Response object for getting a customer
type GetCustomerResponse struct {
    Id            *string             `json:"id"`
    Name          *string             `json:"name"`
    Email         *string             `json:"email"`
    Delinquent    *bool               `json:"delinquent"`
    CreatedAt     *time.Time          `json:"created_at"`
    UpdatedAt     *time.Time          `json:"updated_at"`
    Document      *string             `json:"document"`
    Type          *string             `json:"type"`
    FbAccessToken *string             `json:"fb_access_token"`
    Address       *GetAddressResponse `json:"address"`
    Metadata      map[string]string   `json:"metadata"`
    Phones        *GetPhonesResponse  `json:"phones"`
    FbId          Optional[int64]     `json:"fb_id"`
    // Código de referência do cliente no sistema da loja. Max: 52 caracteres
    Code          *string             `json:"code"`
    DocumentType  *string             `json:"document_type"`
}

// MarshalJSON implements the json.Marshaler interface for GetCustomerResponse.
// It customizes the JSON marshaling process for GetCustomerResponse objects.
func (g *GetCustomerResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetCustomerResponse object to a map representation for JSON marshaling.
func (g *GetCustomerResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["name"] = g.Name
    structMap["email"] = g.Email
    structMap["delinquent"] = g.Delinquent
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    structMap["document"] = g.Document
    structMap["type"] = g.Type
    structMap["fb_access_token"] = g.FbAccessToken
    structMap["address"] = g.Address
    structMap["metadata"] = g.Metadata
    structMap["phones"] = g.Phones
    if g.FbId.IsValueSet() {
        structMap["fb_id"] = g.FbId.Value()
    }
    structMap["code"] = g.Code
    structMap["document_type"] = g.DocumentType
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCustomerResponse.
// It customizes the JSON unmarshaling process for GetCustomerResponse objects.
func (g *GetCustomerResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id            *string             `json:"id"`
        Name          *string             `json:"name"`
        Email         *string             `json:"email"`
        Delinquent    *bool               `json:"delinquent"`
        CreatedAt     *string             `json:"created_at"`
        UpdatedAt     *string             `json:"updated_at"`
        Document      *string             `json:"document"`
        Type          *string             `json:"type"`
        FbAccessToken *string             `json:"fb_access_token"`
        Address       *GetAddressResponse `json:"address"`
        Metadata      map[string]string   `json:"metadata"`
        Phones        *GetPhonesResponse  `json:"phones"`
        FbId          Optional[int64]     `json:"fb_id"`
        Code          *string             `json:"code"`
        DocumentType  *string             `json:"document_type"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Name = temp.Name
    g.Email = temp.Email
    g.Delinquent = temp.Delinquent
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
    g.Document = temp.Document
    g.Type = temp.Type
    g.FbAccessToken = temp.FbAccessToken
    g.Address = temp.Address
    g.Metadata = temp.Metadata
    g.Phones = temp.Phones
    g.FbId = temp.FbId
    g.Code = temp.Code
    g.DocumentType = temp.DocumentType
    return nil
}
