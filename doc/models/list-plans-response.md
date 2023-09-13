
# List Plans Response

Response object for listing plans

## Structure

`ListPlansResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetPlanResponse`](../../doc/models/get-plan-response.md) | Required | The plan objects |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "name": "name5",
      "description": "description5",
      "url": "url9",
      "statement_descriptor": "statement_descriptor5",
      "interval": "interval3",
      "interval_count": 249,
      "billing_type": "billing_type9",
      "payment_methods": [
        "payment_methods0",
        "payment_methods1",
        "payment_methods2"
      ],
      "installments": [
        30,
        31,
        32
      ],
      "status": "status7",
      "currency": "currency5",
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "items": [
        {
          "id": "id2",
          "name": "name2",
          "status": "status4",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "pricing_scheme": {
            "price": 28,
            "scheme_type": "scheme_type4",
            "price_brackets": [
              {
                "start_quantity": 237,
                "price": 175,
                "end_quantity": 245,
                "overage_price": 3
              }
            ],
            "minimum_price": 124,
            "percentage": 132.62
          },
          "description": "description2",
          "plan": {},
          "quantity": 38,
          "cycles": 58,
          "deleted_at": "2016-03-13T12:52:32.123Z"
        }
      ],
      "billing_days": [
        250,
        251
      ],
      "shippable": true,
      "metadata": {
        "key0": "metadata4",
        "key1": "metadata3"
      },
      "trial_period_days": 221,
      "minimum_price": 247,
      "deleted_at": "2016-03-13T12:52:32.123Z"
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

