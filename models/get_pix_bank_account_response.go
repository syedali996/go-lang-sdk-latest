package models

import (
    "encoding/json"
)

// GetPixBankAccountResponse represents a GetPixBankAccountResponse struct.
// Payer's bank details.
type GetPixBankAccountResponse struct {
    BankName      *string `json:"bank_name"`
    Ispb          *string `json:"ispb"`
    BranchCode    *string `json:"branch_code"`
    AccountNumber *string `json:"account_number"`
}

// MarshalJSON implements the json.Marshaler interface for GetPixBankAccountResponse.
// It customizes the JSON marshaling process for GetPixBankAccountResponse objects.
func (g *GetPixBankAccountResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetPixBankAccountResponse object to a map representation for JSON marshaling.
func (g *GetPixBankAccountResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["bank_name"] = g.BankName
    structMap["ispb"] = g.Ispb
    structMap["branch_code"] = g.BranchCode
    structMap["account_number"] = g.AccountNumber
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPixBankAccountResponse.
// It customizes the JSON unmarshaling process for GetPixBankAccountResponse objects.
func (g *GetPixBankAccountResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        BankName      *string `json:"bank_name"`
        Ispb          *string `json:"ispb"`
        BranchCode    *string `json:"branch_code"`
        AccountNumber *string `json:"account_number"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.BankName = temp.BankName
    g.Ispb = temp.Ispb
    g.BranchCode = temp.BranchCode
    g.AccountNumber = temp.AccountNumber
    return nil
}
