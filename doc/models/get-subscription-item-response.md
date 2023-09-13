
# Get Subscription Item Response

## Structure

`GetSubscriptionItemResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*string` | Required | - |
| `Description` | `*string` | Required | - |
| `Status` | `*string` | Required | - |
| `CreatedAt` | `*time.Time` | Required | - |
| `UpdatedAt` | `*time.Time` | Required | - |
| `PricingScheme` | [`*models.GetPricingSchemeResponse`](../../doc/models/get-pricing-scheme-response.md) | Required | - |
| `Discounts` | [`[]models.GetDiscountResponse`](../../doc/models/get-discount-response.md) | Required | - |
| `Increments` | [`[]models.GetIncrementResponse`](../../doc/models/get-increment-response.md) | Required | - |
| `Subscription` | [`*models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md) | Required | - |
| `Name` | `*string` | Required | Item name |
| `Quantity` | `Optional[int]` | Optional | - |
| `Cycles` | `Optional[int]` | Optional | - |
| `DeletedAt` | `Optional[time.Time]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "description": "description0",
  "status": "status8",
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z",
  "pricing_scheme": {
    "price": 166,
    "scheme_type": "scheme_type8",
    "price_brackets": [
      {
        "start_quantity": 119,
        "price": 57,
        "end_quantity": 127,
        "overage_price": 141
      }
    ],
    "minimum_price": 6,
    "percentage": 251.76
  },
  "discounts": [
    {
      "id": "id1",
      "value": 10.23,
      "discount_type": "discount_type9",
      "status": "status3",
      "created_at": "2016-03-13T12:52:32.123Z",
      "cycles": 233,
      "deleted_at": "2016-03-13T12:52:32.123Z",
      "description": "description1",
      "subscription": {
        "id": "id3",
        "code": "code1",
        "start_at": "2016-03-13T12:52:32.123Z",
        "interval": "interval9",
        "interval_count": 43,
        "billing_type": "billing_type3",
        "current_cycle": {
          "start_at": "2016-03-13T12:52:32.123Z",
          "end_at": "2016-03-13T12:52:32.123Z",
          "id": "id1",
          "billing_at": "2016-03-13T12:52:32.123Z",
          "subscription": {},
          "status": "status7",
          "duration": 27,
          "created_at": "created_at9",
          "updated_at": "updated_at7",
          "cycle": 223
        },
        "payment_method": "payment_method7",
        "currency": "currency7",
        "installments": 211,
        "status": "status5",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "customer": {
          "id": "id7",
          "name": "name7",
          "email": "email9",
          "delinquent": true,
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "document": "document9",
          "type": "type3",
          "fb_access_token": "fb_access_token1",
          "address": {
            "id": "id3",
            "street": "street3",
            "number": "number1",
            "complement": "complement9",
            "zip_code": "zip_code7",
            "neighborhood": "neighborhood9",
            "city": "city3",
            "state": "state9",
            "country": "country7",
            "status": "status5",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "customer": {},
            "metadata": {
              "key0": "metadata0"
            },
            "line_1": "line_13",
            "line_2": "line_21",
            "deleted_at": "2016-03-13T12:52:32.123Z"
          },
          "metadata": {
            "key0": "metadata6"
          },
          "phones": {
            "home_phone": {
              "country_code": "country_code1",
              "number": "number3",
              "area_code": "area_code9"
            },
            "mobile_phone": {
              "country_code": "country_code1",
              "number": "number1",
              "area_code": "area_code1"
            }
          },
          "fb_id": 213,
          "code": "code5",
          "document_type": "document_type5"
        },
        "card": {
          "id": "id3",
          "last_four_digits": "last_four_digits9",
          "brand": "brand7",
          "holder_name": "holder_name9",
          "exp_month": 189,
          "exp_year": 107,
          "status": "status5",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "billing_address": {
            "street": "street5",
            "number": "number7",
            "zip_code": "zip_code9",
            "neighborhood": "neighborhood1",
            "city": "city5",
            "state": "state9",
            "country": "country9",
            "complement": "complement9",
            "line_1": "line_11",
            "line_2": "line_23"
          },
          "customer": {
            "id": "id3",
            "name": "name3",
            "email": "email3",
            "delinquent": true,
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "document": "document3",
            "type": "type7",
            "fb_access_token": "fb_access_token7",
            "address": {
              "id": "id9",
              "street": "street9",
              "number": "number7",
              "complement": "complement5",
              "zip_code": "zip_code3",
              "neighborhood": "neighborhood5",
              "city": "city9",
              "state": "state5",
              "country": "country3",
              "status": "status1",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "customer": {},
              "metadata": {
                "key0": "metadata6",
                "key1": "metadata5",
                "key2": "metadata4"
              },
              "line_1": "line_17",
              "line_2": "line_27",
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            "metadata": {
              "key0": "metadata0",
              "key1": "metadata1",
              "key2": "metadata2"
            },
            "phones": {
              "home_phone": {
                "country_code": "country_code5",
                "number": "number7",
                "area_code": "area_code5"
              },
              "mobile_phone": {
                "country_code": "country_code5",
                "number": "number7",
                "area_code": "area_code5"
              }
            },
            "fb_id": 95,
            "code": "code1",
            "document_type": "document_type1"
          },
          "metadata": {
            "key0": "metadata0"
          },
          "type": "type7",
          "holder_document": "holder_document3",
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "first_six_digits": "first_six_digits3",
          "label": "label3"
        },
        "items": [
          {
            "id": "id0",
            "description": "description0",
            "status": "status2",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "pricing_scheme": {
              "price": 110,
              "scheme_type": "scheme_type8",
              "price_brackets": [
                {
                  "start_quantity": 99,
                  "price": 219,
                  "end_quantity": 107,
                  "overage_price": 121
                },
                {
                  "start_quantity": 100,
                  "price": 218,
                  "end_quantity": 108,
                  "overage_price": 122
                },
                {
                  "start_quantity": 101,
                  "price": 217,
                  "end_quantity": 109,
                  "overage_price": 123
                }
              ],
              "minimum_price": 14,
              "percentage": 213.16
            },
            "discounts": [
              {},
              {},
              {}
            ],
            "increments": [
              {
                "id": "id1",
                "value": 86.93,
                "increment_type": "increment_type3",
                "status": "status7",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 243,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description9",
                "subscription": {},
                "subscription_item": {}
              }
            ],
            "subscription": {},
            "name": "name0",
            "quantity": 88,
            "cycles": 148,
            "deleted_at": "2016-03-13T12:52:32.123Z"
          },
          {
            "id": "id1",
            "description": "description1",
            "status": "status3",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "pricing_scheme": {
              "price": 111,
              "scheme_type": "scheme_type7",
              "price_brackets": [
                {
                  "start_quantity": 98,
                  "price": 220,
                  "end_quantity": 106,
                  "overage_price": 120
                },
                {
                  "start_quantity": 99,
                  "price": 219,
                  "end_quantity": 107,
                  "overage_price": 121
                }
              ],
              "minimum_price": 15,
              "percentage": 213.15
            },
            "discounts": [
              {}
            ],
            "increments": [
              {
                "id": "id0",
                "value": 86.92,
                "increment_type": "increment_type2",
                "status": "status8",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 244,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description0",
                "subscription": {},
                "subscription_item": {}
              },
              {
                "id": "id9",
                "value": 86.91,
                "increment_type": "increment_type1",
                "status": "status9",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 245,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description1",
                "subscription": {},
                "subscription_item": {}
              }
            ],
            "subscription": {},
            "name": "name1",
            "quantity": 89,
            "cycles": 147,
            "deleted_at": "2016-03-13T12:52:32.123Z"
          }
        ],
        "statement_descriptor": "statement_descriptor3",
        "metadata": {
          "key0": "metadata0",
          "key1": "metadata1",
          "key2": "metadata2"
        },
        "setup": {
          "id": "id7",
          "description": "description3",
          "amount": 21,
          "status": "status1"
        },
        "gateway_affiliation_id": "gateway_affiliation_id9",
        "next_billing_at": "2016-03-13T12:52:32.123Z",
        "billing_day": 73,
        "minimum_price": 215,
        "increments": [
          {
            "id": "id4",
            "value": 55.46,
            "increment_type": "increment_type6",
            "status": "status4",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 62,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description6",
            "subscription": {},
            "subscription_item": {
              "id": "id0",
              "description": "description0",
              "status": "status8",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "pricing_scheme": {
                "price": 184,
                "scheme_type": "scheme_type8",
                "price_brackets": [
                  {
                    "start_quantity": 25,
                    "price": 219,
                    "end_quantity": 33,
                    "overage_price": 47
                  }
                ],
                "minimum_price": 168,
                "percentage": 171.46
              },
              "discounts": [
                {},
                {}
              ],
              "increments": [
                {},
                {}
              ],
              "subscription": {},
              "name": "name0",
              "quantity": 162,
              "cycles": 74,
              "deleted_at": "2016-03-13T12:52:32.123Z"
            }
          }
        ],
        "split": {
          "enabled": true,
          "rules": [
            {
              "type": "type7",
              "amount": 141,
              "recipient": {
                "id": "id5",
                "name": "name5",
                "email": "email1",
                "document": "document1",
                "description": "description5",
                "type": "type5",
                "status": "status3",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "default_bank_account": {
                  "id": "id3",
                  "holder_name": "holder_name9",
                  "holder_type": "holder_type5",
                  "bank": "bank1",
                  "branch_number": "branch_number9",
                  "branch_check_digit": "branch_check_digit9",
                  "account_number": "account_number3",
                  "account_check_digit": "account_check_digit9",
                  "type": "type7",
                  "status": "status5",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "recipient": {},
                  "metadata": {
                    "key0": "metadata0"
                  },
                  "pix_key": "pix_key3"
                },
                "gateway_recipients": [
                  {
                    "gateway": "gateway9",
                    "status": "status1",
                    "pgid": "pgid5",
                    "created_at": "created_at7",
                    "updated_at": "updated_at5"
                  }
                ],
                "metadata": {
                  "key0": "metadata8"
                },
                "automatic_anticipation_settings": {
                  "enabled": true,
                  "type": "type1",
                  "volume_percentage": 177,
                  "delay": 113,
                  "days": [
                    89,
                    90,
                    91
                  ]
                },
                "transfer_settings": {
                  "transfer_enabled": true,
                  "transfer_interval": "transfer_interval9",
                  "transfer_day": 87
                },
                "code": "code3",
                "payment_mode": "payment_mode1"
              },
              "gateway_id": "gateway_id7",
              "options": {
                "liable": true,
                "charge_processing_fee": true,
                "charge_remainder_fee": "charge_remainder_fee7"
              },
              "id": "id3"
            },
            {
              "type": "type6",
              "amount": 140,
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
                    "key1": "metadata8"
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
                  },
                  {
                    "gateway": "gateway0",
                    "status": "status2",
                    "pgid": "pgid6",
                    "created_at": "created_at8",
                    "updated_at": "updated_at6"
                  }
                ],
                "metadata": {
                  "key0": "metadata9",
                  "key1": "metadata0"
                },
                "automatic_anticipation_settings": {
                  "enabled": false,
                  "type": "type2",
                  "volume_percentage": 178,
                  "delay": 112,
                  "days": [
                    88,
                    89
                  ]
                },
                "transfer_settings": {
                  "transfer_enabled": false,
                  "transfer_interval": "transfer_interval0",
                  "transfer_day": 88
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
          ]
        }
      },
      "subscription_item": {
        "id": "id7",
        "description": "description7",
        "status": "status9",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "pricing_scheme": {
          "price": 19,
          "scheme_type": "scheme_type9",
          "price_brackets": [
            {
              "start_quantity": 190,
              "price": 128,
              "end_quantity": 198,
              "overage_price": 212
            },
            {
              "start_quantity": 191,
              "price": 127,
              "end_quantity": 199,
              "overage_price": 213
            }
          ],
          "minimum_price": 179,
          "percentage": 47.67
        },
        "discounts": [
          {},
          {}
        ],
        "increments": [
          {
            "id": "id6",
            "value": 156.58,
            "increment_type": "increment_type8",
            "status": "status2",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 190,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description4",
            "subscription": {
              "id": "id8",
              "code": "code6",
              "start_at": "2016-03-13T12:52:32.123Z",
              "interval": "interval6",
              "interval_count": 0,
              "billing_type": "billing_type2",
              "current_cycle": {
                "start_at": "2016-03-13T12:52:32.123Z",
                "end_at": "2016-03-13T12:52:32.123Z",
                "id": "id6",
                "billing_at": "2016-03-13T12:52:32.123Z",
                "subscription": {},
                "status": "status8",
                "duration": 240,
                "created_at": "created_at4",
                "updated_at": "updated_at2",
                "cycle": 180
              },
              "payment_method": "payment_method2",
              "currency": "currency8",
              "installments": 168,
              "status": "status0",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "customer": {
                "id": "id8",
                "name": "name8",
                "email": "email8",
                "delinquent": false,
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "document": "document2",
                "type": "type2",
                "fb_access_token": "fb_access_token2",
                "address": {
                  "id": "id4",
                  "street": "street4",
                  "number": "number2",
                  "complement": "complement0",
                  "zip_code": "zip_code8",
                  "neighborhood": "neighborhood0",
                  "city": "city4",
                  "state": "state0",
                  "country": "country8",
                  "status": "status6",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "customer": {},
                  "metadata": {
                    "key0": "metadata5"
                  },
                  "line_1": "line_18",
                  "line_2": "line_22",
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                },
                "metadata": {
                  "key0": "metadata5"
                },
                "phones": {
                  "home_phone": {
                    "country_code": "country_code0",
                    "number": "number8",
                    "area_code": "area_code0"
                  },
                  "mobile_phone": {
                    "country_code": "country_code0",
                    "number": "number2",
                    "area_code": "area_code0"
                  }
                },
                "fb_id": 160,
                "code": "code6",
                "document_type": "document_type6"
              },
              "card": {
                "id": "id2",
                "last_four_digits": "last_four_digits8",
                "brand": "brand6",
                "holder_name": "holder_name8",
                "exp_month": 194,
                "exp_year": 102,
                "status": "status6",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "billing_address": {
                  "street": "street4",
                  "number": "number8",
                  "zip_code": "zip_code8",
                  "neighborhood": "neighborhood0",
                  "city": "city6",
                  "state": "state0",
                  "country": "country8",
                  "complement": "complement0",
                  "line_1": "line_12",
                  "line_2": "line_22"
                },
                "customer": {
                  "id": "id2",
                  "name": "name2",
                  "email": "email4",
                  "delinquent": false,
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "document": "document4",
                  "type": "type8",
                  "fb_access_token": "fb_access_token6",
                  "address": {
                    "id": "id8",
                    "street": "street8",
                    "number": "number6",
                    "complement": "complement4",
                    "zip_code": "zip_code2",
                    "neighborhood": "neighborhood4",
                    "city": "city8",
                    "state": "state4",
                    "country": "country2",
                    "status": "status0",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "updated_at": "2016-03-13T12:52:32.123Z",
                    "customer": {},
                    "metadata": {
                      "key0": "metadata5",
                      "key1": "metadata4",
                      "key2": "metadata3"
                    },
                    "line_1": "line_12",
                    "line_2": "line_26",
                    "deleted_at": "2016-03-13T12:52:32.123Z"
                  },
                  "metadata": {
                    "key0": "metadata1",
                    "key1": "metadata2",
                    "key2": "metadata3"
                  },
                  "phones": {
                    "home_phone": {
                      "country_code": "country_code4",
                      "number": "number8",
                      "area_code": "area_code4"
                    },
                    "mobile_phone": {
                      "country_code": "country_code6",
                      "number": "number6",
                      "area_code": "area_code6"
                    }
                  },
                  "fb_id": 90,
                  "code": "code0",
                  "document_type": "document_type0"
                },
                "metadata": {
                  "key0": "metadata1"
                },
                "type": "type8",
                "holder_document": "holder_document4",
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "first_six_digits": "first_six_digits2",
                "label": "label2"
              },
              "items": [
                {},
                {},
                {}
              ],
              "statement_descriptor": "statement_descriptor8",
              "metadata": {
                "key0": "metadata5",
                "key1": "metadata4"
              },
              "setup": {
                "id": "id2",
                "description": "description2",
                "amount": 192,
                "status": "status4"
              },
              "gateway_affiliation_id": "gateway_affiliation_id4",
              "next_billing_at": "2016-03-13T12:52:32.123Z",
              "billing_day": 140,
              "minimum_price": 254,
              "increments": [
                {},
                {}
              ],
              "split": {
                "enabled": false,
                "rules": [
                  {
                    "type": "type2",
                    "amount": 72,
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
                        "type": "type8",
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
                        "volume_percentage": 66,
                        "delay": 224,
                        "days": [
                          200,
                          201,
                          202
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": false,
                        "transfer_interval": "transfer_interval4",
                        "transfer_day": 232
                      },
                      "code": "code8",
                      "payment_mode": "payment_mode4"
                    },
                    "gateway_id": "gateway_id2",
                    "options": {
                      "liable": false,
                      "charge_processing_fee": false,
                      "charge_remainder_fee": "charge_remainder_fee4"
                    },
                    "id": "id8"
                  },
                  {
                    "type": "type1",
                    "amount": 73,
                    "recipient": {
                      "id": "id1",
                      "name": "name1",
                      "email": "email5",
                      "document": "document5",
                      "description": "description1",
                      "type": "type9",
                      "status": "status7",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "updated_at": "2016-03-13T12:52:32.123Z",
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "default_bank_account": {
                        "id": "id9",
                        "holder_name": "holder_name5",
                        "holder_type": "holder_type1",
                        "bank": "bank7",
                        "branch_number": "branch_number5",
                        "branch_check_digit": "branch_check_digit5",
                        "account_number": "account_number9",
                        "account_check_digit": "account_check_digit5",
                        "type": "type9",
                        "status": "status1",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "updated_at": "2016-03-13T12:52:32.123Z",
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "recipient": {},
                        "metadata": {
                          "key0": "metadata6",
                          "key1": "metadata5",
                          "key2": "metadata4"
                        },
                        "pix_key": "pix_key7"
                      },
                      "gateway_recipients": [
                        {
                          "gateway": "gateway5",
                          "status": "status7",
                          "pgid": "pgid1",
                          "created_at": "created_at3",
                          "updated_at": "updated_at1"
                        },
                        {
                          "gateway": "gateway6",
                          "status": "status8",
                          "pgid": "pgid2",
                          "created_at": "created_at4",
                          "updated_at": "updated_at2"
                        }
                      ],
                      "metadata": {
                        "key0": "metadata2",
                        "key1": "metadata3",
                        "key2": "metadata4"
                      },
                      "automatic_anticipation_settings": {
                        "enabled": true,
                        "type": "type5",
                        "volume_percentage": 65,
                        "delay": 225,
                        "days": [
                          201
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": true,
                        "transfer_interval": "transfer_interval3",
                        "transfer_day": 231
                      },
                      "code": "code9",
                      "payment_mode": "payment_mode5"
                    },
                    "gateway_id": "gateway_id1",
                    "options": {
                      "liable": true,
                      "charge_processing_fee": true,
                      "charge_remainder_fee": "charge_remainder_fee5"
                    },
                    "id": "id9"
                  },
                  {
                    "type": "type0",
                    "amount": 74,
                    "recipient": {
                      "id": "id2",
                      "name": "name2",
                      "email": "email4",
                      "document": "document4",
                      "description": "description2",
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
                        "volume_percentage": 64,
                        "delay": 226,
                        "days": [
                          202,
                          203
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": false,
                        "transfer_interval": "transfer_interval2",
                        "transfer_day": 230
                      },
                      "code": "code0",
                      "payment_mode": "payment_mode6"
                    },
                    "gateway_id": "gateway_id0",
                    "options": {
                      "liable": false,
                      "charge_processing_fee": false,
                      "charge_remainder_fee": "charge_remainder_fee6"
                    },
                    "id": "id0"
                  }
                ]
              }
            },
            "subscription_item": {}
          },
          {
            "id": "id7",
            "value": 156.59,
            "increment_type": "increment_type9",
            "status": "status1",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 189,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description3",
            "subscription": {
              "id": "id7",
              "code": "code5",
              "start_at": "2016-03-13T12:52:32.123Z",
              "interval": "interval5",
              "interval_count": 255,
              "billing_type": "billing_type1",
              "current_cycle": {
                "start_at": "2016-03-13T12:52:32.123Z",
                "end_at": "2016-03-13T12:52:32.123Z",
                "id": "id5",
                "billing_at": "2016-03-13T12:52:32.123Z",
                "subscription": {},
                "status": "status7",
                "duration": 239,
                "created_at": "created_at3",
                "updated_at": "updated_at1",
                "cycle": 179
              },
              "payment_method": "payment_method3",
              "currency": "currency7",
              "installments": 167,
              "status": "status9",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "customer": {
                "id": "id7",
                "name": "name7",
                "email": "email9",
                "delinquent": true,
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "document": "document1",
                "type": "type3",
                "fb_access_token": "fb_access_token1",
                "address": {
                  "id": "id3",
                  "street": "street3",
                  "number": "number1",
                  "complement": "complement9",
                  "zip_code": "zip_code7",
                  "neighborhood": "neighborhood9",
                  "city": "city3",
                  "state": "state9",
                  "country": "country7",
                  "status": "status5",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "customer": {},
                  "metadata": {
                    "key0": "metadata6",
                    "key1": "metadata5",
                    "key2": "metadata4"
                  },
                  "line_1": "line_17",
                  "line_2": "line_21",
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                },
                "metadata": {
                  "key0": "metadata4",
                  "key1": "metadata3"
                },
                "phones": {
                  "home_phone": {
                    "country_code": "country_code9",
                    "number": "number7",
                    "area_code": "area_code9"
                  },
                  "mobile_phone": {
                    "country_code": "country_code1",
                    "number": "number3",
                    "area_code": "area_code9"
                  }
                },
                "fb_id": 159,
                "code": "code5",
                "document_type": "document_type5"
              },
              "card": {
                "id": "id1",
                "last_four_digits": "last_four_digits7",
                "brand": "brand5",
                "holder_name": "holder_name7",
                "exp_month": 195,
                "exp_year": 101,
                "status": "status7",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "billing_address": {
                  "street": "street3",
                  "number": "number9",
                  "zip_code": "zip_code7",
                  "neighborhood": "neighborhood9",
                  "city": "city7",
                  "state": "state1",
                  "country": "country7",
                  "complement": "complement1",
                  "line_1": "line_13",
                  "line_2": "line_21"
                },
                "customer": {
                  "id": "id1",
                  "name": "name1",
                  "email": "email5",
                  "delinquent": true,
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "document": "document5",
                  "type": "type9",
                  "fb_access_token": "fb_access_token5",
                  "address": {
                    "id": "id7",
                    "street": "street7",
                    "number": "number5",
                    "complement": "complement3",
                    "zip_code": "zip_code1",
                    "neighborhood": "neighborhood3",
                    "city": "city7",
                    "state": "state3",
                    "country": "country1",
                    "status": "status9",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "updated_at": "2016-03-13T12:52:32.123Z",
                    "customer": {},
                    "metadata": {
                      "key0": "metadata4"
                    },
                    "line_1": "line_11",
                    "line_2": "line_25",
                    "deleted_at": "2016-03-13T12:52:32.123Z"
                  },
                  "metadata": {
                    "key0": "metadata2"
                  },
                  "phones": {
                    "home_phone": {
                      "country_code": "country_code3",
                      "number": "number9",
                      "area_code": "area_code3"
                    },
                    "mobile_phone": {
                      "country_code": "country_code7",
                      "number": "number5",
                      "area_code": "area_code7"
                    }
                  },
                  "fb_id": 89,
                  "code": "code9",
                  "document_type": "document_type9"
                },
                "metadata": {
                  "key0": "metadata2",
                  "key1": "metadata3"
                },
                "type": "type9",
                "holder_document": "holder_document5",
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "first_six_digits": "first_six_digits1",
                "label": "label1"
              },
              "items": [
                {},
                {}
              ],
              "statement_descriptor": "statement_descriptor7",
              "metadata": {
                "key0": "metadata4",
                "key1": "metadata3",
                "key2": "metadata2"
              },
              "setup": {
                "id": "id1",
                "description": "description1",
                "amount": 191,
                "status": "status3"
              },
              "gateway_affiliation_id": "gateway_affiliation_id3",
              "next_billing_at": "2016-03-13T12:52:32.123Z",
              "billing_day": 139,
              "minimum_price": 253,
              "increments": [
                {}
              ],
              "split": {
                "enabled": true,
                "rules": [
                  {
                    "type": "type3",
                    "amount": 71,
                    "recipient": {
                      "id": "id9",
                      "name": "name9",
                      "email": "email7",
                      "document": "document7",
                      "description": "description9",
                      "type": "type1",
                      "status": "status9",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "updated_at": "2016-03-13T12:52:32.123Z",
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "default_bank_account": {
                        "id": "id7",
                        "holder_name": "holder_name3",
                        "holder_type": "holder_type9",
                        "bank": "bank5",
                        "branch_number": "branch_number3",
                        "branch_check_digit": "branch_check_digit3",
                        "account_number": "account_number7",
                        "account_check_digit": "account_check_digit3",
                        "type": "type7",
                        "status": "status9",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "updated_at": "2016-03-13T12:52:32.123Z",
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "recipient": {},
                        "metadata": {
                          "key0": "metadata4",
                          "key1": "metadata3"
                        },
                        "pix_key": "pix_key9"
                      },
                      "gateway_recipients": [
                        {
                          "gateway": "gateway3",
                          "status": "status5",
                          "pgid": "pgid9",
                          "created_at": "created_at1",
                          "updated_at": "updated_at9"
                        },
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
                        "key0": "metadata4",
                        "key1": "metadata5"
                      },
                      "automatic_anticipation_settings": {
                        "enabled": true,
                        "type": "type7",
                        "volume_percentage": 67,
                        "delay": 223,
                        "days": [
                          199,
                          200
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": true,
                        "transfer_interval": "transfer_interval5",
                        "transfer_day": 233
                      },
                      "code": "code7",
                      "payment_mode": "payment_mode3"
                    },
                    "gateway_id": "gateway_id3",
                    "options": {
                      "liable": true,
                      "charge_processing_fee": true,
                      "charge_remainder_fee": "charge_remainder_fee3"
                    },
                    "id": "id7"
                  },
                  {
                    "type": "type2",
                    "amount": 72,
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
                        "type": "type8",
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
                        "volume_percentage": 66,
                        "delay": 224,
                        "days": [
                          200,
                          201,
                          202
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": false,
                        "transfer_interval": "transfer_interval4",
                        "transfer_day": 232
                      },
                      "code": "code8",
                      "payment_mode": "payment_mode4"
                    },
                    "gateway_id": "gateway_id2",
                    "options": {
                      "liable": false,
                      "charge_processing_fee": false,
                      "charge_remainder_fee": "charge_remainder_fee4"
                    },
                    "id": "id8"
                  }
                ]
              }
            },
            "subscription_item": {}
          },
          {
            "id": "id8",
            "value": 156.6,
            "increment_type": "increment_type0",
            "status": "status0",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 188,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description2",
            "subscription": {
              "id": "id6",
              "code": "code4",
              "start_at": "2016-03-13T12:52:32.123Z",
              "interval": "interval4",
              "interval_count": 254,
              "billing_type": "billing_type0",
              "current_cycle": {
                "start_at": "2016-03-13T12:52:32.123Z",
                "end_at": "2016-03-13T12:52:32.123Z",
                "id": "id4",
                "billing_at": "2016-03-13T12:52:32.123Z",
                "subscription": {},
                "status": "status6",
                "duration": 238,
                "created_at": "created_at2",
                "updated_at": "updated_at0",
                "cycle": 178
              },
              "payment_method": "payment_method4",
              "currency": "currency6",
              "installments": 166,
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
                "address": {
                  "id": "id2",
                  "street": "street2",
                  "number": "number0",
                  "complement": "complement8",
                  "zip_code": "zip_code6",
                  "neighborhood": "neighborhood8",
                  "city": "city2",
                  "state": "state8",
                  "country": "country6",
                  "status": "status4",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "customer": {},
                  "metadata": {
                    "key0": "metadata7",
                    "key1": "metadata6"
                  },
                  "line_1": "line_16",
                  "line_2": "line_20",
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                },
                "metadata": {
                  "key0": "metadata3",
                  "key1": "metadata2",
                  "key2": "metadata1"
                },
                "phones": {
                  "home_phone": {
                    "country_code": "country_code8",
                    "number": "number6",
                    "area_code": "area_code8"
                  },
                  "mobile_phone": {
                    "country_code": "country_code2",
                    "number": "number4",
                    "area_code": "area_code8"
                  }
                },
                "fb_id": 158,
                "code": "code4",
                "document_type": "document_type4"
              },
              "card": {
                "id": "id0",
                "last_four_digits": "last_four_digits6",
                "brand": "brand4",
                "holder_name": "holder_name6",
                "exp_month": 196,
                "exp_year": 100,
                "status": "status8",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "billing_address": {
                  "street": "street2",
                  "number": "number0",
                  "zip_code": "zip_code6",
                  "neighborhood": "neighborhood8",
                  "city": "city8",
                  "state": "state2",
                  "country": "country6",
                  "complement": "complement2",
                  "line_1": "line_14",
                  "line_2": "line_20"
                },
                "customer": {
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
                    "customer": {},
                    "metadata": {
                      "key0": "metadata3",
                      "key1": "metadata2"
                    },
                    "line_1": "line_10",
                    "line_2": "line_24",
                    "deleted_at": "2016-03-13T12:52:32.123Z"
                  },
                  "metadata": {
                    "key0": "metadata3",
                    "key1": "metadata4"
                  },
                  "phones": {
                    "home_phone": {
                      "country_code": "country_code2",
                      "number": "number0",
                      "area_code": "area_code2"
                    },
                    "mobile_phone": {
                      "country_code": "country_code8",
                      "number": "number4",
                      "area_code": "area_code8"
                    }
                  },
                  "fb_id": 88,
                  "code": "code8",
                  "document_type": "document_type8"
                },
                "metadata": {
                  "key0": "metadata3",
                  "key1": "metadata4",
                  "key2": "metadata5"
                },
                "type": "type0",
                "holder_document": "holder_document6",
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "first_six_digits": "first_six_digits0",
                "label": "label0"
              },
              "items": [
                {}
              ],
              "statement_descriptor": "statement_descriptor6",
              "metadata": {
                "key0": "metadata3"
              },
              "setup": {
                "id": "id0",
                "description": "description0",
                "amount": 190,
                "status": "status2"
              },
              "gateway_affiliation_id": "gateway_affiliation_id2",
              "next_billing_at": "2016-03-13T12:52:32.123Z",
              "billing_day": 138,
              "minimum_price": 252,
              "increments": [
                {},
                {},
                {}
              ],
              "split": {
                "enabled": false,
                "rules": [
                  {
                    "type": "type4",
                    "amount": 70,
                    "recipient": {
                      "id": "id8",
                      "name": "name8",
                      "email": "email8",
                      "document": "document8",
                      "description": "description8",
                      "type": "type2",
                      "status": "status0",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "updated_at": "2016-03-13T12:52:32.123Z",
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "default_bank_account": {
                        "id": "id6",
                        "holder_name": "holder_name2",
                        "holder_type": "holder_type8",
                        "bank": "bank4",
                        "branch_number": "branch_number2",
                        "branch_check_digit": "branch_check_digit2",
                        "account_number": "account_number6",
                        "account_check_digit": "account_check_digit2",
                        "type": "type6",
                        "status": "status8",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "updated_at": "2016-03-13T12:52:32.123Z",
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "recipient": {},
                        "metadata": {
                          "key0": "metadata3",
                          "key1": "metadata2",
                          "key2": "metadata1"
                        },
                        "pix_key": "pix_key0"
                      },
                      "gateway_recipients": [
                        {
                          "gateway": "gateway2",
                          "status": "status4",
                          "pgid": "pgid8",
                          "created_at": "created_at0",
                          "updated_at": "updated_at8"
                        },
                        {
                          "gateway": "gateway3",
                          "status": "status5",
                          "pgid": "pgid9",
                          "created_at": "created_at1",
                          "updated_at": "updated_at9"
                        }
                      ],
                      "metadata": {
                        "key0": "metadata5",
                        "key1": "metadata6",
                        "key2": "metadata7"
                      },
                      "automatic_anticipation_settings": {
                        "enabled": false,
                        "type": "type8",
                        "volume_percentage": 68,
                        "delay": 222,
                        "days": [
                          198
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": false,
                        "transfer_interval": "transfer_interval6",
                        "transfer_day": 234
                      },
                      "code": "code6",
                      "payment_mode": "payment_mode2"
                    },
                    "gateway_id": "gateway_id4",
                    "options": {
                      "liable": false,
                      "charge_processing_fee": false,
                      "charge_remainder_fee": "charge_remainder_fee2"
                    },
                    "id": "id6"
                  }
                ]
              }
            },
            "subscription_item": {}
          }
        ],
        "subscription": {
          "id": "id7",
          "code": "code5",
          "start_at": "2016-03-13T12:52:32.123Z",
          "interval": "interval5",
          "interval_count": 55,
          "billing_type": "billing_type9",
          "current_cycle": {
            "start_at": "2016-03-13T12:52:32.123Z",
            "end_at": "2016-03-13T12:52:32.123Z",
            "id": "id5",
            "billing_at": "2016-03-13T12:52:32.123Z",
            "subscription": {},
            "status": "status3",
            "duration": 39,
            "created_at": "created_at3",
            "updated_at": "updated_at1",
            "cycle": 235
          },
          "payment_method": "payment_method3",
          "currency": "currency3",
          "installments": 223,
          "status": "status1",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "customer": {
            "id": "id3",
            "name": "name3",
            "email": "email3",
            "delinquent": true,
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "document": "document3",
            "type": "type7",
            "fb_access_token": "fb_access_token7",
            "address": {
              "id": "id9",
              "street": "street9",
              "number": "number7",
              "complement": "complement5",
              "zip_code": "zip_code3",
              "neighborhood": "neighborhood5",
              "city": "city9",
              "state": "state5",
              "country": "country3",
              "status": "status1",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "customer": {},
              "metadata": {
                "key0": "metadata6",
                "key1": "metadata5"
              },
              "line_1": "line_17",
              "line_2": "line_27",
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            "metadata": {
              "key0": "metadata0",
              "key1": "metadata1"
            },
            "phones": {
              "home_phone": {
                "country_code": "country_code5",
                "number": "number7",
                "area_code": "area_code5"
              },
              "mobile_phone": {
                "country_code": "country_code5",
                "number": "number7",
                "area_code": "area_code5"
              }
            },
            "fb_id": 201,
            "code": "code1",
            "document_type": "document_type1"
          },
          "card": {
            "id": "id9",
            "last_four_digits": "last_four_digits5",
            "brand": "brand3",
            "holder_name": "holder_name5",
            "exp_month": 201,
            "exp_year": 95,
            "status": "status9",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "billing_address": {
              "street": "street1",
              "number": "number1",
              "zip_code": "zip_code5",
              "neighborhood": "neighborhood7",
              "city": "city9",
              "state": "state3",
              "country": "country5",
              "complement": "complement3",
              "line_1": "line_15",
              "line_2": "line_29"
            },
            "customer": {
              "id": "id9",
              "name": "name9",
              "email": "email7",
              "delinquent": true,
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "document": "document7",
              "type": "type1",
              "fb_access_token": "fb_access_token3",
              "address": {
                "id": "id5",
                "street": "street5",
                "number": "number3",
                "complement": "complement1",
                "zip_code": "zip_code9",
                "neighborhood": "neighborhood1",
                "city": "city5",
                "state": "state1",
                "country": "country9",
                "status": "status7",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "customer": {},
                "metadata": {
                  "key0": "metadata2"
                },
                "line_1": "line_19",
                "line_2": "line_23",
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              "metadata": {
                "key0": "metadata4"
              },
              "phones": {
                "home_phone": {
                  "country_code": "country_code1",
                  "number": "number1",
                  "area_code": "area_code1"
                },
                "mobile_phone": {
                  "country_code": "country_code9",
                  "number": "number3",
                  "area_code": "area_code9"
                }
              },
              "fb_id": 83,
              "code": "code7",
              "document_type": "document_type7"
            },
            "metadata": {
              "key0": "metadata4",
              "key1": "metadata5"
            },
            "type": "type1",
            "holder_document": "holder_document7",
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "first_six_digits": "first_six_digits9",
            "label": "label9"
          },
          "items": [
            {},
            {},
            {}
          ],
          "statement_descriptor": "statement_descriptor7",
          "metadata": {
            "key0": "metadata6",
            "key1": "metadata7"
          },
          "setup": {
            "id": "id1",
            "description": "description9",
            "amount": 9,
            "status": "status7"
          },
          "gateway_affiliation_id": "gateway_affiliation_id3",
          "next_billing_at": "2016-03-13T12:52:32.123Z",
          "billing_day": 61,
          "minimum_price": 203,
          "increments": [
            {
              "id": "id0",
              "value": 75.82,
              "increment_type": "increment_type2",
              "status": "status8",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 74,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description0",
              "subscription": {},
              "subscription_item": {}
            },
            {
              "id": "id1",
              "value": 75.83,
              "increment_type": "increment_type3",
              "status": "status7",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 73,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description9",
              "subscription": {},
              "subscription_item": {}
            },
            {
              "id": "id2",
              "value": 75.84,
              "increment_type": "increment_type4",
              "status": "status6",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 72,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description8",
              "subscription": {},
              "subscription_item": {}
            }
          ],
          "split": {
            "enabled": true,
            "rules": [
              {
                "type": "type7",
                "amount": 219,
                "recipient": {
                  "id": "id5",
                  "name": "name5",
                  "email": "email1",
                  "document": "document1",
                  "description": "description5",
                  "type": "type5",
                  "status": "status3",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "default_bank_account": {
                    "id": "id3",
                    "holder_name": "holder_name9",
                    "holder_type": "holder_type5",
                    "bank": "bank1",
                    "branch_number": "branch_number9",
                    "branch_check_digit": "branch_check_digit9",
                    "account_number": "account_number3",
                    "account_check_digit": "account_check_digit9",
                    "type": "type7",
                    "status": "status5",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "updated_at": "2016-03-13T12:52:32.123Z",
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "recipient": {},
                    "metadata": {
                      "key0": "metadata0"
                    },
                    "pix_key": "pix_key3"
                  },
                  "gateway_recipients": [
                    {
                      "gateway": "gateway9",
                      "status": "status1",
                      "pgid": "pgid5",
                      "created_at": "created_at7",
                      "updated_at": "updated_at5"
                    }
                  ],
                  "metadata": {
                    "key0": "metadata8"
                  },
                  "automatic_anticipation_settings": {
                    "enabled": true,
                    "type": "type1",
                    "volume_percentage": 99,
                    "delay": 191,
                    "days": [
                      167,
                      168,
                      169
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": true,
                    "transfer_interval": "transfer_interval9",
                    "transfer_day": 9
                  },
                  "code": "code3",
                  "payment_mode": "payment_mode1"
                },
                "gateway_id": "gateway_id7",
                "options": {
                  "liable": true,
                  "charge_processing_fee": true,
                  "charge_remainder_fee": "charge_remainder_fee7"
                },
                "id": "id3"
              },
              {
                "type": "type8",
                "amount": 220,
                "recipient": {
                  "id": "id6",
                  "name": "name6",
                  "email": "email0",
                  "document": "document0",
                  "description": "description4",
                  "type": "type4",
                  "status": "status2",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "default_bank_account": {
                    "id": "id4",
                    "holder_name": "holder_name0",
                    "holder_type": "holder_type6",
                    "bank": "bank2",
                    "branch_number": "branch_number0",
                    "branch_check_digit": "branch_check_digit0",
                    "account_number": "account_number4",
                    "account_check_digit": "account_check_digit0",
                    "type": "type6",
                    "status": "status6",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "updated_at": "2016-03-13T12:52:32.123Z",
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "recipient": {},
                    "metadata": {
                      "key0": "metadata1",
                      "key1": "metadata0",
                      "key2": "metadata9"
                    },
                    "pix_key": "pix_key2"
                  },
                  "gateway_recipients": [
                    {
                      "gateway": "gateway0",
                      "status": "status2",
                      "pgid": "pgid6",
                      "created_at": "created_at8",
                      "updated_at": "updated_at6"
                    },
                    {
                      "gateway": "gateway1",
                      "status": "status3",
                      "pgid": "pgid7",
                      "created_at": "created_at9",
                      "updated_at": "updated_at7"
                    }
                  ],
                  "metadata": {
                    "key0": "metadata7",
                    "key1": "metadata8",
                    "key2": "metadata9"
                  },
                  "automatic_anticipation_settings": {
                    "enabled": false,
                    "type": "type0",
                    "volume_percentage": 98,
                    "delay": 192,
                    "days": [
                      168
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": false,
                    "transfer_interval": "transfer_interval8",
                    "transfer_day": 8
                  },
                  "code": "code4",
                  "payment_mode": "payment_mode0"
                },
                "gateway_id": "gateway_id8",
                "options": {
                  "liable": false,
                  "charge_processing_fee": false,
                  "charge_remainder_fee": "charge_remainder_fee8"
                },
                "id": "id2"
              },
              {
                "type": "type9",
                "amount": 221,
                "recipient": {
                  "id": "id7",
                  "name": "name7",
                  "email": "email9",
                  "document": "document9",
                  "description": "description3",
                  "type": "type3",
                  "status": "status1",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "default_bank_account": {
                    "id": "id5",
                    "holder_name": "holder_name1",
                    "holder_type": "holder_type7",
                    "bank": "bank3",
                    "branch_number": "branch_number1",
                    "branch_check_digit": "branch_check_digit1",
                    "account_number": "account_number5",
                    "account_check_digit": "account_check_digit1",
                    "type": "type5",
                    "status": "status7",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "updated_at": "2016-03-13T12:52:32.123Z",
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "recipient": {},
                    "metadata": {
                      "key0": "metadata2",
                      "key1": "metadata1"
                    },
                    "pix_key": "pix_key1"
                  },
                  "gateway_recipients": [
                    {
                      "gateway": "gateway1",
                      "status": "status3",
                      "pgid": "pgid7",
                      "created_at": "created_at9",
                      "updated_at": "updated_at7"
                    },
                    {
                      "gateway": "gateway2",
                      "status": "status4",
                      "pgid": "pgid8",
                      "created_at": "created_at0",
                      "updated_at": "updated_at8"
                    },
                    {
                      "gateway": "gateway3",
                      "status": "status5",
                      "pgid": "pgid9",
                      "created_at": "created_at1",
                      "updated_at": "updated_at9"
                    }
                  ],
                  "metadata": {
                    "key0": "metadata6",
                    "key1": "metadata7"
                  },
                  "automatic_anticipation_settings": {
                    "enabled": true,
                    "type": "type9",
                    "volume_percentage": 97,
                    "delay": 193,
                    "days": [
                      169,
                      170
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": true,
                    "transfer_interval": "transfer_interval7",
                    "transfer_day": 7
                  },
                  "code": "code5",
                  "payment_mode": "payment_mode9"
                },
                "gateway_id": "gateway_id9",
                "options": {
                  "liable": true,
                  "charge_processing_fee": true,
                  "charge_remainder_fee": "charge_remainder_fee9"
                },
                "id": "id1"
              }
            ]
          }
        },
        "name": "name7",
        "quantity": 247,
        "cycles": 11,
        "deleted_at": "2016-03-13T12:52:32.123Z"
      }
    }
  ],
  "increments": [
    {
      "id": "id7",
      "value": 124.19,
      "increment_type": "increment_type9",
      "status": "status1",
      "created_at": "2016-03-13T12:52:32.123Z",
      "cycles": 101,
      "deleted_at": "2016-03-13T12:52:32.123Z",
      "description": "description3",
      "subscription": {
        "id": "id7",
        "code": "code5",
        "start_at": "2016-03-13T12:52:32.123Z",
        "interval": "interval5",
        "interval_count": 167,
        "billing_type": "billing_type1",
        "current_cycle": {
          "start_at": "2016-03-13T12:52:32.123Z",
          "end_at": "2016-03-13T12:52:32.123Z",
          "id": "id5",
          "billing_at": "2016-03-13T12:52:32.123Z",
          "subscription": {},
          "status": "status7",
          "duration": 151,
          "created_at": "created_at3",
          "updated_at": "updated_at1",
          "cycle": 91
        },
        "payment_method": "payment_method3",
        "currency": "currency7",
        "installments": 79,
        "status": "status9",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "customer": {
          "id": "id7",
          "name": "name7",
          "email": "email9",
          "delinquent": true,
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "document": "document1",
          "type": "type3",
          "fb_access_token": "fb_access_token1",
          "address": {
            "id": "id3",
            "street": "street3",
            "number": "number1",
            "complement": "complement9",
            "zip_code": "zip_code7",
            "neighborhood": "neighborhood9",
            "city": "city3",
            "state": "state9",
            "country": "country7",
            "status": "status5",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "customer": {},
            "metadata": {
              "key0": "metadata6",
              "key1": "metadata5",
              "key2": "metadata4"
            },
            "line_1": "line_17",
            "line_2": "line_21",
            "deleted_at": "2016-03-13T12:52:32.123Z"
          },
          "metadata": {
            "key0": "metadata4",
            "key1": "metadata3"
          },
          "phones": {
            "home_phone": {
              "country_code": "country_code9",
              "number": "number7",
              "area_code": "area_code9"
            },
            "mobile_phone": {
              "country_code": "country_code1",
              "number": "number3",
              "area_code": "area_code9"
            }
          },
          "fb_id": 71,
          "code": "code5",
          "document_type": "document_type5"
        },
        "card": {
          "id": "id1",
          "last_four_digits": "last_four_digits7",
          "brand": "brand5",
          "holder_name": "holder_name7",
          "exp_month": 27,
          "exp_year": 13,
          "status": "status7",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "billing_address": {
            "street": "street3",
            "number": "number9",
            "zip_code": "zip_code7",
            "neighborhood": "neighborhood9",
            "city": "city7",
            "state": "state1",
            "country": "country7",
            "complement": "complement1",
            "line_1": "line_13",
            "line_2": "line_21"
          },
          "customer": {
            "id": "id1",
            "name": "name1",
            "email": "email5",
            "delinquent": true,
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "document": "document5",
            "type": "type9",
            "fb_access_token": "fb_access_token5",
            "address": {
              "id": "id7",
              "street": "street7",
              "number": "number5",
              "complement": "complement3",
              "zip_code": "zip_code1",
              "neighborhood": "neighborhood3",
              "city": "city7",
              "state": "state3",
              "country": "country1",
              "status": "status9",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "customer": {},
              "metadata": {
                "key0": "metadata4"
              },
              "line_1": "line_11",
              "line_2": "line_25",
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            "metadata": {
              "key0": "metadata2"
            },
            "phones": {
              "home_phone": {
                "country_code": "country_code3",
                "number": "number9",
                "area_code": "area_code3"
              },
              "mobile_phone": {
                "country_code": "country_code7",
                "number": "number5",
                "area_code": "area_code7"
              }
            },
            "fb_id": 1,
            "code": "code9",
            "document_type": "document_type9"
          },
          "metadata": {
            "key0": "metadata2",
            "key1": "metadata3"
          },
          "type": "type9",
          "holder_document": "holder_document5",
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "first_six_digits": "first_six_digits1",
          "label": "label1"
        },
        "items": [
          {
            "id": "id4",
            "description": "description4",
            "status": "status6",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "pricing_scheme": {
              "price": 54,
              "scheme_type": "scheme_type6",
              "price_brackets": [
                {
                  "start_quantity": 155,
                  "price": 163,
                  "end_quantity": 163,
                  "overage_price": 177
                },
                {
                  "start_quantity": 156,
                  "price": 162,
                  "end_quantity": 164,
                  "overage_price": 178
                },
                {
                  "start_quantity": 157,
                  "price": 161,
                  "end_quantity": 165,
                  "overage_price": 179
                }
              ],
              "minimum_price": 214,
              "percentage": 193.24
            },
            "discounts": [
              {
                "id": "id5",
                "value": 190.87,
                "discount_type": "discount_type3",
                "status": "status7",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 167,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description5",
                "subscription": {},
                "subscription_item": {}
              },
              {
                "id": "id6",
                "value": 190.88,
                "discount_type": "discount_type4",
                "status": "status8",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 168,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description6",
                "subscription": {},
                "subscription_item": {}
              },
              {
                "id": "id7",
                "value": 190.89,
                "discount_type": "discount_type5",
                "status": "status9",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 169,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description7",
                "subscription": {},
                "subscription_item": {}
              }
            ],
            "increments": [
              {}
            ],
            "subscription": {},
            "name": "name4",
            "quantity": 212,
            "cycles": 232,
            "deleted_at": "2016-03-13T12:52:32.123Z"
          },
          {
            "id": "id5",
            "description": "description5",
            "status": "status7",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "pricing_scheme": {
              "price": 53,
              "scheme_type": "scheme_type7",
              "price_brackets": [
                {
                  "start_quantity": 156,
                  "price": 162,
                  "end_quantity": 164,
                  "overage_price": 178
                }
              ],
              "minimum_price": 213,
              "percentage": 193.25
            },
            "discounts": [
              {
                "id": "id6",
                "value": 190.88,
                "discount_type": "discount_type4",
                "status": "status8",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 168,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description6",
                "subscription": {},
                "subscription_item": {}
              }
            ],
            "increments": [
              {},
              {}
            ],
            "subscription": {},
            "name": "name5",
            "quantity": 213,
            "cycles": 233,
            "deleted_at": "2016-03-13T12:52:32.123Z"
          }
        ],
        "statement_descriptor": "statement_descriptor7",
        "metadata": {
          "key0": "metadata4",
          "key1": "metadata3",
          "key2": "metadata2"
        },
        "setup": {
          "id": "id1",
          "description": "description1",
          "amount": 103,
          "status": "status3"
        },
        "gateway_affiliation_id": "gateway_affiliation_id3",
        "next_billing_at": "2016-03-13T12:52:32.123Z",
        "billing_day": 51,
        "minimum_price": 165,
        "increments": [
          {}
        ],
        "split": {
          "enabled": true,
          "rules": [
            {
              "type": "type3",
              "amount": 239,
              "recipient": {
                "id": "id9",
                "name": "name9",
                "email": "email7",
                "document": "document7",
                "description": "description9",
                "type": "type1",
                "status": "status9",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "default_bank_account": {
                  "id": "id7",
                  "holder_name": "holder_name3",
                  "holder_type": "holder_type9",
                  "bank": "bank5",
                  "branch_number": "branch_number3",
                  "branch_check_digit": "branch_check_digit3",
                  "account_number": "account_number7",
                  "account_check_digit": "account_check_digit3",
                  "type": "type7",
                  "status": "status9",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "recipient": {},
                  "metadata": {
                    "key0": "metadata4",
                    "key1": "metadata3"
                  },
                  "pix_key": "pix_key1"
                },
                "gateway_recipients": [
                  {
                    "gateway": "gateway3",
                    "status": "status5",
                    "pgid": "pgid9",
                    "created_at": "created_at1",
                    "updated_at": "updated_at9"
                  },
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
                  "key0": "metadata4",
                  "key1": "metadata5"
                },
                "automatic_anticipation_settings": {
                  "enabled": true,
                  "type": "type7",
                  "volume_percentage": 155,
                  "delay": 135,
                  "days": [
                    111,
                    112
                  ]
                },
                "transfer_settings": {
                  "transfer_enabled": true,
                  "transfer_interval": "transfer_interval5",
                  "transfer_day": 65
                },
                "code": "code7",
                "payment_mode": "payment_mode3"
              },
              "gateway_id": "gateway_id3",
              "options": {
                "liable": true,
                "charge_processing_fee": true,
                "charge_remainder_fee": "charge_remainder_fee3"
              },
              "id": "id7"
            },
            {
              "type": "type2",
              "amount": 240,
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
                  "type": "type8",
                  "status": "status0",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "recipient": {},
                  "metadata": {
                    "key0": "metadata5"
                  },
                  "pix_key": "pix_key2"
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
                  "volume_percentage": 154,
                  "delay": 136,
                  "days": [
                    112,
                    113,
                    114
                  ]
                },
                "transfer_settings": {
                  "transfer_enabled": false,
                  "transfer_interval": "transfer_interval4",
                  "transfer_day": 64
                },
                "code": "code8",
                "payment_mode": "payment_mode4"
              },
              "gateway_id": "gateway_id2",
              "options": {
                "liable": false,
                "charge_processing_fee": false,
                "charge_remainder_fee": "charge_remainder_fee4"
              },
              "id": "id8"
            }
          ]
        }
      },
      "subscription_item": {
        "id": "id3",
        "description": "description7",
        "status": "status5",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "pricing_scheme": {
          "price": 145,
          "scheme_type": "scheme_type5",
          "price_brackets": [
            {
              "start_quantity": 64,
              "price": 2,
              "end_quantity": 72,
              "overage_price": 86
            }
          ],
          "minimum_price": 49,
          "percentage": 102.73
        },
        "discounts": [
          {
            "id": "id4",
            "value": 159.26,
            "discount_type": "discount_type2",
            "status": "status6",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 78,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description4",
            "subscription": {
              "id": "id0",
              "code": "code8",
              "start_at": "2016-03-13T12:52:32.123Z",
              "interval": "interval2",
              "interval_count": 176,
              "billing_type": "billing_type6",
              "current_cycle": {
                "start_at": "2016-03-13T12:52:32.123Z",
                "end_at": "2016-03-13T12:52:32.123Z",
                "id": "id8",
                "billing_at": "2016-03-13T12:52:32.123Z",
                "subscription": {},
                "status": "status0",
                "duration": 160,
                "created_at": "created_at6",
                "updated_at": "updated_at4",
                "cycle": 100
              },
              "payment_method": "payment_method0",
              "currency": "currency0",
              "installments": 88,
              "status": "status8",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "customer": {
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
                  "customer": {},
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
                "fb_id": 80,
                "code": "code8",
                "document_type": "document_type8"
              },
              "card": {
                "id": "id6",
                "last_four_digits": "last_four_digits2",
                "brand": "brand0",
                "holder_name": "holder_name2",
                "exp_month": 66,
                "exp_year": 230,
                "status": "status2",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "billing_address": {
                  "street": "street2",
                  "number": "number0",
                  "zip_code": "zip_code6",
                  "neighborhood": "neighborhood8",
                  "city": "city8",
                  "state": "state2",
                  "country": "country6",
                  "complement": "complement2",
                  "line_1": "line_14",
                  "line_2": "line_20"
                },
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
                  "address": {
                    "id": "id2",
                    "street": "street2",
                    "number": "number0",
                    "complement": "complement8",
                    "zip_code": "zip_code6",
                    "neighborhood": "neighborhood8",
                    "city": "city2",
                    "state": "state8",
                    "country": "country6",
                    "status": "status4",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "updated_at": "2016-03-13T12:52:32.123Z",
                    "customer": {},
                    "metadata": {
                      "key0": "metadata9",
                      "key1": "metadata8"
                    },
                    "line_1": "line_14",
                    "line_2": "line_20",
                    "deleted_at": "2016-03-13T12:52:32.123Z"
                  },
                  "metadata": {
                    "key0": "metadata7",
                    "key1": "metadata8"
                  },
                  "phones": {
                    "home_phone": {
                      "country_code": "country_code8",
                      "number": "number4",
                      "area_code": "area_code8"
                    },
                    "mobile_phone": {
                      "country_code": "country_code2",
                      "number": "number0",
                      "area_code": "area_code2"
                    }
                  },
                  "fb_id": 218,
                  "code": "code4",
                  "document_type": "document_type4"
                },
                "metadata": {
                  "key0": "metadata7",
                  "key1": "metadata8",
                  "key2": "metadata9"
                },
                "type": "type4",
                "holder_document": "holder_document0",
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "first_six_digits": "first_six_digits6",
                "label": "label6"
              },
              "items": [
                {}
              ],
              "statement_descriptor": "statement_descriptor0",
              "metadata": {
                "key0": "metadata3"
              },
              "setup": {
                "id": "id4",
                "description": "description6",
                "amount": 144,
                "status": "status4"
              },
              "gateway_affiliation_id": "gateway_affiliation_id6",
              "next_billing_at": "2016-03-13T12:52:32.123Z",
              "billing_day": 196,
              "minimum_price": 82,
              "increments": [
                {},
                {}
              ],
              "split": {
                "enabled": false,
                "rules": [
                  {
                    "type": "type0",
                    "amount": 8,
                    "recipient": {
                      "id": "id8",
                      "name": "name8",
                      "email": "email8",
                      "document": "document8",
                      "description": "description2",
                      "type": "type2",
                      "status": "status0",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "updated_at": "2016-03-13T12:52:32.123Z",
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "default_bank_account": {
                        "id": "id6",
                        "holder_name": "holder_name2",
                        "holder_type": "holder_type8",
                        "bank": "bank4",
                        "branch_number": "branch_number2",
                        "branch_check_digit": "branch_check_digit2",
                        "account_number": "account_number6",
                        "account_check_digit": "account_check_digit2",
                        "type": "type4",
                        "status": "status8",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "updated_at": "2016-03-13T12:52:32.123Z",
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "recipient": {},
                        "metadata": {
                          "key0": "metadata3",
                          "key1": "metadata2",
                          "key2": "metadata1"
                        },
                        "pix_key": "pix_key0"
                      },
                      "gateway_recipients": [
                        {
                          "gateway": "gateway2",
                          "status": "status4",
                          "pgid": "pgid8",
                          "created_at": "created_at0",
                          "updated_at": "updated_at8"
                        },
                        {
                          "gateway": "gateway3",
                          "status": "status5",
                          "pgid": "pgid9",
                          "created_at": "created_at1",
                          "updated_at": "updated_at9"
                        }
                      ],
                      "metadata": {
                        "key0": "metadata5",
                        "key1": "metadata6",
                        "key2": "metadata7"
                      },
                      "automatic_anticipation_settings": {
                        "enabled": false,
                        "type": "type8",
                        "volume_percentage": 54,
                        "delay": 236,
                        "days": [
                          212
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": false,
                        "transfer_interval": "transfer_interval6",
                        "transfer_day": 220
                      },
                      "code": "code6",
                      "payment_mode": "payment_mode8"
                    },
                    "gateway_id": "gateway_id0",
                    "options": {
                      "liable": false,
                      "charge_processing_fee": false,
                      "charge_remainder_fee": "charge_remainder_fee0"
                    },
                    "id": "id0"
                  }
                ]
              }
            },
            "subscription_item": {}
          },
          {
            "id": "id5",
            "value": 159.27,
            "discount_type": "discount_type3",
            "status": "status7",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 79,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description5",
            "subscription": {
              "id": "id1",
              "code": "code9",
              "start_at": "2016-03-13T12:52:32.123Z",
              "interval": "interval1",
              "interval_count": 177,
              "billing_type": "billing_type5",
              "current_cycle": {
                "start_at": "2016-03-13T12:52:32.123Z",
                "end_at": "2016-03-13T12:52:32.123Z",
                "id": "id9",
                "billing_at": "2016-03-13T12:52:32.123Z",
                "subscription": {},
                "status": "status1",
                "duration": 161,
                "created_at": "created_at7",
                "updated_at": "updated_at5",
                "cycle": 101
              },
              "payment_method": "payment_method9",
              "currency": "currency9",
              "installments": 89,
              "status": "status7",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "customer": {
                "id": "id9",
                "name": "name9",
                "email": "email7",
                "delinquent": true,
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "document": "document7",
                "type": "type1",
                "fb_access_token": "fb_access_token3",
                "address": {
                  "id": "id5",
                  "street": "street5",
                  "number": "number3",
                  "complement": "complement1",
                  "zip_code": "zip_code9",
                  "neighborhood": "neighborhood1",
                  "city": "city5",
                  "state": "state1",
                  "country": "country9",
                  "status": "status7",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "customer": {},
                  "metadata": {
                    "key0": "metadata2"
                  },
                  "line_1": "line_11",
                  "line_2": "line_23",
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                },
                "metadata": {
                  "key0": "metadata4"
                },
                "phones": {
                  "home_phone": {
                    "country_code": "country_code9",
                    "number": "number1",
                    "area_code": "area_code1"
                  },
                  "mobile_phone": {
                    "country_code": "country_code9",
                    "number": "number3",
                    "area_code": "area_code9"
                  }
                },
                "fb_id": 79,
                "code": "code7",
                "document_type": "document_type7"
              },
              "card": {
                "id": "id5",
                "last_four_digits": "last_four_digits1",
                "brand": "brand9",
                "holder_name": "holder_name1",
                "exp_month": 67,
                "exp_year": 229,
                "status": "status3",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "billing_address": {
                  "street": "street3",
                  "number": "number9",
                  "zip_code": "zip_code7",
                  "neighborhood": "neighborhood9",
                  "city": "city7",
                  "state": "state1",
                  "country": "country7",
                  "complement": "complement1",
                  "line_1": "line_13",
                  "line_2": "line_21"
                },
                "customer": {
                  "id": "id5",
                  "name": "name5",
                  "email": "email1",
                  "delinquent": true,
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "document": "document1",
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
                      "key0": "metadata8",
                      "key1": "metadata7",
                      "key2": "metadata6"
                    },
                    "line_1": "line_15",
                    "line_2": "line_29",
                    "deleted_at": "2016-03-13T12:52:32.123Z"
                  },
                  "metadata": {
                    "key0": "metadata8",
                    "key1": "metadata9",
                    "key2": "metadata0"
                  },
                  "phones": {
                    "home_phone": {
                      "country_code": "country_code7",
                      "number": "number5",
                      "area_code": "area_code7"
                    },
                    "mobile_phone": {
                      "country_code": "country_code3",
                      "number": "number9",
                      "area_code": "area_code3"
                    }
                  },
                  "fb_id": 217,
                  "code": "code3",
                  "document_type": "document_type3"
                },
                "metadata": {
                  "key0": "metadata8"
                },
                "type": "type5",
                "holder_document": "holder_document1",
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "first_six_digits": "first_six_digits5",
                "label": "label5"
              },
              "items": [
                {},
                {}
              ],
              "statement_descriptor": "statement_descriptor1",
              "metadata": {
                "key0": "metadata2",
                "key1": "metadata3",
                "key2": "metadata4"
              },
              "setup": {
                "id": "id5",
                "description": "description5",
                "amount": 143,
                "status": "status3"
              },
              "gateway_affiliation_id": "gateway_affiliation_id7",
              "next_billing_at": "2016-03-13T12:52:32.123Z",
              "billing_day": 195,
              "minimum_price": 81,
              "increments": [
                {}
              ],
              "split": {
                "enabled": true,
                "rules": [
                  {
                    "type": "type9",
                    "amount": 7,
                    "recipient": {
                      "id": "id7",
                      "name": "name7",
                      "email": "email9",
                      "document": "document9",
                      "description": "description3",
                      "type": "type3",
                      "status": "status1",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "updated_at": "2016-03-13T12:52:32.123Z",
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "default_bank_account": {
                        "id": "id5",
                        "holder_name": "holder_name1",
                        "holder_type": "holder_type7",
                        "bank": "bank3",
                        "branch_number": "branch_number1",
                        "branch_check_digit": "branch_check_digit1",
                        "account_number": "account_number5",
                        "account_check_digit": "account_check_digit1",
                        "type": "type5",
                        "status": "status7",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "updated_at": "2016-03-13T12:52:32.123Z",
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "recipient": {},
                        "metadata": {
                          "key0": "metadata2"
                        },
                        "pix_key": "pix_key1"
                      },
                      "gateway_recipients": [
                        {
                          "gateway": "gateway1",
                          "status": "status3",
                          "pgid": "pgid7",
                          "created_at": "created_at9",
                          "updated_at": "updated_at7"
                        }
                      ],
                      "metadata": {
                        "key0": "metadata6"
                      },
                      "automatic_anticipation_settings": {
                        "enabled": true,
                        "type": "type9",
                        "volume_percentage": 55,
                        "delay": 235,
                        "days": [
                          211,
                          212,
                          213
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": true,
                        "transfer_interval": "transfer_interval7",
                        "transfer_day": 221
                      },
                      "code": "code5",
                      "payment_mode": "payment_mode9"
                    },
                    "gateway_id": "gateway_id9",
                    "options": {
                      "liable": true,
                      "charge_processing_fee": true,
                      "charge_remainder_fee": "charge_remainder_fee9"
                    },
                    "id": "id1"
                  },
                  {
                    "type": "type8",
                    "amount": 6,
                    "recipient": {
                      "id": "id6",
                      "name": "name6",
                      "email": "email0",
                      "document": "document0",
                      "description": "description4",
                      "type": "type4",
                      "status": "status2",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "updated_at": "2016-03-13T12:52:32.123Z",
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "default_bank_account": {
                        "id": "id4",
                        "holder_name": "holder_name0",
                        "holder_type": "holder_type6",
                        "bank": "bank2",
                        "branch_number": "branch_number0",
                        "branch_check_digit": "branch_check_digit0",
                        "account_number": "account_number4",
                        "account_check_digit": "account_check_digit0",
                        "type": "type6",
                        "status": "status6",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "updated_at": "2016-03-13T12:52:32.123Z",
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "recipient": {},
                        "metadata": {
                          "key0": "metadata1",
                          "key1": "metadata0"
                        },
                        "pix_key": "pix_key2"
                      },
                      "gateway_recipients": [
                        {
                          "gateway": "gateway0",
                          "status": "status2",
                          "pgid": "pgid6",
                          "created_at": "created_at8",
                          "updated_at": "updated_at6"
                        },
                        {
                          "gateway": "gateway1",
                          "status": "status3",
                          "pgid": "pgid7",
                          "created_at": "created_at9",
                          "updated_at": "updated_at7"
                        },
                        {
                          "gateway": "gateway2",
                          "status": "status4",
                          "pgid": "pgid8",
                          "created_at": "created_at0",
                          "updated_at": "updated_at8"
                        }
                      ],
                      "metadata": {
                        "key0": "metadata7",
                        "key1": "metadata8"
                      },
                      "automatic_anticipation_settings": {
                        "enabled": false,
                        "type": "type0",
                        "volume_percentage": 56,
                        "delay": 234,
                        "days": [
                          210,
                          211
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": false,
                        "transfer_interval": "transfer_interval8",
                        "transfer_day": 222
                      },
                      "code": "code4",
                      "payment_mode": "payment_mode0"
                    },
                    "gateway_id": "gateway_id8",
                    "options": {
                      "liable": false,
                      "charge_processing_fee": false,
                      "charge_remainder_fee": "charge_remainder_fee8"
                    },
                    "id": "id2"
                  }
                ]
              }
            },
            "subscription_item": {}
          }
        ],
        "increments": [
          {},
          {}
        ],
        "subscription": {
          "id": "id1",
          "code": "code9",
          "start_at": "2016-03-13T12:52:32.123Z",
          "interval": "interval9",
          "interval_count": 179,
          "billing_type": "billing_type5",
          "current_cycle": {
            "start_at": "2016-03-13T12:52:32.123Z",
            "end_at": "2016-03-13T12:52:32.123Z",
            "id": "id9",
            "billing_at": "2016-03-13T12:52:32.123Z",
            "subscription": {},
            "status": "status1",
            "duration": 163,
            "created_at": "created_at7",
            "updated_at": "updated_at5",
            "cycle": 103
          },
          "payment_method": "payment_method9",
          "currency": "currency1",
          "installments": 91,
          "status": "status3",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "customer": {
            "id": "id1",
            "name": "name1",
            "email": "email5",
            "delinquent": true,
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "document": "document5",
            "type": "type9",
            "fb_access_token": "fb_access_token5",
            "address": {
              "id": "id7",
              "street": "street7",
              "number": "number5",
              "complement": "complement3",
              "zip_code": "zip_code1",
              "neighborhood": "neighborhood3",
              "city": "city7",
              "state": "state3",
              "country": "country1",
              "status": "status9",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "customer": {},
              "metadata": {
                "key0": "metadata2"
              },
              "line_1": "line_11",
              "line_2": "line_25",
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            "metadata": {
              "key0": "metadata8"
            },
            "phones": {
              "home_phone": {
                "country_code": "country_code3",
                "number": "number1",
                "area_code": "area_code3"
              },
              "mobile_phone": {
                "country_code": "country_code7",
                "number": "number9",
                "area_code": "area_code3"
              }
            },
            "fb_id": 83,
            "code": "code9",
            "document_type": "document_type9"
          },
          "card": {
            "id": "id5",
            "last_four_digits": "last_four_digits1",
            "brand": "brand9",
            "holder_name": "holder_name1",
            "exp_month": 15,
            "exp_year": 25,
            "status": "status3",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "billing_address": {
              "street": "street7",
              "number": "number5",
              "zip_code": "zip_code1",
              "neighborhood": "neighborhood3",
              "city": "city3",
              "state": "state7",
              "country": "country1",
              "complement": "complement7",
              "line_1": "line_19",
              "line_2": "line_25"
            },
            "customer": {
              "id": "id5",
              "name": "name5",
              "email": "email1",
              "delinquent": true,
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "document": "document1",
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
                  "key0": "metadata8",
                  "key1": "metadata7",
                  "key2": "metadata6"
                },
                "line_1": "line_15",
                "line_2": "line_29",
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              "metadata": {
                "key0": "metadata8",
                "key1": "metadata9",
                "key2": "metadata0"
              },
              "phones": {
                "home_phone": {
                  "country_code": "country_code7",
                  "number": "number5",
                  "area_code": "area_code7"
                },
                "mobile_phone": {
                  "country_code": "country_code3",
                  "number": "number9",
                  "area_code": "area_code3"
                }
              },
              "fb_id": 13,
              "code": "code3",
              "document_type": "document_type3"
            },
            "metadata": {
              "key0": "metadata8"
            },
            "type": "type5",
            "holder_document": "holder_document1",
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "first_six_digits": "first_six_digits5",
            "label": "label5"
          },
          "items": [
            {},
            {},
            {}
          ],
          "statement_descriptor": "statement_descriptor1",
          "metadata": {
            "key0": "metadata8",
            "key1": "metadata7"
          },
          "setup": {
            "id": "id5",
            "description": "description5",
            "amount": 115,
            "status": "status7"
          },
          "gateway_affiliation_id": "gateway_affiliation_id7",
          "next_billing_at": "2016-03-13T12:52:32.123Z",
          "billing_day": 63,
          "minimum_price": 177,
          "increments": [
            {},
            {}
          ],
          "split": {
            "enabled": true,
            "rules": [
              {
                "type": "type9",
                "amount": 251,
                "recipient": {
                  "id": "id3",
                  "name": "name3",
                  "email": "email3",
                  "document": "document3",
                  "description": "description3",
                  "type": "type7",
                  "status": "status5",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "default_bank_account": {
                    "id": "id1",
                    "holder_name": "holder_name7",
                    "holder_type": "holder_type3",
                    "bank": "bank9",
                    "branch_number": "branch_number7",
                    "branch_check_digit": "branch_check_digit7",
                    "account_number": "account_number1",
                    "account_check_digit": "account_check_digit7",
                    "type": "type1",
                    "status": "status3",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "updated_at": "2016-03-13T12:52:32.123Z",
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "recipient": {},
                    "metadata": {
                      "key0": "metadata8"
                    },
                    "pix_key": "pix_key5"
                  },
                  "gateway_recipients": [
                    {
                      "gateway": "gateway7",
                      "status": "status9",
                      "pgid": "pgid3",
                      "created_at": "created_at5",
                      "updated_at": "updated_at3"
                    }
                  ],
                  "metadata": {
                    "key0": "metadata0"
                  },
                  "automatic_anticipation_settings": {
                    "enabled": true,
                    "type": "type3",
                    "volume_percentage": 143,
                    "delay": 147,
                    "days": [
                      123,
                      124,
                      125
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": true,
                    "transfer_interval": "transfer_interval1",
                    "transfer_day": 53
                  },
                  "code": "code1",
                  "payment_mode": "payment_mode7"
                },
                "gateway_id": "gateway_id9",
                "options": {
                  "liable": true,
                  "charge_processing_fee": true,
                  "charge_remainder_fee": "charge_remainder_fee7"
                },
                "id": "id1"
              },
              {
                "type": "type8",
                "amount": 252,
                "recipient": {
                  "id": "id4",
                  "name": "name4",
                  "email": "email2",
                  "document": "document2",
                  "description": "description4",
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
                    "type": "type2",
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
                    "volume_percentage": 142,
                    "delay": 148,
                    "days": [
                      124
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": false,
                    "transfer_interval": "transfer_interval0",
                    "transfer_day": 52
                  },
                  "code": "code2",
                  "payment_mode": "payment_mode8"
                },
                "gateway_id": "gateway_id8",
                "options": {
                  "liable": false,
                  "charge_processing_fee": false,
                  "charge_remainder_fee": "charge_remainder_fee8"
                },
                "id": "id2"
              },
              {
                "type": "type7",
                "amount": 253,
                "recipient": {
                  "id": "id5",
                  "name": "name5",
                  "email": "email1",
                  "document": "document1",
                  "description": "description5",
                  "type": "type5",
                  "status": "status3",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "default_bank_account": {
                    "id": "id3",
                    "holder_name": "holder_name9",
                    "holder_type": "holder_type5",
                    "bank": "bank1",
                    "branch_number": "branch_number9",
                    "branch_check_digit": "branch_check_digit9",
                    "account_number": "account_number3",
                    "account_check_digit": "account_check_digit9",
                    "type": "type3",
                    "status": "status5",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "updated_at": "2016-03-13T12:52:32.123Z",
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "recipient": {},
                    "metadata": {
                      "key0": "metadata0",
                      "key1": "metadata9"
                    },
                    "pix_key": "pix_key3"
                  },
                  "gateway_recipients": [
                    {
                      "gateway": "gateway9",
                      "status": "status1",
                      "pgid": "pgid5",
                      "created_at": "created_at7",
                      "updated_at": "updated_at5"
                    },
                    {
                      "gateway": "gateway0",
                      "status": "status2",
                      "pgid": "pgid6",
                      "created_at": "created_at8",
                      "updated_at": "updated_at6"
                    },
                    {
                      "gateway": "gateway1",
                      "status": "status3",
                      "pgid": "pgid7",
                      "created_at": "created_at9",
                      "updated_at": "updated_at7"
                    }
                  ],
                  "metadata": {
                    "key0": "metadata8",
                    "key1": "metadata9"
                  },
                  "automatic_anticipation_settings": {
                    "enabled": true,
                    "type": "type1",
                    "volume_percentage": 141,
                    "delay": 149,
                    "days": [
                      125,
                      126
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": true,
                    "transfer_interval": "transfer_interval9",
                    "transfer_day": 51
                  },
                  "code": "code3",
                  "payment_mode": "payment_mode9"
                },
                "gateway_id": "gateway_id7",
                "options": {
                  "liable": true,
                  "charge_processing_fee": true,
                  "charge_remainder_fee": "charge_remainder_fee9"
                },
                "id": "id3"
              }
            ]
          }
        },
        "name": "name3",
        "quantity": 123,
        "cycles": 113,
        "deleted_at": "2016-03-13T12:52:32.123Z"
      }
    }
  ],
  "subscription": {
    "id": "id4",
    "code": "code2",
    "start_at": "2016-03-13T12:52:32.123Z",
    "interval": "interval2",
    "interval_count": 234,
    "billing_type": "billing_type8",
    "payment_method": "payment_method4",
    "currency": "currency4",
    "installments": 146,
    "status": "status6",
    "created_at": "2016-03-13T12:52:32.123Z",
    "updated_at": "2016-03-13T12:52:32.123Z",
    "card": {
      "id": "id8",
      "last_four_digits": "last_four_digits4",
      "brand": "brand2",
      "holder_name": "holder_name4",
      "exp_month": 216,
      "exp_year": 80,
      "status": "status0",
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "billing_address": {
        "street": "street0",
        "number": "number2",
        "zip_code": "zip_code4",
        "neighborhood": "neighborhood6",
        "city": "city0",
        "state": "state4",
        "country": "country4",
        "complement": "complement4",
        "line_1": "line_16",
        "line_2": "line_28"
      },
      "customer": {
        "id": "id8",
        "name": "name8",
        "email": "email8",
        "delinquent": false,
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "document": "document8",
        "type": "type2",
        "fb_access_token": "fb_access_token2",
        "address": {
          "id": "id4",
          "street": "street4",
          "number": "number2",
          "complement": "complement0",
          "zip_code": "zip_code8",
          "neighborhood": "neighborhood0",
          "city": "city4",
          "state": "state0",
          "country": "country8",
          "status": "status6",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "customer": {},
          "metadata": {
            "key0": "metadata5"
          },
          "line_1": "line_18",
          "line_2": "line_22",
          "deleted_at": "2016-03-13T12:52:32.123Z"
        },
        "metadata": {
          "key0": "metadata5"
        },
        "phones": {
          "home_phone": {
            "country_code": "country_code0",
            "number": "number2",
            "area_code": "area_code0"
          },
          "mobile_phone": {
            "country_code": "country_code0",
            "number": "number2",
            "area_code": "area_code0"
          }
        },
        "fb_id": 68,
        "code": "code6",
        "document_type": "document_type6"
      },
      "metadata": {
        "key0": "metadata5",
        "key1": "metadata4"
      },
      "type": "type2",
      "holder_document": "holder_document2",
      "deleted_at": "2016-03-13T12:52:32.123Z",
      "first_six_digits": "first_six_digits8",
      "label": "label8"
    },
    "items": [
      {
        "id": "id1",
        "description": "description1",
        "status": "status3",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "pricing_scheme": {
          "price": 243,
          "scheme_type": "scheme_type3",
          "price_brackets": [
            {
              "start_quantity": 222,
              "price": 96,
              "end_quantity": 230,
              "overage_price": 244
            }
          ],
          "minimum_price": 147,
          "percentage": 65.91
        },
        "discounts": [
          {
            "id": "id2",
            "value": 63.54,
            "discount_type": "discount_type0",
            "status": "status4",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 234,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description2",
            "subscription": {},
            "subscription_item": {}
          }
        ],
        "increments": [
          {
            "id": "id0",
            "value": 174.82,
            "increment_type": "increment_type2",
            "status": "status2",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 158,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description0",
            "subscription": {},
            "subscription_item": {}
          }
        ],
        "subscription": {},
        "name": "name1",
        "quantity": 23,
        "cycles": 43,
        "deleted_at": "2016-03-13T12:52:32.123Z"
      }
    ],
    "statement_descriptor": "statement_descriptor4",
    "metadata": {
      "key0": "metadata1",
      "key1": "metadata0",
      "key2": "metadata9"
    },
    "setup": {
      "id": "id8",
      "description": "description8",
      "amount": 170,
      "status": "status0"
    },
    "gateway_affiliation_id": "gateway_affiliation_id0",
    "increments": [
      {
        "id": "id3",
        "value": 204.95,
        "increment_type": "increment_type5",
        "status": "status5",
        "created_at": "2016-03-13T12:52:32.123Z",
        "cycles": 217,
        "deleted_at": "2016-03-13T12:52:32.123Z",
        "description": "description3",
        "subscription": {},
        "subscription_item": {
          "id": "id9",
          "description": "description9",
          "status": "status9",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "pricing_scheme": {
            "price": 29,
            "scheme_type": "scheme_type9",
            "price_brackets": [
              {
                "start_quantity": 180,
                "price": 138,
                "end_quantity": 188,
                "overage_price": 202
              }
            ],
            "minimum_price": 189,
            "percentage": 21.97
          },
          "discounts": [
            {
              "id": "id0",
              "value": 240.02,
              "discount_type": "discount_type8",
              "status": "status2",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 218,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description0",
              "subscription": {},
              "subscription_item": {}
            },
            {
              "id": "id1",
              "value": 240.03,
              "discount_type": "discount_type9",
              "status": "status3",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 219,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description1",
              "subscription": {},
              "subscription_item": {}
            }
          ],
          "increments": [
            {},
            {}
          ],
          "subscription": {},
          "name": "name9",
          "quantity": 7,
          "cycles": 229,
          "deleted_at": "2016-03-13T12:52:32.123Z"
        }
      }
    ],
    "split": {
      "enabled": false,
      "rules": [
        {
          "type": "type4",
          "amount": 50,
          "recipient": {
            "id": "id6",
            "name": "name6",
            "email": "email0",
            "document": "document0",
            "description": "description6",
            "type": "type4",
            "status": "status8",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "default_bank_account": {
              "id": "id4",
              "holder_name": "holder_name0",
              "holder_type": "holder_type6",
              "bank": "bank2",
              "branch_number": "branch_number0",
              "branch_check_digit": "branch_check_digit0",
              "account_number": "account_number4",
              "account_check_digit": "account_check_digit0",
              "type": "type4",
              "status": "status6",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "recipient": {},
              "metadata": {
                "key0": "metadata5",
                "key1": "metadata4",
                "key2": "metadata3"
              },
              "pix_key": "pix_key8"
            },
            "gateway_recipients": [
              {
                "gateway": "gateway0",
                "status": "status2",
                "pgid": "pgid6",
                "created_at": "created_at8",
                "updated_at": "updated_at6"
              },
              {
                "gateway": "gateway1",
                "status": "status3",
                "pgid": "pgid7",
                "created_at": "created_at9",
                "updated_at": "updated_at7"
              },
              {
                "gateway": "gateway2",
                "status": "status4",
                "pgid": "pgid8",
                "created_at": "created_at0",
                "updated_at": "updated_at8"
              }
            ],
            "metadata": {
              "key0": "metadata3",
              "key1": "metadata2"
            },
            "automatic_anticipation_settings": {
              "enabled": false,
              "type": "type0",
              "volume_percentage": 88,
              "delay": 202,
              "days": [
                178,
                179
              ]
            },
            "transfer_settings": {
              "transfer_enabled": false,
              "transfer_interval": "transfer_interval8",
              "transfer_day": 254
            },
            "code": "code4",
            "payment_mode": "payment_mode0"
          },
          "gateway_id": "gateway_id4",
          "options": {
            "liable": false,
            "charge_processing_fee": false,
            "charge_remainder_fee": "charge_remainder_fee0"
          },
          "id": "id4"
        }
      ]
    },
    "boleto": {
      "interest": {
        "days": 2,
        "type": "percentage",
        "amount": 20
      },
      "fine": {
        "days": 2,
        "type": "flat",
        "amount": 10
      },
      "max_days_to_pay_past_due": 2
    },
    "current_cycle": {
      "start_at": "2016-03-13T12:52:32.123Z",
      "end_at": "2016-03-13T12:52:32.123Z",
      "id": "id2",
      "billing_at": "2016-03-13T12:52:32.123Z",
      "subscription": {},
      "status": "status4",
      "duration": 218,
      "created_at": "created_at0",
      "updated_at": "updated_at8",
      "cycle": 158
    },
    "customer": {
      "id": "id4",
      "name": "name4",
      "email": "email2",
      "delinquent": false,
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "document": "document8",
      "type": "type6",
      "fb_access_token": "fb_access_token8",
      "address": {
        "id": "id0",
        "street": "street0",
        "number": "number8",
        "complement": "complement6",
        "zip_code": "zip_code4",
        "neighborhood": "neighborhood6",
        "city": "city0",
        "state": "state6",
        "country": "country4",
        "status": "status2",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "customer": {},
        "metadata": {
          "key0": "metadata9",
          "key1": "metadata8",
          "key2": "metadata7"
        },
        "line_1": "line_14",
        "line_2": "line_28",
        "deleted_at": "2016-03-13T12:52:32.123Z"
      },
      "metadata": {
        "key0": "metadata1",
        "key1": "metadata0"
      },
      "phones": {
        "home_phone": {
          "country_code": "country_code6",
          "number": "number4",
          "area_code": "area_code6"
        },
        "mobile_phone": {
          "country_code": "country_code6",
          "number": "number6",
          "area_code": "area_code6"
        }
      },
      "fb_id": 138,
      "code": "code2",
      "document_type": "document_type2"
    },
    "next_billing_at": "2016-03-13T12:52:32.123Z",
    "billing_day": 118,
    "minimum_price": 232
  },
  "name": "name0",
  "quantity": 68,
  "cycles": 168,
  "deleted_at": "2016-03-13T12:52:32.123Z"
}
```

