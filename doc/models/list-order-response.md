
# List Order Response

Response object for listing order objects

## Structure

`ListOrderResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetOrderResponse`](../../doc/models/get-order-response.md) | Required | The order object |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "code": "code3",
      "currency": "currency5",
      "items": [
        {
          "id": "id2",
          "amount": 180,
          "description": "description2",
          "quantity": 38,
          "category": "category0",
          "code": "code0"
        }
      ],
      "status": "status7",
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "charges": [
        {
          "id": "id1",
          "code": "code9",
          "gateway_id": "gateway_id1",
          "amount": 113,
          "status": "status3",
          "currency": "currency1",
          "payment_method": "payment_method1",
          "due_at": "2016-03-13T12:52:32.123Z",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "last_transaction": {
            "gateway_id": "gateway_id3",
            "amount": 171,
            "status": "status5",
            "success": true,
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "attempt_count": 147,
            "max_attempts": 135,
            "splits": [
              {
                "type": "type5",
                "amount": 103,
                "recipient": {
                  "id": "id7",
                  "name": "name7",
                  "email": "email1",
                  "document": "document1",
                  "description": "description7",
                  "type": "type7",
                  "status": "status9",
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
                      "key0": "metadata6"
                    },
                    "pix_key": "pix_key9"
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
                    "key0": "metadata2"
                  },
                  "automatic_anticipation_settings": {
                    "enabled": true,
                    "type": "type1",
                    "volume_percentage": 221,
                    "delay": 255,
                    "days": [
                      231,
                      232,
                      233
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": true,
                    "transfer_interval": "transfer_interval3",
                    "transfer_day": 91
                  },
                  "code": "code5",
                  "payment_mode": "payment_mode1"
                },
                "gateway_id": "gateway_id5",
                "options": {
                  "liable": true,
                  "charge_processing_fee": true,
                  "charge_remainder_fee": "charge_remainder_fee1"
                },
                "id": "id5"
              }
            ],
            "id": "id3",
            "gateway_response": {
              "code": "code3",
              "errors": [
                {
                  "message": "message0"
                }
              ]
            },
            "antifraud_response": {
              "status": "status3",
              "return_code": "return_code1",
              "return_message": "return_message9",
              "provider_name": "provider_name9",
              "score": "score1"
            },
            "split": [
              {
                "type": "type1",
                "amount": 117,
                "recipient": {
                  "id": "id3",
                  "name": "name3",
                  "email": "email7",
                  "document": "document7",
                  "description": "description3",
                  "type": "type3",
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
                      "key0": "metadata2"
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
                    "key0": "metadata4"
                  },
                  "automatic_anticipation_settings": {
                    "enabled": true,
                    "type": "type7",
                    "volume_percentage": 235,
                    "delay": 13,
                    "days": [
                      245,
                      246,
                      247
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": true,
                    "transfer_interval": "transfer_interval9",
                    "transfer_day": 105
                  },
                  "code": "code1",
                  "payment_mode": "payment_mode7"
                },
                "gateway_id": "gateway_id1",
                "options": {
                  "liable": true,
                  "charge_processing_fee": true,
                  "charge_remainder_fee": "charge_remainder_fee7"
                },
                "id": "id1"
              }
            ],
            "next_attempt": "2016-03-13T12:52:32.123Z",
            "transaction_type": "transaction",
            "metadata": {
              "key0": "metadata4"
            },
            "interest": {
              "days": 77,
              "type": "type3",
              "amount": 151
            },
            "fine": {
              "days": 95,
              "type": "type5",
              "amount": 169
            }
          },
          "metadata": {
            "key0": "metadata8",
            "key1": "metadata7"
          },
          "canceled_amount": 223,
          "paid_amount": 205,
          "recurrency_cycle": "\"first\" or \"subsequent\"",
          "invoice": {
            "id": "id5",
            "code": "code3",
            "url": "url9",
            "amount": 59,
            "status": "status7",
            "payment_method": "payment_method5",
            "created_at": "2016-03-13T12:52:32.123Z",
            "items": [
              {
                "amount": 118,
                "description": "description2",
                "pricing_scheme": {
                  "price": 34,
                  "scheme_type": "scheme_type4",
                  "price_brackets": [
                    {
                      "start_quantity": 175,
                      "price": 113,
                      "end_quantity": 183,
                      "overage_price": 197
                    },
                    {
                      "start_quantity": 176,
                      "price": 114,
                      "end_quantity": 184,
                      "overage_price": 198
                    },
                    {
                      "start_quantity": 177,
                      "price": 115,
                      "end_quantity": 185,
                      "overage_price": 199
                    }
                  ],
                  "minimum_price": 62,
                  "percentage": 162.72
                },
                "price_bracket": {
                  "start_quantity": 92,
                  "price": 30,
                  "end_quantity": 100,
                  "overage_price": 114
                },
                "quantity": 232,
                "name": "name2",
                "subscription_item_id": "subscription_item_id6"
              },
              {
                "amount": 119,
                "description": "description3",
                "pricing_scheme": {
                  "price": 33,
                  "scheme_type": "scheme_type5",
                  "price_brackets": [
                    {
                      "start_quantity": 176,
                      "price": 114,
                      "end_quantity": 184,
                      "overage_price": 198
                    }
                  ],
                  "minimum_price": 63,
                  "percentage": 162.73
                },
                "price_bracket": {
                  "start_quantity": 93,
                  "price": 31,
                  "end_quantity": 101,
                  "overage_price": 115
                },
                "quantity": 233,
                "name": "name3",
                "subscription_item_id": "subscription_item_id7"
              }
            ],
            "customer": {
              "id": "id5",
              "name": "name5",
              "email": "email1",
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
                "key0": "metadata2",
                "key1": "metadata1"
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
              "fb_id": 91,
              "code": "code3",
              "document_type": "document_type3"
            },
            "charge": {},
            "installments": 99,
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
            "subscription": {
              "id": "id1",
              "code": "code9",
              "start_at": "2016-03-13T12:52:32.123Z",
              "interval": "interval1",
              "interval_count": 35,
              "billing_type": "billing_type5",
              "current_cycle": {
                "start_at": "2016-03-13T12:52:32.123Z",
                "end_at": "2016-03-13T12:52:32.123Z",
                "id": "id9",
                "billing_at": "2016-03-13T12:52:32.123Z",
                "subscription": {},
                "status": "status1",
                "duration": 19,
                "created_at": "created_at7",
                "updated_at": "updated_at5",
                "cycle": 215
              },
              "payment_method": "payment_method9",
              "currency": "currency1",
              "installments": 203,
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
                "fb_id": 221,
                "code": "code7",
                "document_type": "document_type7"
              },
              "card": {
                "id": "id5",
                "last_four_digits": "last_four_digits1",
                "brand": "brand9",
                "holder_name": "holder_name1",
                "exp_month": 181,
                "exp_year": 115,
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
                  "fb_id": 103,
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
                {
                  "id": "id8",
                  "description": "description8",
                  "status": "status0",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "pricing_scheme": {
                    "price": 102,
                    "scheme_type": "scheme_type0",
                    "price_brackets": [
                      {
                        "start_quantity": 107,
                        "price": 211,
                        "end_quantity": 115,
                        "overage_price": 129
                      },
                      {
                        "start_quantity": 108,
                        "price": 210,
                        "end_quantity": 116,
                        "overage_price": 130
                      },
                      {
                        "start_quantity": 109,
                        "price": 209,
                        "end_quantity": 117,
                        "overage_price": 131
                      }
                    ],
                    "minimum_price": 6,
                    "percentage": 185.08
                  },
                  "discounts": [
                    {
                      "id": "id9",
                      "value": 76.91,
                      "discount_type": "discount_type7",
                      "status": "status1",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "cycles": 35,
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "description": "description9",
                      "subscription": {},
                      "subscription_item": {}
                    },
                    {
                      "id": "id0",
                      "value": 76.92,
                      "discount_type": "discount_type8",
                      "status": "status2",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "cycles": 36,
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "description": "description0",
                      "subscription": {},
                      "subscription_item": {}
                    },
                    {
                      "id": "id1",
                      "value": 76.93,
                      "discount_type": "discount_type9",
                      "status": "status3",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "cycles": 37,
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "description": "description1",
                      "subscription": {},
                      "subscription_item": {}
                    }
                  ],
                  "increments": [
                    {
                      "id": "id3",
                      "value": 58.85,
                      "increment_type": "increment_type5",
                      "status": "status5",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "cycles": 235,
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "description": "description7",
                      "subscription": {},
                      "subscription_item": {}
                    }
                  ],
                  "subscription": {},
                  "name": "name8",
                  "quantity": 80,
                  "cycles": 156,
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                },
                {
                  "id": "id9",
                  "description": "description9",
                  "status": "status1",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "pricing_scheme": {
                    "price": 103,
                    "scheme_type": "scheme_type9",
                    "price_brackets": [
                      {
                        "start_quantity": 106,
                        "price": 212,
                        "end_quantity": 114,
                        "overage_price": 128
                      },
                      {
                        "start_quantity": 107,
                        "price": 211,
                        "end_quantity": 115,
                        "overage_price": 129
                      }
                    ],
                    "minimum_price": 7,
                    "percentage": 185.07
                  },
                  "discounts": [
                    {
                      "id": "id0",
                      "value": 76.92,
                      "discount_type": "discount_type8",
                      "status": "status2",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "cycles": 36,
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "description": "description0",
                      "subscription": {},
                      "subscription_item": {}
                    }
                  ],
                  "increments": [
                    {
                      "id": "id2",
                      "value": 58.84,
                      "increment_type": "increment_type4",
                      "status": "status6",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "cycles": 236,
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "description": "description8",
                      "subscription": {},
                      "subscription_item": {}
                    },
                    {
                      "id": "id1",
                      "value": 58.83,
                      "increment_type": "increment_type3",
                      "status": "status7",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "cycles": 237,
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "description": "description9",
                      "subscription": {},
                      "subscription_item": {}
                    }
                  ],
                  "subscription": {},
                  "name": "name9",
                  "quantity": 81,
                  "cycles": 155,
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
                "amount": 227,
                "status": "status3"
              },
              "gateway_affiliation_id": "gateway_affiliation_id7",
              "next_billing_at": "2016-03-13T12:52:32.123Z",
              "billing_day": 81,
              "minimum_price": 223,
              "increments": [
                {
                  "id": "id6",
                  "value": 27.38,
                  "increment_type": "increment_type8",
                  "status": "status2",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 54,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description4",
                  "subscription": {},
                  "subscription_item": {
                    "id": "id2",
                    "description": "description8",
                    "status": "status6",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "updated_at": "2016-03-13T12:52:32.123Z",
                    "pricing_scheme": {
                      "price": 192,
                      "scheme_type": "scheme_type6",
                      "price_brackets": [
                        {
                          "start_quantity": 17,
                          "price": 211,
                          "end_quantity": 25,
                          "overage_price": 39
                        }
                      ],
                      "minimum_price": 160,
                      "percentage": 199.54
                    },
                    "discounts": [
                      {
                        "id": "id3",
                        "value": 62.45,
                        "discount_type": "discount_type1",
                        "status": "status5",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 125,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description3",
                        "subscription": {},
                        "subscription_item": {}
                      },
                      {
                        "id": "id4",
                        "value": 62.46,
                        "discount_type": "discount_type2",
                        "status": "status6",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 126,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description4",
                        "subscription": {},
                        "subscription_item": {}
                      }
                    ],
                    "increments": [
                      {},
                      {}
                    ],
                    "subscription": {},
                    "name": "name2",
                    "quantity": 170,
                    "cycles": 66,
                    "deleted_at": "2016-03-13T12:52:32.123Z"
                  }
                }
              ],
              "split": {
                "enabled": true,
                "rules": [
                  {
                    "type": "type9",
                    "amount": 149,
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
                        "volume_percentage": 169,
                        "delay": 121,
                        "days": [
                          97,
                          98,
                          99
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": true,
                        "transfer_interval": "transfer_interval7",
                        "transfer_day": 79
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
                    "amount": 148,
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
                        "volume_percentage": 170,
                        "delay": 120,
                        "days": [
                          96,
                          97
                        ]
                      },
                      "transfer_settings": {
                        "transfer_enabled": false,
                        "transfer_interval": "transfer_interval8",
                        "transfer_day": 80
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
            "cycle": {
              "start_at": "2016-03-13T12:52:32.123Z",
              "end_at": "2016-03-13T12:52:32.123Z",
              "id": "id5",
              "billing_at": "2016-03-13T12:52:32.123Z",
              "subscription": {
                "id": "id1",
                "code": "code9",
                "start_at": "2016-03-13T12:52:32.123Z",
                "interval": "interval9",
                "interval_count": 87,
                "billing_type": "billing_type5",
                "current_cycle": {},
                "payment_method": "payment_method1",
                "currency": "currency1",
                "installments": 255,
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
                  "fb_id": 247,
                  "code": "code9",
                  "document_type": "document_type9"
                },
                "card": {
                  "id": "id5",
                  "last_four_digits": "last_four_digits1",
                  "brand": "brand9",
                  "holder_name": "holder_name1",
                  "exp_month": 107,
                  "exp_year": 189,
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
                    "fb_id": 177,
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
                  {
                    "id": "id8",
                    "description": "description8",
                    "status": "status0",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "updated_at": "2016-03-13T12:52:32.123Z",
                    "pricing_scheme": {
                      "price": 134,
                      "scheme_type": "scheme_type0",
                      "price_brackets": [
                        {
                          "start_quantity": 75,
                          "price": 243,
                          "end_quantity": 83,
                          "overage_price": 97
                        }
                      ],
                      "minimum_price": 38,
                      "percentage": 215.48
                    },
                    "discounts": [
                      {
                        "id": "id9",
                        "value": 213.11,
                        "discount_type": "discount_type7",
                        "status": "status1",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 87,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description9",
                        "subscription": {},
                        "subscription_item": {}
                      }
                    ],
                    "increments": [
                      {
                        "id": "id7",
                        "value": 68.39,
                        "increment_type": "increment_type9",
                        "status": "status9",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 49,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description7",
                        "subscription": {},
                        "subscription_item": {}
                      },
                      {
                        "id": "id8",
                        "value": 68.4,
                        "increment_type": "increment_type0",
                        "status": "status0",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 48,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description8",
                        "subscription": {},
                        "subscription_item": {}
                      }
                    ],
                    "subscription": {},
                    "name": "name8",
                    "quantity": 132,
                    "cycles": 152,
                    "deleted_at": "2016-03-13T12:52:32.123Z"
                  },
                  {
                    "id": "id9",
                    "description": "description9",
                    "status": "status1",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "updated_at": "2016-03-13T12:52:32.123Z",
                    "pricing_scheme": {
                      "price": 133,
                      "scheme_type": "scheme_type1",
                      "price_brackets": [
                        {
                          "start_quantity": 76,
                          "price": 242,
                          "end_quantity": 84,
                          "overage_price": 98
                        },
                        {
                          "start_quantity": 77,
                          "price": 241,
                          "end_quantity": 85,
                          "overage_price": 99
                        }
                      ],
                      "minimum_price": 37,
                      "percentage": 215.49
                    },
                    "discounts": [
                      {
                        "id": "id0",
                        "value": 213.12,
                        "discount_type": "discount_type8",
                        "status": "status2",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 88,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description0",
                        "subscription": {},
                        "subscription_item": {}
                      },
                      {
                        "id": "id1",
                        "value": 213.13,
                        "discount_type": "discount_type9",
                        "status": "status3",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 89,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description1",
                        "subscription": {},
                        "subscription_item": {}
                      }
                    ],
                    "increments": [
                      {
                        "id": "id8",
                        "value": 68.4,
                        "increment_type": "increment_type0",
                        "status": "status0",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 48,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description8",
                        "subscription": {},
                        "subscription_item": {}
                      },
                      {
                        "id": "id9",
                        "value": 68.41,
                        "increment_type": "increment_type1",
                        "status": "status1",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 47,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description9",
                        "subscription": {},
                        "subscription_item": {}
                      },
                      {
                        "id": "id0",
                        "value": 68.42,
                        "increment_type": "increment_type2",
                        "status": "status2",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 46,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description0",
                        "subscription": {},
                        "subscription_item": {}
                      }
                    ],
                    "subscription": {},
                    "name": "name9",
                    "quantity": 133,
                    "cycles": 153,
                    "deleted_at": "2016-03-13T12:52:32.123Z"
                  },
                  {
                    "id": "id0",
                    "description": "description0",
                    "status": "status2",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "updated_at": "2016-03-13T12:52:32.123Z",
                    "pricing_scheme": {
                      "price": 132,
                      "scheme_type": "scheme_type2",
                      "price_brackets": [
                        {
                          "start_quantity": 77,
                          "price": 241,
                          "end_quantity": 85,
                          "overage_price": 99
                        },
                        {
                          "start_quantity": 78,
                          "price": 240,
                          "end_quantity": 86,
                          "overage_price": 100
                        },
                        {
                          "start_quantity": 79,
                          "price": 239,
                          "end_quantity": 87,
                          "overage_price": 101
                        }
                      ],
                      "minimum_price": 36,
                      "percentage": 215.5
                    },
                    "discounts": [
                      {
                        "id": "id1",
                        "value": 213.13,
                        "discount_type": "discount_type9",
                        "status": "status3",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 89,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description1",
                        "subscription": {},
                        "subscription_item": {}
                      },
                      {
                        "id": "id2",
                        "value": 213.14,
                        "discount_type": "discount_type0",
                        "status": "status4",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 90,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description2",
                        "subscription": {},
                        "subscription_item": {}
                      },
                      {
                        "id": "id3",
                        "value": 213.15,
                        "discount_type": "discount_type1",
                        "status": "status5",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 91,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description3",
                        "subscription": {},
                        "subscription_item": {}
                      }
                    ],
                    "increments": [
                      {
                        "id": "id9",
                        "value": 68.41,
                        "increment_type": "increment_type1",
                        "status": "status1",
                        "created_at": "2016-03-13T12:52:32.123Z",
                        "cycles": 47,
                        "deleted_at": "2016-03-13T12:52:32.123Z",
                        "description": "description9",
                        "subscription": {},
                        "subscription_item": {}
                      }
                    ],
                    "subscription": {},
                    "name": "name0",
                    "quantity": 134,
                    "cycles": 154,
                    "deleted_at": "2016-03-13T12:52:32.123Z"
                  }
                ],
                "statement_descriptor": "statement_descriptor1",
                "metadata": {
                  "key0": "metadata8",
                  "key1": "metadata7"
                },
                "setup": {
                  "id": "id5",
                  "description": "description5",
                  "amount": 23,
                  "status": "status7"
                },
                "gateway_affiliation_id": "gateway_affiliation_id7",
                "next_billing_at": "2016-03-13T12:52:32.123Z",
                "billing_day": 227,
                "minimum_price": 85,
                "increments": [
                  {
                    "id": "id0",
                    "value": 98.52,
                    "increment_type": "increment_type2",
                    "status": "status8",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 108,
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
                        "price": 138,
                        "scheme_type": "scheme_type2",
                        "price_brackets": [
                          {
                            "start_quantity": 71,
                            "price": 9,
                            "end_quantity": 79,
                            "overage_price": 93
                          },
                          {
                            "start_quantity": 72,
                            "price": 10,
                            "end_quantity": 80,
                            "overage_price": 94
                          },
                          {
                            "start_quantity": 73,
                            "price": 11,
                            "end_quantity": 81,
                            "overage_price": 95
                          }
                        ],
                        "minimum_price": 42,
                        "percentage": 128.4
                      },
                      "discounts": [
                        {
                          "id": "id7",
                          "value": 133.59,
                          "discount_type": "discount_type5",
                          "status": "status9",
                          "created_at": "2016-03-13T12:52:32.123Z",
                          "cycles": 71,
                          "deleted_at": "2016-03-13T12:52:32.123Z",
                          "description": "description7",
                          "subscription": {},
                          "subscription_item": {}
                        },
                        {
                          "id": "id8",
                          "value": 133.6,
                          "discount_type": "discount_type6",
                          "status": "status0",
                          "created_at": "2016-03-13T12:52:32.123Z",
                          "cycles": 72,
                          "deleted_at": "2016-03-13T12:52:32.123Z",
                          "description": "description8",
                          "subscription": {},
                          "subscription_item": {}
                        },
                        {
                          "id": "id9",
                          "value": 133.61,
                          "discount_type": "discount_type7",
                          "status": "status1",
                          "created_at": "2016-03-13T12:52:32.123Z",
                          "cycles": 73,
                          "deleted_at": "2016-03-13T12:52:32.123Z",
                          "description": "description9",
                          "subscription": {},
                          "subscription_item": {}
                        }
                      ],
                      "increments": [
                        {}
                      ],
                      "subscription": {},
                      "name": "name6",
                      "quantity": 116,
                      "cycles": 120,
                      "deleted_at": "2016-03-13T12:52:32.123Z"
                    }
                  },
                  {
                    "id": "id1",
                    "value": 98.53,
                    "increment_type": "increment_type3",
                    "status": "status7",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 107,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description9",
                    "subscription": {},
                    "subscription_item": {
                      "id": "id7",
                      "description": "description3",
                      "status": "status1",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "updated_at": "2016-03-13T12:52:32.123Z",
                      "pricing_scheme": {
                        "price": 139,
                        "scheme_type": "scheme_type1",
                        "price_brackets": [
                          {
                            "start_quantity": 70,
                            "price": 8,
                            "end_quantity": 78,
                            "overage_price": 92
                          },
                          {
                            "start_quantity": 71,
                            "price": 9,
                            "end_quantity": 79,
                            "overage_price": 93
                          }
                        ],
                        "minimum_price": 43,
                        "percentage": 128.39
                      },
                      "discounts": [
                        {
                          "id": "id8",
                          "value": 133.6,
                          "discount_type": "discount_type6",
                          "status": "status0",
                          "created_at": "2016-03-13T12:52:32.123Z",
                          "cycles": 72,
                          "deleted_at": "2016-03-13T12:52:32.123Z",
                          "description": "description8",
                          "subscription": {},
                          "subscription_item": {}
                        }
                      ],
                      "increments": [
                        {},
                        {},
                        {}
                      ],
                      "subscription": {},
                      "name": "name7",
                      "quantity": 117,
                      "cycles": 119,
                      "deleted_at": "2016-03-13T12:52:32.123Z"
                    }
                  }
                ],
                "split": {
                  "enabled": true,
                  "rules": [
                    {
                      "type": "type9",
                      "amount": 159,
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
                          "volume_percentage": 235,
                          "delay": 55,
                          "days": [
                            31,
                            32,
                            33
                          ]
                        },
                        "transfer_settings": {
                          "transfer_enabled": true,
                          "transfer_interval": "transfer_interval1",
                          "transfer_day": 145
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
                      "amount": 160,
                      "recipient": {
                        "id": "id4",
                        "name": "name4",
                        "email": "email2",
                        "document": "document2",
                        "description": "description4",
                        "type": "type6",
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
                            "key0": "metadata9",
                            "key1": "metadata8",
                            "key2": "metadata7"
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
                          "volume_percentage": 234,
                          "delay": 56,
                          "days": [
                            32
                          ]
                        },
                        "transfer_settings": {
                          "transfer_enabled": false,
                          "transfer_interval": "transfer_interval0",
                          "transfer_day": 144
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
                      "amount": 161,
                      "recipient": {
                        "id": "id5",
                        "name": "name5",
                        "email": "email1",
                        "document": "document1",
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
                            "key0": "metadata0",
                            "key1": "metadata9"
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
                          "volume_percentage": 233,
                          "delay": 57,
                          "days": [
                            33,
                            34
                          ]
                        },
                        "transfer_settings": {
                          "transfer_enabled": true,
                          "transfer_interval": "transfer_interval9",
                          "transfer_day": 143
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
              "status": "status7",
              "duration": 13,
              "created_at": "created_at3",
              "updated_at": "updated_at1",
              "cycle": 209
            },
            "shipping": {
              "amount": 157,
              "description": "description9",
              "recipient_name": "recipient_name7",
              "recipient_phone": "recipient_phone1",
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
                  "address": {},
                  "metadata": {
                    "key0": "metadata6",
                    "key1": "metadata7"
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
                  "fb_id": 205,
                  "code": "code3",
                  "document_type": "document_type3"
                },
                "metadata": {
                  "key0": "metadata6"
                },
                "line_1": "line_19",
                "line_2": "line_23",
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              "max_delivery_date": "2016-03-13T12:52:32.123Z",
              "estimated_delivery_date": "2016-03-13T12:52:32.123Z",
              "type": "type9"
            },
            "metadata": {
              "key0": "metadata4",
              "key1": "metadata3"
            },
            "due_at": "2016-03-13T12:52:32.123Z",
            "canceled_at": "2016-03-13T12:52:32.123Z",
            "billing_at": "2016-03-13T12:52:32.123Z",
            "subscription_id": "subscription_id5"
          },
          "order": {},
          "customer": {
            "id": "id1",
            "name": "name1",
            "email": "email5",
            "delinquent": true,
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "document": "document5",
            "type": "type1",
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
                "key0": "metadata8",
                "key1": "metadata9",
                "key2": "metadata0"
              },
              "line_1": "line_11",
              "line_2": "line_25",
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            "metadata": {
              "key0": "metadata8",
              "key1": "metadata7"
            },
            "phones": {
              "home_phone": {
                "country_code": "country_code3",
                "number": "number1",
                "area_code": "area_code3"
              },
              "mobile_phone": {
                "country_code": "country_code3",
                "number": "number9",
                "area_code": "area_code3"
              }
            },
            "fb_id": 145,
            "code": "code9",
            "document_type": "document_type9"
          },
          "paid_at": "2016-03-13T12:52:32.123Z",
          "canceled_at": "2016-03-13T12:52:32.123Z"
        }
      ],
      "invoice_url": "invoice_url7",
      "shipping": {
        "amount": 219,
        "description": "description9",
        "recipient_name": "recipient_name7",
        "recipient_phone": "recipient_phone1",
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
            "address": {},
            "metadata": {
              "key0": "metadata6",
              "key1": "metadata7"
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
            "fb_id": 11,
            "code": "code3",
            "document_type": "document_type3"
          },
          "metadata": {
            "key0": "metadata6"
          },
          "line_1": "line_19",
          "line_2": "line_23",
          "deleted_at": "2016-03-13T12:52:32.123Z"
        },
        "max_delivery_date": "2016-03-13T12:52:32.123Z",
        "estimated_delivery_date": "2016-03-13T12:52:32.123Z",
        "type": "type9"
      },
      "metadata": {
        "key0": "metadata4",
        "key1": "metadata3"
      },
      "closed": true,
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
      "checkouts": [
        {
          "id": "id2",
          "amount": 130,
          "default_payment_method": "default_payment_method2",
          "success_url": "success_url4",
          "payment_url": "payment_url6",
          "gateway_affiliation_id": "gateway_affiliation_id8",
          "accepted_payment_methods": [
            "accepted_payment_methods5",
            "accepted_payment_methods6",
            "accepted_payment_methods7"
          ],
          "status": "status4",
          "skip_checkout_success_page": false,
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "canceled_at": "2016-03-13T12:52:32.123Z",
          "customer_editable": false,
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
              "key0": "metadata3",
              "key1": "metadata4",
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
            "fb_id": 162,
            "code": "code0",
            "document_type": "document_type0"
          },
          "billingaddress": {
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
            "customer": {
              "id": "id0",
              "name": "name0",
              "email": "email4",
              "delinquent": false,
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "document": "document4",
              "type": "type0",
              "fb_access_token": "fb_access_token4",
              "address": {},
              "metadata": {
                "key0": "metadata1",
                "key1": "metadata2"
              },
              "phones": {
                "home_phone": {
                  "country_code": "country_code2",
                  "number": "number0",
                  "area_code": "area_code2"
                },
                "mobile_phone": {
                  "country_code": "country_code2",
                  "number": "number0",
                  "area_code": "area_code2"
                }
              },
              "fb_id": 244,
              "code": "code8",
              "document_type": "document_type8"
            },
            "metadata": {
              "key0": "metadata1"
            },
            "line_1": "line_14",
            "line_2": "line_28",
            "deleted_at": "2016-03-13T12:52:32.123Z"
          },
          "credit_card": {
            "statementDescriptor": "statementDescriptor6",
            "installments": [
              {
                "number": "number3",
                "total": 87
              }
            ],
            "authentication": {
              "type": "type2",
              "threed_secure": {
                "mpi": "mpi8",
                "eci": "eci0",
                "cavv": "cavv6",
                "transaction_Id": "transaction_Id6",
                "success_url": "success_url2"
              }
            }
          },
          "boleto": {
            "due_at": "2016-03-13T12:52:32.123Z",
            "instructions": "instructions0"
          },
          "billing_address_editable": false,
          "shipping": {
            "amount": 228,
            "description": "description6",
            "recipient_name": "recipient_name4",
            "recipient_phone": "recipient_phone8",
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
                "address": {},
                "metadata": {
                  "key0": "metadata3",
                  "key1": "metadata4"
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
                "fb_id": 20,
                "code": "code0",
                "document_type": "document_type0"
              },
              "metadata": {
                "key0": "metadata3"
              },
              "line_1": "line_16",
              "line_2": "line_20",
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            "max_delivery_date": "2016-03-13T12:52:32.123Z",
            "estimated_delivery_date": "2016-03-13T12:52:32.123Z",
            "type": "type6"
          },
          "shippable": false,
          "closed_at": "2016-03-13T12:52:32.123Z",
          "expires_at": "2016-03-13T12:52:32.123Z",
          "currency": "currency2",
          "accepted_brands": [
            "accepted_brands8",
            "accepted_brands9",
            "accepted_brands0"
          ]
        }
      ],
      "ip": "ip9",
      "session_id": "session_id7",
      "location": {
        "latitude": "latitude3",
        "longitude": "longitude3"
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

