package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetTransferResponse represents a GetTransferResponse struct.
// Transfer response
type GetTransferResponse struct {
    // Id
    Id          *string                 `json:"id"`
    // Transfer amount
    Amount      *int                    `json:"amount"`
    // Transfer status
    Status      *string                 `json:"status"`
    // Transfer creation date
    CreatedAt   *time.Time              `json:"created_at"`
    // Transfer last update date
    UpdatedAt   *time.Time              `json:"updated_at"`
    // Bank account
    BankAccount *GetBankAccountResponse `json:"bank_account"`
    // Metadata
    Metadata    map[string]string       `json:"metadata"`
}

// MarshalJSON implements the json.Marshaler interface for GetTransferResponse.
// It customizes the JSON marshaling process for GetTransferResponse objects.
func (g *GetTransferResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetTransferResponse object to a map representation for JSON marshaling.
func (g *GetTransferResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["amount"] = g.Amount
    structMap["status"] = g.Status
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    structMap["bank_account"] = g.BankAccount
    structMap["metadata"] = g.Metadata
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetTransferResponse.
// It customizes the JSON unmarshaling process for GetTransferResponse objects.
func (g *GetTransferResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id          *string                 `json:"id"`
        Amount      *int                    `json:"amount"`
        Status      *string                 `json:"status"`
        CreatedAt   *string                 `json:"created_at"`
        UpdatedAt   *string                 `json:"updated_at"`
        BankAccount *GetBankAccountResponse `json:"bank_account"`
        Metadata    map[string]string       `json:"metadata"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Amount = temp.Amount
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
    g.BankAccount = temp.BankAccount
    g.Metadata = temp.Metadata
    return nil
}
