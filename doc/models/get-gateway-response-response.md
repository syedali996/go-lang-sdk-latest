
# Get Gateway Response Response

The Transaction Gateway Response

## Structure

`GetGatewayResponseResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `*string` | Required | The error code |
| `Errors` | [`[]models.GetGatewayErrorResponse`](../../doc/models/get-gateway-error-response.md) | Required | The gateway response errors list |

## Example (as JSON)

```json
{
  "code": "code8",
  "errors": [
    {
      "message": "message5"
    }
  ]
}
```

