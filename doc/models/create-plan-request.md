
# Create Plan Request

Request for creating a plan

## Structure

`CreatePlanRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | Plan's name |
| `Description` | `string` | Required | Description |
| `StatementDescriptor` | `string` | Required | Text that will be printed on the credit card's statement |
| `Items` | [`[]models.CreatePlanItemRequest`](../../doc/models/create-plan-item-request.md) | Required | Plan items |
| `Shippable` | `bool` | Required | Indicates if the plan is shippable |
| `PaymentMethods` | `[]string` | Required | Allowed payment methods for the plan |
| `Installments` | `[]int` | Required | Number of installments |
| `Currency` | `string` | Required | Currency |
| `Interval` | `string` | Required | Interval |
| `IntervalCount` | `int` | Required | Interval counts between two charges. For instance, if the interval is 'month' and count is 2, the customer will be charged once every two months. |
| `BillingDays` | `[]int` | Required | Allowed billings days for the subscription, in case the plan type is 'exact_day' |
| `BillingType` | `string` | Required | Billing type |
| `PricingScheme` | [`models.CreatePricingSchemeRequest`](../../doc/models/create-pricing-scheme-request.md) | Required | Plan's pricing scheme |
| `Metadata` | `map[string]string` | Required | Metadata |
| `MinimumPrice` | `*int` | Optional | Minimum price that will be charged |
| `Cycles` | `*int` | Optional | Number of cycles |
| `Quantity` | `*int` | Optional | Quantity |
| `TrialPeriodDays` | `*int` | Optional | Trial period, where the customer will not be charged. |

## Example (as JSON)

```json
{
  "name": "name0",
  "description": "description0",
  "statement_descriptor": "statement_descriptor0",
  "items": [
    {
      "name": "name7",
      "pricing_scheme": {
        "scheme_type": "scheme_type1",
        "price_brackets": [
          {
            "start_quantity": 60,
            "price": 2,
            "end_quantity": 68,
            "overage_price": 82
          },
          {
            "start_quantity": 61,
            "price": 1,
            "end_quantity": 69,
            "overage_price": 83
          },
          {
            "start_quantity": 62,
            "price": 0,
            "end_quantity": 70,
            "overage_price": 84
          }
        ],
        "price": 149,
        "minimum_price": 53,
        "percentage": 25.89
      },
      "id": "id7",
      "description": "description7",
      "cycles": 109,
      "quantity": 127
    }
  ],
  "shippable": false,
  "payment_methods": [
    "payment_methods5",
    "payment_methods6"
  ],
  "installments": [
    119,
    120,
    121
  ],
  "currency": "currency0",
  "interval": "interval2",
  "interval_count": 82,
  "billing_days": [
    143,
    144,
    145
  ],
  "billing_type": "billing_type6",
  "pricing_scheme": {
    "scheme_type": "scheme_type8",
    "price_brackets": [
      {
        "start_quantity": 119,
        "price": 57,
        "end_quantity": 127,
        "overage_price": 141
      },
      {
        "start_quantity": 120,
        "price": 58,
        "end_quantity": 128,
        "overage_price": 142
      },
      {
        "start_quantity": 121,
        "price": 59,
        "end_quantity": 129,
        "overage_price": 143
      }
    ],
    "price": 166,
    "minimum_price": 6,
    "percentage": 251.76
  },
  "metadata": {
    "key0": "metadata3",
    "key1": "metadata4",
    "key2": "metadata5"
  },
  "minimum_price": 176,
  "cycles": 168,
  "quantity": 68,
  "trial_period_days": 54
}
```

