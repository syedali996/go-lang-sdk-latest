
# Get Phones Response

## Structure

`GetPhonesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `HomePhone` | [`*models.GetPhoneResponse`](../../doc/models/get-phone-response.md) | Required | - |
| `MobilePhone` | [`*models.GetPhoneResponse`](../../doc/models/get-phone-response.md) | Required | - |

## Example (as JSON)

```json
{
  "home_phone": {
    "country_code": "country_code0",
    "number": "number2",
    "area_code": "area_code0"
  },
  "mobile_phone": {
    "country_code": "country_code0",
    "number": "number8",
    "area_code": "area_code0"
  }
}
```

