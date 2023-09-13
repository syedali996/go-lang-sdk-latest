package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetTransactionResponse represents a GetTransactionResponse struct.
// Generic response object for getting a transaction.
type GetTransactionResponse struct {
    // Gateway transaction id
    GatewayId           *string                       `json:"gateway_id"`
    // Amount in cents
    Amount              *int                          `json:"amount"`
    // Transaction status
    Status              *string                       `json:"status"`
    // Indicates if the transaction ocurred successfuly
    Success             *bool                         `json:"success"`
    // Creation date
    CreatedAt           *time.Time                    `json:"created_at"`
    // Last update date
    UpdatedAt           *time.Time                    `json:"updated_at"`
    // Number of attempts tried
    AttemptCount        *int                          `json:"attempt_count"`
    // Max attempts
    MaxAttempts         *int                          `json:"max_attempts"`
    // Splits
    Splits              []GetSplitResponse            `json:"splits"`
    // Date and time of the next attempt
    NextAttempt         Optional[time.Time]           `json:"next_attempt"`
    TransactionType     *string                       `json:"transaction_type,omitempty"`
    // Código da transação
    Id                  *string                       `json:"id"`
    // The Gateway Response
    GatewayResponse     *GetGatewayResponseResponse   `json:"gateway_response"`
    AntifraudResponse   *GetAntifraudResponse         `json:"antifraud_response"`
    Metadata            Optional[map[string]string]   `json:"metadata"`
    Split               []GetSplitResponse            `json:"split"`
    Interest            Optional[GetInterestResponse] `json:"interest"`
    Fine                Optional[GetFineResponse]     `json:"fine"`
    MaxDaysToPayPastDue Optional[int]                 `json:"max_days_to_pay_past_due"`
}

