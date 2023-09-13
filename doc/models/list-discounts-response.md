
# List Discounts Response

## Structure

`ListDiscountsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetDiscountResponse`](../../doc/models/get-discount-response.md) | Required | The Discounts response |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "value": 145.67,
      "discount_type": "discount_type3",
      "status": "status7",
      "created_at": "2016-03-13T12:52:32.123Z",
      "cycles": 255,
      "deleted_at": "2016-03-13T12:52:32.123Z",
      "description": "description5",
      "subscription": {
        "id": "id1",
        "code": "code9",
        "start_at": "2016-03-13T12:52:32.123Z",
        "interval": "interval9",
        "interval_count": 97,
        "billing_type": "billing_type5",
        "current_cycle": {
          "start_at": "2016-03-13T12:52:32.123Z",
          "end_at": "2016-03-13T12:52:32.123Z",
          "id": "id9",
          "billing_at": "2016-03-13T12:52:32.123Z",
          "subscription": {},
          "status": "status1",
          "duration": 81,
          "created_at": "created_at7",
          "updated_at": "updated_at5",
          "cycle": 21
        },
        "payment_method": "payment_method9",
        "currency": "currency1",
        "installments": 9,
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
              "key0": "metadata4",
              "key1": "metadata3"
            },
            "line_1": "line_11",
            "line_2": "line_25",
            "deleted_at": "2016-03-13T12:52:32.123Z"
          },
          "metadata": {
            "key0": "metadata2",
            "key1": "metadata3"
          },
          "phones": {
            "home_phone": {
              "country_code": "country_code3",
              "number": "number9",
              "area_code": "area_code3"
            },
            "mobile_phone": {
              "country_code": "country_code3",
              "number": "number5",
              "area_code": "area_code7"
            }
          },
          "fb_id": 1,
          "code": "code9",
          "document_type": "document_type9"
        },
        "card": {
          "id": "id5",
          "last_four_digits": "last_four_digits1",
          "brand": "brand9",
          "holder_name": "holder_name1",
          "exp_month": 97,
          "exp_year": 199,
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
                "key1": "metadata7"
              },
              "line_1": "line_15",
              "line_2": "line_29",
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            "metadata": {
              "key0": "metadata8",
              "key1": "metadata9"
            },
            "phones": {
              "home_phone": {
                "country_code": "country_code3",
                "number": "number5",
                "area_code": "area_code7"
              },
              "mobile_phone": {
                "country_code": "country_code3",
                "number": "number9",
                "area_code": "area_code3"
              }
            },
            "fb_id": 229,
            "code": "code3",
            "document_type": "document_type3"
          },
          "metadata": {
            "key0": "metadata8",
            "key1": "metadata9"
          },
          "type": "type5",
          "holder_document": "holder_document1",
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "first_six_digits": "first_six_digits5",
          "label": "label5"
        },
        "items": [
          {
            "id": "id8",
            "description": "description8",
            "status": "status0",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "pricing_scheme": {
              "price": 124,
              "scheme_type": "scheme_type0",
              "price_brackets": [
                {
                  "start_quantity": 85,
                  "price": 233,
                  "end_quantity": 93,
                  "overage_price": 107
                },
                {
                  "start_quantity": 86,
                  "price": 232,
                  "end_quantity": 94,
                  "overage_price": 108
                },
                {
                  "start_quantity": 87,
                  "price": 231,
                  "end_quantity": 95,
                  "overage_price": 109
                }
              ],
              "minimum_price": 28,
              "percentage": 49.18
            },
            "discounts": [
              {},
              {},
              {}
            ],
            "increments": [
              {
                "id": "id7",
                "value": 158.09,
                "increment_type": "increment_type9",
                "status": "status1",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 39,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description3",
                "subscription": {},
                "subscription_item": {}
              }
            ],
            "subscription": {},
            "name": "name8",
            "quantity": 142,
            "cycles": 162,
            "deleted_at": "2016-03-13T12:52:32.123Z"
          },
          {
            "id": "id9",
            "description": "description9",
            "status": "status1",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "pricing_scheme": {
              "price": 123,
              "scheme_type": "scheme_type1",
              "price_brackets": [
                {
                  "start_quantity": 86,
                  "price": 232,
                  "end_quantity": 94,
                  "overage_price": 108
                }
              ],
              "minimum_price": 27,
              "percentage": 49.19
            },
            "discounts": [
              {}
            ],
            "increments": [
              {
                "id": "id8",
                "value": 158.1,
                "increment_type": "increment_type0",
                "status": "status0",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 38,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description2",
                "subscription": {},
                "subscription_item": {}
              },
              {
                "id": "id9",
                "value": 158.11,
                "increment_type": "increment_type1",
                "status": "status9",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 37,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description1",
                "subscription": {},
                "subscription_item": {}
              }
            ],
            "subscription": {},
            "name": "name9",
            "quantity": 143,
            "cycles": 163,
            "deleted_at": "2016-03-13T12:52:32.123Z"
          }
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
          "amount": 33,
          "status": "status7"
        },
        "gateway_affiliation_id": "gateway_affiliation_id7",
        "next_billing_at": "2016-03-13T12:52:32.123Z",
        "billing_day": 237,
        "minimum_price": 161,
        "increments": [
          {
            "id": "id0",
            "value": 58.82,
            "increment_type": "increment_type2",
            "status": "status8",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 238,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description0",
            "subscription": {},
            "subscription_item": {
              "id": "id6",
              "description": "description4",
              "status": "status2",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "pricing_scheme": {
                "price": 8,
                "scheme_type": "scheme_type2",
                "price_brackets": [
                  {
                    "start_quantity": 201,
                    "price": 139,
                    "end_quantity": 209,
                    "overage_price": 223
                  }
                ],
                "minimum_price": 88,
                "percentage": 168.1
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
              "name": "name6",
              "quantity": 242,
              "cycles": 250,
              "deleted_at": "2016-03-13T12:52:32.123Z"
            }
          }
        ],
        "split": {
          "enabled": true,
          "rules": [
            {
              "type": "type9",
              "amount": 169,
              "recipient": {
                "id": "id3",
                "name": "name3",
                "email": "email3",
                "document": "document3",
                "description": "description7",
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
                  "type": "type9",
                  "status": "status3",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "recipient": {},
                  "metadata": {
                    "key0": "metadata8",
                    "key1": "metadata7"
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
                  },
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
                  "key0": "metadata0",
                  "key1": "metadata1"
                },
                "automatic_anticipation_settings": {
                  "enabled": true,
                  "type": "type3",
                  "volume_percentage": 225,
                  "delay": 65,
                  "days": [
                    41,
                    42
                  ]
                },
                "transfer_settings": {
                  "transfer_enabled": true,
                  "transfer_interval": "transfer_interval1",
                  "transfer_day": 135
                },
                "code": "code1",
                "payment_mode": "payment_mode3"
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
              "amount": 170,
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
                    "key0": "metadata9"
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
                  }
                ],
                "metadata": {
                  "key0": "metadata9"
                },
                "automatic_anticipation_settings": {
                  "enabled": false,
                  "type": "type2",
                  "volume_percentage": 224,
                  "delay": 66,
                  "days": [
                    42,
                    43,
                    44
                  ]
                },
                "transfer_settings": {
                  "transfer_enabled": false,
                  "transfer_interval": "transfer_interval0",
                  "transfer_day": 134
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
          ]
        }
      },
      "subscription_item": {
        "id": "id1",
        "description": "description1",
        "status": "status3",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "pricing_scheme": {
          "price": 43,
          "scheme_type": "scheme_type3",
          "price_brackets": [
            {
              "start_quantity": 166,
              "price": 104,
              "end_quantity": 174,
              "overage_price": 188
            },
            {
              "start_quantity": 167,
              "price": 105,
              "end_quantity": 175,
              "overage_price": 189
            }
          ],
          "minimum_price": 53,
          "percentage": 183.11
        },
        "discounts": [
          {},
          {}
        ],
        "increments": [
          {
            "id": "id0",
            "value": 36.02,
            "increment_type": "increment_type2",
            "status": "status2",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 42,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description0",
            "subscription": {
              "id": "id4",
              "code": "code2",
              "start_at": "2016-03-13T12:52:32.123Z",
              "interval": "interval8",
              "interval_count": 24,
              "billing_type": "billing_type2",
              "current_cycle": {
                "start_at": "2016-03-13T12:52:32.123Z",
                "end_at": "2016-03-13T12:52:32.123Z",
                "id": "id2",
                "billing_at": "2016-03-13T12:52:32.123Z",
                "subscription": {},
                "status": "status6",
                "duration": 8,
                "created_at": "created_at0",
                "updated_at": "updated_at2",
                "cycle": 204
              },
              "payment_method": "payment_method6",
              "currency": "currency6",
              "installments": 192,
              "status": "status4",
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
                "fb_id": 232,
                "code": "code4",
                "document_type": "document_type4"
              },
              "card": {
                "id": "id2",
                "last_four_digits": "last_four_digits8",
                "brand": "brand6",
                "holder_name": "holder_name8",
                "exp_month": 170,
                "exp_year": 126,
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
                      "key0": "metadata5"
                    },
                    "line_1": "line_12",
                    "line_2": "line_26",
                    "deleted_at": "2016-03-13T12:52:32.123Z"
                  },
                  "metadata": {
                    "key0": "metadata1"
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
                  "fb_id": 114,
                  "code": "code0",
                  "document_type": "document_type0"
                },
                "metadata": {
                  "key0": "metadata1",
                  "key1": "metadata2"
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
              "statement_descriptor": "statement_descriptor4",
              "metadata": {
                "key0": "metadata9",
                "key1": "metadata0"
              },
              "setup": {
                "id": "id8",
                "description": "description2",
                "amount": 40,
                "status": "status0"
              },
              "gateway_affiliation_id": "gateway_affiliation_id0",
              "next_billing_at": "2016-03-13T12:52:32.123Z",
              "billing_day": 92,
              "minimum_price": 234,
              "increments": [
                {},
                {},
                {}
              ],
              "split": {
                "enabled": false,
                "rules": [
                  {
                    "type": "type8",
                    "amount": 232,
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
                        "volume_percentage": 86,
                        "delay": 204,
                        "days": [
                          180,
                          181
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": false,
                        "transfer_interval": "transfer_interval8",
                        "transfer_day": 252
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
                    "amount": 233,
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
                        "volume_percentage": 85,
                        "delay": 205,
                        "days": [
                          181,
                          182,
                          183
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": true,
                        "transfer_interval": "transfer_interval7",
                        "transfer_day": 251
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
            "subscription_item": {}
          },
          {
            "id": "id1",
            "value": 36.03,
            "increment_type": "increment_type3",
            "status": "status3",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 43,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description1",
            "subscription": {
              "id": "id3",
              "code": "code1",
              "start_at": "2016-03-13T12:52:32.123Z",
              "interval": "interval9",
              "interval_count": 23,
              "billing_type": "billing_type3",
              "current_cycle": {
                "start_at": "2016-03-13T12:52:32.123Z",
                "end_at": "2016-03-13T12:52:32.123Z",
                "id": "id1",
                "billing_at": "2016-03-13T12:52:32.123Z",
                "subscription": {},
                "status": "status7",
                "duration": 7,
                "created_at": "created_at9",
                "updated_at": "updated_at3",
                "cycle": 203
              },
              "payment_method": "payment_method7",
              "currency": "currency7",
              "installments": 191,
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
                    "country_code": "country_code9",
                    "number": "number3",
                    "area_code": "area_code9"
                  },
                  "mobile_phone": {
                    "country_code": "country_code1",
                    "number": "number1",
                    "area_code": "area_code1"
                  }
                },
                "fb_id": 233,
                "code": "code5",
                "document_type": "document_type5"
              },
              "card": {
                "id": "id3",
                "last_four_digits": "last_four_digits9",
                "brand": "brand7",
                "holder_name": "holder_name9",
                "exp_month": 169,
                "exp_year": 127,
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
                    "line_1": "line_13",
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
                  "fb_id": 115,
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
                {},
                {}
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
                "amount": 41,
                "status": "status1"
              },
              "gateway_affiliation_id": "gateway_affiliation_id9",
              "next_billing_at": "2016-03-13T12:52:32.123Z",
              "billing_day": 93,
              "minimum_price": 235,
              "increments": [
                {}
              ],
              "split": {
                "enabled": true,
                "rules": [
                  {
                    "type": "type9",
                    "amount": 233,
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
                        "volume_percentage": 85,
                        "delay": 205,
                        "days": [
                          181,
                          182,
                          183
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": true,
                        "transfer_interval": "transfer_interval7",
                        "transfer_day": 251
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
                    "type": "type0",
                    "amount": 234,
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
                        "volume_percentage": 84,
                        "delay": 206,
                        "days": [
                          182
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": false,
                        "transfer_interval": "transfer_interval6",
                        "transfer_day": 250
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
                  },
                  {
                    "type": "type1",
                    "amount": 235,
                    "recipient": {
                      "id": "id9",
                      "name": "name9",
                      "email": "email7",
                      "document": "document7",
                      "description": "description1",
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
                        "type": "type3",
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
                        "volume_percentage": 83,
                        "delay": 207,
                        "days": [
                          183,
                          184
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": true,
                        "transfer_interval": "transfer_interval5",
                        "transfer_day": 249
                      },
                      "code": "code7",
                      "payment_mode": "payment_mode7"
                    },
                    "gateway_id": "gateway_id1",
                    "options": {
                      "liable": true,
                      "charge_processing_fee": true,
                      "charge_remainder_fee": "charge_remainder_fee1"
                    },
                    "id": "id9"
                  }
                ]
              }
            },
            "subscription_item": {}
          },
          {
            "id": "id2",
            "value": 36.04,
            "increment_type": "increment_type4",
            "status": "status4",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 44,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description2",
            "subscription": {
              "id": "id2",
              "code": "code0",
              "start_at": "2016-03-13T12:52:32.123Z",
              "interval": "interval0",
              "interval_count": 22,
              "billing_type": "billing_type4",
              "current_cycle": {
                "start_at": "2016-03-13T12:52:32.123Z",
                "end_at": "2016-03-13T12:52:32.123Z",
                "id": "id0",
                "billing_at": "2016-03-13T12:52:32.123Z",
                "subscription": {},
                "status": "status8",
                "duration": 6,
                "created_at": "created_at8",
                "updated_at": "updated_at4",
                "cycle": 202
              },
              "payment_method": "payment_method8",
              "currency": "currency8",
              "installments": 190,
              "status": "status6",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
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
                    "key0": "metadata1",
                    "key1": "metadata0",
                    "key2": "metadata9"
                  },
                  "line_1": "line_12",
                  "line_2": "line_22",
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                },
                "metadata": {
                  "key0": "metadata5",
                  "key1": "metadata6",
                  "key2": "metadata7"
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
                "fb_id": 234,
                "code": "code6",
                "document_type": "document_type6"
              },
              "card": {
                "id": "id4",
                "last_four_digits": "last_four_digits0",
                "brand": "brand8",
                "holder_name": "holder_name0",
                "exp_month": 168,
                "exp_year": 128,
                "status": "status4",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "billing_address": {
                  "street": "street6",
                  "number": "number6",
                  "zip_code": "zip_code0",
                  "neighborhood": "neighborhood2",
                  "city": "city4",
                  "state": "state8",
                  "country": "country0",
                  "complement": "complement8",
                  "line_1": "line_10",
                  "line_2": "line_24"
                },
                "customer": {
                  "id": "id4",
                  "name": "name4",
                  "email": "email2",
                  "delinquent": false,
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "document": "document2",
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
                      "key0": "metadata7",
                      "key1": "metadata6"
                    },
                    "line_1": "line_14",
                    "line_2": "line_28",
                    "deleted_at": "2016-03-13T12:52:32.123Z"
                  },
                  "metadata": {
                    "key0": "metadata9",
                    "key1": "metadata0"
                  },
                  "phones": {
                    "home_phone": {
                      "country_code": "country_code6",
                      "number": "number6",
                      "area_code": "area_code6"
                    },
                    "mobile_phone": {
                      "country_code": "country_code4",
                      "number": "number8",
                      "area_code": "area_code4"
                    }
                  },
                  "fb_id": 116,
                  "code": "code2",
                  "document_type": "document_type2"
                },
                "metadata": {
                  "key0": "metadata9",
                  "key1": "metadata0",
                  "key2": "metadata1"
                },
                "type": "type6",
                "holder_document": "holder_document2",
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "first_six_digits": "first_six_digits4",
                "label": "label4"
              },
              "items": [
                {}
              ],
              "statement_descriptor": "statement_descriptor2",
              "metadata": {
                "key0": "metadata1"
              },
              "setup": {
                "id": "id6",
                "description": "description4",
                "amount": 42,
                "status": "status2"
              },
              "gateway_affiliation_id": "gateway_affiliation_id8",
              "next_billing_at": "2016-03-13T12:52:32.123Z",
              "billing_day": 94,
              "minimum_price": 236,
              "increments": [
                {},
                {}
              ],
              "split": {
                "enabled": false,
                "rules": [
                  {
                    "type": "type0",
                    "amount": 234,
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
                        "volume_percentage": 84,
                        "delay": 206,
                        "days": [
                          182
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": false,
                        "transfer_interval": "transfer_interval6",
                        "transfer_day": 250
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
          }
        ],
        "subscription": {
          "id": "id7",
          "code": "code5",
          "start_at": "2016-03-13T12:52:32.123Z",
          "interval": "interval5",
          "interval_count": 85,
          "billing_type": "billing_type9",
          "current_cycle": {
            "start_at": "2016-03-13T12:52:32.123Z",
            "end_at": "2016-03-13T12:52:32.123Z",
            "id": "id5",
            "billing_at": "2016-03-13T12:52:32.123Z",
            "subscription": {},
            "status": "status7",
            "duration": 69,
            "created_at": "created_at3",
            "updated_at": "updated_at1",
            "cycle": 9
          },
          "payment_method": "payment_method3",
          "currency": "currency7",
          "installments": 253,
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
                "key0": "metadata0",
                "key1": "metadata9",
                "key2": "metadata8"
              },
              "line_1": "line_17",
              "line_2": "line_21",
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            "metadata": {
              "key0": "metadata6",
              "key1": "metadata7",
              "key2": "metadata8"
            },
            "phones": {
              "home_phone": {
                "country_code": "country_code9",
                "number": "number3",
                "area_code": "area_code9"
              },
              "mobile_phone": {
                "country_code": "country_code9",
                "number": "number1",
                "area_code": "area_code1"
              }
            },
            "fb_id": 245,
            "code": "code5",
            "document_type": "document_type5"
          },
          "card": {
            "id": "id1",
            "last_four_digits": "last_four_digits7",
            "brand": "brand5",
            "holder_name": "holder_name7",
            "exp_month": 109,
            "exp_year": 187,
            "status": "status7",
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
                  "key0": "metadata4",
                  "key1": "metadata3"
                },
                "line_1": "line_19",
                "line_2": "line_25",
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              "metadata": {
                "key0": "metadata2",
                "key1": "metadata3"
              },
              "phones": {
                "home_phone": {
                  "country_code": "country_code7",
                  "number": "number9",
                  "area_code": "area_code3"
                },
                "mobile_phone": {
                  "country_code": "country_code7",
                  "number": "number5",
                  "area_code": "area_code7"
                }
              },
              "fb_id": 175,
              "code": "code9",
              "document_type": "document_type9"
            },
            "metadata": {
              "key0": "metadata2",
              "key1": "metadata3",
              "key2": "metadata4"
            },
            "type": "type9",
            "holder_document": "holder_document5",
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "first_six_digits": "first_six_digits1",
            "label": "label1"
          },
          "items": [
            {}
          ],
          "statement_descriptor": "statement_descriptor7",
          "metadata": {
            "key0": "metadata4"
          },
          "setup": {
            "id": "id1",
            "description": "description1",
            "amount": 21,
            "status": "status3"
          },
          "gateway_affiliation_id": "gateway_affiliation_id3",
          "next_billing_at": "2016-03-13T12:52:32.123Z",
          "billing_day": 225,
          "minimum_price": 173,
          "increments": [
            {
              "id": "id4",
              "value": 38.46,
              "increment_type": "increment_type6",
              "status": "status4",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 226,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description6",
              "subscription": {},
              "subscription_item": {}
            },
            {
              "id": "id3",
              "value": 38.45,
              "increment_type": "increment_type5",
              "status": "status5",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 227,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description7",
              "subscription": {},
              "subscription_item": {}
            },
            {
              "id": "id2",
              "value": 38.44,
              "increment_type": "increment_type4",
              "status": "status6",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 228,
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
                "type": "type3",
                "amount": 157,
                "recipient": {
                  "id": "id9",
                  "name": "name9",
                  "email": "email7",
                  "document": "document7",
                  "description": "description1",
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
                    "type": "type3",
                    "status": "status9",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "updated_at": "2016-03-13T12:52:32.123Z",
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "recipient": {},
                    "metadata": {
                      "key0": "metadata4",
                      "key1": "metadata3",
                      "key2": "metadata2"
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
                    }
                  ],
                  "metadata": {
                    "key0": "metadata4",
                    "key1": "metadata5",
                    "key2": "metadata6"
                  },
                  "automatic_anticipation_settings": {
                    "enabled": true,
                    "type": "type7",
                    "volume_percentage": 237,
                    "delay": 53,
                    "days": [
                      29
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": true,
                    "transfer_interval": "transfer_interval5",
                    "transfer_day": 147
                  },
                  "code": "code7",
                  "payment_mode": "payment_mode7"
                },
                "gateway_id": "gateway_id3",
                "options": {
                  "liable": true,
                  "charge_processing_fee": true,
                  "charge_remainder_fee": "charge_remainder_fee3"
                },
                "id": "id7"
              }
            ]
          }
        },
        "name": "name1",
        "quantity": 223,
        "cycles": 243,
        "deleted_at": "2016-03-13T12:52:32.123Z"
      }
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

