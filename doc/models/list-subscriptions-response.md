
# List Subscriptions Response

Response object for listing subscriptions

## Structure

`ListSubscriptionsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md) | Required | The subscription objects |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "code": "code3",
      "start_at": "2016-03-13T12:52:32.123Z",
      "interval": "interval3",
      "interval_count": 249,
      "billing_type": "billing_type9",
      "payment_method": "payment_method5",
      "currency": "currency5",
      "installments": 161,
      "status": "status7",
      "created_at": "2016-03-13T12:52:32.123Z",
      "updated_at": "2016-03-13T12:52:32.123Z",
      "card": {
        "id": "id9",
        "last_four_digits": "last_four_digits5",
        "brand": "brand3",
        "holder_name": "holder_name5",
        "exp_month": 55,
        "exp_year": 95,
        "status": "status1",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
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
        "customer": {
          "id": "id9",
          "name": "name9",
          "email": "email7",
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
              "key0": "metadata6"
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
              "country_code": "country_code1",
              "number": "number1",
              "area_code": "area_code1"
            }
          },
          "fb_id": 83,
          "code": "code7",
          "document_type": "document_type7"
        },
        "metadata": {
          "key0": "metadata0",
          "key1": "metadata9",
          "key2": "metadata8"
        },
        "type": "type9",
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
            "price": 28,
            "scheme_type": "scheme_type4",
            "price_brackets": [
              {
                "start_quantity": 237,
                "price": 175,
                "end_quantity": 245,
                "overage_price": 3
              }
            ],
            "minimum_price": 124,
            "percentage": 132.62
          },
          "discounts": [
            {
              "id": "id3",
              "value": 130.25,
              "discount_type": "discount_type1",
              "status": "status5",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 249,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description3",
              "subscription": {},
              "subscription_item": {}
            }
          ],
          "increments": [
            {
              "id": "id1",
              "value": 241.53,
              "increment_type": "increment_type3",
              "status": "status3",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 113,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description1",
              "subscription": {},
              "subscription_item": {}
            }
          ],
          "subscription": {
            "id": null,
            "code": null,
            "start_at": null,
            "interval": null,
            "interval_count": null,
            "billing_type": null,
            "payment_method": null,
            "currency": null,
            "installments": null,
            "status": null,
            "created_at": null,
            "updated_at": null,
            "card": null,
            "items": null,
            "statement_descriptor": null,
            "metadata": null,
            "setup": null,
            "gateway_affiliation_id": null,
            "increments": [
              null
            ],
            "split": {
              "enabled": null,
              "rules": [
                null
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
            }
          },
          "name": "name2",
          "quantity": 38,
          "cycles": 58,
          "deleted_at": "2016-03-13T12:52:32.123Z"
        }
      ],
      "statement_descriptor": "statement_descriptor5",
      "metadata": {
        "key0": "metadata4",
        "key1": "metadata3"
      },
      "setup": {
        "id": "id9",
        "description": "description9",
        "amount": 185,
        "status": "status1"
      },
      "gateway_affiliation_id": "gateway_affiliation_id1",
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
          "subscription": {},
          "subscription_item": {
            "id": "id0",
            "description": "description0",
            "status": "status2",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "pricing_scheme": {
              "price": 244,
              "scheme_type": "scheme_type2",
              "price_brackets": [
                {
                  "start_quantity": 221,
                  "price": 97,
                  "end_quantity": 229,
                  "overage_price": 243
                },
                {
                  "start_quantity": 222,
                  "price": 96,
                  "end_quantity": 230,
                  "overage_price": 244
                }
              ],
              "minimum_price": 148,
              "percentage": 53.1
            },
            "discounts": [
              {
                "id": "id1",
                "value": 50.73,
                "discount_type": "discount_type9",
                "status": "status3",
                "created_at": "2016-03-13T12:52:32.123Z",
                "cycles": 233,
                "deleted_at": "2016-03-13T12:52:32.123Z",
                "description": "description1",
                "subscription": {},
                "subscription_item": {}
              },
              {
                "id": "id2",
                "value": 50.74,
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
              {},
              {},
              {}
            ],
            "subscription": {},
            "name": "name0",
            "quantity": 22,
            "cycles": 42,
            "deleted_at": "2016-03-13T12:52:32.123Z"
          }
        }
      ],
      "split": {
        "enabled": true,
        "rules": [
          {
            "type": "type5",
            "amount": 65,
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
                  "key0": "metadata6",
                  "key1": "metadata7",
                  "key2": "metadata8"
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
                "key0": "metadata2",
                "key1": "metadata1",
                "key2": "metadata0"
              },
              "automatic_anticipation_settings": {
                "enabled": true,
                "type": "type1",
                "volume_percentage": 183,
                "delay": 217,
                "days": [
                  193,
                  194
                ]
              },
              "transfer_settings": {
                "transfer_enabled": true,
                "transfer_interval": "transfer_interval3",
                "transfer_day": 53
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
        "id": "id3",
        "billing_at": "2016-03-13T12:52:32.123Z",
        "subscription": {},
        "status": "status5",
        "duration": 233,
        "created_at": "created_at1",
        "updated_at": "updated_at9",
        "cycle": 173
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
      "next_billing_at": "2016-03-13T12:52:32.123Z",
      "billing_day": 133,
      "minimum_price": 247
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

