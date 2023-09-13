
# List Invoices Response

Response object for listing invoices

## Structure

`ListInvoicesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md) | Required | The Invoice objects |
| `Paging` | [`*models.PagingResponse`](../../doc/models/paging-response.md) | Required | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id5",
      "code": "code3",
      "url": "url9",
      "amount": 121,
      "status": "status7",
      "payment_method": "payment_method5",
      "created_at": "2016-03-13T12:52:32.123Z",
      "items": [
        {
          "amount": 180,
          "description": "description2",
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
          "price_bracket": {
            "start_quantity": 154,
            "price": 92,
            "end_quantity": 162,
            "overage_price": 176
          },
          "subscription_item_id": "subscription_item_id6",
          "quantity": 38,
          "name": "name2"
        }
      ],
      "charge": {
        "id": "id7",
        "code": "code5",
        "gateway_id": "gateway_id3",
        "amount": 251,
        "status": "status9",
        "currency": "currency7",
        "payment_method": "payment_method3",
        "due_at": "2016-03-13T12:52:32.123Z",
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "last_transaction": {
          "gateway_id": "gateway_id9",
          "amount": 53,
          "status": "status1",
          "success": true,
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "attempt_count": 29,
          "max_attempts": 17,
          "splits": [
            {
              "type": "type1",
              "amount": 241,
              "recipient": {
                "id": "id3",
                "name": "name3",
                "email": "email3",
                "document": "document7",
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
                    "key0": "metadata8",
                    "key1": "metadata7",
                    "key2": "metadata6"
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
                  "key1": "metadata9"
                },
                "automatic_anticipation_settings": {
                  "enabled": true,
                  "type": "type3",
                  "volume_percentage": 153,
                  "delay": 137,
                  "days": [
                    113,
                    114
                  ]
                },
                "transfer_settings": {
                  "transfer_enabled": true,
                  "transfer_interval": "transfer_interval1",
                  "transfer_day": 63
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
          "id": "id9",
          "gateway_response": {
            "code": "code9",
            "errors": [
              {
                "message": "message6"
              }
            ]
          },
          "antifraud_response": {
            "status": "status9",
            "return_code": "return_code7",
            "return_message": "return_message5",
            "provider_name": "provider_name5",
            "score": "score7"
          },
          "split": [
            {
              "type": "type7",
              "amount": 255,
              "recipient": {
                "id": "id9",
                "name": "name9",
                "email": "email3",
                "document": "document3",
                "description": "description9",
                "type": "type9",
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
                    "key0": "metadata8",
                    "key1": "metadata9",
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
                  "key0": "metadata0",
                  "key1": "metadata9",
                  "key2": "metadata8"
                },
                "automatic_anticipation_settings": {
                  "enabled": true,
                  "type": "type3",
                  "volume_percentage": 117,
                  "delay": 151,
                  "days": [
                    127,
                    128
                  ]
                },
                "transfer_settings": {
                  "transfer_enabled": true,
                  "transfer_interval": "transfer_interval5",
                  "transfer_day": 243
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
          ],
          "next_attempt": "2016-03-13T12:52:32.123Z",
          "transaction_type": "transaction",
          "metadata": {
            "key0": "metadata0",
            "key1": "metadata9",
            "key2": "metadata8"
          },
          "interest": {
            "days": 215,
            "type": "type1",
            "amount": 33
          },
          "fine": {
            "days": 233,
            "type": "type1",
            "amount": 51
          }
        },
        "metadata": {
          "key0": "metadata4"
        },
        "canceled_amount": 105,
        "paid_amount": 87,
        "recurrency_cycle": "\"first\" or \"subsequent\"",
        "invoice": {},
        "order": {
          "id": "id9",
          "code": "code7",
          "currency": "currency1",
          "items": [
            {
              "id": "id6",
              "amount": 4,
              "description": "description6",
              "quantity": 118,
              "category": "category4",
              "code": "code4"
            },
            {
              "id": "id7",
              "amount": 5,
              "description": "description7",
              "quantity": 119,
              "category": "category5",
              "code": "code5"
            }
          ],
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
            "fb_id": 183,
            "code": "code9",
            "document_type": "document_type9"
          },
          "status": "status9",
          "created_at": "2016-03-13T12:52:32.123Z",
          "updated_at": "2016-03-13T12:52:32.123Z",
          "charges": [
            {},
            {}
          ],
          "invoice_url": "invoice_url1",
          "shipping": {
            "amount": 43,
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
                "email": "email7",
                "delinquent": true,
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "document": "document3",
                "type": "type9",
                "fb_access_token": "fb_access_token3",
                "address": {},
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
                    "country_code": "country_code1",
                    "number": "number1",
                    "area_code": "area_code1"
                  }
                },
                "fb_id": 91,
                "code": "code7",
                "document_type": "document_type7"
              },
              "metadata": {
                "key0": "metadata6"
              },
              "line_1": "line_13",
              "line_2": "line_27",
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            "max_delivery_date": "2016-03-13T12:52:32.123Z",
            "estimated_delivery_date": "2016-03-13T12:52:32.123Z",
            "type": "type7"
          },
          "metadata": {
            "key0": "metadata4",
            "key1": "metadata5",
            "key2": "metadata6"
          },
          "checkouts": [
            {
              "id": "id6",
              "amount": 210,
              "default_payment_method": "default_payment_method6",
              "success_url": "success_url8",
              "payment_url": "payment_url0",
              "gateway_affiliation_id": "gateway_affiliation_id2",
              "accepted_payment_methods": [
                "accepted_payment_methods9",
                "accepted_payment_methods0",
                "accepted_payment_methods1"
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
                    "key1": "metadata6",
                    "key2": "metadata5"
                  },
                  "line_1": "line_16",
                  "line_2": "line_20",
                  "deleted_at": "2016-03-13T12:52:32.123Z"
                },
                "metadata": {
                  "key0": "metadata3",
                  "key1": "metadata2"
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
                "fb_id": 242,
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
                    "key1": "metadata4"
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
                  "fb_id": 68,
                  "code": "code2",
                  "document_type": "document_type2"
                },
                "metadata": {
                  "key0": "metadata5"
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
                    "total": 167
                  }
                ],
                "authentication": {
                  "type": "type4",
                  "threed_secure": {
                    "mpi": "mpi8",
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
                "amount": 52,
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
                      "key1": "metadata8"
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
                    "fb_id": 100,
                    "code": "code4",
                    "document_type": "document_type4"
                  },
                  "metadata": {
                    "key0": "metadata7"
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
                "accepted_brands2",
                "accepted_brands3",
                "accepted_brands4"
              ]
            }
          ],
          "ip": "ip3",
          "session_id": "session_id9",
          "location": {
            "latitude": "latitude3",
            "longitude": "longitude7"
          },
          "closed": true
        },
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
              "key1": "metadata5"
            },
            "line_1": "line_17",
            "line_2": "line_21",
            "deleted_at": "2016-03-13T12:52:32.123Z"
          },
          "metadata": {
            "key0": "metadata4",
            "key1": "metadata3",
            "key2": "metadata2"
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
          "fb_id": 27,
          "code": "code5",
          "document_type": "document_type5"
        },
        "paid_at": "2016-03-13T12:52:32.123Z",
        "canceled_at": "2016-03-13T12:52:32.123Z"
      },
      "installments": 161,
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
              }
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
                {
                  "id": "id7",
                  "value": 93.89,
                  "discount_type": "discount_type5",
                  "status": "status9",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 197,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description7",
                  "subscription": {},
                  "subscription_item": {}
                },
                {
                  "id": "id8",
                  "value": 93.9,
                  "discount_type": "discount_type6",
                  "status": "status0",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 198,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description8",
                  "subscription": {},
                  "subscription_item": {}
                }
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
      "subscription_id": "subscription_id5",
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
          "interval_count": 149,
          "billing_type": "billing_type5",
          "current_cycle": {},
          "payment_method": "payment_method1",
          "currency": "currency1",
          "installments": 61,
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
                "key0": "metadata8"
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
                "country_code": "country_code3",
                "number": "number9",
                "area_code": "area_code3"
              }
            },
            "fb_id": 53,
            "code": "code9",
            "document_type": "document_type9"
          },
          "card": {
            "id": "id5",
            "last_four_digits": "last_four_digits1",
            "brand": "brand9",
            "holder_name": "holder_name1",
            "exp_month": 45,
            "exp_year": 251,
            "status": "status7",
            "created_at": "2016-03-13T12:52:32.123Z",
            "updated_at": "2016-03-13T12:52:32.123Z",
            "billing_address": {
              "street": "street7",
              "number": "number5",
              "zip_code": "zip_code1",
              "neighborhood": "neighborhood3",
              "city": "city7",
              "state": "state3",
              "country": "country1",
              "complement": "complement3",
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
                  "key0": "metadata8",
                  "key1": "metadata7"
                },
                "line_1": "line_15",
                "line_2": "line_29",
                "deleted_at": "2016-03-13T12:52:32.123Z"
              },
              "metadata": {
                "key0": "metadata2",
                "key1": "metadata1",
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
                  "number": "number5",
                  "area_code": "area_code7"
                }
              },
              "fb_id": 239,
              "code": "code3",
              "document_type": "document_type3"
            },
            "metadata": {
              "key0": "metadata2"
            },
            "type": "type5",
            "holder_document": "holder_document9",
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
                "price": 72,
                "scheme_type": "scheme_type0",
                "price_brackets": [
                  {
                    "start_quantity": 137,
                    "price": 75,
                    "end_quantity": 145,
                    "overage_price": 159
                  }
                ],
                "minimum_price": 24,
                "percentage": 185.38
              },
              "discounts": [
                {
                  "id": "id9",
                  "value": 183.01,
                  "discount_type": "discount_type7",
                  "status": "status1",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 149,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description9",
                  "subscription": {},
                  "subscription_item": {}
                }
              ],
              "increments": [
                {
                  "id": "id7",
                  "value": 38.29,
                  "increment_type": "increment_type9",
                  "status": "status9",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 13,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description7",
                  "subscription": {},
                  "subscription_item": {}
                },
                {
                  "id": "id8",
                  "value": 38.3,
                  "increment_type": "increment_type0",
                  "status": "status0",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 14,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description8",
                  "subscription": {},
                  "subscription_item": {}
                }
              ],
              "subscription": {},
              "name": "name8",
              "quantity": 194,
              "cycles": 214,
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            {
              "id": "id9",
              "description": "description9",
              "status": "status1",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "pricing_scheme": {
                "price": 71,
                "scheme_type": "scheme_type1",
                "price_brackets": [
                  {
                    "start_quantity": 138,
                    "price": 76,
                    "end_quantity": 146,
                    "overage_price": 160
                  },
                  {
                    "start_quantity": 139,
                    "price": 77,
                    "end_quantity": 147,
                    "overage_price": 161
                  }
                ],
                "minimum_price": 25,
                "percentage": 185.39
              },
              "discounts": [
                {
                  "id": "id0",
                  "value": 183.02,
                  "discount_type": "discount_type8",
                  "status": "status2",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 150,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description0",
                  "subscription": {},
                  "subscription_item": {}
                },
                {
                  "id": "id1",
                  "value": 183.03,
                  "discount_type": "discount_type9",
                  "status": "status3",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 151,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description1",
                  "subscription": {},
                  "subscription_item": {}
                }
              ],
              "increments": [
                {
                  "id": "id8",
                  "value": 38.3,
                  "increment_type": "increment_type0",
                  "status": "status0",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 14,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description8",
                  "subscription": {},
                  "subscription_item": {}
                },
                {
                  "id": "id9",
                  "value": 38.31,
                  "increment_type": "increment_type1",
                  "status": "status1",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 15,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description9",
                  "subscription": {},
                  "subscription_item": {}
                },
                {
                  "id": "id0",
                  "value": 38.32,
                  "increment_type": "increment_type2",
                  "status": "status2",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 16,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description0",
                  "subscription": {},
                  "subscription_item": {}
                }
              ],
              "subscription": {},
              "name": "name9",
              "quantity": 195,
              "cycles": 215,
              "deleted_at": "2016-03-13T12:52:32.123Z"
            },
            {
              "id": "id0",
              "description": "description0",
              "status": "status2",
              "created_at": "2016-03-13T12:52:32.123Z",
              "updated_at": "2016-03-13T12:52:32.123Z",
              "pricing_scheme": {
                "price": 70,
                "scheme_type": "scheme_type2",
                "price_brackets": [
                  {
                    "start_quantity": 139,
                    "price": 77,
                    "end_quantity": 147,
                    "overage_price": 161
                  },
                  {
                    "start_quantity": 140,
                    "price": 78,
                    "end_quantity": 148,
                    "overage_price": 162
                  },
                  {
                    "start_quantity": 141,
                    "price": 79,
                    "end_quantity": 149,
                    "overage_price": 163
                  }
                ],
                "minimum_price": 26,
                "percentage": 185.4
              },
              "discounts": [
                {
                  "id": "id1",
                  "value": 183.03,
                  "discount_type": "discount_type9",
                  "status": "status3",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 151,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description1",
                  "subscription": {},
                  "subscription_item": {}
                },
                {
                  "id": "id2",
                  "value": 183.04,
                  "discount_type": "discount_type0",
                  "status": "status4",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 152,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description2",
                  "subscription": {},
                  "subscription_item": {}
                },
                {
                  "id": "id3",
                  "value": 183.05,
                  "discount_type": "discount_type1",
                  "status": "status5",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 153,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description3",
                  "subscription": {},
                  "subscription_item": {}
                }
              ],
              "increments": [
                {
                  "id": "id9",
                  "value": 38.31,
                  "increment_type": "increment_type1",
                  "status": "status1",
                  "created_at": "2016-03-13T12:52:32.123Z",
                  "cycles": 15,
                  "deleted_at": "2016-03-13T12:52:32.123Z",
                  "description": "description9",
                  "subscription": {},
                  "subscription_item": {}
                }
              ],
              "subscription": {},
              "name": "name0",
              "quantity": 196,
              "cycles": 216,
              "deleted_at": "2016-03-13T12:52:32.123Z"
            }
          ],
          "statement_descriptor": "statement_descriptor1",
          "metadata": {
            "key0": "metadata8",
            "key1": "metadata7",
            "key2": "metadata6"
          },
          "setup": {
            "id": "id5",
            "description": "description5",
            "amount": 85,
            "status": "status7"
          },
          "gateway_affiliation_id": "gateway_affiliation_id7",
          "next_billing_at": "2016-03-13T12:52:32.123Z",
          "billing_day": 33,
          "minimum_price": 147,
          "increments": [
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
              "subscription_item": {
                "id": "id6",
                "description": "description6",
                "status": "status8",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "pricing_scheme": {
                  "price": 200,
                  "scheme_type": "scheme_type2",
                  "price_brackets": [
                    {
                      "start_quantity": 9,
                      "price": 53,
                      "end_quantity": 17,
                      "overage_price": 31
                    },
                    {
                      "start_quantity": 10,
                      "price": 52,
                      "end_quantity": 18,
                      "overage_price": 32
                    },
                    {
                      "start_quantity": 11,
                      "price": 51,
                      "end_quantity": 19,
                      "overage_price": 33
                    }
                  ],
                  "minimum_price": 104,
                  "percentage": 158.5
                },
                "discounts": [
                  {
                    "id": "id7",
                    "value": 103.49,
                    "discount_type": "discount_type5",
                    "status": "status9",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 133,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description7",
                    "subscription": {},
                    "subscription_item": {}
                  },
                  {
                    "id": "id8",
                    "value": 103.5,
                    "discount_type": "discount_type6",
                    "status": "status0",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 134,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description8",
                    "subscription": {},
                    "subscription_item": {}
                  },
                  {
                    "id": "id9",
                    "value": 103.51,
                    "discount_type": "discount_type7",
                    "status": "status1",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 135,
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
                "quantity": 178,
                "cycles": 58,
                "deleted_at": "2016-03-13T12:52:32.123Z"
              }
            },
            {
              "id": "id1",
              "value": 68.43,
              "increment_type": "increment_type3",
              "status": "status3",
              "created_at": "2016-03-13T12:52:32.123Z",
              "cycles": 45,
              "deleted_at": "2016-03-13T12:52:32.123Z",
              "description": "description1",
              "subscription": {},
              "subscription_item": {
                "id": "id7",
                "description": "description7",
                "status": "status9",
                "created_at": "2016-03-13T12:52:32.123Z",
                "updated_at": "2016-03-13T12:52:32.123Z",
                "pricing_scheme": {
                  "price": 201,
                  "scheme_type": "scheme_type1",
                  "price_brackets": [
                    {
                      "start_quantity": 8,
                      "price": 54,
                      "end_quantity": 16,
                      "overage_price": 30
                    },
                    {
                      "start_quantity": 9,
                      "price": 53,
                      "end_quantity": 17,
                      "overage_price": 31
                    }
                  ],
                  "minimum_price": 105,
                  "percentage": 158.49
                },
                "discounts": [
                  {
                    "id": "id8",
                    "value": 103.5,
                    "discount_type": "discount_type6",
                    "status": "status0",
                    "created_at": "2016-03-13T12:52:32.123Z",
                    "cycles": 134,
                    "deleted_at": "2016-03-13T12:52:32.123Z",
                    "description": "description8",
                    "subscription": {},
                    "subscription_item": {}
                  }
                ],
                "increments": [
                  {},
                  {}
                ],
                "subscription": {},
                "name": "name7",
                "quantity": 179,
                "cycles": 57,
                "deleted_at": "2016-03-13T12:52:32.123Z"
              }
            }
          ],
          "split": {
            "enabled": true,
            "rules": [
              {
                "type": "type1",
                "amount": 221,
                "recipient": {
                  "id": "id3",
                  "name": "name3",
                  "email": "email3",
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
                    "type": "type7",
                    "volume_percentage": 173,
                    "delay": 117,
                    "days": [
                      93,
                      94,
                      95
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": true,
                    "transfer_interval": "transfer_interval1",
                    "transfer_day": 83
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
              },
              {
                "type": "type2",
                "amount": 222,
                "recipient": {
                  "id": "id4",
                  "name": "name4",
                  "email": "email2",
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
                      "key0": "metadata7",
                      "key1": "metadata6"
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
                    "key0": "metadata1",
                    "key1": "metadata0",
                    "key2": "metadata9"
                  },
                  "automatic_anticipation_settings": {
                    "enabled": false,
                    "type": "type8",
                    "volume_percentage": 172,
                    "delay": 118,
                    "days": [
                      94
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": false,
                    "transfer_interval": "transfer_interval0",
                    "transfer_day": 82
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
                "amount": 223,
                "recipient": {
                  "id": "id5",
                  "name": "name5",
                  "email": "email1",
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
                      "key0": "metadata6",
                      "key1": "metadata5",
                      "key2": "metadata4"
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
                    "key0": "metadata2",
                    "key1": "metadata1"
                  },
                  "automatic_anticipation_settings": {
                    "enabled": true,
                    "type": "type9",
                    "volume_percentage": 171,
                    "delay": 119,
                    "days": [
                      95,
                      96
                    ]
                  },
                  "transfer_settings": {
                    "transfer_enabled": true,
                    "transfer_interval": "transfer_interval9",
                    "transfer_day": 81
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
        "status": "status7",
        "duration": 75,
        "created_at": "created_at3",
        "updated_at": "updated_at1",
        "cycle": 15
      },
      "due_at": "2016-03-13T12:52:32.123Z",
      "canceled_at": "2016-03-13T12:52:32.123Z",
      "billing_at": "2016-03-13T12:52:32.123Z"
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

