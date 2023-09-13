package models

import (
    "encoding/json"
    "fmt"
)

// GetCreditCardTransactionResponse represents a GetCreditCardTransactionResponse struct.
// Response object for getting a credit card transaction
type GetCreditCardTransactionResponse struct {
    GetTransactionResponse
    // Text that will appear on the credit card's statement
    StatementDescriptor     *string          `json:"statement_descriptor"`
    // Acquirer name
    AcquirerName            string           `json:"acquirer_name"`
    // Aquirer affiliation code
    AcquirerAffiliationCode *string          `json:"acquirer_affiliation_code"`
    // Acquirer TID
    AcquirerTid             string           `json:"acquirer_tid"`
    // Acquirer NSU
    AcquirerNsu             string           `json:"acquirer_nsu"`
    // Acquirer authorization code
    AcquirerAuthCode        *string          `json:"acquirer_auth_code"`
    // Operation type
    OperationType           *string          `json:"operation_type"`
    // Card data
    Card                    *GetCardResponse `json:"card"`
    // Acquirer message
    AcquirerMessage         *string          `json:"acquirer_message"`
    // Acquirer Return Code
    AcquirerReturnCode      *string          `json:"acquirer_return_code"`
    // Number of installments
    Installments            Optional[int]    `json:"installments"`
    // 3D-S authentication Url
    ThreedAuthenticationUrl *string          `json:"threed_authentication_url"`
}

// MarshalJSON implements the json.Marshaler interface for GetCreditCardTransactionResponse.
// It customizes the JSON marshaling process for GetCreditCardTransactionResponse objects.
func (g *GetCreditCardTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetCreditCardTransactionResponse object to a map representation for JSON marshaling.
func (g *GetCreditCardTransactionResponse) toMap() map[string]any {
    structMap := g.GetTransactionResponse.toMap()
    if g.TransactionType != nil {
        structMap["transaction_type"] = *g.TransactionType
    } else {
        structMap["transaction_type"] = "credit_card"
    }
    structMap["statement_descriptor"] = g.StatementDescriptor
    structMap["acquirer_name"] = g.AcquirerName
    structMap["acquirer_affiliation_code"] = g.AcquirerAffiliationCode
    structMap["acquirer_tid"] = g.AcquirerTid
    structMap["acquirer_nsu"] = g.AcquirerNsu
    structMap["acquirer_auth_code"] = g.AcquirerAuthCode
    structMap["operation_type"] = g.OperationType
    structMap["card"] = g.Card
    structMap["acquirer_message"] = g.AcquirerMessage
    structMap["acquirer_return_code"] = g.AcquirerReturnCode
    if g.Installments.IsValueSet() {
        structMap["installments"] = g.Installments.Value()
    }
    structMap["threed_authentication_url"] = g.ThreedAuthenticationUrl
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCreditCardTransactionResponse.
// It customizes the JSON unmarshaling process for GetCreditCardTransactionResponse objects.
func (g *GetCreditCardTransactionResponse) UnmarshalJSON(input []byte) error {
    g.GetTransactionResponse.UnmarshalJSON(input)
    temp := &struct {
        StatementDescriptor     *string          `json:"statement_descriptor"`
        AcquirerName            string           `json:"acquirer_name"`
        AcquirerAffiliationCode *string          `json:"acquirer_affiliation_code"`
        AcquirerTid             string           `json:"acquirer_tid"`
        AcquirerNsu             string           `json:"acquirer_nsu"`
        AcquirerAuthCode        *string          `json:"acquirer_auth_code"`
        OperationType           *string          `json:"operation_type"`
        Card                    *GetCardResponse `json:"card"`
        AcquirerMessage         *string          `json:"acquirer_message"`
        AcquirerReturnCode      *string          `json:"acquirer_return_code"`
        Installments            Optional[int]    `json:"installments"`
        ThreedAuthenticationUrl *string          `json:"threed_authentication_url"`
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
    g.OperationType = temp.OperationType
    g.Card = temp.Card
    g.AcquirerMessage = temp.AcquirerMessage
    g.AcquirerReturnCode = temp.AcquirerReturnCode
    g.Installments = temp.Installments
    g.ThreedAuthenticationUrl = temp.ThreedAuthenticationUrl
    return nil
}

// statement_descriptor returns the statement_descriptor from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetStatementDescriptor() *string {
    return g.StatementDescriptor
}

// acquirer_name returns the acquirer_name from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerName() string {
    return g.AcquirerName
}

// acquirer_affiliation_code returns the acquirer_affiliation_code from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerAffiliationCode() *string {
    return g.AcquirerAffiliationCode
}

