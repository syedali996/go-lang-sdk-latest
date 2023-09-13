
# List Cards Response

Response object for listing cards

## Structure

`ListCardsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetCardResponse`](../../doc/models/get-card-response.md) | Required | The card objects |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "last_four_digits": "last_four_digits1",
      "brand": "brand9",
      "holder_name": "holder_name1",
      "exp_month": 125,
      "exp_year": 165,
      "status": "status7",
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "billing_address": {
        "street": "street7",
        "number": "number5",
        "zip_code": "zip_code1",
        "neighborhood": "neighborhood3",
        "city": "city7",
        "state": "state3",
        "country": "country1",
        "complement": "complement3",
        "line_1": "line_11",
        "line_2": "line_25"
      },
      "customer": {
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
            "key1": "metadata3",
            "key2": "metadata4"
          },
          "line_1": "line_15",
          "line_2": "line_29",
          "deleted_at": "2016-03-13T12:52:32.123Z"
        },
        "metadata": {
          "key0": "metadata4",
          "key1": "metadata3",
          "key2": "metadata2"
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
        "fb_id": 153,
        "code": "code3",
        "document_type": "document_type3"
      },
      "metadata": {
        "key0": "metadata4",
        "key1": "metadata3"
      },
      "type": "type5",
      "holder_document": "holder_document9",
      "deleted_at": "2016-03-13T12:52:32.123Z",
      "first_six_digits": "first_six_digits5",
      "label": "label5"
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

