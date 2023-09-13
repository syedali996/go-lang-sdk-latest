
# Create Card Request

Card data

## Structure

`CreateCardRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Number` | `string` | Required | Credit card number |
| `HolderName` | `string` | Required | Holder name, as written on the card |
| `ExpMonth` | `int` | Required | The expiration month |
| `ExpYear` | `int` | Required | The expiration year, that can be informed with 2 or 4 digits |
| `Cvv` | `string` | Required | The card's security code |
| `BillingAddress` | [`models.CreateAddressRequest`](../../doc/models/create-address-request.md) | Required | Card's billing address |
| `Brand` | `string` | Required | Card brand |
| `BillingAddressId` | `string` | Required | The address id for the billing address |
| `Metadata` | `map[string]string` | Required | Metadata |
| `Type` | `string` | Required | Card type<br>**Default**: `"credit"` |
| `Options` | [`models.CreateCardOptionsRequest`](../../doc/models/create-card-options-request.md) | Required | Options for creating the card |
| `HolderDocument` | `*string` | Optional | Document number for the card's holder |
| `PrivateLabel` | `bool` | Required | Indicates whether it is a private label card |
| `Label` | `string` | Required | - |
| `Id` | `*string` | Optional | Identifier |
| `Token` | `*string` | Optional | token identifier |

## Example (as JSON)

```json
{
  "number": "number2",
  "holder_name": "holder_name4",
  "exp_month": 42,
  "exp_year": 254,
  "cvv": "cvv2",
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
  "brand": "brand4",
  "billing_address_id": "billing_address_id6",
  "metadata": {
    "key0": "metadata3",
    "key1": "metadata4",
    "key2": "metadata5"
  },
  "type": "credit",
  "options": {
    "verify_card": false
  },
  "private_label": false,
  "label": "label0",
  "holder_document": "holder_document6",
  "id": "id0",
  "token": "token6"
}
```

