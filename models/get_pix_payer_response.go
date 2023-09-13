package models

import (
    "encoding/json"
)

// GetPixPayerResponse represents a GetPixPayerResponse struct.
// Pix payer data.
type GetPixPayerResponse struct {
    Name         *string                    `json:"name"`
    Document     *string                    `json:"document"`
    DocumentType *string                    `json:"document_type"`
    BankAccount  *GetPixBankAccountResponse `json:"bank_account"`
}

// MarshalJSON implements the json.Marshaler interface for GetPixPayerResponse.
// It customizes the JSON marshaling process for GetPixPayerResponse objects.
func (g *GetPixPayerResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetPixPayerResponse object to a map representation for JSON marshaling.
func (g *GetPixPayerResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["name"] = g.Name
    structMap["document"] = g.Document
    structMap["document_type"] = g.DocumentType
    structMap["bank_account"] = g.BankAccount
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPixPayerResponse.
// It customizes the JSON unmarshaling process for GetPixPayerResponse objects.
func (g *GetPixPayerResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Name         *string                    `json:"name"`
        Document     *string                    `json:"document"`
        DocumentType *string                    `json:"document_type"`
        BankAccount  *GetPixBankAccountResponse `json:"bank_account"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Name = temp.Name
    g.Document = temp.Document
    g.DocumentType = temp.DocumentType
    g.BankAccount = temp.BankAccount
    return nil
}
