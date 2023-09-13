
# List Anticipation Response

Anticipations

## Structure

`ListAnticipationResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetAnticipationResponse`](../../doc/models/get-anticipation-response.md) | Required | Anticipations |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "requested_amount": 157,
      "approved_amount": 211,
      "recipient": {
        "id": "id7",
        "name": "name7",
        "email": "email1",
        "document": "document1",
        "description": "description7",
        "type": "type7",
        "status": "status9",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "deleted_at": "2016-03-13T12:52:32.123Z",
        "default_bank_account": {
          "id": "id5",
          "holder_name": "holder_name1",
          "holder_type": "holder_type7",
          "bank": "bank3",
          "branch_number": "branch_number1",
          "branch_check_digit": "branch_check_digit1",
          "account_number": "account_number5",
          "account_check_digit": "account_check_digit1",
          "type": "type5",
          "status": "status7",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "recipient": {},
          "metadata": {
            "key0": "metadata6",
            "key1": "metadata7"
          },
          "pix_key": "pix_key9"
        },
        "gateway_recipients": [
          {
            "gateway": "gateway1",
            "status": "status3",
            "pgid": "pgid7",
            "created_at": "created_at9",
            "updated_at": "updated_at7"
          },
          {
            "gateway": "gateway2",
            "status": "status4",
            "pgid": "pgid8",
            "created_at": "created_at0",
            "updated_at": "updated_at8"
          }
        ],
        "metadata": {
          "key0": "metadata2",
          "key1": "metadata1"
        },
        "automatic_anticipation_settings": {
          "enabled": true,
          "type": "type1",
          "volume_percentage": 239,
          "delay": 17,
          "days": [
            249
          ]
        },
        "transfer_settings": {
          "transfer_enabled": true,
          "transfer_interval": "transfer_interval3",
          "transfer_day": 109
        },
        "code": "code5",
        "payment_mode": "payment_mode1"
      },
      "pgid": "pgid1",
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "payment_date": "2016-03-13T12:52:32.123Z",
      "status": "status7",
      "timeframe": "timeframe7"
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

