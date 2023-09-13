
# List Customers Response

Response for listing the customers

## Structure

`ListCustomersResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetCustomerResponse`](../../doc/models/get-customer-response.md) | Required | The customer object |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "name": "name5",
      "email": "email9",
      "delinquent": true,
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "document": "document9",
      "type": "type5",
      "fb_access_token": "fb_access_token9",
      "address": {
        "id": "id1",
        "street": "street1",
        "number": "number9",
        "complement": "complement7",
        "zip_code": "zip_code5",
        "neighborhood": "neighborhood7",
        "city": "city1",
        "state": "state7",
        "country": "country5",
        "status": "status3",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "customer": {},
        "metadata": {
          "key0": "metadata2",
          "key1": "metadata3"
        },
        "line_1": "line_15",
        "line_2": "line_29",
        "deleted_at": "2016-03-13T12:52:32.123Z"
      },
      "metadata": {
        "key0": "metadata4",
        "key1": "metadata3"
      },
      "phones": {
        "home_phone": {
          "country_code": "country_code7",
          "number": "number5",
          "area_code": "area_code7"
        },
        "mobile_phone": {
          "country_code": "country_code7",
          "number": "number5",
          "area_code": "area_code7"
        }
      },
      "fb_id": 119,
      "code": "code3",
      "document_type": "document_type3"
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

