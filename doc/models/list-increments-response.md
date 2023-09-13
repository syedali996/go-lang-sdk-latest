
# List Increments Response

## Structure

`ListIncrementsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetIncrementResponse`](../../doc/models/get-increment-response.md) | Required | The Increments response |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "value": 145.67,
      "increment_type": "increment_type7",
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
              {
                "id": "id9",
                "value": 46.81,
                "discount_type": "discount_type7",
                "status": "status1",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 97,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description9",
                "subscription": {},
                "subscription_item": {}
              },
              {
                "id": "id0",
                "value": 46.82,
                "discount_type": "discount_type8",
                "status": "status2",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 98,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description0",
                "subscription": {},
                "subscription_item": {}
              },
              {
                "id": "id1",
                "value": 46.83,
                "discount_type": "discount_type9",
                "status": "status3",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 99,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description1",
                "subscription": {},
                "subscription_item": {}
              }
            ],
            "increments": [
              {}
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
              {
                "id": "id0",
                "value": 46.82,
                "discount_type": "discount_type8",
                "status": "status2",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 98,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description0",
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
          {}
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
          {
            "id": "id2",
            "value": 180.74,
            "discount_type": "discount_type0",
            "status": "status4",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 178,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description2",
            "subscription": {
              "id": "id8",
              "code": "code6",
              "start_at": "2016-03-13T12:52:32.123Z",
              "interval": "interval6",
              "interval_count": 20,
              "billing_type": "billing_type2",
              "current_cycle": {
                "start_at": "2016-03-13T12:52:32.123Z",
                "end_at": "2016-03-13T12:52:32.123Z",
                "id": "id6",
                "billing_at": "2016-03-13T12:52:32.123Z",
                "subscription": {},
                "status": "status8",
                "duration": 4,
                "created_at": "created_at4",
                "updated_at": "updated_at2",
                "cycle": 200
              },
              "payment_method": "payment_method8",
              "currency": "currency8",
              "installments": 188,
              "status": "status0",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "customer": {
                "id": "id8",
                "name": "name8",
                "email": "email2",
                "delinquent": false,
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "document": "document2",
                "type": "type8",
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
                    "key0": "metadata5",
                    "key1": "metadata6"
                  },
                  "line_1": "line_18",
                  "line_2": "line_22",
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                },
                "metadata": {
                  "key0": "metadata1",
                  "key1": "metadata0"
                },
                "phones": {
                  "home_phone": {
                    "country_code": "country_code0",
                    "number": "number8",
                    "area_code": "area_code0"
                  },
                  "mobile_phone": {
                    "country_code": "country_code0",
                    "number": "number8",
                    "area_code": "area_code0"
                  }
                },
                "fb_id": 180,
                "code": "code6",
                "document_type": "document_type6"
              },
              "card": {
                "id": "id2",
                "last_four_digits": "last_four_digits8",
                "brand": "brand6",
                "holder_name": "holder_name8",
                "exp_month": 82,
                "exp_year": 122,
                "status": "status4",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "billing_address": {
                  "street": "street4",
                  "number": "number2",
                  "zip_code": "zip_code8",
                  "neighborhood": "neighborhood0",
                  "city": "city4",
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
                  "document": "document6",
                  "type": "type2",
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
                      "key0": "metadata1",
                      "key1": "metadata0",
                      "key2": "metadata9"
                    },
                    "line_1": "line_12",
                    "line_2": "line_26",
                    "deleted_at": "2016-03-13T12:52:32.123Z"
                  },
                  "metadata": {
                    "key0": "metadata9",
                    "key1": "metadata8"
                  },
                  "phones": {
                    "home_phone": {
                      "country_code": "country_code4",
                      "number": "number2",
                      "area_code": "area_code4"
                    },
                    "mobile_phone": {
                      "country_code": "country_code4",
                      "number": "number8",
                      "area_code": "area_code4"
                    }
                  },
                  "fb_id": 110,
                  "code": "code0",
                  "document_type": "document_type0"
                },
                "metadata": {
                  "key0": "metadata9",
                  "key1": "metadata8",
                  "key2": "metadata7"
                },
                "type": "type2",
                "holder_document": "holder_document6",
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "first_six_digits": "first_six_digits2",
                "label": "label2"
              },
              "items": [
                {}
              ],
              "statement_descriptor": "statement_descriptor8",
              "metadata": {
                "key0": "metadata1"
              },
              "setup": {
                "id": "id2",
                "description": "description2",
                "amount": 212,
                "status": "status4"
              },
              "gateway_affiliation_id": "gateway_affiliation_id4",
              "next_billing_at": "2016-03-13T12:52:32.123Z",
              "billing_day": 160,
              "minimum_price": 18,
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
                    "amount": 92,
                    "recipient": {
                      "id": "id0",
                      "name": "name0",
                      "email": "email4",
                      "document": "document4",
                      "description": "description0",
                      "type": "type0",
                      "status": "status2",
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
                          "key0": "metadata9",
                          "key1": "metadata0"
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
                        "key0": "metadata9",
                        "key1": "metadata8"
                      },
                      "automatic_anticipation_settings": {
                        "enabled": false,
                        "type": "type4",
                        "volume_percentage": 210,
                        "delay": 244,
                        "days": [
                          220
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": false,
                        "transfer_interval": "transfer_interval6",
                        "transfer_day": 80
                      },
                      "code": "code8",
                      "payment_mode": "payment_mode4"
                    },
                    "gateway_id": "gateway_id8",
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
            "id": "id3",
            "value": 180.75,
            "discount_type": "discount_type1",
            "status": "status5",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 179,
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "description": "description3",
            "subscription": {
              "id": "id9",
              "code": "code7",
              "start_at": "2016-03-13T12:52:32.123Z",
              "interval": "interval7",
              "interval_count": 21,
              "billing_type": "billing_type3",
              "current_cycle": {
                "start_at": "2016-03-13T12:52:32.123Z",
                "end_at": "2016-03-13T12:52:32.123Z",
                "id": "id7",
                "billing_at": "2016-03-13T12:52:32.123Z",
                "subscription": {},
                "status": "status9",
                "duration": 5,
                "created_at": "created_at5",
                "updated_at": "updated_at3",
                "cycle": 201
              },
              "payment_method": "payment_method9",
              "currency": "currency9",
              "installments": 189,
              "status": "status1",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "customer": {
                "id": "id9",
                "name": "name9",
                "email": "email3",
                "delinquent": true,
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "document": "document3",
                "type": "type9",
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
                    "key0": "metadata6",
                    "key1": "metadata7",
                    "key2": "metadata8"
                  },
                  "line_1": "line_19",
                  "line_2": "line_23",
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                },
                "metadata": {
                  "key0": "metadata0",
                  "key1": "metadata9",
                  "key2": "metadata8"
                },
                "phones": {
                  "home_phone": {
                    "country_code": "country_code1",
                    "number": "number9",
                    "area_code": "area_code1"
                  },
                  "mobile_phone": {
                    "country_code": "country_code1",
                    "number": "number9",
                    "area_code": "area_code1"
                  }
                },
                "fb_id": 181,
                "code": "code7",
                "document_type": "document_type7"
              },
              "card": {
                "id": "id3",
                "last_four_digits": "last_four_digits9",
                "brand": "brand7",
                "holder_name": "holder_name9",
                "exp_month": 83,
                "exp_year": 123,
                "status": "status5",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "billing_address": {
                  "street": "street5",
                  "number": "number3",
                  "zip_code": "zip_code9",
                  "neighborhood": "neighborhood1",
                  "city": "city5",
                  "state": "state1",
                  "country": "country9",
                  "complement": "complement1",
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
                  "document": "document7",
                  "type": "type3",
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
                      "key0": "metadata0"
                    },
                    "line_1": "line_13",
                    "line_2": "line_27",
                    "deleted_at": "2016-03-13T12:52:32.123Z"
                  },
                  "metadata": {
                    "key0": "metadata0"
                  },
                  "phones": {
                    "home_phone": {
                      "country_code": "country_code5",
                      "number": "number3",
                      "area_code": "area_code5"
                    },
                    "mobile_phone": {
                      "country_code": "country_code5",
                      "number": "number7",
                      "area_code": "area_code5"
                    }
                  },
                  "fb_id": 111,
                  "code": "code1",
                  "document_type": "document_type1"
                },
                "metadata": {
                  "key0": "metadata0",
                  "key1": "metadata9"
                },
                "type": "type3",
                "holder_document": "holder_document7",
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "first_six_digits": "first_six_digits3",
                "label": "label3"
              },
              "items": [
                {},
                {}
              ],
              "statement_descriptor": "statement_descriptor9",
              "metadata": {
                "key0": "metadata0",
                "key1": "metadata9"
              },
              "setup": {
                "id": "id3",
                "description": "description3",
                "amount": 213,
                "status": "status5"
              },
              "gateway_affiliation_id": "gateway_affiliation_id5",
              "next_billing_at": "2016-03-13T12:52:32.123Z",
              "billing_day": 161,
              "minimum_price": 19,
              "increments": [
                {}
              ],
              "split": {
                "enabled": true,
                "rules": [
                  {
                    "type": "type9",
                    "amount": 93,
                    "recipient": {
                      "id": "id1",
                      "name": "name1",
                      "email": "email5",
                      "document": "document5",
                      "description": "description1",
                      "type": "type1",
                      "status": "status3",
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
                          "key0": "metadata0",
                          "key1": "metadata1",
                          "key2": "metadata2"
                        },
                        "pix_key": "pix_key3"
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
                        },
                        {
                          "gateway": "gateway7",
                          "status": "status9",
                          "pgid": "pgid3",
                          "created_at": "created_at5",
                          "updated_at": "updated_at3"
                        }
                      ],
                      "metadata": {
                        "key0": "metadata8",
                        "key1": "metadata7",
                        "key2": "metadata6"
                      },
                      "automatic_anticipation_settings": {
                        "enabled": true,
                        "type": "type5",
                        "volume_percentage": 211,
                        "delay": 245,
                        "days": [
                          221,
                          222
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": true,
                        "transfer_interval": "transfer_interval7",
                        "transfer_day": 81
                      },
                      "code": "code9",
                      "payment_mode": "payment_mode5"
                    },
                    "gateway_id": "gateway_id9",
                    "options": {
                      "liable": true,
                      "charge_processing_fee": true,
                      "charge_remainder_fee": "charge_remainder_fee5"
                    },
                    "id": "id9"
                  },
                  {
                    "type": "type0",
                    "amount": 94,
                    "recipient": {
                      "id": "id2",
                      "name": "name2",
                      "email": "email6",
                      "document": "document6",
                      "description": "description2",
                      "type": "type2",
                      "status": "status4",
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
                          "key0": "metadata1"
                        },
                        "pix_key": "pix_key4"
                      },
                      "gateway_recipients": [
                        {
                          "gateway": "gateway6",
                          "status": "status8",
                          "pgid": "pgid2",
                          "created_at": "created_at4",
                          "updated_at": "updated_at2"
                        }
                      ],
                      "metadata": {
                        "key0": "metadata7"
                      },
                      "automatic_anticipation_settings": {
                        "enabled": false,
                        "type": "type6",
                        "volume_percentage": 212,
                        "delay": 246,
                        "days": [
                          222,
                          223,
                          224
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": false,
                        "transfer_interval": "transfer_interval8",
                        "transfer_day": 82
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
          }
        ],
        "increments": [
          {},
          {},
          {}
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
            {},
            {},
            {}
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

