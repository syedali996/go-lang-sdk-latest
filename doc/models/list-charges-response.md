
# List Charges Response

Response object for listing charges

## Structure

`ListChargesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetChargeResponse`](../../doc/models/get-charge-response.md) | Required | The charge objects |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "code": "code3",
      "gateway_id": "gateway_id5",
      "amount": 121,
      "status": "status7",
      "currency": "currency5",
      "payment_method": "payment_method5",
      "due_at": "2016-03-13T12:52:32.123Z",
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "last_transaction": {
        "gateway_id": "gateway_id7",
        "amount": 179,
        "status": "status9",
        "success": true,
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "attempt_count": 155,
        "max_attempts": 143,
        "splits": [
          {
            "type": "type9",
            "amount": 111,
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
                  "key0": "metadata0"
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
                }
              ],
              "metadata": {
                "key0": "metadata8"
              },
              "automatic_anticipation_settings": {
                "enabled": true,
                "type": "type5",
                "volume_percentage": 229,
                "delay": 7,
                "days": [
                  239,
                  240,
                  241
                ]
              },
              "transfer_settings": {
                "transfer_enabled": true,
                "transfer_interval": "transfer_interval7",
                "transfer_day": 99
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
          }
        ],
        "id": "id7",
        "gateway_response": {
          "code": "code7",
          "errors": [
            {
              "message": "message4"
            }
          ]
        },
        "antifraud_response": {
          "status": "status7",
          "return_code": "return_code5",
          "return_message": "return_message3",
          "provider_name": "provider_name3",
          "score": "score5"
        },
        "split": [
          {
            "type": "type5",
            "amount": 125,
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
                "key0": "metadata8"
              },
              "automatic_anticipation_settings": {
                "enabled": true,
                "type": "type1",
                "volume_percentage": 243,
                "delay": 21,
                "days": [
                  253,
                  254,
                  255
                ]
              },
              "transfer_settings": {
                "transfer_enabled": true,
                "transfer_interval": "transfer_interval3",
                "transfer_day": 113
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
        "next_attempt": "2016-03-13T12:52:32.123Z",
        "transaction_type": "transaction",
        "metadata": {
          "key0": "metadata8"
        },
        "interest": {
          "days": 85,
          "type": "type7",
          "amount": 159
        },
        "fine": {
          "days": 103,
          "type": "type9",
          "amount": 177
        }
      },
      "metadata": {
        "key0": "metadata4",
        "key1": "metadata3"
      },
      "canceled_amount": 231,
      "paid_amount": 213,
      "recurrency_cycle": "\"first\" or \"subsequent\"",
      "invoice": {
        "id": "id9",
        "code": "code7",
        "url": "url3",
        "amount": 67,
        "status": "status1",
        "payment_method": "payment_method9",
        "created_at": "2016-03-13T12:52:32.123Z",
        "items": [
          {
            "amount": 126,
            "description": "description6",
            "pricing_scheme": {
              "price": 230,
              "scheme_type": "scheme_type8",
              "price_brackets": [
                {
                  "start_quantity": 183,
                  "price": 121,
                  "end_quantity": 191,
                  "overage_price": 205
                },
                {
                  "start_quantity": 184,
                  "price": 122,
                  "end_quantity": 192,
                  "overage_price": 206
                },
                {
                  "start_quantity": 185,
                  "price": 123,
                  "end_quantity": 193,
                  "overage_price": 207
                }
              ],
              "minimum_price": 70,
              "percentage": 101.36
            },
            "price_bracket": {
              "start_quantity": 100,
              "price": 38,
              "end_quantity": 108,
              "overage_price": 122
            },
            "quantity": 240,
            "name": "name6",
            "subscription_item_id": "subscription_item_id0"
          },
          {
            "amount": 127,
            "description": "description7",
            "pricing_scheme": {
              "price": 231,
              "scheme_type": "scheme_type9",
              "price_brackets": [
                {
                  "start_quantity": 184,
                  "price": 122,
                  "end_quantity": 192,
                  "overage_price": 206
                }
              ],
              "minimum_price": 71,
              "percentage": 101.37
            },
            "price_bracket": {
              "start_quantity": 101,
              "price": 39,
              "end_quantity": 109,
              "overage_price": 123
            },
            "quantity": 241,
            "name": "name7",
            "subscription_item_id": "subscription_item_id1"
          }
        ],
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
          "fb_id": 99,
          "code": "code7",
          "document_type": "document_type7"
        },
        "charge": {},
        "installments": 107,
        "billing_address": {
          "street": "street1",
          "number": "number9",
          "zip_code": "zip_code5",
          "neighborhood": "neighborhood7",
          "city": "city1",
          "state": "state7",
          "country": "country5",
          "complement": "complement7",
          "line_1": "line_15",
          "line_2": "line_29"
        },
        "subscription": {
          "id": "id5",
          "code": "code3",
          "start_at": "2016-03-13T12:52:32.123Z",
          "interval": "interval3",
          "interval_count": 43,
          "billing_type": "billing_type1",
          "current_cycle": {
            "start_at": "2016-03-13T12:52:32.123Z",
            "end_at": "2016-03-13T12:52:32.123Z",
            "id": "id3",
            "billing_at": "2016-03-13T12:52:32.123Z",
            "subscription": {},
            "status": "status5",
            "duration": 27,
            "created_at": "created_at1",
            "updated_at": "updated_at9",
            "cycle": 223
          },
          "payment_method": "payment_method5",
          "currency": "currency5",
          "installments": 211,
          "status": "status7",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
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
            "fb_id": 203,
            "code": "code3",
            "document_type": "document_type3"
          },
          "card": {
            "id": "id9",
            "last_four_digits": "last_four_digits5",
            "brand": "brand3",
            "holder_name": "holder_name5",
            "exp_month": 151,
            "exp_year": 145,
            "status": "status9",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "billing_address": {
              "street": "street9",
              "number": "number3",
              "zip_code": "zip_code3",
              "neighborhood": "neighborhood5",
              "city": "city1",
              "state": "state5",
              "country": "country3",
              "complement": "complement5",
              "line_1": "line_17",
              "line_2": "line_27"
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
              "fb_id": 27,
              "code": "code9",
              "document_type": "document_type9"
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
            {
              "id": "id2",
              "description": "description2",
              "status": "status4",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "pricing_scheme": {
                "price": 178,
                "scheme_type": "scheme_type4",
                "price_brackets": [
                  {
                    "start_quantity": 31,
                    "price": 31,
                    "end_quantity": 39,
                    "overage_price": 53
                  },
                  {
                    "start_quantity": 32,
                    "price": 30,
                    "end_quantity": 40,
                    "overage_price": 54
                  },
                  {
                    "start_quantity": 33,
                    "price": 29,
                    "end_quantity": 41,
                    "overage_price": 55
                  }
                ],
                "minimum_price": 82,
                "percentage": 17.92
              },
              "discounts": [
                {
                  "id": "id3",
                  "value": 15.55,
                  "discount_type": "discount_type1",
                  "status": "status5",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 43,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description3",
                  "subscription": {},
                  "subscription_item": {}
                },
                {
                  "id": "id4",
                  "value": 15.56,
                  "discount_type": "discount_type2",
                  "status": "status6",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 44,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description4",
                  "subscription": {},
                  "subscription_item": {}
                },
                {
                  "id": "id5",
                  "value": 15.57,
                  "discount_type": "discount_type3",
                  "status": "status7",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 45,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description5",
                  "subscription": {},
                  "subscription_item": {}
                }
              ],
              "increments": [
                {
                  "id": "id1",
                  "value": 126.83,
                  "increment_type": "increment_type3",
                  "status": "status7",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 93,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description9",
                  "subscription": {},
                  "subscription_item": {}
                }
              ],
              "subscription": {},
              "name": "name2",
              "quantity": 88,
              "cycles": 148,
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            {
              "id": "id3",
              "description": "description3",
              "status": "status5",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "pricing_scheme": {
                "price": 177,
                "scheme_type": "scheme_type5",
                "price_brackets": [
                  {
                    "start_quantity": 32,
                    "price": 30,
                    "end_quantity": 40,
                    "overage_price": 54
                  }
                ],
                "minimum_price": 81,
                "percentage": 17.93
              },
              "discounts": [
                {
                  "id": "id4",
                  "value": 15.56,
                  "discount_type": "discount_type2",
                  "status": "status6",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 44,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description4",
                  "subscription": {},
                  "subscription_item": {}
                }
              ],
              "increments": [
                {
                  "id": "id2",
                  "value": 126.84,
                  "increment_type": "increment_type4",
                  "status": "status6",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 92,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description8",
                  "subscription": {},
                  "subscription_item": {}
                },
                {
                  "id": "id3",
                  "value": 126.85,
                  "increment_type": "increment_type5",
                  "status": "status5",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 91,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description7",
                  "subscription": {},
                  "subscription_item": {}
                }
              ],
              "subscription": {},
              "name": "name3",
              "quantity": 89,
              "cycles": 147,
              "deleted_at": "2016-03-13T12:52:32.123Z"
            }
          ],
          "statement_descriptor": "statement_descriptor5",
          "metadata": {
            "key0": "metadata8",
            "key1": "metadata9",
            "key2": "metadata0"
          },
          "setup": {
            "id": "id9",
            "description": "description9",
            "amount": 235,
            "status": "status9"
          },
          "gateway_affiliation_id": "gateway_affiliation_id1",
          "next_billing_at": "2016-03-13T12:52:32.123Z",
          "billing_day": 183,
          "minimum_price": 215,
          "increments": [
            {
              "id": "id6",
              "value": 90.08,
              "increment_type": "increment_type8",
              "status": "status2",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 184,
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
                  "price": 62,
                  "scheme_type": "scheme_type6",
                  "price_brackets": [
                    {
                      "start_quantity": 147,
                      "price": 85,
                      "end_quantity": 155,
                      "overage_price": 169
                    }
                  ],
                  "minimum_price": 222,
                  "percentage": 136.84
                },
                "discounts": [
                  {
                    "id": "id3",
                    "value": 125.15,
                    "discount_type": "discount_type1",
                    "status": "status5",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 251,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description3",
                    "subscription": {},
                    "subscription_item": {}
                  },
                  {
                    "id": "id4",
                    "value": 125.16,
                    "discount_type": "discount_type2",
                    "status": "status6",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 252,
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
                "quantity": 40,
                "cycles": 196,
                "deleted_at": "2016-03-13T12:52:32.123Z"
              }
            }
          ],
          "split": {
            "enabled": true,
            "rules": [
              {
                "type": "type5",
                "amount": 115,
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
                    "volume_percentage": 23,
                    "delay": 11,
                    "days": [
                      243,
                      244
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": true,
                    "transfer_interval": "transfer_interval7",
                    "transfer_day": 189
                  },
                  "code": "code5",
                  "payment_mode": "payment_mode9"
                },
                "gateway_id": "gateway_id5",
                "options": {
                  "liable": true,
                  "charge_processing_fee": true,
                  "charge_remainder_fee": "charge_remainder_fee1"
                },
                "id": "id5"
              },
              {
                "type": "type4",
                "amount": 116,
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
                    "volume_percentage": 22,
                    "delay": 12,
                    "days": [
                      244,
                      245,
                      246
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": false,
                    "transfer_interval": "transfer_interval6",
                    "transfer_day": 188
                  },
                  "code": "code6",
                  "payment_mode": "payment_mode8"
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
        "cycle": {
          "start_at": "2016-03-13T12:52:32.123Z",
          "end_at": "2016-03-13T12:52:32.123Z",
          "id": "id9",
          "billing_at": "2016-03-13T12:52:32.123Z",
          "subscription": {
            "id": "id5",
            "code": "code3",
            "start_at": "2016-03-13T12:52:32.123Z",
            "interval": "interval3",
            "interval_count": 95,
            "billing_type": "billing_type9",
            "current_cycle": {},
            "payment_method": "payment_method5",
            "currency": "currency5",
            "installments": 7,
            "status": "status7",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
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
                  "key0": "metadata2"
                },
                "line_1": "line_15",
                "line_2": "line_29",
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              "metadata": {
                "key0": "metadata2"
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
              "fb_id": 255,
              "code": "code3",
              "document_type": "document_type3"
            },
            "card": {
              "id": "id9",
              "last_four_digits": "last_four_digits5",
              "brand": "brand3",
              "holder_name": "holder_name5",
              "exp_month": 99,
              "exp_year": 197,
              "status": "status1",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "billing_address": {
                "street": "street1",
                "number": "number1",
                "zip_code": "zip_code5",
                "neighborhood": "neighborhood7",
                "city": "city1",
                "state": "state7",
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
                "document": "document3",
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
                    "key0": "metadata4",
                    "key1": "metadata3"
                  },
                  "line_1": "line_19",
                  "line_2": "line_23",
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                },
                "metadata": {
                  "key0": "metadata6",
                  "key1": "metadata5",
                  "key2": "metadata4"
                },
                "phones": {
                  "home_phone": {
                    "country_code": "country_code1",
                    "number": "number9",
                    "area_code": "area_code1"
                  },
                  "mobile_phone": {
                    "country_code": "country_code9",
                    "number": "number1",
                    "area_code": "area_code1"
                  }
                },
                "fb_id": 185,
                "code": "code7",
                "document_type": "document_type7"
              },
              "metadata": {
                "key0": "metadata6"
              },
              "type": "type1",
              "holder_document": "holder_document3",
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "first_six_digits": "first_six_digits9",
              "label": "label9"
            },
            "items": [
              {
                "id": "id2",
                "description": "description2",
                "status": "status4",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "pricing_scheme": {
                  "price": 126,
                  "scheme_type": "scheme_type4",
                  "price_brackets": [
                    {
                      "start_quantity": 83,
                      "price": 21,
                      "end_quantity": 91,
                      "overage_price": 105
                    }
                  ],
                  "minimum_price": 226,
                  "percentage": 154.12
                },
                "discounts": [
                  {
                    "id": "id3",
                    "value": 151.75,
                    "discount_type": "discount_type1",
                    "status": "status5",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 95,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description3",
                    "subscription": {},
                    "subscription_item": {}
                  }
                ],
                "increments": [
                  {
                    "id": "id1",
                    "value": 7.03,
                    "increment_type": "increment_type3",
                    "status": "status3",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 41,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description1",
                    "subscription": {},
                    "subscription_item": {}
                  },
                  {
                    "id": "id2",
                    "value": 7.04,
                    "increment_type": "increment_type4",
                    "status": "status4",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 40,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description2",
                    "subscription": {},
                    "subscription_item": {}
                  }
                ],
                "subscription": {},
                "name": "name2",
                "quantity": 140,
                "cycles": 160,
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              {
                "id": "id3",
                "description": "description3",
                "status": "status5",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "pricing_scheme": {
                  "price": 125,
                  "scheme_type": "scheme_type5",
                  "price_brackets": [
                    {
                      "start_quantity": 84,
                      "price": 22,
                      "end_quantity": 92,
                      "overage_price": 106
                    },
                    {
                      "start_quantity": 85,
                      "price": 23,
                      "end_quantity": 93,
                      "overage_price": 107
                    }
                  ],
                  "minimum_price": 227,
                  "percentage": 154.13
                },
                "discounts": [
                  {
                    "id": "id4",
                    "value": 151.76,
                    "discount_type": "discount_type2",
                    "status": "status6",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 96,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description4",
                    "subscription": {},
                    "subscription_item": {}
                  },
                  {
                    "id": "id5",
                    "value": 151.77,
                    "discount_type": "discount_type3",
                    "status": "status7",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 97,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description5",
                    "subscription": {},
                    "subscription_item": {}
                  }
                ],
                "increments": [
                  {
                    "id": "id2",
                    "value": 7.04,
                    "increment_type": "increment_type4",
                    "status": "status4",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 40,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description2",
                    "subscription": {},
                    "subscription_item": {}
                  },
                  {
                    "id": "id3",
                    "value": 7.05,
                    "increment_type": "increment_type5",
                    "status": "status5",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 39,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description3",
                    "subscription": {},
                    "subscription_item": {}
                  },
                  {
                    "id": "id4",
                    "value": 7.06,
                    "increment_type": "increment_type6",
                    "status": "status6",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 38,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description4",
                    "subscription": {},
                    "subscription_item": {}
                  }
                ],
                "subscription": {},
                "name": "name3",
                "quantity": 141,
                "cycles": 161,
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              {
                "id": "id4",
                "description": "description4",
                "status": "status6",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "pricing_scheme": {
                  "price": 124,
                  "scheme_type": "scheme_type6",
                  "price_brackets": [
                    {
                      "start_quantity": 85,
                      "price": 23,
                      "end_quantity": 93,
                      "overage_price": 107
                    },
                    {
                      "start_quantity": 86,
                      "price": 24,
                      "end_quantity": 94,
                      "overage_price": 108
                    },
                    {
                      "start_quantity": 87,
                      "price": 25,
                      "end_quantity": 95,
                      "overage_price": 109
                    }
                  ],
                  "minimum_price": 228,
                  "percentage": 154.14
                },
                "discounts": [
                  {
                    "id": "id5",
                    "value": 151.77,
                    "discount_type": "discount_type3",
                    "status": "status7",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 97,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description5",
                    "subscription": {},
                    "subscription_item": {}
                  },
                  {
                    "id": "id6",
                    "value": 151.78,
                    "discount_type": "discount_type4",
                    "status": "status8",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 98,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description6",
                    "subscription": {},
                    "subscription_item": {}
                  },
                  {
                    "id": "id7",
                    "value": 151.79,
                    "discount_type": "discount_type5",
                    "status": "status9",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 99,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description7",
                    "subscription": {},
                    "subscription_item": {}
                  }
                ],
                "increments": [
                  {
                    "id": "id3",
                    "value": 7.05,
                    "increment_type": "increment_type5",
                    "status": "status5",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 39,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description3",
                    "subscription": {},
                    "subscription_item": {}
                  }
                ],
                "subscription": {},
                "name": "name4",
                "quantity": 142,
                "cycles": 162,
                "deleted_at": "2016-03-13T12:52:32.123Z"
              }
            ],
            "statement_descriptor": "statement_descriptor5",
            "metadata": {
              "key0": "metadata4",
              "key1": "metadata3",
              "key2": "metadata2"
            },
            "setup": {
              "id": "id9",
              "description": "description9",
              "amount": 31,
              "status": "status1"
            },
            "gateway_affiliation_id": "gateway_affiliation_id1",
            "next_billing_at": "2016-03-13T12:52:32.123Z",
            "billing_day": 235,
            "minimum_price": 93,
            "increments": [
              {
                "id": "id4",
                "value": 37.16,
                "increment_type": "increment_type6",
                "status": "status6",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 100,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description4",
                "subscription": {},
                "subscription_item": {
                  "id": "id0",
                  "description": "description0",
                  "status": "status2",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "pricing_scheme": {
                    "price": 146,
                    "scheme_type": "scheme_type8",
                    "price_brackets": [
                      {
                        "start_quantity": 63,
                        "price": 255,
                        "end_quantity": 71,
                        "overage_price": 85
                      },
                      {
                        "start_quantity": 64,
                        "price": 254,
                        "end_quantity": 72,
                        "overage_price": 86
                      },
                      {
                        "start_quantity": 65,
                        "price": 253,
                        "end_quantity": 73,
                        "overage_price": 87
                      }
                    ],
                    "minimum_price": 50,
                    "percentage": 189.76
                  },
                  "discounts": [
                    {
                      "id": "id1",
                      "value": 72.23,
                      "discount_type": "discount_type9",
                      "status": "status3",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "cycles": 79,
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "description": "description1",
                      "subscription": {},
                      "subscription_item": {}
                    },
                    {
                      "id": "id2",
                      "value": 72.24,
                      "discount_type": "discount_type0",
                      "status": "status4",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "cycles": 80,
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "description": "description2",
                      "subscription": {},
                      "subscription_item": {}
                    },
                    {
                      "id": "id3",
                      "value": 72.25,
                      "discount_type": "discount_type1",
                      "status": "status5",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "cycles": 81,
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "description": "description3",
                      "subscription": {},
                      "subscription_item": {}
                    }
                  ],
                  "increments": [
                    {}
                  ],
                  "subscription": {},
                  "name": "name0",
                  "quantity": 124,
                  "cycles": 112,
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                }
              },
              {
                "id": "id5",
                "value": 37.17,
                "increment_type": "increment_type7",
                "status": "status7",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 99,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description5",
                "subscription": {},
                "subscription_item": {
                  "id": "id1",
                  "description": "description1",
                  "status": "status3",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "pricing_scheme": {
                    "price": 147,
                    "scheme_type": "scheme_type7",
                    "price_brackets": [
                      {
                        "start_quantity": 62,
                        "price": 0,
                        "end_quantity": 70,
                        "overage_price": 84
                      },
                      {
                        "start_quantity": 63,
                        "price": 255,
                        "end_quantity": 71,
                        "overage_price": 85
                      }
                    ],
                    "minimum_price": 51,
                    "percentage": 189.75
                  },
                  "discounts": [
                    {
                      "id": "id2",
                      "value": 72.24,
                      "discount_type": "discount_type0",
                      "status": "status4",
                      "created_at": "2016-03-13T12:52:32.123Z",
                      "cycles": 80,
                      "deleted_at": "2016-03-13T12:52:32.123Z",
                      "description": "description2",
                      "subscription": {},
                      "subscription_item": {}
                    }
                  ],
                  "increments": [
                    {},
                    {}
                  ],
                  "subscription": {},
                  "name": "name1",
                  "quantity": 125,
                  "cycles": 111,
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                }
              }
            ],
            "split": {
              "enabled": true,
              "rules": [
                {
                  "type": "type5",
                  "amount": 167,
                  "recipient": {
                    "id": "id7",
                    "name": "name7",
                    "email": "email9",
                    "document": "document1",
                    "description": "description7",
                    "type": "type3",
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
                        "key0": "metadata4"
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
                      "key0": "metadata4"
                    },
                    "automatic_anticipation_settings": {
                      "enabled": true,
                      "type": "type1",
                      "volume_percentage": 227,
                      "delay": 63,
                      "days": [
                        39,
                        40,
                        41
                      ]
                    },
                    "transfer_settings": {
                      "transfer_enabled": true,
                      "transfer_interval": "transfer_interval7",
                      "transfer_day": 137
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
                },
                {
                  "type": "type6",
                  "amount": 168,
                  "recipient": {
                    "id": "id8",
                    "name": "name8",
                    "email": "email8",
                    "document": "document2",
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
                        "key1": "metadata2"
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
                      "key1": "metadata4",
                      "key2": "metadata3"
                    },
                    "automatic_anticipation_settings": {
                      "enabled": false,
                      "type": "type2",
                      "volume_percentage": 226,
                      "delay": 64,
                      "days": [
                        40
                      ]
                    },
                    "transfer_settings": {
                      "transfer_enabled": false,
                      "transfer_interval": "transfer_interval6",
                      "transfer_day": 136
                    },
                    "code": "code6",
                    "payment_mode": "payment_mode2"
                  },
                  "gateway_id": "gateway_id6",
                  "options": {
                    "liable": false,
                    "charge_processing_fee": false,
                    "charge_remainder_fee": "charge_remainder_fee2"
                  },
                  "id": "id6"
                },
                {
                  "type": "type7",
                  "amount": 169,
                  "recipient": {
                    "id": "id9",
                    "name": "name9",
                    "email": "email7",
                    "document": "document3",
                    "description": "description9",
                    "type": "type1",
                    "status": "status1",
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
                        "key0": "metadata2",
                        "key1": "metadata1",
                        "key2": "metadata0"
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
                      "key0": "metadata6",
                      "key1": "metadata5"
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
                      "transfer_interval": "transfer_interval5",
                      "transfer_day": 135
                    },
                    "code": "code7",
                    "payment_mode": "payment_mode3"
                  },
                  "gateway_id": "gateway_id7",
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
          "status": "status1",
          "duration": 21,
          "created_at": "created_at7",
          "updated_at": "updated_at5",
          "cycle": 217
        },
        "shipping": {
          "amount": 165,
          "description": "description3",
          "recipient_name": "recipient_name1",
          "recipient_phone": "recipient_phone5",
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
              "address": {},
              "metadata": {
                "key0": "metadata0",
                "key1": "metadata1"
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
              "fb_id": 213,
              "code": "code7",
              "document_type": "document_type7"
            },
            "metadata": {
              "key0": "metadata0"
            },
            "line_1": "line_13",
            "line_2": "line_27",
            "deleted_at": "2016-03-13T12:52:32.123Z"
          },
          "max_delivery_date": "2016-03-13T12:52:32.123Z",
          "estimated_delivery_date": "2016-03-13T12:52:32.123Z",
          "type": "type3"
        },
        "metadata": {
          "key0": "metadata0",
          "key1": "metadata9"
        },
        "due_at": "2016-03-13T12:52:32.123Z",
        "canceled_at": "2016-03-13T12:52:32.123Z",
        "billing_at": "2016-03-13T12:52:32.123Z",
        "subscription_id": "subscription_id9"
      },
      "order": {
        "id": "id9",
        "code": "code7",
        "currency": "currency9",
        "items": [
          {
            "id": "id6",
            "amount": 148,
            "description": "description6",
            "quantity": 6,
            "category": "category4",
            "code": "code4"
          },
          {
            "id": "id7",
            "amount": 149,
            "description": "description7",
            "quantity": 7,
            "category": "category5",
            "code": "code5"
          },
          {
            "id": "id8",
            "amount": 150,
            "description": "description8",
            "quantity": 8,
            "category": "category6",
            "code": "code6"
          }
        ],
        "customer": {
          "id": "id9",
          "name": "name9",
          "email": "email7",
          "delinquent": true,
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "document": "document3",
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
              "key0": "metadata4"
            },
            "line_1": "line_19",
            "line_2": "line_23",
            "deleted_at": "2016-03-13T12:52:32.123Z"
          },
          "metadata": {
            "key0": "metadata6"
          },
          "phones": {
            "home_phone": {
              "country_code": "country_code1",
              "number": "number9",
              "area_code": "area_code1"
            },
            "mobile_phone": {
              "country_code": "country_code9",
              "number": "number1",
              "area_code": "area_code1"
            }
          },
          "fb_id": 121,
          "code": "code7",
          "document_type": "document_type7"
        },
        "status": "status1",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "charges": [
          {}
        ],
        "invoice_url": "invoice_url1",
        "shipping": {
          "amount": 187,
          "description": "description3",
          "recipient_name": "recipient_name1",
          "recipient_phone": "recipient_phone5",
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
              "address": {},
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
              "fb_id": 235,
              "code": "code7",
              "document_type": "document_type7"
            },
            "metadata": {
              "key0": "metadata0",
              "key1": "metadata1"
            },
            "line_1": "line_13",
            "line_2": "line_27",
            "deleted_at": "2016-03-13T12:52:32.123Z"
          },
          "max_delivery_date": "2016-03-13T12:52:32.123Z",
          "estimated_delivery_date": "2016-03-13T12:52:32.123Z",
          "type": "type3"
        },
        "metadata": {
          "key0": "metadata6",
          "key1": "metadata5"
        },
        "checkouts": [
          {
            "id": "id6",
            "amount": 98,
            "default_payment_method": "default_payment_method6",
            "success_url": "success_url8",
            "payment_url": "payment_url0",
            "gateway_affiliation_id": "gateway_affiliation_id2",
            "accepted_payment_methods": [
              "accepted_payment_methods9"
            ],
            "status": "status8",
            "skip_checkout_success_page": false,
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "canceled_at": "2016-03-13T12:52:32.123Z",
            "customer_editable": false,
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
                  "key0": "metadata3"
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
                  "number": "number6",
                  "area_code": "area_code8"
                }
              },
              "fb_id": 130,
              "code": "code4",
              "document_type": "document_type4"
            },
            "billingaddress": {
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
              "customer": {
                "id": "id4",
                "name": "name4",
                "email": "email8",
                "delinquent": false,
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "document": "document8",
                "type": "type4",
                "fb_access_token": "fb_access_token8",
                "address": {},
                "metadata": {
                  "key0": "metadata5",
                  "key1": "metadata6",
                  "key2": "metadata7"
                },
                "phones": {
                  "home_phone": {
                    "country_code": "country_code6",
                    "number": "number4",
                    "area_code": "area_code6"
                  },
                  "mobile_phone": {
                    "country_code": "country_code6",
                    "number": "number4",
                    "area_code": "area_code6"
                  }
                },
                "fb_id": 212,
                "code": "code2",
                "document_type": "document_type2"
              },
              "metadata": {
                "key0": "metadata5",
                "key1": "metadata6"
              },
              "line_1": "line_18",
              "line_2": "line_22",
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            "credit_card": {
              "statementDescriptor": "statementDescriptor0",
              "installments": [
                {
                  "number": "number7",
                  "total": 55
                },
                {
                  "number": "number8",
                  "total": 56
                }
              ],
              "authentication": {
                "type": "type6",
                "threed_secure": {
                  "mpi": "mpi2",
                  "eci": "eci4",
                  "cavv": "cavv0",
                  "transaction_Id": "transaction_Id0",
                  "success_url": "success_url6"
                }
              }
            },
            "boleto": {
              "due_at": "2016-03-13T12:52:32.123Z",
              "instructions": "instructions4"
            },
            "billing_address_editable": false,
            "shipping": {
              "amount": 196,
              "description": "description0",
              "recipient_name": "recipient_name8",
              "recipient_phone": "recipient_phone2",
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
                  "type": "type6",
                  "fb_access_token": "fb_access_token0",
                  "address": {},
                  "metadata": {
                    "key0": "metadata7",
                    "key1": "metadata8",
                    "key2": "metadata9"
                  },
                  "phones": {
                    "home_phone": {
                      "country_code": "country_code8",
                      "number": "number6",
                      "area_code": "area_code8"
                    },
                    "mobile_phone": {
                      "country_code": "country_code8",
                      "number": "number6",
                      "area_code": "area_code8"
                    }
                  },
                  "fb_id": 244,
                  "code": "code4",
                  "document_type": "document_type4"
                },
                "metadata": {
                  "key0": "metadata7",
                  "key1": "metadata8"
                },
                "line_1": "line_10",
                "line_2": "line_24",
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              "max_delivery_date": "2016-03-13T12:52:32.123Z",
              "estimated_delivery_date": "2016-03-13T12:52:32.123Z",
              "type": "type0"
            },
            "shippable": false,
            "closed_at": "2016-03-13T12:52:32.123Z",
            "expires_at": "2016-03-13T12:52:32.123Z",
            "currency": "currency6",
            "accepted_brands": [
              "accepted_brands2"
            ]
          },
          {
            "id": "id7",
            "amount": 99,
            "default_payment_method": "default_payment_method7",
            "success_url": "success_url9",
            "payment_url": "payment_url1",
            "gateway_affiliation_id": "gateway_affiliation_id3",
            "accepted_payment_methods": [
              "accepted_payment_methods0",
              "accepted_payment_methods1"
            ],
            "status": "status9",
            "skip_checkout_success_page": true,
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "canceled_at": "2016-03-13T12:52:32.123Z",
            "customer_editable": true,
            "customer": {
              "id": "id7",
              "name": "name7",
              "email": "email1",
              "delinquent": true,
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "document": "document1",
              "type": "type7",
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
                  "key0": "metadata4",
                  "key1": "metadata5"
                },
                "line_1": "line_17",
                "line_2": "line_21",
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              "metadata": {
                "key0": "metadata2",
                "key1": "metadata1"
              },
              "phones": {
                "home_phone": {
                  "country_code": "country_code9",
                  "number": "number7",
                  "area_code": "area_code9"
                },
                "mobile_phone": {
                  "country_code": "country_code9",
                  "number": "number7",
                  "area_code": "area_code9"
                }
              },
              "fb_id": 131,
              "code": "code5",
              "document_type": "document_type5"
            },
            "billingaddress": {
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
                  "key0": "metadata6"
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
                "fb_id": 213,
                "code": "code3",
                "document_type": "document_type3"
              },
              "metadata": {
                "key0": "metadata6",
                "key1": "metadata7",
                "key2": "metadata8"
              },
              "line_1": "line_19",
              "line_2": "line_23",
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            "credit_card": {
              "statementDescriptor": "statementDescriptor1",
              "installments": [
                {
                  "number": "number8",
                  "total": 56
                },
                {
                  "number": "number9",
                  "total": 57
                },
                {
                  "number": "number0",
                  "total": 58
                }
              ],
              "authentication": {
                "type": "type7",
                "threed_secure": {
                  "mpi": "mpi3",
                  "eci": "eci5",
                  "cavv": "cavv1",
                  "transaction_Id": "transaction_Id1",
                  "success_url": "success_url7"
                }
              }
            },
            "boleto": {
              "due_at": "2016-03-13T12:52:32.123Z",
              "instructions": "instructions5"
            },
            "billing_address_editable": true,
            "shipping": {
              "amount": 197,
              "description": "description1",
              "recipient_name": "recipient_name9",
              "recipient_phone": "recipient_phone3",
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
                "customer": {
                  "id": "id7",
                  "name": "name7",
                  "email": "email1",
                  "delinquent": true,
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "updated_at": "2016-03-13T12:52:32.123Z",
                  "document": "document1",
                  "type": "type7",
                  "fb_access_token": "fb_access_token1",
                  "address": {},
                  "metadata": {
                    "key0": "metadata8"
                  },
                  "phones": {
                    "home_phone": {
                      "country_code": "country_code9",
                      "number": "number7",
                      "area_code": "area_code9"
                    },
                    "mobile_phone": {
                      "country_code": "country_code9",
                      "number": "number7",
                      "area_code": "area_code9"
                    }
                  },
                  "fb_id": 245,
                  "code": "code5",
                  "document_type": "document_type5"
                },
                "metadata": {
                  "key0": "metadata8",
                  "key1": "metadata9",
                  "key2": "metadata0"
                },
                "line_1": "line_11",
                "line_2": "line_25",
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              "max_delivery_date": "2016-03-13T12:52:32.123Z",
              "estimated_delivery_date": "2016-03-13T12:52:32.123Z",
              "type": "type1"
            },
            "shippable": true,
            "closed_at": "2016-03-13T12:52:32.123Z",
            "expires_at": "2016-03-13T12:52:32.123Z",
            "currency": "currency7",
            "accepted_brands": [
              "accepted_brands3",
              "accepted_brands4"
            ]
          }
        ],
        "ip": "ip3",
        "session_id": "session_id1",
        "location": {
          "latitude": "latitude7",
          "longitude": "longitude7"
        },
        "closed": true
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
      "paid_at": "2016-03-13T12:52:32.123Z",
      "canceled_at": "2016-03-13T12:52:32.123Z"
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

