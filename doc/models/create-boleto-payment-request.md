
# Create Boleto Payment Request

Contains the settings for creating a boleto payment

## Structure

`CreateBoletoPaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Retries` | `int` | Required | Number of retries |
| `Bank` | `string` | Required | The bank code, containing three characters. The available codes are on the API specification |
| `Instructions` | `string` | Required | The instructions field that will be printed on the boleto. |
| `DueAt` | `*time.Time` | Optional | Boleto due date |
| `BillingAddress` | [`models.CreateAddressRequest`](../../doc/models/create-address-request.md) | Required | Card's billing address |
| `BillingAddressId` | `string` | Required | The address id for the billing address |
| `NossoNumero` | `*string` | Optional | Customer identification number with the bank |
| `DocumentNumber` | `string` | Required | Boleto identification |
| `StatementDescriptor` | `string` | Required | Soft Descriptor |
| `Interest` | [`*models.CreateInterestRequest`](../../doc/models/create-interest-request.md) | Optional | - |
| `Fine` | [`*models.CreateFineRequest`](../../doc/models/create-fine-request.md) | Optional | - |
| `MaxDaysToPayPastDue` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "retries": 230,
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
      "key0": "metadata5",
      "key1": "metadata6"
    },
    "line_1": "line_18",
    "line_2": "line_26"
  },
  "billing_address_id": "billing_address_id6",
  "nosso_numero": "nosso_numero0",
  "document_number": "document_number6",
  "statement_descriptor": "statement_descriptor0",
  "interest": {
    "days": 156,
    "type": "type0",
    "amount": 230
  },
  "fine": {
    "days": 138,
    "type": "type2",
    "amount": 212
  },
  "max_days_to_pay_past_due": 122
}
```

