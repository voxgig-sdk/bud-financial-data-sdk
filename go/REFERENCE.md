# BudFinancialData Golang SDK Reference

Complete API reference for the BudFinancialData Golang SDK.


## BudFinancialDataSDK

### Constructor

```go
func NewBudFinancialDataSDK(options map[string]any) *BudFinancialDataSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *BudFinancialDataSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *BudFinancialDataSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `CorrectFinancialData(data map[string]any) BudFinancialDataEntity`

Create a new `CorrectFinancialData` entity instance. Pass `nil` for no initial data.

#### `CustomerMerchantCorrection(data map[string]any) BudFinancialDataEntity`

Create a new `CustomerMerchantCorrection` entity instance. Pass `nil` for no initial data.

#### `Label(data map[string]any) BudFinancialDataEntity`

Create a new `Label` entity instance. Pass `nil` for no initial data.

#### `ListLabel(data map[string]any) BudFinancialDataEntity`

Create a new `ListLabel` entity instance. Pass `nil` for no initial data.

#### `ListTransaction(data map[string]any) BudFinancialDataEntity`

Create a new `ListTransaction` entity instance. Pass `nil` for no initial data.

#### `ManageFinancialData(data map[string]any) BudFinancialDataEntity`

Create a new `ManageFinancialData` entity instance. Pass `nil` for no initial data.

#### `ManageTransactionLabel(data map[string]any) BudFinancialDataEntity`

Create a new `ManageTransactionLabel` entity instance. Pass `nil` for no initial data.

#### `RetrieveFinancialData(data map[string]any) BudFinancialDataEntity`

Create a new `RetrieveFinancialData` entity instance. Pass `nil` for no initial data.

#### `Similar(data map[string]any) BudFinancialDataEntity`

Create a new `Similar` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## CorrectFinancialDataEntity

```go
correctFinancialData := client.CorrectFinancialData(nil)
fmt.Println(correctFinancialData.GetName()) // "correct_financial_data"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes | Date that the custom merchant was created, compliant with RFC3339. |
| `custom_merchant_id` | `string` | Yes | UUID representing the custom merchant. |
| `data` | `[]any` | Yes | List of merchant that matched the query. |
| `frequency` | `string` | Yes | The frequency to assign. |
| `include_similar` | `bool` | No | Apply the correction to the whole group of transactions similar to the specified transaction, rather than just the transaction itself |
| `logo_feedback` | `string` | Yes | The type of feedback for the merchant. |
| `metadata` | `map[string]any` | Yes | Metadata associated with the response schema |
| `name` | `string` | Yes | The display name of the custom merchant to be created |
| `online_or_billing_only` | `bool` | No | This is used to indicate if the custom merchant being created is an exclusively online or billing merchant. |
| `operation_id` | `string` | Yes | A unique identifier/reference associated with a given endpoint/operation |
| `reference_transaction_id` | `string` | No | The transaction whose group the specified transactions should join or leave. |
| `rule_definition` | `map[string]any` | Yes | The definition of the rule, mirroring the request body that created it. |
| `rule_type` | `string` | Yes | The type of correction rule |
| `similar` | `bool` | No | Whether the specified transactions are similar to each other (or to `reference_transaction_id`, if supplied) |
| `suggested_logo` | `string` | No |  |
| `suggested_url` | `string` | No | A suggested URL for the merchant, to help bud identify it and expand our merchant database. |
| `transaction_id` | `string` | Yes | The unique identifier for the transaction |
| `transaction_ids` | `[]any` | Yes | The transactions the rule applies to |

### Field Usage by Operation

