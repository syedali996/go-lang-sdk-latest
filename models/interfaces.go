package models

import (
    "time"
)

// Response object for listing transactions
type ListTransactionsResponseInterface interface {
    GetData() []GetTransactionResponseInterface
    GetPaging() *PagingResponse
}

// Response object for getting a charge
type GetChargeResponseInterface interface {
    GetId() *string
    GetCode() *string
    GetGatewayId() *string
    GetAmount() *int
    GetStatus() *string
    GetCurrency() *string
    GetPaymentMethod() *string
    GetDueAt() *time.Time
    GetCreatedAt() *time.Time
    GetUpdatedAt() *time.Time
    GetLastTransaction() GetTransactionResponseInterface
    GetInvoice() Optional[GetInvoiceResponse]
    GetOrder() Optional[GetOrderResponse]
    GetCustomer() Optional[GetCustomerResponse]
    GetMetadata() map[string]string
    GetPaidAt() Optional[time.Time]
    GetCanceledAt() Optional[time.Time]
    GetCanceledAmount() *int
    GetPaidAmount() *int
    GetInterestAndFinePaid() Optional[int]
    GetRecurrencyCycle() Optional[string]
}

// Response object for getting a bank transfer transaction
type GetBankTransferTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetUrl() string
    GetBankTid() string
    GetBank() string
    GetPaidAt() *time.Time
    GetPaidAmount() *int
}

// Response object for getting a safety pay transaction
type GetSafetyPayTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetUrl() *string
    GetBankTid() *string
    GetPaidAt() Optional[time.Time]
    GetPaidAmount() Optional[int]
}

// Response for voucher transactions
type GetVoucherTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetStatementDescriptor() *string
    GetAcquirerName() *string
    GetAcquirerAffiliationCode() *string
    GetAcquirerTid() *string
    GetAcquirerNsu() *string
    GetAcquirerAuthCode() *string
    GetAcquirerMessage() *string
    GetAcquirerReturnCode() *string
    GetOperationType() *string
    GetCard() *GetCardResponse
}

// Response object for getting a boleto transaction
type GetBoletoTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetUrl() *string
    GetBarcode() *string
    GetNossoNumero() *string
    GetBank() *string
    GetDocumentNumber() *string
    GetInstructions() *string
    GetBillingAddress() *GetBillingAddressResponse
    GetDueAt() Optional[time.Time]
    GetQrCode() *string
    GetLine() *string
    GetPdfPassword() *string
    GetPdf() *string
    GetPaidAt() Optional[time.Time]
    GetPaidAmount() *string
    GetType() *string
    GetCreditAt() Optional[time.Time]
    GetStatementDescriptor() *string
}

// Response object for getting a debit card transaction
type GetDebitCardTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetStatementDescriptor() *string
    GetAcquirerName() *string
    GetAcquirerAffiliationCode() *string
    GetAcquirerTid() *string
    GetAcquirerNsu() *string
    GetAcquirerAuthCode() *string
    GetOperationType() *string
    GetCard() *GetCardResponse
    GetAcquirerMessage() *string
    GetAcquirerReturnCode() *string
    GetMpi() *string
    GetEci() *string
    GetAuthenticationType() *string
    GetThreedAuthenticationUrl() *string
}

// Generic response object for getting a transaction.
type GetTransactionResponseInterface interface {
    GetGatewayId() *string
    GetAmount() *int
    GetStatus() *string
    GetSuccess() *bool
    GetCreatedAt() *time.Time
    GetUpdatedAt() *time.Time
    GetAttemptCount() *int
    GetMaxAttempts() *int
    GetSplits() []GetSplitResponse
    GetNextAttempt() Optional[time.Time]
    GetTransactionType() *string
    GetId() *string
    GetGatewayResponse() *GetGatewayResponseResponse
    GetAntifraudResponse() *GetAntifraudResponse
    GetMetadata() Optional[map[string]string]
    GetSplit() []GetSplitResponse
    GetInterest() Optional[GetInterestResponse]
    GetFine() Optional[GetFineResponse]
    GetMaxDaysToPayPastDue() Optional[int]
}

// Response object for getting a private label transaction
type GetPrivateLabelTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetStatementDescriptor() *string
    GetAcquirerName() *string
    GetAcquirerAffiliationCode() *string
    GetAcquirerTid() *string
    GetAcquirerNsu() *string
    GetAcquirerAuthCode() *string
    GetOperationType() *string
    GetCard() *GetCardResponse
    GetAcquirerMessage() *string
    GetAcquirerReturnCode() *string
    GetInstallments() Optional[int]
}

// Response object for getting a cash transaction
type GetCashTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetDescription() *string
}

// Response object for getting a credit card transaction
type GetCreditCardTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetStatementDescriptor() *string
    GetAcquirerName() string
    GetAcquirerAffiliationCode() *string
    GetAcquirerTid() string
    GetAcquirerNsu() string
    GetAcquirerAuthCode() *string
    GetOperationType() *string
    GetCard() *GetCardResponse
    GetAcquirerMessage() *string
    GetAcquirerReturnCode() *string
    GetInstallments() Optional[int]
    GetThreedAuthenticationUrl() *string
}

// Response object for listing charge transactions
type ListChargeTransactionsResponseInterface interface {
    GetData() []GetTransactionResponseInterface
    GetPaging() *PagingResponse
}

// Response object when getting a pix transaction
type GetPixTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetQrCode() *string
    GetQrCodeUrl() *string
    GetExpiresAt() *time.Time
    GetAdditionalInformation() []PixAdditionalInformation
    GetEndToEndId() *string
    GetPayer() *GetPixPayerResponse
}
