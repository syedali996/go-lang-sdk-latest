
# Get Plan Item Response

Response object for getting a plan item

## Structure

`GetPlanItemResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*string` | Required | - |
| `Name` | `*string` | Required | - |
| `Status` | `*string` | Required | - |
| `CreatedAt` | `*time.Time` | Required | - |
| `UpdatedAt` | `*time.Time` | Required | - |
| `PricingScheme` | [`*models.GetPricingSchemeResponse`](../../doc/models/get-pricing-scheme-response.md) | Required | - |
| `Description` | `*string` | Required | - |
| `Plan` | [`*models.GetPlanResponse`](../../doc/models/get-plan-response.md) | Required | - |
| `Quantity` | `Optional[int]` | Optional | - |
| `Cycles` | `Optional[int]` | Optional | - |
| `DeletedAt` | `Optional[time.Time]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "name": "name0",
  "status": "status8",
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z",
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
  "description": "description0",
  "plan": {
    "id": "id4",
    "name": "name4",
    "description": "description4",
    "url": "url8",
    "statement_descriptor": "statement_descriptor4",
    "interval": "interval2",
    "interval_count": 148,
    "billing_type": "billing_type2",
    "payment_methods": [
      "payment_methods9",
      "payment_methods8"
    ],
    "installments": [
      185,
      186,
      187
    ],
    "status": "status6",
    "currency": "currency4",
    "created_at": "2016-03-13T12:52:32.123Z",
    "updated_at": "2016-03-13T12:52:32.123Z",
    "items": [
      {
        "id": "id1",
        "name": "name1",
        "status": "status3",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "pricing_scheme": {
          "price": 73,
          "scheme_type": "scheme_type3",
          "price_brackets": [
            {
              "start_quantity": 136,
              "price": 182,
              "end_quantity": 144,
              "overage_price": 158
            }
          ],
          "minimum_price": 233,
          "percentage": 42.01
        },
        "description": "description1",
        "plan": {},
        "quantity": 193,
        "cycles": 213,
        "deleted_at": "2016-03-13T12:52:32.123Z"
      }
    ],
    "billing_days": [
      149,
      150
    ],
    "shippable": false,
    "metadata": {
      "key0": "metadata9",
      "key1": "metadata0",
      "key2": "metadata1"
    },
    "trial_period_days": 120,
    "minimum_price": 110,
    "deleted_at": "2016-03-13T12:52:32.123Z"
  },
  "quantity": 68,
  "cycles": 168,
  "deleted_at": "2016-03-13T12:52:32.123Z"
}
```

