
# Get Plan Response

Response object for getting a plan

## Structure

`GetPlanResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*string` | Required | - |
| `Name` | `*string` | Required | - |
| `Description` | `*string` | Required | - |
| `Url` | `*string` | Required | - |
| `StatementDescriptor` | `*string` | Required | - |
| `Interval` | `*string` | Required | - |
| `IntervalCount` | `*int` | Required | - |
| `BillingType` | `*string` | Required | - |
| `PaymentMethods` | `[]string` | Required | - |
| `Installments` | `[]int` | Required | - |
| `Status` | `*string` | Required | - |
| `Currency` | `*string` | Required | - |
| `CreatedAt` | `*time.Time` | Required | - |
| `UpdatedAt` | `*time.Time` | Required | - |
| `Items` | [`[]models.GetPlanItemResponse`](../../doc/models/get-plan-item-response.md) | Required | - |
| `BillingDays` | `[]int` | Required | - |
| `Shippable` | `*bool` | Required | - |
| `Metadata` | `map[string]string` | Required | - |
| `TrialPeriodDays` | `Optional[int]` | Optional | - |
| `MinimumPrice` | `Optional[int]` | Optional | - |
| `DeletedAt` | `Optional[time.Time]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "name": "name0",
  "description": "description0",
  "url": "url4",
  "statement_descriptor": "statement_descriptor0",
  "interval": "interval2",
  "interval_count": 82,
  "billing_type": "billing_type6",
  "payment_methods": [
    "payment_methods5",
    "payment_methods6"
  ],
  "installments": [
    119,
    120,
    121
  ],
  "status": "status8",
  "currency": "currency0",
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z",
  "items": [
    {
      "id": "id7",
      "name": "name7",
      "status": "status1",
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "pricing_scheme": {
        "price": 149,
        "scheme_type": "scheme_type1",
        "price_brackets": [
          {
            "start_quantity": 60,
            "price": 2,
            "end_quantity": 68,
            "overage_price": 82
          }
        ],
        "minimum_price": 53,
        "percentage": 25.89
      },
      "description": "description7",
      "plan": {
        "id": "id7",
        "name": "name7",
        "description": "description3",
        "url": "url1",
        "statement_descriptor": "statement_descriptor7",
        "interval": "interval5",
        "interval_count": 89,
        "billing_type": "billing_type9",
        "payment_methods": [
          "payment_methods8",
          "payment_methods9"
        ],
        "installments": [
          126,
          127,
          128
        ],
        "status": "status1",
        "currency": "currency7",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "items": [
          {},
          {}
        ],
        "billing_days": [
          120,
          119,
          118
        ],
        "shippable": true,
        "metadata": {
          "key0": "metadata6",
          "key1": "metadata7",
          "key2": "metadata8"
        },
        "trial_period_days": 61,
        "minimum_price": 169,
        "deleted_at": "2016-03-13T12:52:32.123Z"
      },
      "quantity": 127,
      "cycles": 109,
      "deleted_at": "2016-03-13T12:52:32.123Z"
    }
  ],
  "billing_days": [
    143,
    144,
    145
  ],
  "shippable": false,
  "metadata": {
    "key0": "metadata3",
    "key1": "metadata4",
    "key2": "metadata5"
  },
  "trial_period_days": 54,
  "minimum_price": 176,
  "deleted_at": "2016-03-13T12:52:32.123Z"
}
```