| Field | load | list | create | remove |
| --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - |
| `custom_merchant_id` | - | - | - | - |
| `data` | - | - | - | - |
| `frequency` | - | - | - | - |
| `include_similar` | - | - | - | - |
| `logo_feedback` | - | - | - | - |
| `metadata` | - | - | - | - |
| `name` | - | - | - | - |
| `online_or_billing_only` | - | Yes | - | - |
| `operation_id` | - | - | - | - |
| `reference_transaction_id` | - | - | - | - |
| `rule_definition` | - | - | - | - |
| `rule_type` | - | - | - | - |
| `similar` | - | - | - | - |
| `suggested_logo` | - | - | - | - |
| `suggested_url` | - | Yes | - | - |
| `transaction_id` | - | - | - | - |
| `transaction_ids` | - | - | - | - |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.CorrectFinancialData(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.CorrectFinancialData(nil).Load(map[string]any{"rule_id": "rule_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CorrectFinancialData(nil).Create(map[string]any{
    "merchant_id": "example_merchant_id",
    "created_at": "example_created_at",
    "custom_merchant_id": "example_custom_merchant_id",
    "data": []any{},
    "frequency": "example_frequency",
    "logo_feedback": "example_logo_feedback",
    "metadata": map[string]any{},
    "name": "example_name",
    "operation_id": "example_operation_id",
    "rule_definition": map[string]any{},
    "rule_type": "example_rule_type",
    "transaction_id": "example_transaction_id",
    "transaction_ids": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.CorrectFinancialData(nil).Remove(map[string]any{"rule_id": "rule_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CorrectFinancialDataEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CustomerMerchantCorrectionEntity

```go
customerMerchantCorrection := client.CustomerMerchantCorrection(nil)
fmt.Println(customerMerchantCorrection.GetName()) // "customer_merchant_correction"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `[]any` | Yes |  |
| `metadata` | `map[string]any` | Yes | Metadata associated with the merchant correction response schema |
| `operation_id` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CustomerMerchantCorrection(nil).Create(map[string]any{
    "data": []any{},
    "metadata": map[string]any{},
    "operation_id": "example_operation_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CustomerMerchantCorrectionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## LabelEntity

```go
label := client.Label(nil)
fmt.Println(label.GetName()) // "label"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes | RFC3339 timestamp at which the label was created. |
| `id` | `string` | Yes | Unique identifier of the label. |
| `name` | `string` | Yes | Display name for the new label. |
| `updated_at` | `string` | Yes | RFC3339 timestamp at which the label was last updated. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Label(nil).Create(map[string]any{
    "created_at": "example_created_at",
    "id": "example_id",
    "name": "example_name",
    "updated_at": "example_updated_at",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Label(nil).Update(map[string]any{
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `LabelEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ListLabelEntity

```go
listLabel := client.ListLabel(nil)
fmt.Println(listLabel.GetName()) // "list_label"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes | RFC3339 timestamp at which the label was created. |
| `id` | `string` | Yes | Unique identifier of the label. |
| `name` | `string` | Yes | Display name of the label. |
| `updated_at` | `string` | Yes | RFC3339 timestamp at which the label was last updated. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.ListLabel(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ListLabelEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ListTransactionEntity

```go
listTransaction := client.ListTransaction(nil)
fmt.Println(listTransaction.GetName()) // "list_transaction"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | Yes | Identifier for the account associated with the transaction. |
| `amount` | `map[string]any` | Yes | The monetary amount. |
| `client_attributes` | `map[string]any` | No | An optional map of key-value string pairs that allows clients to attach arbitrary metadata to the transaction. |
| `counterparty` | `map[string]any` | No | An object containing details of the counterparty in this transaction |
| `credit_debit_indicator` | `string` | Yes | Credit/Debit Indicator |
| `date_time` | `string` | Yes | Date that the transaction occured compliant with RFC3339. |
| `description` | `string` | Yes | Description of the transaction. |
| `enrichments` | `map[string]any` | No | Contextual enrichments associated with a Transaction |
| `labels` | `[]any` | No | Customer-defined labels currently attached to this transaction. |
| `merchant_category_code` | `string` | No | Merchant category code conforming to ISO 18245, related to the type of services or goods the merchant provides for the transaction. |
| `posted_date_time` | `string` | No | Date the assets involved in the transaction transferred compliant with RFC3339. |
| `provider` | `string` | No | Name of the transaction source provider. |
| `running_balance` | `map[string]any` | Yes | The running balance for the account that the transaction takes place against |
| `running_balance_credit_debit_indicator` | `string` | No | Credit/Debit Indicator for the running balance field |
| `status` | `string` | Yes | Status of the transaction. |
| `suggested_description` | `string` | Yes | The description Bud suggests client apps show for a given transaction. |
| `suggested_logo` | `string` | No | The logo Bud suggests client apps show for a given transaction. |
| `tags` | `[]any` | No | A list of potential tags associated with the transaction after contextual enrichment, which can be used for filtering. |
| `transaction_id` | `string` | Yes | Unique identifier for the transaction. |
| `transaction_type` | `map[string]any` | No | The code and a description of the transaction type. |
| `value_date_time` | `string` | No | Date the assets involved in the transaction transferred compliant with RFC3339. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.ListTransaction(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ListTransactionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ManageFinancialDataEntity

```go
manageFinancialData := client.ManageFinancialData(nil)
fmt.Println(manageFinancialData.GetName()) // "manage_financial_data"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `label_id` | `string` | Yes | Identifier of the label to attach. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ManageFinancialData(nil).Create(map[string]any{
    "transaction_id": "example_transaction_id",
    "label_id": "example_label_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.ManageFinancialData(nil).Remove(map[string]any{"provider": "provider"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ManageFinancialDataEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ManageTransactionLabelEntity

```go
manageTransactionLabel := client.ManageTransactionLabel(nil)
fmt.Println(manageTransactionLabel.GetName()) // "manage_transaction_label"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.ManageTransactionLabel(nil).Remove(map[string]any{"id": "id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ManageTransactionLabelEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RetrieveFinancialDataEntity

```go
retrieveFinancialData := client.RetrieveFinancialData(nil)
fmt.Println(retrieveFinancialData.GetName()) // "retrieve_financial_data"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_category` | `string` | No | The account category that the account type is a part of Currently supported values include (but are not necessarily limited to): - `credit` - `depository` - `insurance` - `investment` - `loan` - `savings` - `other` |
| `account_id` | `string` | Yes | A unique id associated with the account. |
| `account_name` | `string` | No | The name associated with the account, this is often a friendly name assigned to the account to make it easier to refer to. |
| `account_type` | `string` | No | The type of account. |
| `balances` | `map[string]any` | No | The balances ingested for the account. |
| `booked` | `any` | Yes | The current balance of the account |
| `closed_at` | `string` | No | The datetime the account was closed in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `credit_lines` | `map[string]any` | No | The latest credit limit available for the account. |
| `currency` | `string` | Yes | The three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) of the monetary amount. |
| `data` | `map[string]any` | Yes |  |
| `date` | `string` | Yes | The date of the given balance in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `details` | `[]any` | No | Provides more details for the authorised payment. |
| `first_transaction_date` | `string` | Yes | The date of the first transaction that bud has stored for this account. |
| `frequency` | `string` | No | The frequency of the authorised payment. |
| `holder` | `map[string]any` | No | The account holder. |
| `holders` | `[]any` | No | The account holder(s). |
| `id` | `string` | Yes | The id for the authorised payment. |
| `identifiers` | `map[string]any` | No | These are wellknown fields which uniquely identify the account. |
| `last_transaction_date` | `string` | Yes | The date of the last transaction that bud has stored for this account. |
| `metadata` | `map[string]any` | Yes |  |
| `name` | `string` | Yes | The name of the payment. |
| `opening_date_time` | `string` | No | The datetime the account was opened in the format [RFC 3339] |
| `operation_id` | `string` | Yes |  |
| `pending` | `any` | Yes | The balance of the account if all pending transactions have settled. |
| `provider` | `string` | No | The account's provider, this is usually the bank or building society the account is held with. |
| `provider_display_name` | `string` | No | The display name for the account's provider. |
| `provider_logo` | `string` | No | A link to an image for the accounts provider's logo. |
| `reference` | `string` | Yes | The reference associated with the authorised payment. |
| `restriction` | `string` | No | The restriction on the account, if any. |
| `status` | `string` | No | The status of the account. |
| `suggested_name` | `string` | Yes | The name Bud suggests client apps show for the account. |
| `transaction_windows` | `[]any` | No | The transaction windows indicate for which periods we have full coverage of transactions ingested for the account. |
| `type` | `string` | Yes | The type of the authorised payment. |
| `usage_type` | `string` | No | The intended usage of the account. |

### Field Usage by Operation

| Field | load | list |
| --- | --- | --- |
| `account_category` | - | - |
| `account_id` | - | Yes |
| `account_name` | - | - |
| `account_type` | - | - |
| `balances` | - | Yes |
| `booked` | - | - |
| `closed_at` | - | - |
| `credit_lines` | - | - |
| `currency` | - | - |
| `data` | - | - |
| `date` | - | - |
| `details` | - | - |
| `first_transaction_date` | - | - |
| `frequency` | - | - |
| `holder` | - | - |
| `holders` | - | - |
| `id` | - | - |
| `identifiers` | - | - |
| `last_transaction_date` | - | - |
| `metadata` | - | - |
| `name` | - | - |
| `opening_date_time` | - | - |
| `operation_id` | - | - |
| `pending` | - | - |
| `provider` | - | - |
| `provider_display_name` | - | - |
| `provider_logo` | - | - |
| `reference` | - | - |
| `restriction` | - | - |
| `status` | - | Yes |
| `suggested_name` | - | - |
| `transaction_windows` | - | - |
| `type` | - | - |
| `usage_type` | - | - |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.RetrieveFinancialData(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.RetrieveFinancialData(nil).Load(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RetrieveFinancialDataEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SimilarEntity

```go
similar := client.Similar(nil)
fmt.Println(similar.GetName()) // "similar"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `[]any` | Yes |  |
| `id` | `string` | No |  |
| `metadata` | `map[string]any` | Yes |  |
| `operation_id` | `string` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Similar(nil).Load(map[string]any{"id": "similar_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SimilarEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `debug` | 0.0.1 | Debug capture |
| `idempotency` | 0.0.1 | Idempotency |
| `metrics` | 0.0.1 | Metrics |
| `paging` | 0.0.1 | Paging |
| `ratelimit` | 0.0.1 | Rate limiting |
| `retry` | 0.0.1 | Retry |
| `test` | 0.0.1 | Test transport |
| `timeout` | 0.0.1 | Timeout |


Features are activated via the `feature` option:

```go
client := sdk.NewBudFinancialDataSDK(map[string]any{
    "feature": map[string]any{
        "debug": map[string]any{"active": true},
        "idempotency": map[string]any{"active": true},
        "metrics": map[string]any{"active": true},
        "paging": map[string]any{"active": true},
        "ratelimit": map[string]any{"active": true},
        "retry": map[string]any{"active": true},
        "test": map[string]any{"active": true},
        "timeout": map[string]any{"active": true},
    },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### Ordering

`ratelimit`, `retry`, `timeout` wrap the transport. Each
wraps whatever is already installed, so **activation order is nesting order**:
a feature activated later sits OUTSIDE one activated earlier, and sees the call
first.

That decides behaviour, not just sequence: a feature that short-circuits the
call, such as a cache serving a hit, stops every feature nested inside it from
ever seeing that call.

`debug`, `idempotency`, `metrics`, `paging`, `test` attach to pipeline hooks
rather than the transport, so their order does not affect what they observe.

#### `debug`

Debug capture.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `max` | `100` |
| `redact` | `['authorization', 'cookie', 'set-cookie', 'api-key', 'apikey', 'x-api-key', 'idempotency-key']` |

| Option | Type |
|---|---|
| `now` | function |
| `onEntry` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.debug.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `idempotency`

Idempotency.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `header` | `'Idempotency-Key'` |
| `methods` | `['POST', 'PUT', 'PATCH', 'DELETE']` |
| `ops` | `['create', 'update', 'remove']` |

| Option | Type |
|---|---|
| `keygen` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.idempotency.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `metrics`

Metrics.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `now` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.metrics.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `paging`

Paging.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `afterVar` | `'after'` |
| `cursorParam` | `'cursor'` |
| `firstVar` | `'first'` |
| `limitParam` | `'limit'` |
| `pageParam` | `'page'` |
| `startPage` | `1` |

| Option | Type |
|---|---|
| `limit` | number |
| `ops` | list |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.paging.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `ratelimit`

Rate limiting.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

| Option | Type |
|---|---|
| `now` | function |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.ratelimit.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `retry`

Retry.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

| Option | Type |
|---|---|
| `jitter` | boolean |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.retry.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `test`

Test transport.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

#### `timeout`

Timeout.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

| Option | Type |
|---|---|
| `clearTimer` | function |
| `setTimer` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.timeout.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

