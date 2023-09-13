package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetRecipientResponse represents a GetRecipientResponse struct.
// Recipient response
type GetRecipientResponse struct {
    // Id
    Id                            *string                                    `json:"id"`
    // Name
    Name                          *string                                    `json:"name"`
    // Email
    Email                         *string                                    `json:"email"`
    // Document
    Document                      *string                                    `json:"document"`
    // Description
    Description                   *string                                    `json:"description"`
    // Type
    Type                          *string                                    `json:"type"`
    // Status
    Status                        *string                                    `json:"status"`
    // Creation date
    CreatedAt                     *time.Time                                 `json:"created_at"`
    // Last update date
    UpdatedAt                     *time.Time                                 `json:"updated_at"`
    // Deletion date
    DeletedAt                     *time.Time                                 `json:"deleted_at"`
    // Default bank account
    DefaultBankAccount            *GetBankAccountResponse                    `json:"default_bank_account"`
    // Info about the recipient on the gateway
    GatewayRecipients             []GetGatewayRecipientResponse              `json:"gateway_recipients"`
    // Metadata
    Metadata                      map[string]string                          `json:"metadata"`
    AutomaticAnticipationSettings Optional[GetAutomaticAnticipationResponse] `json:"automatic_anticipation_settings"`
    TransferSettings              Optional[GetTransferSettingsResponse]      `json:"transfer_settings"`
    // Recipient code
    Code                          *string                                    `json:"code"`
    // Payment mode
    PaymentMode                   *string                                    `json:"payment_mode"`
}

// MarshalJSON implements the json.Marshaler interface for GetRecipientResponse.
// It customizes the JSON marshaling process for GetRecipientResponse objects.
func (g *GetRecipientResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetRecipientResponse object to a map representation for JSON marshaling.
func (g *GetRecipientResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["name"] = g.Name
    structMap["email"] = g.Email
    structMap["document"] = g.Document
    structMap["description"] = g.Description
    structMap["type"] = g.Type
    structMap["status"] = g.Status
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    structMap["deleted_at"] = g.DeletedAt.Format(time.RFC3339)
    structMap["default_bank_account"] = g.DefaultBankAccount
    structMap["gateway_recipients"] = g.GatewayRecipients
    structMap["metadata"] = g.Metadata
    if g.AutomaticAnticipationSettings.IsValueSet() {
        structMap["automatic_anticipation_settings"] = g.AutomaticAnticipationSettings.Value()
    }
    if g.TransferSettings.IsValueSet() {
        structMap["transfer_settings"] = g.TransferSettings.Value()
    }
    structMap["code"] = g.Code
    structMap["payment_mode"] = g.PaymentMode
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetRecipientResponse.
// It customizes the JSON unmarshaling process for GetRecipientResponse objects.
func (g *GetRecipientResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                            *string                                    `json:"id"`
        Name                          *string                                    `json:"name"`
        Email                         *string                                    `json:"email"`
        Document                      *string                                    `json:"document"`
        Description                   *string                                    `json:"description"`
        Type                          *string                                    `json:"type"`
        Status                        *string                                    `json:"status"`
        CreatedAt                     *string                                    `json:"created_at"`
        UpdatedAt                     *string                                    `json:"updated_at"`
        DeletedAt                     *string                                    `json:"deleted_at"`
        DefaultBankAccount            *GetBankAccountResponse                    `json:"default_bank_account"`
        GatewayRecipients             []GetGatewayRecipientResponse              `json:"gateway_recipients"`
        Metadata                      map[string]string                          `json:"metadata"`
        AutomaticAnticipationSettings Optional[GetAutomaticAnticipationResponse] `json:"automatic_anticipation_settings"`
        TransferSettings              Optional[GetTransferSettingsResponse]      `json:"transfer_settings"`
        Code                          *string                                    `json:"code"`
        PaymentMode                   *string                                    `json:"payment_mode"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Name = temp.Name
    g.Email = temp.Email
    g.Document = temp.Document
    g.Description = temp.Description
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
    g.DefaultBankAccount = temp.DefaultBankAccount
    g.GatewayRecipients = temp.GatewayRecipients
    g.Metadata = temp.Metadata
    g.AutomaticAnticipationSettings = temp.AutomaticAnticipationSettings
    g.TransferSettings = temp.TransferSettings
    g.Code = temp.Code
    g.PaymentMode = temp.PaymentMode
    return nil
}
