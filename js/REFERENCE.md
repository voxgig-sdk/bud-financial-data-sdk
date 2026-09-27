# BudFinancialData JavaScript SDK Reference

Complete API reference for the BudFinancialData JavaScript SDK.


## BudFinancialDataSDK

### Constructor

```ts
new BudFinancialDataSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `BudFinancialDataSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = BudFinancialDataSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `BudFinancialDataSDK` instance in test mode.


### Instance Methods

#### `CorrectFinancialData(data?: object)`

Create a new `CorrectFinancialData` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CorrectFinancialDataEntity` instance.

#### `CustomerMerchantCorrection(data?: object)`

Create a new `CustomerMerchantCorrection` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CustomerMerchantCorrectionEntity` instance.

#### `Label(data?: object)`

Create a new `Label` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `LabelEntity` instance.

#### `ListLabel(data?: object)`

Create a new `ListLabel` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ListLabelEntity` instance.

#### `ListTransaction(data?: object)`

Create a new `ListTransaction` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ListTransactionEntity` instance.

#### `ManageFinancialData(data?: object)`

Create a new `ManageFinancialData` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ManageFinancialDataEntity` instance.

#### `ManageTransactionLabel(data?: object)`

Create a new `ManageTransactionLabel` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ManageTransactionLabelEntity` instance.

#### `RetrieveFinancialData(data?: object)`

