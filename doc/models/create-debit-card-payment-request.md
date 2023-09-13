
# Create Debit Card Payment Request

The settings for creating a debit card payment

## Structure

`CreateDebitCardPaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StatementDescriptor` | `*string` | Optional | The text that will be shown on the debit card's statement |
| `Card` | [`*models.CreateCardRequest`](../../doc/models/create-card-request.md) | Optional | Debit card data |
| `CardId` | `*string` | Optional | The debit card id |
| `CardToken` | `*string` | Optional | The debit card token |
| `Recurrence` | `*bool` | Optional | Indicates a recurrence |
| `Authentication` | [`*models.CreatePaymentAuthenticationRequest`](../../doc/models/create-payment-authentication-request.md) | Optional | The payment authentication request |
| `Token` | [`*models.CreateCardPaymentContactlessRequest`](../../doc/models/create-card-payment-contactless-request.md) | Optional | The Debit card payment token request |

## Example (as JSON)

```json
{
  "statement_descriptor": "statement_descriptor0",
  "card": {
    "number": "number6",
    "holder_name": "holder_name2",
    "exp_month": 228,
    "exp_year": 68,
    "cvv": "cvv4",
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
        "key0": "metadata5",
        "key1": "metadata6"
      },
      "line_1": "line_18",
      "line_2": "line_26"
    },
    "brand": "brand0",
    "billing_address_id": "billing_address_id2",
    "metadata": {
      "key0": "metadata7"
    },
    "type": "type4",
    "options": {
      "verify_card": false
    },
    "holder_document": "holder_document0",
    "private_label": false,
    "label": "label6",
    "id": "id6",
    "token": "token0"
  },
  "card_id": "card_id4",
  "card_token": "card_token0",
  "recurrence": false
}
```

