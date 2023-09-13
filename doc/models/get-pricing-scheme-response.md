
# Get Pricing Scheme Response

Response object for getting a pricing scheme

## Structure

`GetPricingSchemeResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Price` | `*int` | Required | - |
| `SchemeType` | `*string` | Required | - |
| `PriceBrackets` | [`[]models.GetPriceBracketResponse`](../../doc/models/get-price-bracket-response.md) | Required | - |
| `MinimumPrice` | `Optional[int]` | Optional | - |
| `Percentage` | `Optional[float64]` | Optional | percentual value used in pricing_scheme Percent |

## Example (as JSON)

```json
{
  "price": 16,
  "scheme_type": "scheme_type0",
  "price_brackets": [
    {
      "start_quantity": 193,
      "price": 125,
      "end_quantity": 201,
      "overage_price": 215
    }
  ],
  "minimum_price": 176,
  "percentage": 4.18
}
```

