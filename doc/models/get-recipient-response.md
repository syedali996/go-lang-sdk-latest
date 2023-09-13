
# Get Recipient Response

Recipient response

## Structure

`GetRecipientResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*string` | Required | Id |
| `Name` | `*string` | Required | Name |
| `Email` | `*string` | Required | Email |
| `Document` | `*string` | Required | Document |
| `Description` | `*string` | Required | Description |
| `Type` | `*string` | Required | Type |
| `Status` | `*string` | Required | Status |
| `CreatedAt` | `*time.Time` | Required | Creation date |
| `UpdatedAt` | `*time.Time` | Required | Last update date |
| `DeletedAt` | `*time.Time` | Required | Deletion date |
| `DefaultBankAccount` | [`*models.GetBankAccountResponse`](../../doc/models/get-bank-account-response.md) | Required | Default bank account |
| `GatewayRecipients` | [`[]models.GetGatewayRecipientResponse`](../../doc/models/get-gateway-recipient-response.md) | Required | Info about the recipient on the gateway |
| `Metadata` | `map[string]string` | Required | Metadata |
| `AutomaticAnticipationSettings` | [`Optional[models.GetAutomaticAnticipationResponse]`](../../doc/models/get-automatic-anticipation-response.md) | Optional | - |
| `TransferSettings` | [`Optional[models.GetTransferSettingsResponse]`](../../doc/models/get-transfer-settings-response.md) | Optional | - |
| `Code` | `*string` | Required | Recipient code |
| `PaymentMode` | `*string` | Required | Payment mode<br>**Default**: `"bank_transfer"` |

## Example (as JSON)

```json
{
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
      "default_bank_account": {},
      "gateway_recipients": [
        {
          "gateway": "gateway4",
          "status": "status6",
          "pgid": "pgid0",
          "created_at": "created_at2",
          "updated_at": "updated_at0"
        },
        {
          "gateway": "gateway5",
          "status": "status7",
          "pgid": "pgid1",
          "created_at": "created_at3",
          "updated_at": "updated_at1"
        }
      ],
      "metadata": {
        "key0": "metadata3",
        "key1": "metadata4",
        "key2": "metadata5"
      },
      "automatic_anticipation_settings": {
        "enabled": false,
        "type": "type6",
        "volume_percentage": 30,
        "delay": 4,
        "days": [
          236
        ]
      },
      "transfer_settings": {
        "transfer_enabled": false,
        "transfer_interval": "transfer_interval4",
        "transfer_day": 196
      },
      "code": "code8",
      "payment_mode": "payment_mode6"
    },
    "metadata": {
      "key0": "metadata5",
      "key1": "metadata4",
      "key2": "metadata3"
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
    "key0": "metadata3",
    "key1": "metadata4",
    "key2": "metadata5"
  },
  "code": "code8",
  "payment_mode": "bank_transfer",
  "automatic_anticipation_settings": {
    "enabled": false,
    "type": "type6",
    "volume_percentage": 100,
    "delay": 190,
    "days": [
      166
    ]
  },
  "transfer_settings": {
    "transfer_enabled": false,
    "transfer_interval": "transfer_interval4",
    "transfer_day": 10
  }
}
```

