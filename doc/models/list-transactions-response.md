
# List Transactions Response

Response object for listing transactions

## Structure

`ListTransactionsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetTransactionResponseInterface`](../../doc/models/get-transaction-response.md) | Required | The transaction objects |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "gateway_id": "gateway_id5",
      "amount": 121,
      "status": "status7",
      "success": true,
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "attempt_count": 97,
      "max_attempts": 85,
      "splits": [
        {
          "type": "type7",
          "amount": 53,
          "recipient": {
            "id": "id9",
            "name": "name9",
            "email": "email7",
            "document": "document3",
            "description": "description9",
            "type": "type1",
            "status": "status1",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "default_bank_account": {
              "id": "id7",
              "holder_name": "holder_name3",
              "holder_type": "holder_type9",
              "bank": "bank5",
              "branch_number": "branch_number3",
              "branch_check_digit": "branch_check_digit3",
              "account_number": "account_number7",
              "account_check_digit": "account_check_digit3",
              "type": "type7",
              "status": "status9",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "recipient": {},
              "metadata": {
                "key0": "metadata2",
                "key1": "metadata1"
              },
              "pix_key": "pix_key1"
            },
            "gateway_recipients": [
              {
                "gateway": "gateway3",
                "status": "status5",
                "pgid": "pgid9",
                "created_at": "created_at1",
                "updated_at": "updated_at9"
              },
              {
                "gateway": "gateway4",
                "status": "status6",
                "pgid": "pgid0",
                "created_at": "created_at2",
                "updated_at": "updated_at0"
              }
            ],
            "metadata": {
              "key0": "metadata6",
              "key1": "metadata5",
              "key2": "metadata4"
            },
            "automatic_anticipation_settings": {
              "enabled": true,
              "type": "type7",
              "volume_percentage": 85,
              "delay": 205,
              "days": [
                181
              ]
            },
            "transfer_settings": {
              "transfer_enabled": true,
              "transfer_interval": "transfer_interval5",
              "transfer_day": 251
            },
            "code": "code7",
            "payment_mode": "payment_mode3"
          },
          "gateway_id": "gateway_id3",
          "options": {
            "liable": true,
            "charge_processing_fee": true,
            "charge_remainder_fee": "charge_remainder_fee3"
          },
          "id": "id7"
        }
      ],
      "id": "id5",
      "gateway_response": {
        "code": "code5",
        "errors": [
          {
            "message": "message2"
          }
        ]
      },
      "antifraud_response": {
        "status": "status5",
        "return_code": "return_code3",
        "return_message": "return_message1",
        "provider_name": "provider_name1",
        "score": "score3"
      },
      "split": [
        {
          "type": "type3",
          "amount": 67,
          "recipient": {
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
              },
              {
                "gateway": "gateway0",
                "status": "status2",
                "pgid": "pgid6",
                "created_at": "created_at8",
                "updated_at": "updated_at6"
              }
            ],
            "metadata": {
              "key0": "metadata4",
              "key1": "metadata3"
            },
            "automatic_anticipation_settings": {
              "enabled": true,
              "type": "type9",
              "volume_percentage": 185,
              "delay": 219,
              "days": [
                195
              ]
            },
            "transfer_settings": {
              "transfer_enabled": true,
              "transfer_interval": "transfer_interval1",
              "transfer_day": 55
            },
            "code": "code3",
            "payment_mode": "payment_mode9"
          },
          "gateway_id": "gateway_id3",
          "options": {
            "liable": true,
            "charge_processing_fee": true,
            "charge_remainder_fee": "charge_remainder_fee9"
          },
          "id": "id3"
        }
      ],
      "next_attempt": "2016-03-13T12:52:32.123Z",
      "transaction_type": "transaction",
      "metadata": {
        "key0": "metadata4",
        "key1": "metadata3"
      },
      "interest": {
        "days": 27,
        "type": "type5",
        "amount": 101
      },
      "fine": {
        "days": 45,
        "type": "type7",
        "amount": 119
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

