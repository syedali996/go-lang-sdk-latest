
# Create Usage Request

Request for creating a usage

## Structure

`CreateUsageRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Quantity` | `int` | Required | - |
| `Description` | `string` | Required | - |
| `UsedAt` | `time.Time` | Required | - |
| `Code` | `*string` | Optional | Identification code in the client system |
| `Group` | `*string` | Optional | identification group in the client system |
| `Amount` | `*int` | Optional | Field used in item scheme type 'Percent' |

## Example (as JSON)

```json
{
  "quantity": 68,
  "description": "description0",
  "used_at": "2016-03-13T12:52:32.123Z",
  "code": "code8",
  "group": "group8",
  "amount": 46
}
```

