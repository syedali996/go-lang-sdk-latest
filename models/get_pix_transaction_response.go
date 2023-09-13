package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// GetPixTransactionResponse represents a GetPixTransactionResponse struct.
// Response object when getting a pix transaction
type GetPixTransactionResponse struct {
    GetTransactionResponse
    QrCode                 *string                    `json:"qr_code"`
    QrCodeUrl              *string                    `json:"qr_code_url"`
    ExpiresAt              *time.Time                 `json:"expires_at"`
    AdditionalInformation  []PixAdditionalInformation `json:"additional_information"`
    EndToEndId             *string                    `json:"end_to_end_id"`
    Payer                  *GetPixPayerResponse       `json:"payer"`
}

// MarshalJSON implements the json.Marshaler interface for GetPixTransactionResponse.
// It customizes the JSON marshaling process for GetPixTransactionResponse objects.
func (g *GetPixTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetPixTransactionResponse object to a map representation for JSON marshaling.
func (g *GetPixTransactionResponse) toMap() map[string]any {
    structMap := g.GetTransactionResponse.toMap()
    if g.TransactionType != nil {
        structMap["transaction_type"] = *g.TransactionType
    } else {
        structMap["transaction_type"] = "pix"
    }
    structMap["qr_code"] = g.QrCode
    structMap["qr_code_url"] = g.QrCodeUrl
    structMap["expires_at"] = g.ExpiresAt.Format(time.RFC3339)
    structMap["additional_information"] = g.AdditionalInformation
    structMap["end_to_end_id"] = g.EndToEndId
    structMap["payer"] = g.Payer
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPixTransactionResponse.
// It customizes the JSON unmarshaling process for GetPixTransactionResponse objects.
func (g *GetPixTransactionResponse) UnmarshalJSON(input []byte) error {
    g.GetTransactionResponse.UnmarshalJSON(input)
    temp := &struct {
        QrCode                *string                    `json:"qr_code"`
        QrCodeUrl             *string                    `json:"qr_code_url"`
        ExpiresAt             *string                    `json:"expires_at"`
        AdditionalInformation []PixAdditionalInformation `json:"additional_information"`
        EndToEndId            *string                    `json:"end_to_end_id"`
        Payer                 *GetPixPayerResponse       `json:"payer"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.QrCode = temp.QrCode
    g.QrCodeUrl = temp.QrCodeUrl
    if temp.ExpiresAt != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        g.ExpiresAt = &ExpiresAtVal
    }
    g.AdditionalInformation = temp.AdditionalInformation
    g.EndToEndId = temp.EndToEndId
    g.Payer = temp.Payer
    return nil
}

// qr_code returns the qr_code from the GetPixTransactionResponse.
func (g *GetPixTransactionResponse) GetQrCode() *string {
    return g.QrCode
}

// qr_code_url returns the qr_code_url from the GetPixTransactionResponse.
func (g *GetPixTransactionResponse) GetQrCodeUrl() *string {
    return g.QrCodeUrl
}

// expires_at returns the expires_at from the GetPixTransactionResponse.
func (g *GetPixTransactionResponse) GetExpiresAt() *time.Time {
    return g.ExpiresAt
}

// additional_information returns the additional_information from the GetPixTransactionResponse.
func (g *GetPixTransactionResponse) GetAdditionalInformation() []PixAdditionalInformation {
    return g.AdditionalInformation
}

// end_to_end_id returns the end_to_end_id from the GetPixTransactionResponse.
func (g *GetPixTransactionResponse) GetEndToEndId() *string {
    return g.EndToEndId
}

// payer returns the payer from the GetPixTransactionResponse.
func (g *GetPixTransactionResponse) GetPayer() *GetPixPayerResponse {
    return g.Payer
}

// UnmarshalGetPixTransactionResponse is a utility function used to unmarshal JSON into a GetPixTransactionResponseInterface.
// It takes a JSON input []byte and returns the corresponding GetPixTransactionResponseInterface and an error, if any.
func UnmarshalGetPixTransactionResponse(input []byte) (
    GetPixTransactionResponseInterface,
    error) {
    getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
    if err != nil {
        return nil, err
    }
    
    if getPixTransactionResponse, ok := getTransactionResponse.(GetPixTransactionResponseInterface); ok {
        return getPixTransactionResponse, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal getPixTransactionResponse %v", getTransactionResponse)
}

// ToGetPixTransactionResponseArray converts an array of GetPixTransactionResponseField to an array of GetPixTransactionResponseInterface.
func ToGetPixTransactionResponseArray(getPixTransactionResponse []GetPixTransactionResponseField) []GetPixTransactionResponseInterface {
    var items []GetPixTransactionResponseInterface
    for _, temp := range getPixTransactionResponse {
        items = append(items, temp.GetPixTransactionResponseInterface)
    }
    return items
}

// GetPixTransactionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetPixTransactionResponseField struct {
    GetPixTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPixTransactionResponseField.
// It customizes the JSON unmarshaling process for GetPixTransactionResponseField objects.
func (g *GetPixTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetPixTransactionResponseInterface, err = UnmarshalGetPixTransactionResponse(input)
    return err
}

// GetPixTransactionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetPixTransactionResponseSlice struct {
    Slice []GetPixTransactionResponseInterface 
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPixTransactionResponseSlice.
// It customizes the JSON unmarshaling process for GetPixTransactionResponseSlice objects.
func (ga *GetPixTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
    	ga.Slice = []GetPixTransactionResponseInterface{}
    	for _, getPixTransactionResponse := range temp {
    		if getPixTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetPixTransactionResponse
    		err := json.Unmarshal(getPixTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}
