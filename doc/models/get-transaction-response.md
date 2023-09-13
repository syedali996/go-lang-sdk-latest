
# Get Transaction Response

Generic response object for getting a transaction.

## Structure

`GetTransactionResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `GatewayId` | `*string` | Required | Gateway transaction id |
| `Amount` | `*int` | Required | Amount in cents |
| `Status` | `*string` | Required | Transaction status |
| `Success` | `*bool` | Required | Indicates if the transaction ocurred successfuly |
| `CreatedAt` | `*time.Time` | Required | Creation date |
| `UpdatedAt` | `*time.Time` | Required | Last update date |
| `AttemptCount` | `*int` | Required | Number of attempts tried |
| `MaxAttempts` | `*int` | Required | Max attempts |
| `Splits` | [`[]models.GetSplitResponse`](../../doc/models/get-split-response.md) | Required | Splits |
| `NextAttempt` | `Optional[time.Time]` | Optional | Date and time of the next attempt |
| `TransactionType` | `*string` | Optional | - |
| `Id` | `*string` | Required | Código da transação |
| `GatewayResponse` | [`*models.GetGatewayResponseResponse`](../../doc/models/get-gateway-response-response.md) | Required | The Gateway Response |
| `AntifraudResponse` | [`*models.GetAntifraudResponse`](../../doc/models/get-antifraud-response.md) | Required | - |
| `Metadata` | `Optional[map[string]string]` | Optional | - |
| `Split` | [`[]models.GetSplitResponse`](../../doc/models/get-split-response.md) | Required | - |
| `Interest` | [`Optional[models.GetInterestResponse]`](../../doc/models/get-interest-response.md) | Optional | - |
| `Fine` | [`Optional[models.GetFineResponse]`](../../doc/models/get-fine-response.md) | Optional | - |
| `MaxDaysToPayPastDue` | `Optional[int]` | Optional | - |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id0",
  "amount": 46,
  "status": "status8",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z",
  "attempt_count": 70,
  "max_attempts": 174,
  "splits": [
    {
      "type": "type4",
      "amount": 62,
      "recipient": {
        "id": "id2",
        "name": "name2",
        "email": "email4",
        "document": "document4",
        "description": "description8",
        "type": "type8",
        "status": "status6",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "deleted_at": "2016-03-13T12:52:32.123Z",
        "default_bank_account": {
          "id": "id0",
          "holder_name": "holder_name6",
          "holder_type": "holder_type2",
          "bank": "bank8",
          "branch_number": "branch_number6",
          "branch_check_digit": "branch_check_digit6",
          "account_number": "account_number0",
          "account_check_digit": "account_check_digit6",
          "type": "type0",
          "status": "status2",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "recipient": {},
          "metadata": {
            "key0": "metadata7",
            "key1": "metadata6"
          },
          "pix_key": "pix_key6"
        },
        "gateway_recipients": [
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
          },
          {
            "gateway": "gateway8",
            "status": "status0",
            "pgid": "pgid4",
            "created_at": "created_at6",
            "updated_at": "updated_at4"
          }
        ],
        "metadata": {
          "key0": "metadata1",
          "key1": "metadata2"
        },
        "automatic_anticipation_settings": {
          "enabled": false,
          "type": "type4",
          "volume_percentage": 124,
          "delay": 166,
          "days": [
            142,
            143
          ]
        },
        "transfer_settings": {
          "transfer_enabled": false,
          "transfer_interval": "transfer_interval2",
          "transfer_day": 34
        },
        "code": "code0",
        "payment_mode": "payment_mode4"
      },
      "gateway_id": "gateway_id4",
      "options": {
        "liable": false,
        "charge_processing_fee": false,
        "charge_remainder_fee": "charge_remainder_fee4"
      },
      "id": "id6"
    }
  ],
  "id": "id0",
  "gateway_response": {
    "code": "code6",
    "errors": [
      {
        "message": "message3"
      }
    ]
  },
  "antifraud_response": {
    "status": "status0",
    "return_code": "return_code8",
    "return_message": "return_message4",
    "provider_name": "provider_name4",
    "score": "score8"
  },
  "split": [
    {
      "type": "type6",
      "amount": 28,
      "recipient": {
        "id": "id4",
        "name": "name4",
        "email": "email2",
        "document": "document2",
        "description": "description6",
        "type": "type6",
        "status": "status4",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "deleted_at": "2016-03-13T12:52:32.123Z",
        "default_bank_account": {
          "id": "id2",
          "holder_name": "holder_name8",
          "holder_type": "holder_type4",
          "bank": "bank0",
          "branch_number": "branch_number8",
          "branch_check_digit": "branch_check_digit8",
          "account_number": "account_number2",
          "account_check_digit": "account_check_digit8",
          "type": "type8",
          "status": "status4",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "recipient": {},
          "metadata": {
            "key0": "metadata9"
          },
          "pix_key": "pix_key4"
        },
        "gateway_recipients": [
          {
            "gateway": "gateway8",
            "status": "status0",
            "pgid": "pgid4",
            "created_at": "created_at6",
            "updated_at": "updated_at4"
          }
        ],
        "metadata": {
          "key0": "metadata9"
        },
        "automatic_anticipation_settings": {
          "enabled": false,
          "type": "type2",
          "volume_percentage": 34,
          "delay": 0,
          "days": [
            232,
            233,
            234
          ]
        },
        "transfer_settings": {
          "transfer_enabled": false,
          "transfer_interval": "transfer_interval0",
          "transfer_day": 200
        },
        "code": "code2",
        "payment_mode": "payment_mode2"
      },
      "gateway_id": "gateway_id6",
      "options": {
        "liable": false,
        "charge_processing_fee": false,
        "charge_remainder_fee": "charge_remainder_fee6"
      },
      "id": "id4"
    }
  ],
  "next_attempt": "2016-03-13T12:52:32.123Z",
  "transaction_type": "transaction",
  "metadata": {
    "key0": "metadata3",
    "key1": "metadata4",
    "key2": "metadata5"
  },
  "interest": {
    "days": 156,
    "type": "type0",
    "amount": 230
  },
  "fine": {
    "days": 138,
    "type": "type2",
    "amount": 212
  }
}
```

