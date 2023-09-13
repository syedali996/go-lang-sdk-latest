
# List Usages Response

Response model for listing the usages from a subscription item

## Structure

`ListUsagesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetUsageResponse`](../../doc/models/get-usage-response.md) | Required | The usage objects |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "quantity": 235,
      "description": "description5",
      "used_at": "2016-03-13T12:52:32.123Z",
      "created_at": "2016-03-13T12:52:32.123Z",
      "status": "status7",
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
                {
                  "id": "id7",
                  "value": 223.29,
                  "increment_type": "increment_type9",
                  "status": "status9",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 175,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description7",
                  "subscription": {},
                  "subscription_item": {}
                },
                {
                  "id": "id8",
                  "value": 223.3,
                  "increment_type": "increment_type0",
                  "status": "status0",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 174,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description8",
                  "subscription": {},
                  "subscription_item": {}
                },
                {
                  "id": "id9",
                  "value": 223.31,
                  "increment_type": "increment_type1",
                  "status": "status1",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 173,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description9",
                  "subscription": {},
                  "subscription_item": {}
                }
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
          }
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
          }
        ],
        "subscription": {
          "id": "id7",
          "code": "code5",
          "start_at": "2016-03-13T12:52:32.123Z",
          "interval": "interval5",
          "interval_count": 85,
          "billing_type": "billing_type9",
          "payment_method": "payment_method3",
          "currency": "currency7",
          "installments": 253,
          "status": "status9",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
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
            "id": "id5",
            "billing_at": "2016-03-13T12:52:32.123Z",
            "subscription": {},
            "status": "status7",
            "duration": 69,
            "created_at": "created_at3",
            "updated_at": "updated_at1",
            "cycle": 9
          },
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
          "next_billing_at": "2016-03-13T12:52:32.123Z",
          "billing_day": 225,
          "minimum_price": 173
        },
        "name": "name1",
        "quantity": 223,
        "cycles": 243,
        "deleted_at": "2016-03-13T12:52:32.123Z"
      },
      "deleted_at": "2016-03-13T12:52:32.123Z",
      "code": "code3",
      "group": "group3",
      "amount": 121
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

