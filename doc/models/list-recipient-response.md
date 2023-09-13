
# List Recipient Response

Response for the listing recipient method

## Structure

`ListRecipientResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetRecipientResponse`](../../doc/models/get-recipient-response.md) | Required | Recipients |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "name": "name5",
      "email": "email9",
      "document": "document9",
      "description": "description5",
      "type": "type5",
      "status": "status7",
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "deleted_at": "2016-03-13T12:52:32.123Z",
      "default_bank_account": {
        "id": "id3",
        "holder_name": "holder_name9",
        "holder_type": "holder_type5",
        "bank": "bank1",
        "branch_number": "branch_number9",
        "branch_check_digit": "branch_check_digit9",
        "account_number": "account_number3",
        "account_check_digit": "account_check_digit9",
        "type": "type3",
        "status": "status5",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "deleted_at": "2016-03-13T12:52:32.123Z",
        "recipient": {},
        "metadata": {
          "key0": "metadata4",
          "key1": "metadata5"
        },
        "pix_key": "pix_key7"
      },
      "gateway_recipients": [
        {
          "gateway": "gateway9",
          "status": "status1",
          "pgid": "pgid5",
          "created_at": "created_at7",
          "updated_at": "updated_at5"
        }
      ],
      "metadata": {
        "key0": "metadata4",
        "key1": "metadata3"
      },
      "code": "code3",
      "payment_mode": "bank_transfer",
      "automatic_anticipation_settings": {
        "enabled": true,
        "type": "type9",
        "volume_percentage": 67,
        "delay": 101,
        "days": [
          77
        ]
      },
      "transfer_settings": {
        "transfer_enabled": true,
        "transfer_interval": "transfer_interval1",
        "transfer_day": 193
      }
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

