
# Get Withdraw Response

## Structure

`GetWithdrawResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*string` | Required | - |
| `GatewayId` | `*string` | Required | - |
| `Amount` | `*int` | Required | - |
| `Status` | `*string` | Required | - |
| `CreatedAt` | `*time.Time` | Required | - |
| `UpdatedAt` | `*time.Time` | Required | - |
| `Metadata` | `Optional[[]string]` | Optional | - |
| `Fee` | `Optional[int]` | Optional | - |
| `FundingDate` | `Optional[time.Time]` | Optional | - |
| `FundingEstimatedDate` | `Optional[time.Time]` | Optional | - |
| `Type` | `*string` | Required | - |
| `Source` | [`*models.GetWithdrawSourceResponse`](../../doc/models/get-withdraw-source-response.md) | Required | - |
| `Target` | [`*models.GetWithdrawTargetResponse`](../../doc/models/get-withdraw-target-response.md) | Required | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "gateway_id": "gateway_id0",
  "amount": 46,
  "status": "status8",
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z",
  "metadata": [
    "metadata3",
    "metadata4",
    "metadata5"
  ],
  "fee": 168,
  "funding_date": "2016-03-13T12:52:32.123Z",
  "funding_estimated_date": "2016-03-13T12:52:32.123Z",
  "type": "type0",
  "source": {
    "source_id": "source_id8",
    "type": "type6"
  },
  "target": {
    "target_id": "target_id2",
    "type": "type8"
  }
}
```

