
# Update Charge Payment Method Request

Request for updating the payment method of a charge

## Structure

`UpdateChargePaymentMethodRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `UpdateSubscription` | `bool` | Required | Indicates if the payment method from the subscription must also be updated |
| `PaymentMethod` | `string` | Required | The new payment method |
| `CreditCard` | [`models.CreateCreditCardPaymentRequest`](../../doc/models/create-credit-card-payment-request.md) | Required | Credit card data |
| `DebitCard` | [`models.CreateDebitCardPaymentRequest`](../../doc/models/create-debit-card-payment-request.md) | Required | Debit card data |
| `Boleto` | [`models.CreateBoletoPaymentRequest`](../../doc/models/create-boleto-payment-request.md) | Required | Boleto data |
| `Voucher` | [`models.CreateVoucherPaymentRequest`](../../doc/models/create-voucher-payment-request.md) | Required | Voucher data |
| `Cash` | [`models.CreateCashPaymentRequest`](../../doc/models/create-cash-payment-request.md) | Required | Cash data |
| `BankTransfer` | [`models.CreateBankTransferPaymentRequest`](../../doc/models/create-bank-transfer-payment-request.md) | Required | Bank Transfer data |
| `PrivateLabel` | [`models.CreatePrivateLabelPaymentRequest`](../../doc/models/create-private-label-payment-request.md) | Required | - |

## Example (as JSON)

```json
{
  "update_subscription": false,
  "payment_method": "payment_method0",
  "credit_card": {
    "installments": 1,
    "capture": true,
    "recurrency_cycle": "\"first\" or \"subsequent\"",
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
  "voucher": {
    "recurrency_cycle": "\"first\" or \"subsequent\"",
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
    }
  },
  "cash": {
    "description": "description0",
    "confirm": false
  },
  "bank_transfer": {
    "bank": "bank0",
    "retries": 236
  },
  "private_label": {
    "installments": 1,
    "capture": true,
    "recurrency_cycle": "\"first\" or \"subsequent\"",
    "statement_descriptor": "statement_descriptor0",
    "card": {
      "number": "number2",
      "holder_name": "holder_name0",
      "exp_month": 116,
      "exp_year": 180,
      "cvv": "cvv2",
      "billing_address": {
        "street": "street6",
        "number": "number6",
        "zip_code": "zip_code0",
        "neighborhood": "neighborhood2",
        "city": "city6",
        "state": "state2",
        "country": "country0",
        "complement": "complement8",
        "metadata": {
          "key0": "metadata7",
          "key1": "metadata8",
          "key2": "metadata9"
        },
        "line_1": "line_10",
        "line_2": "line_24"
      },
      "brand": "brand8",
      "billing_address_id": "billing_address_id0",
      "metadata": {
        "key0": "metadata1",
        "key1": "metadata0"
      },
      "type": "type6",
      "options": {
        "verify_card": false
      },
      "holder_document": "holder_document8",
      "private_label": false,
      "label": "label4",
      "id": "id4",
      "token": "token2"
    },
    "card_id": "card_id6",
    "card_token": "card_token0"
  }
}
```

