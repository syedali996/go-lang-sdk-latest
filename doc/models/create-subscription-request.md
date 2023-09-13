
# Create Subscription Request

Request for creating a subcription

## Structure

`CreateSubscriptionRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Customer` | [`models.CreateCustomerRequest`](../../doc/models/create-customer-request.md) | Required | Customer |
| `Card` | [`models.CreateCardRequest`](../../doc/models/create-card-request.md) | Required | Card |
| `Code` | `string` | Required | Subscription code |
| `PaymentMethod` | `string` | Required | Payment method |
| `BillingType` | `string` | Required | Billing type |
| `StatementDescriptor` | `string` | Required | Statement descriptor for credit card subscriptions |
| `Description` | `string` | Required | Subscription description |
| `Currency` | `string` | Required | Currency |
| `Interval` | `string` | Required | Interval |
| `IntervalCount` | `int` | Required | Interval count |
| `PricingScheme` | [`models.CreatePricingSchemeRequest`](../../doc/models/create-pricing-scheme-request.md) | Required | Subscription pricing scheme |
| `Items` | [`[]models.CreateSubscriptionItemRequest`](../../doc/models/create-subscription-item-request.md) | Required | Subscription items |
| `Shipping` | [`models.CreateShippingRequest`](../../doc/models/create-shipping-request.md) | Required | Shipping |
| `Discounts` | [`[]models.CreateDiscountRequest`](../../doc/models/create-discount-request.md) | Required | Discounts |
| `Metadata` | `map[string]string` | Required | Metadata |
| `Setup` | [`*models.CreateSetupRequest`](../../doc/models/create-setup-request.md) | Optional | Setup data |
| `PlanId` | `*string` | Optional | Plan id |
| `CustomerId` | `*string` | Optional | Customer id |
| `CardId` | `*string` | Optional | Card id |
| `BillingDay` | `*int` | Optional | Billing day |
| `Installments` | `*int` | Optional | Number of installments |
| `StartAt` | `*time.Time` | Optional | Subscription start date |
| `MinimumPrice` | `*int` | Optional | Subscription minimum price |
| `Cycles` | `*int` | Optional | Number of cycles |
| `CardToken` | `*string` | Optional | Card token |
| `GatewayAffiliationId` | `*string` | Optional | Gateway Affiliation code |
| `Quantity` | `*int` | Optional | Quantity |
| `BoletoDueDays` | `*int` | Optional | Days until boleto expires |
| `Increments` | [`[]models.CreateIncrementRequest`](../../doc/models/create-increment-request.md) | Required | Increments |
| `Period` | [`*models.CreatePeriodRequest`](../../doc/models/create-period-request.md) | Optional | - |
| `Submerchant` | [`*models.CreateSubMerchantRequest`](../../doc/models/create-sub-merchant-request.md) | Optional | SubMerchant |
| `Split` | [`*models.CreateSubscriptionSplitRequest`](../../doc/models/create-subscription-split-request.md) | Optional | Subscription's split |
| `Boleto` | [`*models.CreateSubscriptionBoletoRequest`](../../doc/models/create-subscription-boleto-request.md) | Optional | Information about fines and interest on the "boleto" used from payment |

## Example (as JSON)

