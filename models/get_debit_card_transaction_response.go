package models

import (
    "encoding/json"
    "fmt"
)

// GetDebitCardTransactionResponse represents a GetDebitCardTransactionResponse struct.
// Response object for getting a debit card transaction
type GetDebitCardTransactionResponse struct {
    GetTransactionResponse
    // Text that will appear on the debit card's statement
    StatementDescriptor     *string          `json:"statement_descriptor"`
    // Acquirer name
    AcquirerName            *string          `json:"acquirer_name"`
    // Aquirer affiliation code
    AcquirerAffiliationCode *string          `json:"acquirer_affiliation_code"`
    // Acquirer TID
    AcquirerTid             *string          `json:"acquirer_tid"`
    // Acquirer NSU
    AcquirerNsu             *string          `json:"acquirer_nsu"`
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
    // Merchant Plugin
    Mpi                     *string          `json:"mpi"`
    // Electronic Commerce Indicator (ECI)
    Eci                     *string          `json:"eci"`
    // Authentication type
    AuthenticationType      *string          `json:"authentication_type"`
    // 3D-S Authentication Url
    ThreedAuthenticationUrl *string          `json:"threed_authentication_url"`
}

// MarshalJSON implements the json.Marshaler interface for GetDebitCardTransactionResponse.
// It customizes the JSON marshaling process for GetDebitCardTransactionResponse objects.
func (g *GetDebitCardTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetDebitCardTransactionResponse object to a map representation for JSON marshaling.
func (g *GetDebitCardTransactionResponse) toMap() map[string]any {
    structMap := g.GetTransactionResponse.toMap()
    if g.TransactionType != nil {
        structMap["transaction_type"] = *g.TransactionType
    } else {
        structMap["transaction_type"] = "debit_card"
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
    structMap["mpi"] = g.Mpi
    structMap["eci"] = g.Eci
    structMap["authentication_type"] = g.AuthenticationType
    structMap["threed_authentication_url"] = g.ThreedAuthenticationUrl
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetDebitCardTransactionResponse.
// It customizes the JSON unmarshaling process for GetDebitCardTransactionResponse objects.
func (g *GetDebitCardTransactionResponse) UnmarshalJSON(input []byte) error {
    g.GetTransactionResponse.UnmarshalJSON(input)
    temp := &struct {
        StatementDescriptor     *string          `json:"statement_descriptor"`
        AcquirerName            *string          `json:"acquirer_name"`
        AcquirerAffiliationCode *string          `json:"acquirer_affiliation_code"`
        AcquirerTid             *string          `json:"acquirer_tid"`
        AcquirerNsu             *string          `json:"acquirer_nsu"`
        AcquirerAuthCode        *string          `json:"acquirer_auth_code"`
        OperationType           *string          `json:"operation_type"`
        Card                    *GetCardResponse `json:"card"`
        AcquirerMessage         *string          `json:"acquirer_message"`
        AcquirerReturnCode      *string          `json:"acquirer_return_code"`
        Mpi                     *string          `json:"mpi"`
        Eci                     *string          `json:"eci"`
        AuthenticationType      *string          `json:"authentication_type"`
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
    g.Mpi = temp.Mpi
    g.Eci = temp.Eci
    g.AuthenticationType = temp.AuthenticationType
    g.ThreedAuthenticationUrl = temp.ThreedAuthenticationUrl
    return nil
}

// statement_descriptor returns the statement_descriptor from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetStatementDescriptor() *string {
    return g.StatementDescriptor
}

// acquirer_name returns the acquirer_name from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerName() *string {
    return g.AcquirerName
}

// acquirer_affiliation_code returns the acquirer_affiliation_code from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerAffiliationCode() *string {
    return g.AcquirerAffiliationCode
}

// acquirer_tid returns the acquirer_tid from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerTid() *string {
    return g.AcquirerTid
}

// acquirer_nsu returns the acquirer_nsu from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerNsu() *string {
    return g.AcquirerNsu
}

// acquirer_auth_code returns the acquirer_auth_code from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerAuthCode() *string {
    return g.AcquirerAuthCode
}

// operation_type returns the operation_type from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetOperationType() *string {
    return g.OperationType
}

// card returns the card from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetCard() *GetCardResponse {
    return g.Card
}

// acquirer_message returns the acquirer_message from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerMessage() *string {
    return g.AcquirerMessage
}

// acquirer_return_code returns the acquirer_return_code from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerReturnCode() *string {
    return g.AcquirerReturnCode
}

// mpi returns the mpi from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetMpi() *string {
    return g.Mpi
}

// eci returns the eci from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetEci() *string {
    return g.Eci
}

// authentication_type returns the authentication_type from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAuthenticationType() *string {
    return g.AuthenticationType
}

// threed_authentication_url returns the threed_authentication_url from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetThreedAuthenticationUrl() *string {
    return g.ThreedAuthenticationUrl
}

// UnmarshalGetDebitCardTransactionResponse is a utility function used to unmarshal JSON into a GetDebitCardTransactionResponseInterface.
// It takes a JSON input []byte and returns the corresponding GetDebitCardTransactionResponseInterface and an error, if any.
func UnmarshalGetDebitCardTransactionResponse(input []byte) (
    GetDebitCardTransactionResponseInterface,
    error) {
    getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
    if err != nil {
        return nil, err
    }
    
    if getDebitCardTransactionResponse, ok := getTransactionResponse.(GetDebitCardTransactionResponseInterface); ok {
        return getDebitCardTransactionResponse, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal getDebitCardTransactionResponse %v", getTransactionResponse)
}

// ToGetDebitCardTransactionResponseArray converts an array of GetDebitCardTransactionResponseField to an array of GetDebitCardTransactionResponseInterface.
func ToGetDebitCardTransactionResponseArray(getDebitCardTransactionResponse []GetDebitCardTransactionResponseField) []GetDebitCardTransactionResponseInterface {
    var items []GetDebitCardTransactionResponseInterface
    for _, temp := range getDebitCardTransactionResponse {
        items = append(items, temp.GetDebitCardTransactionResponseInterface)
    }
    return items
}

// GetDebitCardTransactionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetDebitCardTransactionResponseField struct {
    GetDebitCardTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetDebitCardTransactionResponseField.
// It customizes the JSON unmarshaling process for GetDebitCardTransactionResponseField objects.
func (g *GetDebitCardTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetDebitCardTransactionResponseInterface, err = UnmarshalGetDebitCardTransactionResponse(input)
    return err
}

// GetDebitCardTransactionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetDebitCardTransactionResponseSlice struct {
    Slice []GetDebitCardTransactionResponseInterface 
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetDebitCardTransactionResponseSlice.
// It customizes the JSON unmarshaling process for GetDebitCardTransactionResponseSlice objects.
func (ga *GetDebitCardTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
    	ga.Slice = []GetDebitCardTransactionResponseInterface{}
    	for _, getDebitCardTransactionResponse := range temp {
    		if getDebitCardTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetDebitCardTransactionResponse
    		err := json.Unmarshal(getDebitCardTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}
