
# Get Usage Response

Response object for getting a usage

## Structure

`GetUsageResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*string` | Required | Id |
| `Quantity` | `*int` | Required | Quantity |
| `Description` | `*string` | Required | Description |
| `UsedAt` | `*time.Time` | Required | Used at |
| `CreatedAt` | `*time.Time` | Required | Creation date |
| `Status` | `*string` | Required | Status |
| `DeletedAt` | `Optional[time.Time]` | Optional | - |
| `SubscriptionItem` | [`*models.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md) | Required | Subscription item |
| `Code` | `Optional[string]` | Optional | Identification code in the client system |
| `Group` | `Optional[string]` | Optional | Identification group in the client system |
| `Amount` | `Optional[int]` | Optional | Field used in item scheme type 'Percent' |

## Example (as JSON)

```json
{
  "id": "id0",
  "quantity": 68,
  "description": "description0",
  "used_at": "2016-03-13T12:52:32.123Z",
  "created_at": "2016-03-13T12:52:32.123Z",
  "status": "status8",
  "subscription_item": {
    "id": "id6",
    "description": "description4",
    "status": "status2",
    "created_at": "2016-03-13T12:52:32.123Z",
    "updated_at": "2016-03-13T12:52:32.123Z",
    "pricing_scheme": {
      "price": 178,
      "scheme_type": "scheme_type2",
      "price_brackets": [
        {
          "start_quantity": 131,
          "price": 69,
          "end_quantity": 139,
          "overage_price": 153
        }
      ],
      "minimum_price": 18,
      "percentage": 231.4
    },
    "discounts": [
      {
        "id": "id7",
        "value": 30.59,
        "discount_type": "discount_type5",
        "status": "status9",
        "created_at": "2016-03-13T12:52:32.123Z",
        "cycles": 11,
        "deleted_at": "2016-03-13T12:52:32.123Z",
        "description": "description7",
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
        "subscription_item": {}
      }
    ],
    "increments": [
      {
        "id": "id1",
        "value": 103.83,
        "increment_type": "increment_type3",
        "status": "status7",
        "created_at": "2016-03-13T12:52:32.123Z",
        "cycles": 89,
        "deleted_at": "2016-03-13T12:52:32.123Z",
        "description": "description9",
        "subscription": {
          "id": "id3",
          "code": "code1",
          "start_at": "2016-03-13T12:52:32.123Z",
          "interval": "interval1",
          "interval_count": 155,
          "billing_type": "billing_type7",
          "current_cycle": {
            "start_at": "2016-03-13T12:52:32.123Z",
            "end_at": "2016-03-13T12:52:32.123Z",
            "id": "id1",
            "billing_at": "2016-03-13T12:52:32.123Z",
            "subscription": {},
            "status": "status3",
            "duration": 139,
            "created_at": "created_at9",
            "updated_at": "updated_at7",
            "cycle": 79
          },
          "payment_method": "payment_method3",
          "currency": "currency3",
          "installments": 67,
          "status": "status5",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "customer": {
            "id": "id3",
            "name": "name3",
            "email": "email3",
            "delinquent": true,
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "document": "document7",
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
                "key0": "metadata0",
                "key1": "metadata9"
              },
              "line_1": "line_13",
              "line_2": "line_27",
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            "metadata": {
              "key0": "metadata0",
              "key1": "metadata9",
              "key2": "metadata8"
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
            "fb_id": 59,
            "code": "code1",
            "document_type": "document_type1"
          },
          "card": {
            "id": "id7",
            "last_four_digits": "last_four_digits3",
            "brand": "brand1",
            "holder_name": "holder_name3",
            "exp_month": 39,
            "exp_year": 1,
            "status": "status1",
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
                  "key1": "metadata9"
                },
                "line_1": "line_17",
                "line_2": "line_21",
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              "metadata": {
                "key0": "metadata6",
                "key1": "metadata7"
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
              "fb_id": 245,
              "code": "code5",
              "document_type": "document_type5"
            },
            "metadata": {
              "key0": "metadata6",
              "key1": "metadata7",
              "key2": "metadata8"
            },
            "type": "type3",
            "holder_document": "holder_document9",
            "deleted_at": "2016-03-13T12:52:32.123Z",
            "first_six_digits": "first_six_digits7",
            "label": "label7"
          },
          "items": [
            {}
          ],
          "statement_descriptor": "statement_descriptor3",
          "metadata": {
            "key0": "metadata0"
          },
          "setup": {
            "id": "id7",
            "description": "description7",
            "amount": 91,
            "status": "status9"
          },
          "gateway_affiliation_id": "gateway_affiliation_id9",
          "next_billing_at": "2016-03-13T12:52:32.123Z",
          "billing_day": 39,
          "minimum_price": 153,
          "increments": [
            {},
            {},
            {}
          ],
          "split": {
            "enabled": true,
            "rules": [
              {
                "type": "type7",
                "amount": 227,
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
                      "key1": "metadata9",
                      "key2": "metadata8"
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
                    }
                  ],
                  "metadata": {
                    "key0": "metadata8",
                    "key1": "metadata9",
                    "key2": "metadata0"
                  },
                  "automatic_anticipation_settings": {
                    "enabled": true,
                    "type": "type1",
                    "volume_percentage": 167,
                    "delay": 123,
                    "days": [
                      99
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": true,
                    "transfer_interval": "transfer_interval9",
                    "transfer_day": 77
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
        "subscription_item": {}
      }
    ],
    "subscription": {
      "id": "id8",
      "code": "code6",
      "start_at": "2016-03-13T12:52:32.123Z",
      "interval": "interval6",
      "interval_count": 246,
      "billing_type": "billing_type2",
      "payment_method": "payment_method8",
      "currency": "currency8",
      "installments": 158,
      "status": "status0",
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "card": {
        "id": "id2",
        "last_four_digits": "last_four_digits8",
        "brand": "brand6",
        "holder_name": "holder_name8",
        "exp_month": 204,
        "exp_year": 92,
        "status": "status4",
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
              "country_code": "country_code4",
              "number": "number6",
              "area_code": "area_code6"
            }
          },
          "fb_id": 80,
          "code": "code0",
          "document_type": "document_type0"
        },
        "metadata": {
          "key0": "metadata9"
        },
        "type": "type8",
        "holder_document": "holder_document4",
        "deleted_at": "2016-03-13T12:52:32.123Z",
        "first_six_digits": "first_six_digits2",
        "label": "label2"
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
      "statement_descriptor": "statement_descriptor8",
      "metadata": {
        "key0": "metadata5",
        "key1": "metadata4"
      },
      "setup": {
        "id": "id2",
        "description": "description2",
        "amount": 182,
        "status": "status4"
      },
      "gateway_affiliation_id": "gateway_affiliation_id4",
      "increments": [
        {
          "id": "id7",
          "value": 184.59,
          "increment_type": "increment_type9",
          "status": "status1",
          "created_at": "2016-03-13T12:52:32.123Z",
          "cycles": 205,
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "description": "description3",
          "subscription": {},
          "subscription_item": {}
        }
      ],
      "split": {
        "enabled": false,
        "rules": [
          {
            "type": "type8",
            "amount": 62,
            "recipient": {
              "id": "id0",
              "name": "name0",
              "email": "email6",
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
                  "key0": "metadata1"
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
                "key0": "metadata7"
              },
              "automatic_anticipation_settings": {
                "enabled": false,
                "type": "type6",
                "volume_percentage": 76,
                "delay": 214,
                "days": [
                  190,
                  191,
                  192
                ]
              },
              "transfer_settings": {
                "transfer_enabled": false,
                "transfer_interval": "transfer_interval4",
                "transfer_day": 242
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
        "id": "id6",
        "billing_at": "2016-03-13T12:52:32.123Z",
        "subscription": {},
        "status": "status8",
        "duration": 230,
        "created_at": "created_at4",
        "updated_at": "updated_at2",
        "cycle": 170
      },
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
        "fb_id": 150,
        "code": "code6",
        "document_type": "document_type6"
      },
      "next_billing_at": "2016-03-13T12:52:32.123Z",
      "billing_day": 130,
      "minimum_price": 244
    },
    "name": "name6",
    "quantity": 56,
    "cycles": 180,
    "deleted_at": "2016-03-13T12:52:32.123Z"
  },
  "deleted_at": "2016-03-13T12:52:32.123Z",
  "code": "code8",
  "group": "group8",
  "amount": 46
}
```