Create a new `RetrieveFinancialData` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RetrieveFinancialDataEntity` instance.

#### `Similar(data?: object)`

Create a new `Similar` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SimilarEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `BudFinancialDataSDK.test()`.

**Returns:** `BudFinancialDataSDK` instance in test mode.


---

## CorrectFinancialDataEntity

```ts
const correct_financial_data = client.CorrectFinancialData()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes | Date that the custom merchant was created, compliant with RFC3339. |
| `custom_merchant_id` | `string` | Yes | UUID representing the custom merchant. |
| `data` | `Array` | Yes | List of merchant that matched the query. |
| `frequency` | `string` | Yes | The frequency to assign. |
| `include_similar` | `boolean` | No | Apply the correction to the whole group of transactions similar to the specified transaction, rather than just the transaction itself |
| `logo_feedback` | `string` | Yes | The type of feedback for the merchant. |
| `metadata` | `Object` | Yes | Metadata associated with the response schema |
| `name` | `string` | Yes | The display name of the custom merchant to be created |
| `online_or_billing_only` | `boolean` | No | This is used to indicate if the custom merchant being created is an exclusively online or billing merchant. |
| `operation_id` | `string` | Yes | A unique identifier/reference associated with a given endpoint/operation |
| `reference_transaction_id` | `string` | No | The transaction whose group the specified transactions should join or leave. |
| `rule_definition` | `Object` | Yes | The definition of the rule, mirroring the request body that created it. |
| `rule_type` | `string` | Yes | The type of correction rule |
| `similar` | `boolean` | No | Whether the specified transactions are similar to each other (or to `reference_transaction_id`, if supplied) |
| `suggested_logo` | `string` | No |  |
| `suggested_url` | `string` | No | A suggested URL for the merchant, to help bud identify it and expand our merchant database. |
| `transaction_id` | `string` | Yes | The unique identifier for the transaction |
| `transaction_ids` | `Array` | Yes | The transactions the rule applies to |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CorrectFinancialData().create({
  merchant_id: 'example_merchant_id',
  created_at: 'example_created_at',
  custom_merchant_id: 'example_custom_merchant_id',
  data: [],
  frequency: 'example_frequency',
  logo_feedback: 'example_logo_feedback',
  metadata: {},
  name: 'example_name',
  operation_id: 'example_operation_id',
  rule_definition: {},
  rule_type: 'example_rule_type',
  transaction_id: 'example_transaction_id',
  transaction_ids: [],
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.CorrectFinancialData().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.CorrectFinancialData().load({ rule_id: 'rule_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.CorrectFinancialData().remove({ rule_id: 'rule_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CorrectFinancialDataEntity` instance with the same client and
options.

#### `client()`

Return the parent `BudFinancialDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CustomerMerchantCorrectionEntity

```ts
const customer_merchant_correction = client.CustomerMerchantCorrection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `Array` | Yes |  |
| `metadata` | `Object` | Yes | Metadata associated with the merchant correction response schema |
| `operation_id` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CustomerMerchantCorrection().create({
  data: [],
  metadata: {},
  operation_id: 'example_operation_id',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CustomerMerchantCorrectionEntity` instance with the same client and
options.

#### `client()`

Return the parent `BudFinancialDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## LabelEntity

```ts
const label = client.Label()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes | RFC3339 timestamp at which the label was created. |
| `id` | `string` | Yes | Unique identifier of the label. |
| `name` | `string` | Yes | Display name for the new label. |
| `updated_at` | `string` | Yes | RFC3339 timestamp at which the label was last updated. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Label().create({
  created_at: 'example_created_at',
  id: 'example_id',
  name: 'example_name',
  updated_at: 'example_updated_at',
})
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Label().update({
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `LabelEntity` instance with the same client and
options.

#### `client()`

Return the parent `BudFinancialDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ListLabelEntity

```ts
const list_label = client.ListLabel()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes | RFC3339 timestamp at which the label was created. |
| `id` | `string` | Yes | Unique identifier of the label. |
| `name` | `string` | Yes | Display name of the label. |
| `updated_at` | `string` | Yes | RFC3339 timestamp at which the label was last updated. |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.ListLabel().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ListLabelEntity` instance with the same client and
options.

#### `client()`

Return the parent `BudFinancialDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ListTransactionEntity

```ts
const list_transaction = client.ListTransaction()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | Yes | Identifier for the account associated with the transaction. |
| `amount` | `Object` | Yes | The monetary amount. |
| `client_attributes` | `Object` | No | An optional map of key-value string pairs that allows clients to attach arbitrary metadata to the transaction. |
| `counterparty` | `Object` | No | An object containing details of the counterparty in this transaction |
| `credit_debit_indicator` | `string` | Yes | Credit/Debit Indicator |
| `date_time` | `string` | Yes | Date that the transaction occured compliant with RFC3339. |
| `description` | `string` | Yes | Description of the transaction. |
| `enrichments` | `Object` | No | Contextual enrichments associated with a Transaction |
| `labels` | `Array` | No | Customer-defined labels currently attached to this transaction. |
| `merchant_category_code` | `string` | No | Merchant category code conforming to ISO 18245, related to the type of services or goods the merchant provides for the transaction. |
| `posted_date_time` | `string` | No | Date the assets involved in the transaction transferred compliant with RFC3339. |
| `provider` | `string` | No | Name of the transaction source provider. |
| `running_balance` | `Object` | Yes | The running balance for the account that the transaction takes place against |
| `running_balance_credit_debit_indicator` | `string` | No | Credit/Debit Indicator for the running balance field |
| `status` | `string` | Yes | Status of the transaction. |
| `suggested_description` | `string` | Yes | The description Bud suggests client apps show for a given transaction. |
| `suggested_logo` | `string` | No | The logo Bud suggests client apps show for a given transaction. |
| `tags` | `Array` | No | A list of potential tags associated with the transaction after contextual enrichment, which can be used for filtering. |
| `transaction_id` | `string` | Yes | Unique identifier for the transaction. |
| `transaction_type` | `Object` | No | The code and a description of the transaction type. |
| `value_date_time` | `string` | No | Date the assets involved in the transaction transferred compliant with RFC3339. |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.ListTransaction().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ListTransactionEntity` instance with the same client and
options.

#### `client()`

Return the parent `BudFinancialDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ManageFinancialDataEntity

```ts
const manage_financial_data = client.ManageFinancialData()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `label_id` | `string` | Yes | Identifier of the label to attach. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ManageFinancialData().create({
  transaction_id: 'example_transaction_id',
  label_id: 'example_label_id',
})
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.ManageFinancialData().remove({ provider: 'provider' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ManageFinancialDataEntity` instance with the same client and
options.

#### `client()`

Return the parent `BudFinancialDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ManageTransactionLabelEntity

```ts
const manage_transaction_label = client.ManageTransactionLabel()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |

### Operations

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.ManageTransactionLabel().remove({ id: 'id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ManageTransactionLabelEntity` instance with the same client and
options.

#### `client()`

Return the parent `BudFinancialDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RetrieveFinancialDataEntity

```ts
const retrieve_financial_data = client.RetrieveFinancialData()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_category` | `string` | No | The account category that the account type is a part of Currently supported values include (but are not necessarily limited to): - `credit` - `depository` - `insurance` - `investment` - `loan` - `savings` - `other` |
| `account_id` | `string` | Yes | A unique id associated with the account. |
| `account_name` | `string` | No | The name associated with the account, this is often a friendly name assigned to the account to make it easier to refer to. |
| `account_type` | `string` | No | The type of account. |
| `balances` | `Object` | No | The balances ingested for the account. |
| `booked` | `*` | Yes | The current balance of the account |
| `closed_at` | `string` | No | The datetime the account was closed in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `credit_lines` | `Object` | No | The latest credit limit available for the account. |
| `currency` | `string` | Yes | The three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) of the monetary amount. |
| `data` | `Object` | Yes |  |
| `date` | `string` | Yes | The date of the given balance in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `details` | `Array` | No | Provides more details for the authorised payment. |
| `first_transaction_date` | `string` | Yes | The date of the first transaction that bud has stored for this account. |
| `frequency` | `string` | No | The frequency of the authorised payment. |
| `holder` | `Object` | No | The account holder. |
| `holders` | `Array` | No | The account holder(s). |
| `id` | `string` | Yes | The id for the authorised payment. |
| `identifiers` | `Object` | No | These are wellknown fields which uniquely identify the account. |
| `last_transaction_date` | `string` | Yes | The date of the last transaction that bud has stored for this account. |
| `metadata` | `Object` | Yes |  |
| `name` | `string` | Yes | The name of the payment. |
| `opening_date_time` | `string` | No | The datetime the account was opened in the format [RFC 3339] |
| `operation_id` | `string` | Yes |  |
| `pending` | `*` | Yes | The balance of the account if all pending transactions have settled. |
| `provider` | `string` | No | The account's provider, this is usually the bank or building society the account is held with. |
| `provider_display_name` | `string` | No | The display name for the account's provider. |
| `provider_logo` | `string` | No | A link to an image for the accounts provider's logo. |
| `reference` | `string` | Yes | The reference associated with the authorised payment. |
| `restriction` | `string` | No | The restriction on the account, if any. |
| `status` | `string` | No | The status of the account. |
| `suggested_name` | `string` | Yes | The name Bud suggests client apps show for the account. |
| `transaction_windows` | `Array` | No | The transaction windows indicate for which periods we have full coverage of transactions ingested for the account. |
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.RetrieveFinancialData().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.RetrieveFinancialData().load({ account_id: 'account_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RetrieveFinancialDataEntity` instance with the same client and
options.

#### `client()`

Return the parent `BudFinancialDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SimilarEntity

```ts
const similar = client.Similar()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `Array` | Yes |  |
| `id` | `string` | No |  |
| `metadata` | `Object` | Yes |  |
| `operation_id` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Similar().load({ id: 'similar_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SimilarEntity` instance with the same client and
options.

#### `client()`

Return the parent `BudFinancialDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


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

```ts
const client = new BudFinancialDataSDK({
  feature: {
    debug: { active: true },
    idempotency: { active: true },
    metrics: { active: true },
    paging: { active: true },
    ratelimit: { active: true },
    retry: { active: true },
    test: { active: true },
    timeout: { active: true },
  }
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

