
# Client Class Documentation

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `httpConfiguration` | [`HttpConfiguration`](http-configuration.md) | Configurable http client options like timeout and retries. |
| `basicAuthUserName` | `string` | The username to use with basic authentication |
| `basicAuthPassword` | `string` | The password to use with basic authentication |

The API client can be initialized as follows:

```go
config := pagarmeapisdk.CreateConfiguration(
    pagarmeapisdk.WithHttpConfiguration(
        pagarmeapisdk.CreateHttpConfiguration(
            pagarmeapisdk.WithTimeout(0),
            pagarmeapisdk.WithTransport(http.DefaultTransport),
            pagarmeapisdk.WithRetryConfiguration(
                pagarmeapisdk.CreateRetryConfiguration(
                    pagarmeapisdk.WithMaxRetryAttempts(0),
                    pagarmeapisdk.WithRetryOnTimeout(true),
                    pagarmeapisdk.WithRetryInterval(1),
                    pagarmeapisdk.WithMaximumRetryWaitTime(0),
                    pagarmeapisdk.WithBackoffFactor(2),
                    pagarmeapisdk.WithHttpStatusCodesToRetry([]int64{408, 413, 429, 500, 502, 503, 504, 521, 522, 524}),
                    pagarmeapisdk.WithHttpMethodsToRetry([]string{"GET", "PUT"}),
                ),
            ),
        ),
    ),
    pagarmeapisdk.WithBasicAuthUserName("BasicAuthUserName"),
    pagarmeapisdk.WithBasicAuthPassword("BasicAuthPassword"),
)
client := pagarmeapisdk.NewClient(config)
```

## PagarmeApiSDK Client

The gateway for the SDK. This class acts as a factory for the Controllers and also holds the configuration of the SDK.

## Controllers

| Name | Description |
|  --- | --- |
| subscriptions | Gets SubscriptionsController |
| orders | Gets OrdersController |
| plans | Gets PlansController |
| invoices | Gets InvoicesController |
| customers | Gets CustomersController |
| charges | Gets ChargesController |
| recipients | Gets RecipientsController |
| tokens | Gets TokensController |
| transactions | Gets TransactionsController |
| transfers | Gets TransfersController |

