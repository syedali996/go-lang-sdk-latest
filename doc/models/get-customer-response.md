
# Get Customer Response

Response object for getting a customer

## Structure

`GetCustomerResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*string` | Required | - |
| `Name` | `*string` | Required | - |
| `Email` | `*string` | Required | - |
| `Delinquent` | `*bool` | Required | - |
| `CreatedAt` | `*time.Time` | Required | - |
| `UpdatedAt` | `*time.Time` | Required | - |
| `Document` | `*string` | Required | - |
| `Type` | `*string` | Required | - |
| `FbAccessToken` | `*string` | Required | - |
| `Address` | [`*models.GetAddressResponse`](../../doc/models/get-address-response.md) | Required | - |
| `Metadata` | `map[string]string` | Required | - |
| `Phones` | [`*models.GetPhonesResponse`](../../doc/models/get-phones-response.md) | Required | - |
| `FbId` | `Optional[int64]` | Optional | - |
| `Code` | `*string` | Required | Código de referência do cliente no sistema da loja. Max: 52 caracteres |
| `DocumentType` | `*string` | Required | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "name": "name0",
  "email": "email6",
  "delinquent": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z",
  "document": "document6",
  "type": "type0",
  "fb_access_token": "fb_access_token4",
  "address": {
    "id": "id6",
    "street": "street6",
    "number": "number4",
    "complement": "complement2",
    "zip_code": "zip_code0",
    "neighborhood": "neighborhood2",
    "city": "city6",
    "state": "state2",
    "country": "country0",
    "status": "status8",
    "created_at": "2016-03-13T12:52:32.123Z",
    "updated_at": "2016-03-13T12:52:32.123Z",
    "customer": {
      "id": "id6",
      "name": "name6",
      "email": "email0",
      "delinquent": false,
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "document": "document0",
      "type": "type4",
      "fb_access_token": "fb_access_token0",
      "address": {},
      "metadata": {
        "key0": "metadata7",
        "key1": "metadata8"
      },
      "phones": {
        "home_phone": {
          "country_code": "country_code8",
          "number": "number6",
          "area_code": "area_code8"
        },
        "mobile_phone": {
          "country_code": "country_code8",
          "number": "number0",
          "area_code": "area_code2"
        }
      },
      "fb_id": 2,
      "code": "code4",
      "document_type": "document_type4"
    },
    "metadata": {
      "key0": "metadata3",
      "key1": "metadata2",
      "key2": "metadata1"
    },
    "line_1": "line_10",
    "line_2": "line_24",
    "deleted_at": "2016-03-13T12:52:32.123Z"
  },
  "metadata": {
    "key0": "metadata3",
    "key1": "metadata4",
    "key2": "metadata5"
  },
  "phones": {
    "home_phone": {
      "country_code": "country_code8",
      "number": "number0",
      "area_code": "area_code2"
    },
    "mobile_phone": {
      "country_code": "country_code8",
      "number": "number4",
      "area_code": "area_code8"
    }
  },
  "fb_id": 208,
  "code": "code8",
  "document_type": "document_type8"
}
```

