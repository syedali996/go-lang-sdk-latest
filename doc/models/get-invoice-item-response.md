
# Get Invoice Item Response

Response object for getting an invoice item

## Structure

`GetInvoiceItemResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | `*int` | Required | - |
| `Description` | `*string` | Required | - |
| `PricingScheme` | [`*models.GetPricingSchemeResponse`](../../doc/models/get-pricing-scheme-response.md) | Required | - |
| `PriceBracket` | [`*models.GetPriceBracketResponse`](../../doc/models/get-price-bracket-response.md) | Required | - |
| `Quantity` | `Optional[int]` | Optional | - |
| `Name` | `Optional[string]` | Optional | - |
| `SubscriptionItemId` | `*string` | Required | Subscription Item Id |

## Example (as JSON)

```json
{
  "amount": 46,
  "description": "description0",
  "pricing_scheme": {
    "price": 166,
    "scheme_type": "scheme_type8",
    "price_brackets": [
      {
        "start_quantity": 119,
        "price": 57,
        "end_quantity": 127,
        "overage_price": 141
      }
    ],
    "minimum_price": 6,
    "percentage": 251.76
  },
  "price_bracket": {
    "start_quantity": 164,
    "price": 154,
    "end_quantity": 172,
    "overage_price": 186
  },
  "subscription_item_id": "subscription_item_id4",
  "quantity": 68,
  "name": "name0"
}
```

