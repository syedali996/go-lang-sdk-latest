package models

import (
    "encoding/json"
    "fmt"
)

// GetVoucherTransactionResponse represents a GetVoucherTransactionResponse struct.
// Response for voucher transactions
type GetVoucherTransactionResponse struct {
    GetTransactionResponse
    // Text that will appear on the voucher's statement
    StatementDescriptor     *string          `json:"statement_descriptor"`
    // Acquirer name
    AcquirerName            *string          `json:"acquirer_name"`
    // Acquirer affiliation code
    AcquirerAffiliationCode *string          `json:"acquirer_affiliation_code"`
    // Acquirer TID
    AcquirerTid             *string          `json:"acquirer_tid"`
    // Acquirer NSU
    AcquirerNsu             *string          `json:"acquirer_nsu"`
    // Acquirer authorization code
    AcquirerAuthCode        *string          `json:"acquirer_auth_code"`
    // acquirer_message
    AcquirerMessage         *string          `json:"acquirer_message"`
    // Acquirer return code
    AcquirerReturnCode      *string          `json:"acquirer_return_code"`
    // Operation type
    OperationType           *string          `json:"operation_type"`
    // Card data
    Card                    *GetCardResponse `json:"card"`
}

// MarshalJSON implements the json.Marshaler interface for GetVoucherTransactionResponse.
// It customizes the JSON marshaling process for GetVoucherTransactionResponse objects.
func (g *GetVoucherTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetVoucherTransactionResponse object to a map representation for JSON marshaling.
func (g *GetVoucherTransactionResponse) toMap() map[string]any {
    structMap := g.GetTransactionResponse.toMap()
    if g.TransactionType != nil {
        structMap["transaction_type"] = *g.TransactionType
    } else {
        structMap["transaction_type"] = "voucher"
    }
    structMap["statement_descriptor"] = g.StatementDescriptor
    structMap["acquirer_name"] = g.AcquirerName
    structMap["acquirer_affiliation_code"] = g.AcquirerAffiliationCode
    structMap["acquirer_tid"] = g.AcquirerTid
    structMap["acquirer_nsu"] = g.AcquirerNsu
    structMap["acquirer_auth_code"] = g.AcquirerAuthCode
    structMap["acquirer_message"] = g.AcquirerMessage
    structMap["acquirer_return_code"] = g.AcquirerReturnCode
    structMap["operation_type"] = g.OperationType
    structMap["card"] = g.Card
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetVoucherTransactionResponse.
// It customizes the JSON unmarshaling process for GetVoucherTransactionResponse objects.
func (g *GetVoucherTransactionResponse) UnmarshalJSON(input []byte) error {
    g.GetTransactionResponse.UnmarshalJSON(input)
    temp := &struct {
        StatementDescriptor     *string          `json:"statement_descriptor"`
        AcquirerName            *string          `json:"acquirer_name"`
        AcquirerAffiliationCode *string          `json:"acquirer_affiliation_code"`
        AcquirerTid             *string          `json:"acquirer_tid"`
        AcquirerNsu             *string          `json:"acquirer_nsu"`
        AcquirerAuthCode        *string          `json:"acquirer_auth_code"`
        AcquirerMessage         *string          `json:"acquirer_message"`
        AcquirerReturnCode      *string          `json:"acquirer_return_code"`
        OperationType           *string          `json:"operation_type"`
        Card                    *GetCardResponse `json:"card"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.StatementDescriptor = temp.StatementDescriptor
    g.AcquirerName = temp.AcquirerName
    g.AcquirerAffiliationCode = temp.AcquirerAffiliationCode
    g.AcquirerTid = temp.AcquirerTid
    g.AcquirerNsu = temp.AcquirerNsu
    g.AcquirerAuthCode = temp.AcquirerAuthCode
    g.AcquirerMessage = temp.AcquirerMessage
    g.AcquirerReturnCode = temp.AcquirerReturnCode
    g.OperationType = temp.OperationType
    g.Card = temp.Card
    return nil
}

// statement_descriptor returns the statement_descriptor from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetStatementDescriptor() *string {
    return g.StatementDescriptor
}

// acquirer_name returns the acquirer_name from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerName() *string {
    return g.AcquirerName
}

// acquirer_affiliation_code returns the acquirer_affiliation_code from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerAffiliationCode() *string {
    return g.AcquirerAffiliationCode
}

// acquirer_tid returns the acquirer_tid from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerTid() *string {
    return g.AcquirerTid
}

// acquirer_nsu returns the acquirer_nsu from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerNsu() *string {
    return g.AcquirerNsu
}

// acquirer_auth_code returns the acquirer_auth_code from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerAuthCode() *string {
    return g.AcquirerAuthCode
}

// acquirer_message returns the acquirer_message from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerMessage() *string {
    return g.AcquirerMessage
}

// acquirer_return_code returns the acquirer_return_code from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerReturnCode() *string {
    return g.AcquirerReturnCode
}

// operation_type returns the operation_type from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetOperationType() *string {
    return g.OperationType
}

// card returns the card from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetCard() *GetCardResponse {
    return g.Card
}

// UnmarshalGetVoucherTransactionResponse is a utility function used to unmarshal JSON into a GetVoucherTransactionResponseInterface.
// It takes a JSON input []byte and returns the corresponding GetVoucherTransactionResponseInterface and an error, if any.
func UnmarshalGetVoucherTransactionResponse(input []byte) (
    GetVoucherTransactionResponseInterface,
    error) {
    getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
    if err != nil {
        return nil, err
    }
    
    if getVoucherTransactionResponse, ok := getTransactionResponse.(GetVoucherTransactionResponseInterface); ok {
        return getVoucherTransactionResponse, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal getVoucherTransactionResponse %v", getTransactionResponse)
}

// ToGetVoucherTransactionResponseArray converts an array of GetVoucherTransactionResponseField to an array of GetVoucherTransactionResponseInterface.
func ToGetVoucherTransactionResponseArray(getVoucherTransactionResponse []GetVoucherTransactionResponseField) []GetVoucherTransactionResponseInterface {
    var items []GetVoucherTransactionResponseInterface
    for _, temp := range getVoucherTransactionResponse {
        items = append(items, temp.GetVoucherTransactionResponseInterface)
    }
    return items
}

// GetVoucherTransactionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetVoucherTransactionResponseField struct {
    GetVoucherTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetVoucherTransactionResponseField.
// It customizes the JSON unmarshaling process for GetVoucherTransactionResponseField objects.
func (g *GetVoucherTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetVoucherTransactionResponseInterface, err = UnmarshalGetVoucherTransactionResponse(input)
    return err
}

// GetVoucherTransactionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetVoucherTransactionResponseSlice struct {
    Slice []GetVoucherTransactionResponseInterface 
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetVoucherTransactionResponseSlice.
// It customizes the JSON unmarshaling process for GetVoucherTransactionResponseSlice objects.
func (ga *GetVoucherTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
    	ga.Slice = []GetVoucherTransactionResponseInterface{}
    	for _, getVoucherTransactionResponse := range temp {
    		if getVoucherTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetVoucherTransactionResponse
    		err := json.Unmarshal(getVoucherTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}
