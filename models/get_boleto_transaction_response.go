package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// GetBoletoTransactionResponse represents a GetBoletoTransactionResponse struct.
// Response object for getting a boleto transaction
type GetBoletoTransactionResponse struct {
    GetTransactionResponse
    Url                    *string                    `json:"url"`
    Barcode                *string                    `json:"barcode"`
    NossoNumero            *string                    `json:"nosso_numero"`
    Bank                   *string                    `json:"bank"`
    DocumentNumber         *string                    `json:"document_number"`
    Instructions           *string                    `json:"instructions"`
    BillingAddress         *GetBillingAddressResponse `json:"billing_address"`
    DueAt                  Optional[time.Time]        `json:"due_at"`
    QrCode                 *string                    `json:"qr_code"`
    Line                   *string                    `json:"line"`
    PdfPassword            *string                    `json:"pdf_password"`
    Pdf                    *string                    `json:"pdf"`
    PaidAt                 Optional[time.Time]        `json:"paid_at"`
    PaidAmount             *string                    `json:"paid_amount"`
    Type                   *string                    `json:"type"`
    CreditAt               Optional[time.Time]        `json:"credit_at"`
    // Soft Descriptor
    StatementDescriptor    *string                    `json:"statement_descriptor"`
}

// MarshalJSON implements the json.Marshaler interface for GetBoletoTransactionResponse.
// It customizes the JSON marshaling process for GetBoletoTransactionResponse objects.
func (g *GetBoletoTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetBoletoTransactionResponse object to a map representation for JSON marshaling.
func (g *GetBoletoTransactionResponse) toMap() map[string]any {
    structMap := g.GetTransactionResponse.toMap()
    if g.TransactionType != nil {
        structMap["transaction_type"] = *g.TransactionType
    } else {
        structMap["transaction_type"] = "boleto"
    }
    structMap["url"] = g.Url
    structMap["barcode"] = g.Barcode
    structMap["nosso_numero"] = g.NossoNumero
    structMap["bank"] = g.Bank
    structMap["document_number"] = g.DocumentNumber
    structMap["instructions"] = g.Instructions
    structMap["billing_address"] = g.BillingAddress
    if g.DueAt.IsValueSet() {
        var DueAtVal *string = nil
        if g.DueAt.Value() != nil {
            val := g.DueAt.Value().Format(time.RFC3339)
            DueAtVal = &val
        }
        structMap["due_at"] = DueAtVal
    }
    structMap["qr_code"] = g.QrCode
    structMap["line"] = g.Line
    structMap["pdf_password"] = g.PdfPassword
    structMap["pdf"] = g.Pdf
    if g.PaidAt.IsValueSet() {
        var PaidAtVal *string = nil
        if g.PaidAt.Value() != nil {
            val := g.PaidAt.Value().Format(time.RFC3339)
            PaidAtVal = &val
        }
        structMap["paid_at"] = PaidAtVal
    }
    structMap["paid_amount"] = g.PaidAmount
    structMap["type"] = g.Type
    if g.CreditAt.IsValueSet() {
        var CreditAtVal *string = nil
        if g.CreditAt.Value() != nil {
            val := g.CreditAt.Value().Format(time.RFC3339)
            CreditAtVal = &val
        }
        structMap["credit_at"] = CreditAtVal
    }
    structMap["statement_descriptor"] = g.StatementDescriptor
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetBoletoTransactionResponse.
// It customizes the JSON unmarshaling process for GetBoletoTransactionResponse objects.
func (g *GetBoletoTransactionResponse) UnmarshalJSON(input []byte) error {
    g.GetTransactionResponse.UnmarshalJSON(input)
    temp := &struct {
        Url                 *string                    `json:"url"`
        Barcode             *string                    `json:"barcode"`
        NossoNumero         *string                    `json:"nosso_numero"`
        Bank                *string                    `json:"bank"`
        DocumentNumber      *string                    `json:"document_number"`
        Instructions        *string                    `json:"instructions"`
        BillingAddress      *GetBillingAddressResponse `json:"billing_address"`
        DueAt               Optional[string]           `json:"due_at"`
        QrCode              *string                    `json:"qr_code"`
        Line                *string                    `json:"line"`
        PdfPassword         *string                    `json:"pdf_password"`
        Pdf                 *string                    `json:"pdf"`
        PaidAt              Optional[string]           `json:"paid_at"`
        PaidAmount          *string                    `json:"paid_amount"`
        Type                *string                    `json:"type"`
        CreditAt            Optional[string]           `json:"credit_at"`
        StatementDescriptor *string                    `json:"statement_descriptor"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Url = temp.Url
    g.Barcode = temp.Barcode
    g.NossoNumero = temp.NossoNumero
    g.Bank = temp.Bank
    g.DocumentNumber = temp.DocumentNumber
    g.Instructions = temp.Instructions
    g.BillingAddress = temp.BillingAddress
    g.DueAt.ShouldSetValue(temp.DueAt.IsValueSet())
    if temp.DueAt.Value() != nil {
        DueAtVal, err := time.Parse(time.RFC3339, (*temp.DueAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
        }
        g.DueAt.SetValue(&DueAtVal)
    }
    g.QrCode = temp.QrCode
    g.Line = temp.Line
    g.PdfPassword = temp.PdfPassword
    g.Pdf = temp.Pdf
    g.PaidAt.ShouldSetValue(temp.PaidAt.IsValueSet())
    if temp.PaidAt.Value() != nil {
        PaidAtVal, err := time.Parse(time.RFC3339, (*temp.PaidAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse paid_at as % s format.", time.RFC3339)
        }
        g.PaidAt.SetValue(&PaidAtVal)
    }
    g.PaidAmount = temp.PaidAmount
    g.Type = temp.Type
    g.CreditAt.ShouldSetValue(temp.CreditAt.IsValueSet())
    if temp.CreditAt.Value() != nil {
        CreditAtVal, err := time.Parse(time.RFC3339, (*temp.CreditAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse credit_at as % s format.", time.RFC3339)
        }
        g.CreditAt.SetValue(&CreditAtVal)
    }
    g.StatementDescriptor = temp.StatementDescriptor
    return nil
}

// url returns the url from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetUrl() *string {
    return g.Url
}

// barcode returns the barcode from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetBarcode() *string {
    return g.Barcode
}

// nosso_numero returns the nosso_numero from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetNossoNumero() *string {
    return g.NossoNumero
}

// bank returns the bank from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetBank() *string {
    return g.Bank
}

// document_number returns the document_number from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetDocumentNumber() *string {
    return g.DocumentNumber
}

// instructions returns the instructions from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetInstructions() *string {
    return g.Instructions
}

// billing_address returns the billing_address from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetBillingAddress() *GetBillingAddressResponse {
    return g.BillingAddress
}

// due_at returns the due_at from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetDueAt() Optional[time.Time] {
    return g.DueAt
}

// qr_code returns the qr_code from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetQrCode() *string {
    return g.QrCode
}

// line returns the line from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetLine() *string {
    return g.Line
}

// pdf_password returns the pdf_password from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetPdfPassword() *string {
    return g.PdfPassword
}

// pdf returns the pdf from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetPdf() *string {
    return g.Pdf
}

// paid_at returns the paid_at from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetPaidAt() Optional[time.Time] {
    return g.PaidAt
}

// paid_amount returns the paid_amount from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetPaidAmount() *string {
    return g.PaidAmount
}

// type returns the type from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetType() *string {
    return g.Type
}

// credit_at returns the credit_at from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetCreditAt() Optional[time.Time] {
    return g.CreditAt
}

// statement_descriptor returns the statement_descriptor from the GetBoletoTransactionResponse.
func (g *GetBoletoTransactionResponse) GetStatementDescriptor() *string {
    return g.StatementDescriptor
}

// UnmarshalGetBoletoTransactionResponse is a utility function used to unmarshal JSON into a GetBoletoTransactionResponseInterface.
// It takes a JSON input []byte and returns the corresponding GetBoletoTransactionResponseInterface and an error, if any.
func UnmarshalGetBoletoTransactionResponse(input []byte) (
    GetBoletoTransactionResponseInterface,
    error) {
    getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
    if err != nil {
        return nil, err
    }
    
    if getBoletoTransactionResponse, ok := getTransactionResponse.(GetBoletoTransactionResponseInterface); ok {
        return getBoletoTransactionResponse, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal getBoletoTransactionResponse %v", getTransactionResponse)
}

// ToGetBoletoTransactionResponseArray converts an array of GetBoletoTransactionResponseField to an array of GetBoletoTransactionResponseInterface.
func ToGetBoletoTransactionResponseArray(getBoletoTransactionResponse []GetBoletoTransactionResponseField) []GetBoletoTransactionResponseInterface {
    var items []GetBoletoTransactionResponseInterface
    for _, temp := range getBoletoTransactionResponse {
        items = append(items, temp.GetBoletoTransactionResponseInterface)
    }
    return items
}

// GetBoletoTransactionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetBoletoTransactionResponseField struct {
    GetBoletoTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetBoletoTransactionResponseField.
// It customizes the JSON unmarshaling process for GetBoletoTransactionResponseField objects.
func (g *GetBoletoTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetBoletoTransactionResponseInterface, err = UnmarshalGetBoletoTransactionResponse(input)
    return err
}

// GetBoletoTransactionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetBoletoTransactionResponseSlice struct {
    Slice []GetBoletoTransactionResponseInterface 
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetBoletoTransactionResponseSlice.
// It customizes the JSON unmarshaling process for GetBoletoTransactionResponseSlice objects.
func (ga *GetBoletoTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
    	ga.Slice = []GetBoletoTransactionResponseInterface{}
    	for _, getBoletoTransactionResponse := range temp {
    		if getBoletoTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetBoletoTransactionResponse
    		err := json.Unmarshal(getBoletoTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}
