
# Get Pix Transaction Response

Response object when getting a pix transaction

## Structure

`GetPixTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `QrCode` | `*string` | Required | - |
| `QrCodeUrl` | `*string` | Required | - |
| `ExpiresAt` | `*time.Time` | Required | - |
| `AdditionalInformation` | [`[]models.PixAdditionalInformation`](../../doc/models/pix-additional-information.md) | Required | - |
| `EndToEndId` | `*string` | Required | - |
| `Payer` | [`*models.GetPixPayerResponse`](../../doc/models/get-pix-payer-response.md) | Required | - |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id6",
  "amount": 190,
  "status": "status4",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z",
  "attempt_count": 90,
  "max_attempts": 154,
  "splits": [
    {
      "type": "type2",
      "amount": 82,
      "recipient": {
        "id": "id0",
        "name": "name0",
        "email": "email6",
        "document": "document6",
        "description": "description0",
        "type": "type0",
        "status": "status8",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "deleted_at": "2016-03-13T12:52:32.123Z",
        "default_bank_account": {
          "id": "id8",
          "holder_name": "holder_name4",
          "holder_type": "holder_type0",
          "bank": "bank6",
          "branch_number": "branch_number4",
          "branch_check_digit": "branch_check_digit4",
          "account_number": "account_number8",
          "account_check_digit": "account_check_digit4",
          "type": "type2",
          "status": "status0",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "recipient": {},
          "metadata": {
            "key0": "metadata5"
          },
          "pix_key": "pix_key8"
        },
        "gateway_recipients": [
          {
            "gateway": "gateway4",
            "status": "status6",
            "pgid": "pgid0",
            "created_at": "created_at2",
            "updated_at": "updated_at0"
          }
        ],
        "metadata": {
          "key0": "metadata3"
        },
        "automatic_anticipation_settings": {
          "enabled": false,
          "type": "type6",
          "volume_percentage": 144,
          "delay": 146,
          "days": [
            122,
            123,
            124
          ]
        },
        "transfer_settings": {
          "transfer_enabled": false,
          "transfer_interval": "transfer_interval4",
          "transfer_day": 54
        },
        "code": "code8",
        "payment_mode": "payment_mode6"
      },
      "gateway_id": "gateway_id2",
      "options": {
        "liable": false,
        "charge_processing_fee": false,
        "charge_remainder_fee": "charge_remainder_fee2"
      },
      "id": "id8"
    }
  ],
  "id": "id4",
  "gateway_response": {
    "code": "code2",
    "errors": [
      {
        "message": "message9"
      }
    ]
  },
  "antifraud_response": {
    "status": "status4",
    "return_code": "return_code2",
    "return_message": "return_message0",
    "provider_name": "provider_name0",
    "score": "score2"
  },
  "split": [
    {
      "type": "type8",
      "amount": 136,
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
            "key0": "metadata9",
            "key1": "metadata8",
            "key2": "metadata7"
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
          },
          {
            "gateway": "gateway9",
            "status": "status1",
            "pgid": "pgid5",
            "created_at": "created_at7",
            "updated_at": "updated_at5"
          }
        ],
        "metadata": {
          "key0": "metadata9",
          "key1": "metadata0",
          "key2": "metadata1"
        },
        "automatic_anticipation_settings": {
          "enabled": false,
          "type": "type2",
          "volume_percentage": 2,
          "delay": 32,
          "days": [
            8
          ]
        },
        "transfer_settings": {
          "transfer_enabled": false,
          "transfer_interval": "transfer_interval0",
          "transfer_day": 168
        },
        "code": "code2",
        "payment_mode": "payment_mode2"
      },
      "gateway_id": "gateway_id8",
      "options": {
        "liable": false,
        "charge_processing_fee": false,
        "charge_remainder_fee": "charge_remainder_fee8"
      },
      "id": "id2"
    }
  ],
  "qr_code": "qr_code6",
  "qr_code_url": "qr_code_url2",
  "expires_at": "2016-03-13T12:52:32.123Z",
  "additional_information": [
    {
      "Name": "Name5",
      "Value": "Value7"
    }
  ],
  "end_to_end_id": "end_to_end_id0",
  "payer": {
    "name": "name8",
    "document": "document2",
    "document_type": "document_type6",
    "bank_account": {
      "bank_name": "bank_name8",
      "ispb": "ispb6",
      "branch_code": "branch_code0",
      "account_number": "account_number2"
    }
  },
  "next_attempt": "2016-03-13T12:52:32.123Z",
  "transaction_type": "pix",
  "metadata": {
    "key0": "metadata9",
    "key1": "metadata0",
    "key2": "metadata1"
  },
  "interest": {
    "days": 176,
    "type": "type4",
    "amount": 250
  },
  "fine": {
    "days": 114,
    "type": "type4",
    "amount": 188
  }
}
```