// MarshalJSON implements the json.Marshaler interface for GetTransactionResponse.
// It customizes the JSON marshaling process for GetTransactionResponse objects.
func (g *GetTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetTransactionResponse object to a map representation for JSON marshaling.
func (g *GetTransactionResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    if g.TransactionType != nil {
        structMap["transaction_type"] = *g.TransactionType
    } else {
        structMap["transaction_type"] = "transaction"
    }
    structMap["gateway_id"] = g.GatewayId
    structMap["amount"] = g.Amount
    structMap["status"] = g.Status
    structMap["success"] = g.Success
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
    structMap["attempt_count"] = g.AttemptCount
    structMap["max_attempts"] = g.MaxAttempts
    structMap["splits"] = g.Splits
    if g.NextAttempt.IsValueSet() {
        var NextAttemptVal *string = nil
        if g.NextAttempt.Value() != nil {
            val := g.NextAttempt.Value().Format(time.RFC3339)
            NextAttemptVal = &val
        }
        structMap["next_attempt"] = NextAttemptVal
    }
    structMap["id"] = g.Id
    structMap["gateway_response"] = g.GatewayResponse
    structMap["antifraud_response"] = g.AntifraudResponse
    if g.Metadata.IsValueSet() {
        structMap["metadata"] = g.Metadata.Value()
    }
    structMap["split"] = g.Split
    if g.Interest.IsValueSet() {
        structMap["interest"] = g.Interest.Value()
    }
    if g.Fine.IsValueSet() {
        structMap["fine"] = g.Fine.Value()
    }
    if g.MaxDaysToPayPastDue.IsValueSet() {
        structMap["max_days_to_pay_past_due"] = g.MaxDaysToPayPastDue.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetTransactionResponse.
// It customizes the JSON unmarshaling process for GetTransactionResponse objects.
func (g *GetTransactionResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        GatewayId           *string                       `json:"gateway_id"`
        Amount              *int                          `json:"amount"`
        Status              *string                       `json:"status"`
        Success             *bool                         `json:"success"`
        CreatedAt           *string                       `json:"created_at"`
        UpdatedAt           *string                       `json:"updated_at"`
        AttemptCount        *int                          `json:"attempt_count"`
        MaxAttempts         *int                          `json:"max_attempts"`
        Splits              []GetSplitResponse            `json:"splits"`
        NextAttempt         Optional[string]              `json:"next_attempt"`
        TransactionType     *string                       `json:"transaction_type,omitempty"`
        Id                  *string                       `json:"id"`
        GatewayResponse     *GetGatewayResponseResponse   `json:"gateway_response"`
        AntifraudResponse   *GetAntifraudResponse         `json:"antifraud_response"`
        Metadata            Optional[map[string]string]   `json:"metadata"`
        Split               []GetSplitResponse            `json:"split"`
        Interest            Optional[GetInterestResponse] `json:"interest"`
        Fine                Optional[GetFineResponse]     `json:"fine"`
        MaxDaysToPayPastDue Optional[int]                 `json:"max_days_to_pay_past_due"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.GatewayId = temp.GatewayId
    g.Amount = temp.Amount
    g.Status = temp.Status
    g.Success = temp.Success
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
    g.AttemptCount = temp.AttemptCount
    g.MaxAttempts = temp.MaxAttempts
    g.Splits = temp.Splits
    g.NextAttempt.ShouldSetValue(temp.NextAttempt.IsValueSet())
    if temp.NextAttempt.Value() != nil {
        NextAttemptVal, err := time.Parse(time.RFC3339, (*temp.NextAttempt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse next_attempt as % s format.", time.RFC3339)
        }
        g.NextAttempt.SetValue(&NextAttemptVal)
    }
    g.TransactionType = temp.TransactionType
    g.Id = temp.Id
    g.GatewayResponse = temp.GatewayResponse
    g.AntifraudResponse = temp.AntifraudResponse
    g.Metadata = temp.Metadata
    g.Split = temp.Split
    g.Interest = temp.Interest
    g.Fine = temp.Fine
    g.MaxDaysToPayPastDue = temp.MaxDaysToPayPastDue
    return nil
}

// gateway_id returns the gateway_id from the GetTransactionResponse.
func (g *GetTransactionResponse) GetGatewayId() *string {
    return g.GatewayId
}

// amount returns the amount from the GetTransactionResponse.
func (g *GetTransactionResponse) GetAmount() *int {
    return g.Amount
}

// status returns the status from the GetTransactionResponse.
func (g *GetTransactionResponse) GetStatus() *string {
    return g.Status
}

// success returns the success from the GetTransactionResponse.
func (g *GetTransactionResponse) GetSuccess() *bool {
    return g.Success
}

// created_at returns the created_at from the GetTransactionResponse.
func (g *GetTransactionResponse) GetCreatedAt() *time.Time {
    return g.CreatedAt
}

// updated_at returns the updated_at from the GetTransactionResponse.
func (g *GetTransactionResponse) GetUpdatedAt() *time.Time {
    return g.UpdatedAt
}

// attempt_count returns the attempt_count from the GetTransactionResponse.
func (g *GetTransactionResponse) GetAttemptCount() *int {
    return g.AttemptCount
}

// max_attempts returns the max_attempts from the GetTransactionResponse.
func (g *GetTransactionResponse) GetMaxAttempts() *int {
    return g.MaxAttempts
}

// splits returns the splits from the GetTransactionResponse.
func (g *GetTransactionResponse) GetSplits() []GetSplitResponse {
    return g.Splits
}

// next_attempt returns the next_attempt from the GetTransactionResponse.
func (g *GetTransactionResponse) GetNextAttempt() Optional[time.Time] {
    return g.NextAttempt
}

// transaction_type returns the transaction_type from the GetTransactionResponse.
func (g *GetTransactionResponse) GetTransactionType() *string {
    return g.TransactionType
}

// id returns the id from the GetTransactionResponse.
func (g *GetTransactionResponse) GetId() *string {
    return g.Id
}

// gateway_response returns the gateway_response from the GetTransactionResponse.
func (g *GetTransactionResponse) GetGatewayResponse() *GetGatewayResponseResponse {
    return g.GatewayResponse
}

// antifraud_response returns the antifraud_response from the GetTransactionResponse.
func (g *GetTransactionResponse) GetAntifraudResponse() *GetAntifraudResponse {
    return g.AntifraudResponse
}

// metadata returns the metadata from the GetTransactionResponse.
func (g *GetTransactionResponse) GetMetadata() Optional[map[string]string] {
    return g.Metadata
}

// split returns the split from the GetTransactionResponse.
func (g *GetTransactionResponse) GetSplit() []GetSplitResponse {
    return g.Split
}

// interest returns the interest from the GetTransactionResponse.
func (g *GetTransactionResponse) GetInterest() Optional[GetInterestResponse] {
    return g.Interest
}

// fine returns the fine from the GetTransactionResponse.
func (g *GetTransactionResponse) GetFine() Optional[GetFineResponse] {
    return g.Fine
}

// max_days_to_pay_past_due returns the max_days_to_pay_past_due from the GetTransactionResponse.
func (g *GetTransactionResponse) GetMaxDaysToPayPastDue() Optional[int] {
    return g.MaxDaysToPayPastDue
}

// UnmarshalGetTransactionResponse is a utility function used to unmarshal JSON into a GetTransactionResponseInterface.
// It takes a JSON input []byte and returns the corresponding GetTransactionResponseInterface and an error, if any.
func UnmarshalGetTransactionResponse(input []byte) (
    GetTransactionResponseInterface,
    error) {
    discrim := &struct {
        TransactionType *string
    }{}
    
    err := json.Unmarshal(input, &discrim)
    if err != nil {
        return nil, err
    }
    if discrim.TransactionType == nil {
        transactionType := "transaction"
        discrim.TransactionType = &transactionType
    }
    
    var res GetTransactionResponseInterface
    switch *discrim.TransactionType {
    case "bank_transfer":
        res = &GetBankTransferTransactionResponse{}
    case "safetypay":
        res = &GetSafetyPayTransactionResponse{}
    case "voucher":
        res = &GetVoucherTransactionResponse{}
    case "boleto":
        res = &GetBoletoTransactionResponse{}
    case "debit_card":
        res = &GetDebitCardTransactionResponse{}
    case "private_label":
        res = &GetPrivateLabelTransactionResponse{}
    case "cash":
        res = &GetCashTransactionResponse{}
    case "credit_card":
        res = &GetCreditCardTransactionResponse{}
    case "pix":
        res = &GetPixTransactionResponse{}
    default:
        res = &GetTransactionResponse{}
    }
    json.Unmarshal(input, res)
    return res, nil
}

// ToGetTransactionResponseArray converts an array of GetTransactionResponseField to an array of GetTransactionResponseInterface.
func ToGetTransactionResponseArray(getTransactionResponse []GetTransactionResponseField) []GetTransactionResponseInterface {
    var items []GetTransactionResponseInterface
    for _, temp := range getTransactionResponse {
        items = append(items, temp.GetTransactionResponseInterface)
    }
    return items
}

// GetTransactionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetTransactionResponseField struct {
    GetTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetTransactionResponseField.
// It customizes the JSON unmarshaling process for GetTransactionResponseField objects.
func (g *GetTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetTransactionResponseInterface, err = UnmarshalGetTransactionResponse(input)
    return err
}

// GetTransactionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetTransactionResponseSlice struct {
    Slice []GetTransactionResponseInterface 
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetTransactionResponseSlice.
// It customizes the JSON unmarshaling process for GetTransactionResponseSlice objects.
func (ga *GetTransactionResponseSlice) UnmarshalJSON(input []byte) error {
    tempStruct := struct {
    	Slice []json.RawMessage
    }{}
    temp := []json.RawMessage{}
    err := json.Unmarshal(input, &tempStruct)
    if err != nil {
    	err := json.Unmarshal(input, &temp)
    	if err != nil {
    		return err
    	}
    } else {
    	temp = tempStruct.Slice
    }
    
    ga.Slice = nil
    if temp != nil {
    	ga.Slice = []GetTransactionResponseInterface{}
    	for _, getTransactionResponse := range temp {
    		if getTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetTransactionResponse
    		err := json.Unmarshal(getTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}
