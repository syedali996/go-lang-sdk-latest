package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetBankAccountResponse represents a GetBankAccountResponse struct.
type GetBankAccountResponse struct {
    // Id
    Id                *string                        `json:"id"`
    // Holder name
    HolderName        *string                        `json:"holder_name"`
    // Holder type
    HolderType        *string                        `json:"holder_type"`
    // Bank
    Bank              *string                        `json:"bank"`
    // Branch number
    BranchNumber      *string                        `json:"branch_number"`
    // Branch check digit
    BranchCheckDigit  *string                        `json:"branch_check_digit"`
    // Account number
    AccountNumber     *string                        `json:"account_number"`
    // Account check digit
    AccountCheckDigit *string                        `json:"account_check_digit"`
    // Bank account type
    Type              *string                        `json:"type"`
    // Bank account status
    Status            *string                        `json:"status"`
    // Creation date
    CreatedAt         *time.Time                     `json:"created_at"`
    // Last update date
    UpdatedAt         *time.Time                     `json:"updated_at"`
    // Deletion date
    DeletedAt         *time.Time                     `json:"deleted_at"`
    // Recipient
    Recipient         Optional[GetRecipientResponse] `json:"recipient"`
    // Metadata
    Metadata          map[string]string              `json:"metadata"`
    // Pix Key
    PixKey            *string                        `json:"pix_key"`
}

// MarshalJSON implements the json.Marshaler interface for GetBankAccountResponse.
// It customizes the JSON marshaling process for GetBankAccountResponse objects.
func (g *GetBankAccountResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetBankAccountResponse object to a map representation for JSON marshaling.
func (g *GetBankAccountResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["holder_name"] = g.HolderName
    structMap["holder_type"] = g.HolderType
    structMap["bank"] = g.Bank
    structMap["branch_number"] = g.BranchNumber
    structMap["branch_check_digit"] = g.BranchCheckDigit
    structMap["account_number"] = g.AccountNumber
    structMap["account_check_digit"] = g.AccountCheckDigit
    structMap["type"] = g.Type
    structMap["status"] = g.Status
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    structMap["deleted_at"] = g.DeletedAt.Format(time.RFC3339)
    if g.Recipient.IsValueSet() {
        structMap["recipient"] = g.Recipient.Value()
    }
    structMap["metadata"] = g.Metadata
    structMap["pix_key"] = g.PixKey
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetBankAccountResponse.
// It customizes the JSON unmarshaling process for GetBankAccountResponse objects.
func (g *GetBankAccountResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                *string                        `json:"id"`
        HolderName        *string                        `json:"holder_name"`
        HolderType        *string                        `json:"holder_type"`
        Bank              *string                        `json:"bank"`
        BranchNumber      *string                        `json:"branch_number"`
        BranchCheckDigit  *string                        `json:"branch_check_digit"`
        AccountNumber     *string                        `json:"account_number"`
        AccountCheckDigit *string                        `json:"account_check_digit"`
        Type              *string                        `json:"type"`
        Status            *string                        `json:"status"`
        CreatedAt         *string                        `json:"created_at"`
        UpdatedAt         *string                        `json:"updated_at"`
        DeletedAt         *string                        `json:"deleted_at"`
        Recipient         Optional[GetRecipientResponse] `json:"recipient"`
        Metadata          map[string]string              `json:"metadata"`
        PixKey            *string                        `json:"pix_key"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.HolderName = temp.HolderName
    g.HolderType = temp.HolderType
    g.Bank = temp.Bank
    g.BranchNumber = temp.BranchNumber
    g.BranchCheckDigit = temp.BranchCheckDigit
    g.AccountNumber = temp.AccountNumber
    g.AccountCheckDigit = temp.AccountCheckDigit
    g.Type = temp.Type
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
    if temp.DeletedAt != nil {
        DeletedAtVal, err := time.Parse(time.RFC3339, *temp.DeletedAt)
        if err != nil {
            log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
        }
        g.DeletedAt = &DeletedAtVal
    }
    g.Recipient = temp.Recipient
    g.Metadata = temp.Metadata
    g.PixKey = temp.PixKey
    return nil
}