```json
{
  "customer": {
    "name": "{\n    \"name\": \"Tony Stark\"\n}",
    "email": "email6",
    "document": "document6",
    "type": "type0",
    "address": {
      "street": "street6",
      "number": "number4",
      "zip_code": "zip_code0",
      "neighborhood": "neighborhood2",
      "city": "city6",
      "state": "state2",
      "country": "country0",
      "complement": "complement2",
      "metadata": {
        "key0": "metadata3"
      },
      "line_1": "line_10",
      "line_2": "line_24"
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
    "code": "code8",
    "gender": "gender6",
    "document_type": "document_type8"
  },
  "card": {
    "number": "number6",
    "holder_name": "holder_name2",
    "exp_month": 228,
    "exp_year": 68,
    "cvv": "cvv4",
    "billing_address": {
      "street": "street8",
      "number": "number4",
      "zip_code": "zip_code2",
      "neighborhood": "neighborhood4",
      "city": "city2",
      "state": "state6",
      "country": "country2",
      "complement": "complement6",
      "metadata": {
        "key0": "metadata5",
        "key1": "metadata6"
      },
      "line_1": "line_18",
      "line_2": "line_26"
    },
    "brand": "brand0",
    "billing_address_id": "billing_address_id2",
    "metadata": {
      "key0": "metadata7"
    },
    "type": "credit",
    "options": {
      "verify_card": false
    },
    "private_label": false,
    "label": "label6",
    "holder_document": "holder_document0",
    "id": "id6",
    "token": "token0"
  },
  "code": "code8",
  "payment_method": "payment_method0",
  "billing_type": "billing_type6",
  "statement_descriptor": "statement_descriptor0",
  "description": "description0",
  "currency": "currency0",
  "interval": "interval2",
  "interval_count": 82,
  "pricing_scheme": {
    "scheme_type": "scheme_type8",
    "price_brackets": [
      {
        "start_quantity": 119,
        "price": 57,
        "end_quantity": 127,
        "overage_price": 141
      },
      {
        "start_quantity": 120,
        "price": 58,
        "end_quantity": 128,
        "overage_price": 142
      },
      {
        "start_quantity": 121,
        "price": 59,
        "end_quantity": 129,
        "overage_price": 143
      }
    ],
    "price": 166,
    "minimum_price": 6,
    "percentage": 251.76
  },
  "items": [
    {
      "description": "description7",
      "pricing_scheme": {
        "scheme_type": "scheme_type1",
        "price_brackets": [
          {
            "start_quantity": 60,
            "price": 2,
            "end_quantity": 68,
            "overage_price": 82
          },
          {
            "start_quantity": 61,
            "price": 1,
            "end_quantity": 69,
            "overage_price": 83
          },
          {
            "start_quantity": 62,
            "price": 0,
            "end_quantity": 70,
            "overage_price": 84
          }
        ],
        "price": 149,
        "minimum_price": 53,
        "percentage": 25.89
      },
      "id": "id7",
      "plan_item_id": "plan_item_id7",
      "discounts": [
        {
          "value": 236.1,
          "discount_type": "discount_type6",
          "item_id": "item_id8",
          "cycles": 82,
          "description": "description8"
        }
      ],
      "name": "name7",
      "cycles": 109,
      "quantity": 127,
      "minimum_price": 117
    }
  ],
  "shipping": {
    "amount": 52,
    "description": "description6",
    "recipient_name": "recipient_name2",
    "recipient_phone": "recipient_phone6",
    "address_id": "address_id6",
    "address": {
      "street": "street0",
      "number": "number8",
      "zip_code": "zip_code4",
      "neighborhood": "neighborhood6",
      "city": "city0",
      "state": "state6",
      "country": "country4",
      "complement": "complement6",
      "metadata": {
        "key0": "metadata7"
      },
      "line_1": "line_14",
      "line_2": "line_28"
    },
    "max_delivery_date": "2016-03-13T12:52:32.123Z",
    "estimated_delivery_date": "2016-03-13T12:52:32.123Z",
    "type": "type6"
  },
  "discounts": [
    {
      "value": 10.23,
      "discount_type": "discount_type9",
      "item_id": "item_id1",
      "cycles": 233,
      "description": "description1"
    }
  ],
  "metadata": {
    "key0": "metadata3",
    "key1": "metadata4",
    "key2": "metadata5"
  },
  "increments": [
    {
      "value": 124.19,
      "increment_type": "increment_type9",
      "item_id": "item_id3",
      "cycles": 101,
      "description": "description3"
    }
  ],
  "setup": {
    "amount": 110,
    "description": "description4",
    "payment": {
      "payment_method": "payment_method8",
      "credit_card": {
        "installments": 228,
        "statement_descriptor": "statement_descriptor0",
        "card": {
          "number": "number2",
          "holder_name": "holder_name0",
          "exp_month": 134,
          "exp_year": 162,
          "cvv": "cvv2",
          "billing_address": {
            "street": "street6",
            "number": "number6",
            "zip_code": "zip_code0",
            "neighborhood": "neighborhood2",
            "city": "city6",
            "state": "state2",
            "country": "country0",
            "complement": "complement8",
            "metadata": {
              "key0": "metadata7",
              "key1": "metadata8"
            },
            "line_1": "line_10",
            "line_2": "line_24"
          },
          "brand": "brand8",
          "billing_address_id": "billing_address_id0",
          "metadata": {
            "key0": "metadata1"
          },
          "type": "type6",
          "options": {
            "verify_card": false
          },
          "holder_document": "holder_document8",
          "private_label": false,
          "label": "label4",
          "id": "id4",
          "token": "token2"
        },
        "card_id": "card_id6",
        "card_token": "card_token0"
      },
      "debit_card": {
        "statement_descriptor": "statement_descriptor6",
        "card": {
          "number": "number2",
          "holder_name": "holder_name6",
          "exp_month": 60,
          "exp_year": 236,
          "cvv": "cvv8",
          "billing_address": {
            "street": "street2",
            "number": "number0",
            "zip_code": "zip_code6",
            "neighborhood": "neighborhood8",
            "city": "city8",
            "state": "state2",
            "country": "country6",
            "complement": "complement2",
            "metadata": {
              "key0": "metadata1"
            },
            "line_1": "line_14",
            "line_2": "line_20"
          },
          "brand": "brand4",
          "billing_address_id": "billing_address_id6",
          "metadata": {
            "key0": "metadata3",
            "key1": "metadata4",
            "key2": "metadata5"
          },
          "type": "type0",
          "options": {
            "verify_card": false
          },
          "holder_document": "holder_document6",
          "private_label": false,
          "label": "label0",
          "id": "id0",
          "token": "token6"
        },
        "card_id": "card_id2",
        "card_token": "card_token4",
        "recurrence": false
      },
      "boleto": {
        "retries": 154,
        "bank": "bank0",
        "instructions": "instructions0",
        "due_at": "2016-03-13T12:52:32.123Z",
        "billing_address": {
          "street": "street4",
          "number": "number8",
          "zip_code": "zip_code8",
          "neighborhood": "neighborhood0",
          "city": "city6",
          "state": "state0",
          "country": "country8",
          "complement": "complement0",
          "metadata": {
            "key0": "metadata9",
            "key1": "metadata0",
            "key2": "metadata1"
          },
          "line_1": "line_12",
          "line_2": "line_22"
        },
        "billing_address_id": "billing_address_id8",
        "nosso_numero": "nosso_numero2",
        "document_number": "document_number4",
        "statement_descriptor": "statement_descriptor2",
        "interest": {
          "days": 232,
          "type": "type2",
          "amount": 50
        },
        "fine": {
          "days": 58,
          "type": "type6",
          "amount": 132
        },
        "max_days_to_pay_past_due": 46
      },
      "currency": "currency2",
      "voucher": {
        "statement_descriptor": "statement_descriptor0",
        "card_id": "card_id4",
        "card_token": "card_token0",
        "Card": {
          "number": "number8",
          "holder_name": "holder_name6",
          "exp_month": 22,
          "exp_year": 62,
          "cvv": "cvv8",
          "billing_address": {
            "street": "street2",
            "number": "number0",
            "zip_code": "zip_code6",
            "neighborhood": "neighborhood8",
            "city": "city2",
            "state": "state8",
            "country": "country6",
            "complement": "complement8",
            "metadata": {
              "key0": "metadata9",
              "key1": "metadata8"
            },
            "line_1": "line_14",
            "line_2": "line_20"
          },
          "brand": "brand4",
          "billing_address_id": "billing_address_id6",
          "metadata": {
            "key0": "metadata7"
          },
          "type": "type0",
          "options": {
            "verify_card": false
          },
          "holder_document": "holder_document4",
          "private_label": false,
          "label": "label0",
          "id": "id0",
          "token": "token4"
        },
        "recurrency_cycle": "recurrency_cycle4"
      }
    }
  },
  "plan_id": "plan_id8",
  "customer_id": "customer_id8",
  "card_id": "card_id4",
  "billing_day": 34
}
```

