
# List Transfer Response

List of paginated transfer objects

## Structure

`ListTransferResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetTransferResponse`](../../doc/models/get-transfer-response.md) | Required | Transfers |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "amount": 121,
      "status": "status7",
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "bank_account": {
        "id": "id9",
        "holder_name": "holder_name5",
        "holder_type": "holder_type1",
        "bank": "bank7",
        "branch_number": "branch_number5",
        "branch_check_digit": "branch_check_digit5",
        "account_number": "account_number9",
        "account_check_digit": "account_check_digit5",
        "type": "type9",
        "status": "status1",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "deleted_at": "2016-03-13T12:52:32.123Z",
        "recipient": {
          "id": "id1",
          "name": "name1",
          "email": "email5",
          "document": "document5",
          "description": "description1",
          "type": "type1",
          "status": "status3",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "default_bank_account": {},
          "gateway_recipients": [
            {
              "gateway": "gateway5",
              "status": "status7",
              "pgid": "pgid1",
              "created_at": "created_at3",
              "updated_at": "updated_at1"
            },
            {
              "gateway": "gateway6",
              "status": "status8",
              "pgid": "pgid2",
              "created_at": "created_at4",
              "updated_at": "updated_at2"
            },
            {
              "gateway": "gateway7",
              "status": "status9",
              "pgid": "pgid3",
              "created_at": "created_at5",
              "updated_at": "updated_at3"
            }
          ],
          "metadata": {
            "key0": "metadata2",
            "key1": "metadata3",
            "key2": "metadata4"
          },
          "automatic_anticipation_settings": {
            "enabled": true,
            "type": "type5",
            "volume_percentage": 103,
            "delay": 137,
            "days": [
              113,
              114
            ]
          },
          "transfer_settings": {
            "transfer_enabled": true,
            "transfer_interval": "transfer_interval7",
            "transfer_day": 229
          },
          "code": "code9",
          "payment_mode": "payment_mode5"
        },
        "metadata": {
          "key0": "metadata0",
          "key1": "metadata1",
          "key2": "metadata2"
        },
        "pix_key": "pix_key3"
      },
      "metadata": {
        "key0": "metadata4",
        "key1": "metadata3"
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

