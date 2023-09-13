
# Get Subscription Response

## Structure

`GetSubscriptionResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*string` | Required | - |
| `Code` | `*string` | Required | - |
| `StartAt` | `*time.Time` | Required | - |
| `Interval` | `*string` | Required | - |
| `IntervalCount` | `*int` | Required | - |
| `BillingType` | `*string` | Required | - |
| `CurrentCycle` | [`Optional[models.GetPeriodResponse]`](../../doc/models/get-period-response.md) | Optional | - |
| `PaymentMethod` | `*string` | Required | - |
| `Currency` | `*string` | Required | - |
| `Installments` | `*int` | Required | - |
| `Status` | `*string` | Required | - |
| `CreatedAt` | `*time.Time` | Required | - |
| `UpdatedAt` | `*time.Time` | Required | - |
| `Customer` | [`Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | - |
| `Card` | [`*models.GetCardResponse`](../../doc/models/get-card-response.md) | Required | - |
| `Items` | [`[]models.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md) | Required | - |
| `StatementDescriptor` | `*string` | Required | - |
| `Metadata` | `map[string]string` | Required | - |
| `Setup` | [`*models.GetSetupResponse`](../../doc/models/get-setup-response.md) | Required | - |
| `GatewayAffiliationId` | `*string` | Required | Affiliation Code |
| `NextBillingAt` | `Optional[time.Time]` | Optional | - |
| `BillingDay` | `Optional[int]` | Optional | - |
| `MinimumPrice` | `Optional[int]` | Optional | - |
| `CanceledAt` | `Optional[time.Time]` | Optional | - |
| `Discounts` | [`Optional[[]models.GetDiscountResponse]`](../../doc/models/get-discount-response.md) | Optional | Subscription discounts |
| `Increments` | [`[]models.GetIncrementResponse`](../../doc/models/get-increment-response.md) | Required | Subscription increments |
| `BoletoDueDays` | `Optional[int]` | Optional | Days until boleto expires |
| `Split` | [`*models.GetSubscriptionSplitResponse`](../../doc/models/get-subscription-split-response.md) | Required | Subscription's split response |
| `Boleto` | [`Optional[models.GetSubscriptionBoletoResponse]`](../../doc/models/get-subscription-boleto-response.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "code": "code8",
  "start_at": "2016-03-13T12:52:32.123Z",
  "interval": "interval2",
  "interval_count": 82,
  "billing_type": "billing_type6",
  "payment_method": "payment_method0",
  "currency": "currency0",
  "installments": 250,
  "status": "status8",
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z",
  "card": {
    "id": "id6",
    "last_four_digits": "last_four_digits2",
    "brand": "brand0",
    "holder_name": "holder_name2",
    "exp_month": 228,
    "exp_year": 68,
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
        "line_1": "line_16",
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
      "fb_id": 56,
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
    {
      "id": "id7",
      "description": "description7",
      "status": "status1",
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "pricing_scheme": {
        "price": 149,
        "scheme_type": "scheme_type1",
        "price_brackets": [
          {
            "start_quantity": 60,
            "price": 2,
            "end_quantity": 68,
            "overage_price": 82
          }
        ],
        "minimum_price": 53,
        "percentage": 25.89
      },
      "discounts": [
        {
          "id": "id8",
          "value": 236.1,
          "discount_type": "discount_type6",
          "status": "status0",
          "created_at": "2016-03-13T12:52:32.123Z",
          "cycles": 82,
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "description": "description8",
          "subscription": {
            "id": "id4",
            "code": "code2",
            "start_at": "2016-03-13T12:52:32.123Z",
            "interval": "interval2",
            "interval_count": 180,
            "billing_type": "billing_type2",
            "current_cycle": {
              "start_at": "2016-03-13T12:52:32.123Z",
              "end_at": "2016-03-13T12:52:32.123Z",
              "id": "id2",
              "billing_at": "2016-03-13T12:52:32.123Z",
              "subscription": {},
              "status": "status4",
              "duration": 164,
              "created_at": "created_at0",
              "updated_at": "updated_at8",
              "cycle": 104
            },
            "payment_method": "payment_method6",
            "currency": "currency4",
            "installments": 92,
            "status": "status4",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
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
                "line_1": "line_16",
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
              "fb_id": 84,
              "code": "code2",
              "document_type": "document_type2"
            },
            "card": {
              "id": "id2",
              "last_four_digits": "last_four_digits8",
              "brand": "brand6",
              "holder_name": "holder_name8",
              "exp_month": 70,
              "exp_year": 226,
              "status": "status6",
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
                "fb_id": 202,
                "code": "code6",
                "document_type": "document_type6"
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
              {}
            ],
            "statement_descriptor": "statement_descriptor4",
            "metadata": {
              "key0": "metadata9",
              "key1": "metadata0",
              "key2": "metadata1"
            },
            "setup": {
              "id": "id8",
              "description": "description2",
              "amount": 116,
              "status": "status0"
            },
            "gateway_affiliation_id": "gateway_affiliation_id0",
            "next_billing_at": "2016-03-13T12:52:32.123Z",
            "billing_day": 192,
            "minimum_price": 78,
            "increments": [
              {
                "id": "id7",
                "value": 23.99,
                "increment_type": "increment_type9",
                "status": "status1",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 137,
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
                  "type": "type6",
                  "amount": 252,
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
                      "volume_percentage": 58,
                      "delay": 232,
                      "days": [
                        208,
                        209,
                        210
                      ]
                    },
                    "transfer_settings": {
                      "transfer_enabled": false,
                      "transfer_interval": "transfer_interval0",
                      "transfer_day": 224
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
                },
                {
                  "type": "type5",
                  "amount": 253,
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
                      "volume_percentage": 59,
                      "delay": 231,
                      "days": [
                        207,
                        208
                      ]
                    },
                    "transfer_settings": {
                      "transfer_enabled": true,
                      "transfer_interval": "transfer_interval1",
                      "transfer_day": 225
                    },
                    "code": "code1",
                    "payment_mode": "payment_mode3"
                  },
                  "gateway_id": "gateway_id5",
                  "options": {
                    "liable": true,
                    "charge_processing_fee": true,
                    "charge_remainder_fee": "charge_remainder_fee5"
                  },
                  "id": "id5"
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
          "value": 92.72,
          "increment_type": "increment_type2",
          "status": "status8",
          "created_at": "2016-03-13T12:52:32.123Z",
          "cycles": 176,
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "description": "description0",
          "subscription": {
            "id": "id4",
            "code": "code2",
            "start_at": "2016-03-13T12:52:32.123Z",
            "interval": "interval2",
            "interval_count": 242,
            "billing_type": "billing_type8",
            "current_cycle": {
              "start_at": "2016-03-13T12:52:32.123Z",
              "end_at": "2016-03-13T12:52:32.123Z",
              "id": "id2",
              "billing_at": "2016-03-13T12:52:32.123Z",
              "subscription": {},
              "status": "status4",
              "duration": 226,
              "created_at": "created_at0",
              "updated_at": "updated_at8",
              "cycle": 166
            },
            "payment_method": "payment_method4",
            "currency": "currency4",
            "installments": 154,
            "status": "status6",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
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
                  "country_code": "country_code4",
                  "number": "number6",
                  "area_code": "area_code6"
                }
              },
              "fb_id": 146,
              "code": "code2",
              "document_type": "document_type2"
            },
            "card": {
              "id": "id8",
              "last_four_digits": "last_four_digits4",
              "brand": "brand2",
              "holder_name": "holder_name4",
              "exp_month": 208,
              "exp_year": 88,
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
                    "key0": "metadata1"
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
                "fb_id": 76,
                "code": "code6",
                "document_type": "document_type6"
              },
              "metadata": {
                "key0": "metadata5",
                "key1": "metadata6"
              },
              "type": "type2",
              "holder_document": "holder_document8",
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "first_six_digits": "first_six_digits8",
              "label": "label8"
            },
            "items": [
              {},
              {}
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
              "amount": 178,
              "status": "status0"
            },
            "gateway_affiliation_id": "gateway_affiliation_id0",
            "next_billing_at": "2016-03-13T12:52:32.123Z",
            "billing_day": 126,
            "minimum_price": 240,
            "increments": [
              {}
            ],
            "split": {
              "enabled": false,
              "rules": [
                {
                  "type": "type6",
                  "amount": 58,
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
                        "key0": "metadata1",
                        "key1": "metadata0"
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
                      "key0": "metadata7",
                      "key1": "metadata8"
                    },
                    "automatic_anticipation_settings": {
                      "enabled": false,
                      "type": "type0",
                      "volume_percentage": 80,
                      "delay": 210,
                      "days": [
                        186,
                        187
                      ]
                    },
                    "transfer_settings": {
                      "transfer_enabled": false,
                      "transfer_interval": "transfer_interval8",
                      "transfer_day": 246
                    },
                    "code": "code4",
                    "payment_mode": "payment_mode0"
                  },
                  "gateway_id": "gateway_id6",
                  "options": {
                    "liable": false,
                    "charge_processing_fee": false,
                    "charge_remainder_fee": "charge_remainder_fee0"
                  },
                  "id": "id4"
                },
                {
                  "type": "type5",
                  "amount": 59,
                  "recipient": {
                    "id": "id7",
                    "name": "name7",
                    "email": "email9",
                    "document": "document9",
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
                        "key0": "metadata2"
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
                      "key0": "metadata6"
                    },
                    "automatic_anticipation_settings": {
                      "enabled": true,
                      "type": "type9",
                      "volume_percentage": 79,
                      "delay": 211,
                      "days": [
                        187,
                        188,
                        189
                      ]
                    },
                    "transfer_settings": {
                      "transfer_enabled": true,
                      "transfer_interval": "transfer_interval7",
                      "transfer_day": 245
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
        "interval_count": 175,
        "billing_type": "billing_type1",
        "payment_method": "payment_method3",
        "currency": "currency7",
        "installments": 87,
        "status": "status9",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "card": {
          "id": "id1",
          "last_four_digits": "last_four_digits7",
          "brand": "brand5",
          "holder_name": "holder_name7",
          "exp_month": 19,
          "exp_year": 21,
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
                "key0": "metadata4"
              },
              "line_1": "line_19",
              "line_2": "line_25",
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            "metadata": {
              "key0": "metadata2"
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
            "fb_id": 9,
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
          "amount": 111,
          "status": "status3"
        },
        "gateway_affiliation_id": "gateway_affiliation_id3",
        "increments": [
          {
            "id": "id4",
            "value": 11.96,
            "increment_type": "increment_type4",
            "status": "status4",
            "created_at": "2016-03-13T12:52:32.123Z",
            "cycles": 60,
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
              "amount": 247,
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
                  "volume_percentage": 147,
                  "delay": 143,
                  "days": [
                    119,
                    120
                  ]
                },
                "transfer_settings": {
                  "transfer_enabled": true,
                  "transfer_interval": "transfer_interval5",
                  "transfer_day": 57
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
          "duration": 159,
          "created_at": "created_at3",
          "updated_at": "updated_at1",
          "cycle": 99
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
              "key0": "metadata6",
              "key1": "metadata5",
              "key2": "metadata4"
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
              "country_code": "country_code9",
              "number": "number1",
              "area_code": "area_code1"
            }
          },
          "fb_id": 79,
          "code": "code5",
          "document_type": "document_type5"
        },
        "next_billing_at": "2016-03-13T12:52:32.123Z",
        "billing_day": 59,
        "minimum_price": 83
      },
      "name": "name7",
      "quantity": 127,
      "cycles": 109,
      "deleted_at": "2016-03-13T12:52:32.123Z"
    }
  ],
  "statement_descriptor": "statement_descriptor0",
  "metadata": {
    "key0": "metadata3",
    "key1": "metadata4",
    "key2": "metadata5"
  },
  "setup": {
    "id": "id6",
    "description": "description4",
    "amount": 110,
    "status": "status2"
  },
  "gateway_affiliation_id": "gateway_affiliation_id4",
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
  "split": {
    "enabled": false,
    "rules": [
      {
        "type": "type4",
        "amount": 246,
        "recipient": {
          "id": "id2",
          "name": "name2",
          "email": "email4",
          "document": "document4",
          "description": "description8",
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
              "key0": "metadata7"
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
            }
          ],
          "metadata": {
            "key0": "metadata1"
          },
          "automatic_anticipation_settings": {
            "enabled": false,
            "type": "type4",
            "volume_percentage": 72,
            "delay": 218,
            "days": [
              194,
              195,
              196
            ]
          },
          "transfer_settings": {
            "transfer_enabled": false,
            "transfer_interval": "transfer_interval2",
            "transfer_day": 238
          },
          "code": "code0",
          "payment_mode": "payment_mode4"
        },
        "gateway_id": "gateway_id4",
        "options": {
          "liable": false,
          "charge_processing_fee": false,
          "charge_remainder_fee": "charge_remainder_fee4"
        },
        "id": "id6"
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
    "id": "id8",
    "billing_at": "2016-03-13T12:52:32.123Z",
    "subscription": {
      "id": "id6",
      "code": "code4",
      "start_at": "2016-03-13T12:52:32.123Z",
      "interval": "interval4",
      "interval_count": 24,
      "billing_type": "billing_type0",
      "current_cycle": {},
      "payment_method": "payment_method4",
      "currency": "currency6",
      "installments": 192,
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
            "country_code": "country_code2",
            "number": "number4",
            "area_code": "area_code8"
          }
        },
        "fb_id": 184,
        "code": "code4",
        "document_type": "document_type4"
      },
      "card": {
        "id": "id0",
        "last_four_digits": "last_four_digits6",
        "brand": "brand4",
        "holder_name": "holder_name6",
        "exp_month": 170,
        "exp_year": 126,
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
          "fb_id": 114,
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
        {
          "id": "id3",
          "description": "description3",
          "status": "status5",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "pricing_scheme": {
            "price": 197,
            "scheme_type": "scheme_type5",
            "price_brackets": [
              {
                "start_quantity": 12,
                "price": 50,
                "end_quantity": 20,
                "overage_price": 34
              },
              {
                "start_quantity": 13,
                "price": 49,
                "end_quantity": 21,
                "overage_price": 35
              }
            ],
            "minimum_price": 101,
            "percentage": 132.93
          },
          "discounts": [
            {
              "id": "id4",
              "value": 130.56,
              "discount_type": "discount_type2",
              "status": "status6",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 24,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description4",
              "subscription": {},
              "subscription_item": {}
            },
            {
              "id": "id5",
              "value": 130.57,
              "discount_type": "discount_type3",
              "status": "status7",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 25,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description5",
              "subscription": {},
              "subscription_item": {}
            }
          ],
          "increments": [
            {
              "id": "id2",
              "value": 241.84,
              "increment_type": "increment_type4",
              "status": "status6",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 112,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description2",
              "subscription": {},
              "subscription_item": {}
            },
            {
              "id": "id3",
              "value": 241.85,
              "increment_type": "increment_type5",
              "status": "status5",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 111,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description3",
              "subscription": {},
              "subscription_item": {}
            },
            {
              "id": "id4",
              "value": 241.86,
              "increment_type": "increment_type6",
              "status": "status4",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 110,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description4",
              "subscription": {},
              "subscription_item": {}
            }
          ],
          "subscription": {},
          "name": "name3",
          "quantity": 69,
          "cycles": 89,
          "deleted_at": "2016-03-13T12:52:32.123Z"
        }
      ],
      "statement_descriptor": "statement_descriptor6",
      "metadata": {
        "key0": "metadata3"
      },
      "setup": {
        "id": "id0",
        "description": "description0",
        "amount": 216,
        "status": "status2"
      },
      "gateway_affiliation_id": "gateway_affiliation_id2",
      "next_billing_at": "2016-03-13T12:52:32.123Z",
      "billing_day": 164,
      "minimum_price": 22,
      "increments": [
        {
          "id": "id5",
          "value": 15.97,
          "increment_type": "increment_type3",
          "status": "status3",
          "created_at": "2016-03-13T12:52:32.123Z",
          "cycles": 171,
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "description": "description5",
          "subscription": {},
          "subscription_item": {
            "id": "id1",
            "description": "description9",
            "status": "status7",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "pricing_scheme": {
              "price": 75,
              "scheme_type": "scheme_type7",
              "price_brackets": [
                {
                  "start_quantity": 134,
                  "price": 72,
                  "end_quantity": 142,
                  "overage_price": 156
                },
                {
                  "start_quantity": 135,
                  "price": 73,
                  "end_quantity": 143,
                  "overage_price": 157
                }
              ],
              "minimum_price": 21,
              "percentage": 210.95
            },
            "discounts": [
              {
                "id": "id2",
                "value": 51.04,
                "discount_type": "discount_type0",
                "status": "status4",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 8,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description2",
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
            "name": "name1",
            "quantity": 53,
            "cycles": 183,
            "deleted_at": "2016-03-13T12:52:32.123Z"
          }
        },
        {
          "id": "id6",
          "value": 15.98,
          "increment_type": "increment_type2",
          "status": "status2",
          "created_at": "2016-03-13T12:52:32.123Z",
          "cycles": 170,
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
              "price": 76,
              "scheme_type": "scheme_type6",
              "price_brackets": [
                {
                  "start_quantity": 133,
                  "price": 71,
                  "end_quantity": 141,
                  "overage_price": 155
                }
              ],
              "minimum_price": 20,
              "percentage": 210.94
            },
            "discounts": [
              {
                "id": "id3",
                "value": 51.05,
                "discount_type": "discount_type1",
                "status": "status5",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 9,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description3",
                "subscription": {},
                "subscription_item": {}
              },
              {
                "id": "id4",
                "value": 51.06,
                "discount_type": "discount_type2",
                "status": "status6",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 10,
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
            "quantity": 54,
            "cycles": 182,
            "deleted_at": "2016-03-13T12:52:32.123Z"
          }
        },
        {
          "id": "id7",
          "value": 15.99,
          "increment_type": "increment_type1",
          "status": "status1",
          "created_at": "2016-03-13T12:52:32.123Z",
          "cycles": 169,
          "deleted_at": "2016-03-13T12:52:32.123Z",
          "description": "description3",
          "subscription": {},
          "subscription_item": {
            "id": "id3",
            "description": "description7",
            "status": "status5",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "pricing_scheme": {
              "price": 77,
              "scheme_type": "scheme_type5",
              "price_brackets": [
                {
                  "start_quantity": 132,
                  "price": 70,
                  "end_quantity": 140,
                  "overage_price": 154
                },
                {
                  "start_quantity": 133,
                  "price": 71,
                  "end_quantity": 141,
                  "overage_price": 155
                },
                {
                  "start_quantity": 134,
                  "price": 72,
                  "end_quantity": 142,
                  "overage_price": 156
                }
              ],
              "minimum_price": 19,
              "percentage": 210.93
            },
            "discounts": [
              {
                "id": "id4",
                "value": 51.06,
                "discount_type": "discount_type2",
                "status": "status6",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 10,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description4",
                "subscription": {},
                "subscription_item": {}
              },
              {
                "id": "id5",
                "value": 51.07,
                "discount_type": "discount_type3",
                "status": "status7",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 11,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description5",
                "subscription": {},
                "subscription_item": {}
              },
              {
                "id": "id6",
                "value": 51.08,
                "discount_type": "discount_type4",
                "status": "status8",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 12,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description6",
                "subscription": {},
                "subscription_item": {}
              }
            ],
            "increments": [
              {}
            ],
            "subscription": {},
            "name": "name3",
            "quantity": 55,
            "cycles": 181,
            "deleted_at": "2016-03-13T12:52:32.123Z"
          }
        }
      ],
      "split": {
        "enabled": false,
        "rules": [
          {
            "type": "type4",
            "amount": 96,
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
                "volume_percentage": 42,
                "delay": 248,
                "days": [
                  224
                ]
              },
              "transfer_settings": {
                "transfer_enabled": false,
                "transfer_interval": "transfer_interval6",
                "transfer_day": 208
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
    "status": "status0",
    "duration": 66,
    "created_at": "created_at6",
    "updated_at": "updated_at6",
    "cycle": 6
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
    "fb_id": 174,
    "code": "code8",
    "document_type": "document_type8"
  },
  "next_billing_at": "2016-03-13T12:52:32.123Z",
  "billing_day": 34,
  "minimum_price": 176
}
```

