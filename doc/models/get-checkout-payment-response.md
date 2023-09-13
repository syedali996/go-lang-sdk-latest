
# Get Checkout Payment Response

Resposta das configurações de pagamento do checkout

## Structure

`GetCheckoutPaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*string` | Required | - |
| `Amount` | `Optional[int]` | Optional | Valor em centavos |
| `DefaultPaymentMethod` | `*string` | Required | Meio de pagamento padrão no checkout |
| `SuccessUrl` | `*string` | Required | Url de redirecionamento de sucesso após o checkou |
| `PaymentUrl` | `*string` | Required | Url para pagamento usando o checkout |
| `GatewayAffiliationId` | `*string` | Required | Código da afiliação onde o pagamento será processado no gateway |
| `AcceptedPaymentMethods` | `[]string` | Required | Meios de pagamento aceitos no checkout |
| `Status` | `*string` | Required | Status do checkout |
| `SkipCheckoutSuccessPage` | `*bool` | Required | Pular tela de sucesso pós-pagamento? |
| `CreatedAt` | `*time.Time` | Required | Data de criação |
| `UpdatedAt` | `*time.Time` | Required | Data de atualização |
| `CanceledAt` | `Optional[time.Time]` | Optional | Data de cancelamento |
| `CustomerEditable` | `*bool` | Required | Torna o objeto customer editável |
| `Customer` | [`Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | Dados do comprador |
| `Billingaddress` | [`*models.GetAddressResponse`](../../doc/models/get-address-response.md) | Required | Dados do endereço de cobrança |
| `CreditCard` | [`*models.GetCheckoutCreditCardPaymentResponse`](../../doc/models/get-checkout-credit-card-payment-response.md) | Required | Configurações de cartão de crédito |
| `Boleto` | [`*models.GetCheckoutBoletoPaymentResponse`](../../doc/models/get-checkout-boleto-payment-response.md) | Required | Configurações de boleto |
| `BillingAddressEditable` | `*bool` | Required | Indica se o billing address poderá ser editado |
| `Shipping` | [`*models.GetShippingResponse`](../../doc/models/get-shipping-response.md) | Required | Configurações  de entrega |
| `Shippable` | `*bool` | Required | Indica se possui entrega |
| `ClosedAt` | `Optional[time.Time]` | Optional | Data de fechamento |
| `ExpiresAt` | `Optional[time.Time]` | Optional | Data de expiração |
| `Currency` | `*string` | Required | Moeda |
| `DebitCard` | [`Optional[models.GetCheckoutDebitCardPaymentResponse]`](../../doc/models/get-checkout-debit-card-payment-response.md) | Optional | Configurações de cartão de débito |
| `BankTransfer` | [`Optional[models.GetCheckoutBankTransferPaymentResponse]`](../../doc/models/get-checkout-bank-transfer-payment-response.md) | Optional | Bank transfer payment response |
| `AcceptedBrands` | `[]string` | Required | Accepted Brands |
| `Pix` | [`Optional[models.GetCheckoutPixPaymentResponse]`](../../doc/models/get-checkout-pix-payment-response.md) | Optional | Pix payment response |

## Example (as JSON)

```json
{
  "id": "id0",
  "default_payment_method": "default_payment_method0",
  "success_url": "success_url2",
  "payment_url": "payment_url6",
  "gateway_affiliation_id": "gateway_affiliation_id4",
  "accepted_payment_methods": [
    "accepted_payment_methods3",
    "accepted_payment_methods4",
    "accepted_payment_methods5"
  ],
  "status": "status8",
  "skip_checkout_success_page": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z",
  "customer_editable": false,
  "billingaddress": {
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
      "address": {},
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
      "fb_id": 68,
      "code": "code6",
      "document_type": "document_type6"
    },
    "metadata": {
      "key0": "metadata5"
    },
    "line_1": "line_18",
    "line_2": "line_26",
    "deleted_at": "2016-03-13T12:52:32.123Z"
  },
  "credit_card": {
    "statementDescriptor": "statementDescriptor4",
    "installments": [
      {
        "number": "number1",
        "total": 167
      }
    ],
    "authentication": {
      "type": "type0",
      "threed_secure": {
        "mpi": "mpi0",
        "eci": "eci2",
        "cavv": "cavv8",
        "transaction_Id": "transaction_Id2",
        "success_url": "success_url4"
      }
    }
  },
  "boleto": {
    "due_at": "2016-03-13T12:52:32.123Z",
    "instructions": "instructions2"
  },
  "billing_address_editable": false,
  "shipping": {
    "amount": 52,
    "description": "description6",
    "recipient_name": "recipient_name2",
    "recipient_phone": "recipient_phone6",
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
      "customer": {
        "id": "id0",
        "name": "name0",
        "email": "email6",
        "delinquent": false,
        "created_at": "2016-03-13T12:52:32.123Z",
        "updated_at": "2016-03-13T12:52:32.123Z",
        "document": "document4",
        "type": "type0",
        "fb_access_token": "fb_access_token4",
        "address": {},
        "metadata": {
          "key0": "metadata7",
          "key1": "metadata6",
          "key2": "metadata5"
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
        "fb_id": 100,
        "code": "code8",
        "document_type": "document_type8"
      },
      "metadata": {
        "key0": "metadata7"
      },
      "line_1": "line_14",
      "line_2": "line_28",
      "deleted_at": "2016-03-13T12:52:32.123Z"
    },
    "max_delivery_date": "2016-03-13T12:52:32.123Z",
    "estimated_delivery_date": "2016-03-13T12:52:32.123Z",
    "type": "type6"
  },
  "shippable": false,
  "currency": "currency0",
  "accepted_brands": [
    "accepted_brands8",
    "accepted_brands9"
  ],
  "amount": 46,
  "canceled_at": "2016-03-13T12:52:32.123Z",
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
  "closed_at": "2016-03-13T12:52:32.123Z",
  "expires_at": "2016-03-13T12:52:32.123Z"
}
```

