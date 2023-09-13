
# List Subscription Items Response

Response model for listing subscription items

## Structure

`ListSubscriptionItemsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md) | Required | The subscription items |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "description": "description5",
      "status": "status7",
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "pricing_scheme": {
        "price": 31,
        "scheme_type": "scheme_type7",
        "price_brackets": [
          {
            "start_quantity": 178,
            "price": 116,
            "end_quantity": 186,
            "overage_price": 200
          }
        ],
        "minimum_price": 65,
        "percentage": 162.75
      },
      "discounts": [
        {
          "id": "id6",
          "value": 160.38,
          "discount_type": "discount_type4",
          "status": "status8",
          "created_at": "2016-03-13T12:52:32.123Z",
          "cycles": 190,
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "description": "description6",
          "subscription": {
            "id": "id2",
            "code": "code0",
            "start_at": "2016-03-13T12:52:32.123Z",
            "interval": "interval0",
            "interval_count": 32,
            "billing_type": "billing_type6",
            "current_cycle": {
              "start_at": "2016-03-13T12:52:32.123Z",
              "end_at": "2016-03-13T12:52:32.123Z",
              "id": "id0",
              "billing_at": "2016-03-13T12:52:32.123Z",
              "subscription": {},
              "status": "status2",
              "duration": 16,
              "created_at": "created_at8",
              "updated_at": "updated_at6",
              "cycle": 212
            },
            "payment_method": "payment_method2",
            "currency": "currency2",
            "installments": 200,
            "status": "status4",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "customer": {
              "id": "id2",
              "name": "name2",
              "email": "email6",
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
                  "key0": "metadata9",
                  "key1": "metadata0",
                  "key2": "metadata1"
                },
                "line_1": "line_12",
                "line_2": "line_26",
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              "metadata": {
                "key0": "metadata7",
                "key1": "metadata6",
                "key2": "metadata5"
              },
              "phones": {
                "home_phone": {
                  "country_code": "country_code4",
                  "number": "number2",
                  "area_code": "area_code4"
                },
                "mobile_phone": {
                  "country_code": "country_code4",
                  "number": "number2",
                  "area_code": "area_code4"
                }
              },
              "fb_id": 192,
              "code": "code0",
              "document_type": "document_type0"
            },
            "card": {
              "id": "id6",
              "last_four_digits": "last_four_digits2",
              "brand": "brand0",
              "holder_name": "holder_name2",
              "exp_month": 94,
              "exp_year": 134,
              "status": "status8",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "billing_address": {
                "street": "street8",
                "number": "number6",
                "zip_code": "zip_code2",
                "neighborhood": "neighborhood4",
                "city": "city8",
                "state": "state4",
                "country": "country2",
                "complement": "complement4",
                "line_1": "line_18",
                "line_2": "line_26"
              },
              "customer": {
                "id": "id6",
                "name": "name6",
                "email": "email0",
                "delinquent": false,
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "document": "document0",
                "type": "type6",
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
                    "key0": "metadata7"
                  },
                  "line_1": "line_16",
                  "line_2": "line_20",
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                },
                "metadata": {
                  "key0": "metadata3"
                },
                "phones": {
                  "home_phone": {
                    "country_code": "country_code8",
                    "number": "number6",
                    "area_code": "area_code8"
                  },
                  "mobile_phone": {
                    "country_code": "country_code8",
                    "number": "number4",
                    "area_code": "area_code8"
                  }
                },
                "fb_id": 122,
                "code": "code4",
                "document_type": "document_type4"
              },
              "metadata": {
                "key0": "metadata3",
                "key1": "metadata2"
              },
              "type": "type6",
              "holder_document": "holder_document0",
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "first_six_digits": "first_six_digits6",
              "label": "label6"
            },
            "items": [
              {},
              {}
            ],
            "statement_descriptor": "statement_descriptor2",
            "metadata": {
              "key0": "metadata7",
              "key1": "metadata6"
            },
            "setup": {
              "id": "id6",
              "description": "description6",
              "amount": 224,
              "status": "status8"
            },
            "gateway_affiliation_id": "gateway_affiliation_id8",
            "next_billing_at": "2016-03-13T12:52:32.123Z",
            "billing_day": 172,
            "minimum_price": 30,
            "increments": [
              {
                "id": "id1",
                "value": 202.93,
                "increment_type": "increment_type3",
                "status": "status3",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 163,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description1",
                "subscription": {},
                "subscription_item": {}
              }
            ],
            "split": {
              "enabled": false,
              "rules": [
                {
                  "type": "type2",
                  "amount": 104,
                  "recipient": {
                    "id": "id4",
                    "name": "name4",
                    "email": "email8",
                    "document": "document8",
                    "description": "description4",
                    "type": "type4",
                    "status": "status6",
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
                        "key0": "metadata3",
                        "key1": "metadata4",
                        "key2": "metadata5"
                      },
                      "pix_key": "pix_key6"
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
                      "key0": "metadata5",
                      "key1": "metadata4",
                      "key2": "metadata3"
                    },
                    "automatic_anticipation_settings": {
                      "enabled": false,
                      "type": "type8",
                      "volume_percentage": 222,
                      "delay": 0,
                      "days": [
                        232,
                        233
                      ]
                    },
                    "transfer_settings": {
                      "transfer_enabled": false,
                      "transfer_interval": "transfer_interval0",
                      "transfer_day": 92
                    },
                    "code": "code2",
                    "payment_mode": "payment_mode8"
                  },
                  "gateway_id": "gateway_id2",
                  "options": {
                    "liable": false,
                    "charge_processing_fee": false,
                    "charge_remainder_fee": "charge_remainder_fee8"
                  },
                  "id": "id2"
                },
                {
                  "type": "type3",
                  "amount": 105,
                  "recipient": {
                    "id": "id5",
                    "name": "name5",
                    "email": "email9",
                    "document": "document9",
                    "description": "description5",
                    "type": "type5",
                    "status": "status7",
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
                        "key0": "metadata4"
                      },
                      "pix_key": "pix_key7"
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
                      "key0": "metadata4"
                    },
                    "automatic_anticipation_settings": {
                      "enabled": true,
                      "type": "type9",
                      "volume_percentage": 223,
                      "delay": 1,
                      "days": [
                        233,
                        234,
                        235
                      ]
                    },
                    "transfer_settings": {
                      "transfer_enabled": true,
                      "transfer_interval": "transfer_interval1",
                      "transfer_day": 93
                    },
                    "code": "code3",
                    "payment_mode": "payment_mode9"
                  },
                  "gateway_id": "gateway_id3",
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
          "subscription_item": {}
        }
      ],
      "increments": [
        {
          "id": "id4",
          "value": 15.66,
          "increment_type": "increment_type6",
          "status": "status6",
          "created_at": "2016-03-13T12:52:32.123Z",
          "cycles": 202,
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "description": "description4",
          "subscription": {
            "id": "id0",
            "code": "code8",
            "start_at": "2016-03-13T12:52:32.123Z",
            "interval": "interval2",
            "interval_count": 12,
            "billing_type": "billing_type6",
            "current_cycle": {
              "start_at": "2016-03-13T12:52:32.123Z",
              "end_at": "2016-03-13T12:52:32.123Z",
              "id": "id8",
              "billing_at": "2016-03-13T12:52:32.123Z",
              "subscription": {},
              "status": "status0",
              "duration": 252,
              "created_at": "created_at6",
              "updated_at": "updated_at4",
              "cycle": 192
            },
            "payment_method": "payment_method0",
            "currency": "currency0",
            "installments": 180,
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
                  "key0": "metadata3"
                },
                "line_1": "line_10",
                "line_2": "line_24",
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              "metadata": {
                "key0": "metadata3"
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
              "fb_id": 244,
              "code": "code8",
              "document_type": "document_type8"
            },
            "card": {
              "id": "id6",
              "last_four_digits": "last_four_digits2",
              "brand": "brand0",
              "holder_name": "holder_name2",
              "exp_month": 158,
              "exp_year": 138,
              "status": "status2",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "billing_address": {
                "street": "street8",
                "number": "number4",
                "zip_code": "zip_code2",
                "neighborhood": "neighborhood4",
                "city": "city2",
                "state": "state6",
                "country": "country2",
                "complement": "complement6",
                "line_1": "line_18",
                "line_2": "line_26"
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
                    "key1": "metadata8",
                    "key2": "metadata7"
                  },
                  "line_1": "line_14",
                  "line_2": "line_20",
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                },
                "metadata": {
                  "key0": "metadata7",
                  "key1": "metadata8",
                  "key2": "metadata9"
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
                "fb_id": 126,
                "code": "code4",
                "document_type": "document_type4"
              },
              "metadata": {
                "key0": "metadata7"
              },
              "type": "type4",
              "holder_document": "holder_document0",
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "first_six_digits": "first_six_digits6",
              "label": "label6"
            },
            "items": [
              {},
              {}
            ],
            "statement_descriptor": "statement_descriptor0",
            "metadata": {
              "key0": "metadata3",
              "key1": "metadata4",
              "key2": "metadata5"
            },
            "setup": {
              "id": "id4",
              "description": "description6",
              "amount": 52,
              "status": "status4"
            },
            "gateway_affiliation_id": "gateway_affiliation_id6",
            "next_billing_at": "2016-03-13T12:52:32.123Z",
            "billing_day": 104,
            "minimum_price": 246,
            "increments": [
              {}
            ],
            "split": {
              "enabled": false,
              "rules": [
                {
                  "type": "type0",
                  "amount": 172,
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
                        "key0": "metadata3"
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
                      }
                    ],
                    "metadata": {
                      "key0": "metadata5"
                    },
                    "automatic_anticipation_settings": {
                      "enabled": false,
                      "type": "type8",
                      "volume_percentage": 146,
                      "delay": 144,
                      "days": [
                        120,
                        121,
                        122
                      ]
                    },
                    "transfer_settings": {
                      "transfer_enabled": false,
                      "transfer_interval": "transfer_interval6",
                      "transfer_day": 56
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
                  "type": "type9",
                  "amount": 171,
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
                      "volume_percentage": 147,
                      "delay": 143,
                      "days": [
                        119,
                        120
                      ]
                    },
                    "transfer_settings": {
                      "transfer_enabled": true,
                      "transfer_interval": "transfer_interval7",
                      "transfer_day": 57
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
        }
      ],
      "subscription": {
        "id": "id1",
        "code": "code9",
        "start_at": "2016-03-13T12:52:32.123Z",
        "interval": "interval9",
        "interval_count": 97,
        "billing_type": "billing_type5",
        "payment_method": "payment_method9",
        "currency": "currency1",
        "installments": 9,
        "status": "status3",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
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
            "id": null,
            "description": null,
            "status": null,
            "created_at": null,
            "updated_at": null,
            "pricing_scheme": {
              "price": null,
              "scheme_type": null,
              "price_brackets": [
                null
              ]
            },
            "discounts": [
              null
            ],
            "increments": [
              null
            ],
            "subscription": null,
            "name": null
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
            "subscription_item": {}
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
          "id": "id9",
          "billing_at": "2016-03-13T12:52:32.123Z",
          "subscription": {},
          "status": "status1",
          "duration": 81,
          "created_at": "created_at7",
          "updated_at": "updated_at5",
          "cycle": 21
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
        "next_billing_at": "2016-03-13T12:52:32.123Z",
        "billing_day": 237,
        "minimum_price": 161
      },
      "name": "name5",
      "quantity": 235,
      "cycles": 255,
      "deleted_at": "2016-03-13T12:52:32.123Z"
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

