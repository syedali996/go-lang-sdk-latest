
# Create Payment Request

Payment data

## Structure

`CreatePaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentMethod` | `string` | Required | Payment method |
| `CreditCard` | [`*models.CreateCreditCardPaymentRequest`](../../doc/models/create-credit-card-payment-request.md) | Optional | Settings for credit card payment |
| `DebitCard` | [`*models.CreateDebitCardPaymentRequest`](../../doc/models/create-debit-card-payment-request.md) | Optional | Settings for debit card payment |
| `Boleto` | [`*models.CreateBoletoPaymentRequest`](../../doc/models/create-boleto-payment-request.md) | Optional | Settings for boleto payment |
| `Currency` | `*string` | Optional | Currency. Must be informed using 3 characters |
| `Voucher` | [`*models.CreateVoucherPaymentRequest`](../../doc/models/create-voucher-payment-request.md) | Optional | Settings for voucher payment |
| `Split` | [`[]models.CreateSplitRequest`](../../doc/models/create-split-request.md) | Optional | Splits |
| `BankTransfer` | [`*models.CreateBankTransferPaymentRequest`](../../doc/models/create-bank-transfer-payment-request.md) | Optional | Settings for bank transfer payment |
| `GatewayAffiliationId` | `*string` | Optional | Gateway affiliation code |
| `Amount` | `*int` | Optional | The amount of the payment, in cents |
| `Checkout` | [`*models.CreateCheckoutPaymentRequest`](../../doc/models/create-checkout-payment-request.md) | Optional | Settings for checkout payment |
| `CustomerId` | `*string` | Optional | Customer Id |
| `Customer` | [`*models.CreateCustomerRequest`](../../doc/models/create-customer-request.md) | Optional | Customer |
| `Metadata` | `map[string]string` | Optional | Metadata |
| `Cash` | [`*models.CreateCashPaymentRequest`](../../doc/models/create-cash-payment-request.md) | Optional | Settings for cash payment |
| `PrivateLabel` | [`*models.CreatePrivateLabelPaymentRequest`](../../doc/models/create-private-label-payment-request.md) | Optional | Settings for private label payment |
| `Pix` | [`*models.CreatePixPaymentRequest`](../../doc/models/create-pix-payment-request.md) | Optional | Settings for pix payment |

## Example (as JSON)

```json
{
  "payment_method": "payment_method0",
  "credit_card": {
    "installments": 52,
    "statement_descriptor": "statement_descriptor8",
    "card": {
      "number": "number0",
      "holder_name": "holder_name8",
      "exp_month": 54,
      "exp_year": 242,
      "cvv": "cvv0",
      "billing_address": {
        "street": "street6",
        "number": "number6",
        "zip_code": "zip_code0",
        "neighborhood": "neighborhood2",
        "city": "city4",
        "state": "state8",
        "country": "country0",
        "complement": "complement8",
        "metadata": {
          "key0": "metadata7"
        },
        "line_1": "line_10",
        "line_2": "line_24"
      },
      "brand": "brand6",
      "billing_address_id": "billing_address_id8",
      "metadata": {
        "key0": "metadata1"
      },
      "type": "type8",
      "options": {
        "verify_card": false
      },
      "holder_document": "holder_document4",
      "private_label": false,
      "label": "label2",
      "id": "id2",
      "token": "token4"
    },
    "card_id": "card_id4",
    "card_token": "card_token2"
  },
  "debit_card": {
    "statement_descriptor": "statement_descriptor4",
    "card": {
      "number": "number0",
      "holder_name": "holder_name8",
      "exp_month": 104,
      "exp_year": 192,
      "cvv": "cvv0",
      "billing_address": {
        "street": "street6",
        "number": "number6",
        "zip_code": "zip_code0",
        "neighborhood": "neighborhood2",
        "city": "city4",
        "state": "state8",
        "country": "country0",
        "complement": "complement8",
        "metadata": {
          "key0": "metadata7",
          "key1": "metadata8"
        },
        "line_1": "line_10",
        "line_2": "line_24"
      },
      "brand": "brand6",
      "billing_address_id": "billing_address_id8",
      "metadata": {
        "key0": "metadata1",
        "key1": "metadata2",
        "key2": "metadata3"
      },
      "type": "type8",
      "options": {
        "verify_card": false
      },
      "holder_document": "holder_document4",
      "private_label": false,
      "label": "label2",
      "id": "id2",
      "token": "token4"
    },
    "card_id": "card_id0",
    "card_token": "card_token6",
    "recurrence": false
  },
  "boleto": {
    "retries": 226,
    "bank": "bank8",
    "instructions": "instructions2",
    "due_at": "2016-03-13T12:52:32.123Z",
    "billing_address": {
      "street": "street8",
      "number": "number4",
      "zip_code": "zip_code2",
      "neighborhood": "neighborhood4",
      "city": "city2",
      "state": "state6",
      "country": "country2",
      "complement": "complement6",
      "metadata": {
        "key0": "metadata5"
      },
      "line_1": "line_18",
      "line_2": "line_26"
    },
    "billing_address_id": "billing_address_id6",
    "nosso_numero": "nosso_numero0",
    "document_number": "document_number6",
    "statement_descriptor": "statement_descriptor0",
    "interest": {
      "days": 160,
      "type": "type0",
      "amount": 234
    },
    "fine": {
      "days": 130,
      "type": "type8",
      "amount": 52
    },
    "max_days_to_pay_past_due": 118
  },
  "currency": "currency0",
  "voucher": {
    "statement_descriptor": "statement_descriptor2",
    "card_id": "card_id8",
    "card_token": "card_token8",
    "Card": {
      "number": "number0",
      "holder_name": "holder_name8",
      "exp_month": 198,
      "exp_year": 238,
      "cvv": "cvv0",
      "billing_address": {
        "street": "street4",
        "number": "number2",
        "zip_code": "zip_code8",
        "neighborhood": "neighborhood0",
        "city": "city4",
        "state": "state0",
        "country": "country8",
        "complement": "complement0",
        "metadata": {
          "key0": "metadata5",
          "key1": "metadata4",
          "key2": "metadata3"
        },
        "line_1": "line_18",
        "line_2": "line_22"
      },
      "brand": "brand6",
      "billing_address_id": "billing_address_id8",
      "metadata": {
        "key0": "metadata7"
      },
      "type": "type2",
      "options": {
        "verify_card": false
      },
      "holder_document": "holder_document6",
      "private_label": false,
      "label": "label2",
      "id": "id2",
      "token": "token6"
    },
    "recurrency_cycle": "recurrency_cycle6"
  }
}
```