// acquirer_tid returns the acquirer_tid from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerTid() string {
    return g.AcquirerTid
}

// acquirer_nsu returns the acquirer_nsu from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerNsu() string {
    return g.AcquirerNsu
}

// acquirer_auth_code returns the acquirer_auth_code from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerAuthCode() *string {
    return g.AcquirerAuthCode
}

// operation_type returns the operation_type from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetOperationType() *string {
    return g.OperationType
}

// card returns the card from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetCard() *GetCardResponse {
    return g.Card
}

// acquirer_message returns the acquirer_message from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerMessage() *string {
    return g.AcquirerMessage
}

// acquirer_return_code returns the acquirer_return_code from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerReturnCode() *string {
    return g.AcquirerReturnCode
}

// installments returns the installments from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetInstallments() Optional[int] {
    return g.Installments
}

// threed_authentication_url returns the threed_authentication_url from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetThreedAuthenticationUrl() *string {
    return g.ThreedAuthenticationUrl
}

// UnmarshalGetCreditCardTransactionResponse is a utility function used to unmarshal JSON into a GetCreditCardTransactionResponseInterface.
// It takes a JSON input []byte and returns the corresponding GetCreditCardTransactionResponseInterface and an error, if any.
func UnmarshalGetCreditCardTransactionResponse(input []byte) (
    GetCreditCardTransactionResponseInterface,
    error) {
    getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
    if err != nil {
        return nil, err
    }
    
    if getCreditCardTransactionResponse, ok := getTransactionResponse.(GetCreditCardTransactionResponseInterface); ok {
        return getCreditCardTransactionResponse, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal getCreditCardTransactionResponse %v", getTransactionResponse)
}

// ToGetCreditCardTransactionResponseArray converts an array of GetCreditCardTransactionResponseField to an array of GetCreditCardTransactionResponseInterface.
func ToGetCreditCardTransactionResponseArray(getCreditCardTransactionResponse []GetCreditCardTransactionResponseField) []GetCreditCardTransactionResponseInterface {
    var items []GetCreditCardTransactionResponseInterface
    for _, temp := range getCreditCardTransactionResponse {
        items = append(items, temp.GetCreditCardTransactionResponseInterface)
    }
    return items
}

// GetCreditCardTransactionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetCreditCardTransactionResponseField struct {
    GetCreditCardTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCreditCardTransactionResponseField.
// It customizes the JSON unmarshaling process for GetCreditCardTransactionResponseField objects.
func (g *GetCreditCardTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetCreditCardTransactionResponseInterface, err = UnmarshalGetCreditCardTransactionResponse(input)
    return err
}

// GetCreditCardTransactionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetCreditCardTransactionResponseSlice struct {
    Slice []GetCreditCardTransactionResponseInterface 
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCreditCardTransactionResponseSlice.
// It customizes the JSON unmarshaling process for GetCreditCardTransactionResponseSlice objects.
func (ga *GetCreditCardTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
    	ga.Slice = []GetCreditCardTransactionResponseInterface{}
    	for _, getCreditCardTransactionResponse := range temp {
    		if getCreditCardTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetCreditCardTransactionResponse
    		err := json.Unmarshal(getCreditCardTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}
