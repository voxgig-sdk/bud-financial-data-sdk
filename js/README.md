# BudFinancialData JavaScript SDK



The JavaScript SDK for the BudFinancialData API — an entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.CorrectFinancialData()` — each with a small set of operations (`list`, `load`, `create`, `update`, `remove`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```js
npm install bud-financial-data
```
## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.


### Create a Client

```js
const { BudFinancialDataSDK } = require('@voxgig-sdk/bud-financial-data-sdk-js')

const client = new BudFinancialDataSDK({
  apikey: process.env.BUD_FINANCIAL_DATA_APIKEY,
})
```

### Load a CorrectFinancialData

```js
const correct_financial_data = await client.CorrectFinancialData().load({ rule_id: 'example_rule_id' })
console.log(correct_financial_data)
```

### List CorrectFinancialData Records

```js
const correct_financial_datas = await client.CorrectFinancialData().list()
for (const correct_financial_data of correct_financial_datas) {
  console.log(correct_financial_data)
}
```

### Create a CorrectFinancialData

```js
const created = await client.CorrectFinancialData().create({
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
console.log(created)
```

### Remove a CorrectFinancialData

```js
await client.CorrectFinancialData().remove({ rule_id: 'example_rule_id' })
```

### Direct API Access

Use `client.direct()` to call any API endpoint directly:

```js
const result = await client.direct({
  path: '/custom/endpoint/{id}',
  method: 'GET',
  params: { id: 'abc123' },
})

if (result.ok) {
  console.log(result.data)
}
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const listtransactions = await client.ListTransaction().list()
  console.log(listtransactions)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```js
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```js
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```js
const client = BudFinancialDataSDK.test()

const listtransaction = await client.ListTransaction().list()
// listtransaction is the entity, populated with mock response data
// — call listtransaction.data() for the record itself
console.log(listtransaction)
```

You can also use the instance method:

```js
const client = new BudFinancialDataSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```js
const entity = client.ListTransaction()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data)
```

### Add custom middleware

Pass features via the `extend` option:

```js
const logger = {
  hooks: {
    PreRequest: (ctx) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new BudFinancialDataSDK({
  apikey: '...',
  extend: [logger],
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
BUD_FINANCIAL_DATA_TEST_LIVE=TRUE
BUD_FINANCIAL_DATA_APIKEY=<your-key>
```

Then run:

```bash
cd js && npm test
```


## Reference

### BudFinancialDataSDK

#### Constructor

```js
new BudFinancialDataSDK(options?)
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `CorrectFinancialData(data?)` | `CorrectFinancialDataEntity` | Create a CorrectFinancialData entity instance. |
| `CustomerMerchantCorrection(data?)` | `CustomerMerchantCorrectionEntity` | Create a CustomerMerchantCorrection entity instance. |
| `Label(data?)` | `LabelEntity` | Create a Label entity instance. |
| `ListLabel(data?)` | `ListLabelEntity` | Create a ListLabel entity instance. |
| `ListTransaction(data?)` | `ListTransactionEntity` | Create a ListTransaction entity instance. |
| `ManageFinancialData(data?)` | `ManageFinancialDataEntity` | Create a ManageFinancialData entity instance. |
| `ManageTransactionLabel(data?)` | `ManageTransactionLabelEntity` | Create a ManageTransactionLabel entity instance. |
| `RetrieveFinancialData(data?)` | `RetrieveFinancialDataEntity` | Create a RetrieveFinancialData entity instance. |
| `Similar(data?)` | `SimilarEntity` | Create a Similar entity instance. |
| `tester(testopts?, sdkopts?)` | `BudFinancialDataSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `BudFinancialDataSDK.test(testopts?, sdkopts?)` | `BudFinancialDataSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `update` | `update(reqdata?, ctrl?): Promise<Entity>` | Update an existing entity. |
| `remove` | `remove(reqmatch?, ctrl?): Promise<void>` | Remove an entity. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): BudFinancialDataSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load`, `create` and `update` resolve to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).
- `remove` resolves to `undefined`.

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```js
{
  ok: true,
  status: 200,
  headers: {},
  data: {}
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```js
{
  url: 'string',
  method: 'string',
  headers: {},
  body: undefined
}
```

### Entities

#### CorrectFinancialData

| Field | Description |
| --- | --- |
| `created_at` | Date that the custom merchant was created, compliant with RFC3339. |
| `custom_merchant_id` | UUID representing the custom merchant. |
| `data` | List of merchant that matched the query. |
| `frequency` | The frequency to assign. |
| `include_similar` | Apply the correction to the whole group of transactions similar to the specified transaction, rather than just the transaction itself |
| `logo_feedback` | The type of feedback for the merchant. |
| `metadata` | Metadata associated with the response schema |
| `name` | The display name of the custom merchant to be created |
| `online_or_billing_only` | This is used to indicate if the custom merchant being created is an exclusively online or billing merchant. |
| `operation_id` | A unique identifier/reference associated with a given endpoint/operation |
| `reference_transaction_id` | The transaction whose group the specified transactions should join or leave. |
| `rule_definition` | The definition of the rule, mirroring the request body that created it. |
| `rule_type` | The type of correction rule |
| `similar` | Whether the specified transactions are similar to each other (or to `reference_transaction_id`, if supplied) |
| `suggested_logo` |  |
| `suggested_url` | A suggested URL for the merchant, to help bud identify it and expand our merchant database. |
| `transaction_id` | The unique identifier for the transaction |
| `transaction_ids` | The transactions the rule applies to |

Operations: create, list, load, remove.

API path: `/corrections/v2/categories`

#### CustomerMerchantCorrection

| Field | Description |
| --- | --- |
| `data` |  |
| `metadata` | Metadata associated with the merchant correction response schema |
| `operation_id` |  |

Operations: create.

API path: `/corrections/v2/merchants`

#### Label

| Field | Description |
| --- | --- |
| `created_at` | RFC3339 timestamp at which the label was created. |
| `id` | Unique identifier of the label. |
| `name` | Display name for the new label. |
| `updated_at` | RFC3339 timestamp at which the label was last updated. |

Operations: create, update.

API path: `/financial/v2/transactions/labels`

#### ListLabel

| Field | Description |
| --- | --- |
| `created_at` | RFC3339 timestamp at which the label was created. |
| `id` | Unique identifier of the label. |
| `name` | Display name of the label. |
| `updated_at` | RFC3339 timestamp at which the label was last updated. |

Operations: list.

API path: `/financial/v2/transactions/labels`

#### ListTransaction

| Field | Description |
| --- | --- |
| `account_id` | Identifier for the account associated with the transaction. |
| `amount` | The monetary amount. |
| `client_attributes` | An optional map of key-value string pairs that allows clients to attach arbitrary metadata to the transaction. |
| `counterparty` | An object containing details of the counterparty in this transaction |
| `credit_debit_indicator` | Credit/Debit Indicator |
| `date_time` | Date that the transaction occured compliant with RFC3339. |
| `description` | Description of the transaction. |
| `enrichments` | Contextual enrichments associated with a Transaction |
| `labels` | Customer-defined labels currently attached to this transaction. |
| `merchant_category_code` | Merchant category code conforming to ISO 18245, related to the type of services or goods the merchant provides for the transaction. |
| `posted_date_time` | Date the assets involved in the transaction transferred compliant with RFC3339. |
| `provider` | Name of the transaction source provider. |
| `running_balance` | The running balance for the account that the transaction takes place against |
| `running_balance_credit_debit_indicator` | Credit/Debit Indicator for the running balance field |
| `status` | Status of the transaction. |
| `suggested_description` | The description Bud suggests client apps show for a given transaction. |
| `suggested_logo` | The logo Bud suggests client apps show for a given transaction. |
| `tags` | A list of potential tags associated with the transaction after contextual enrichment, which can be used for filtering. |
| `transaction_id` | Unique identifier for the transaction. |
| `transaction_type` | The code and a description of the transaction type. |
| `value_date_time` | Date the assets involved in the transaction transferred compliant with RFC3339. |

Operations: list.

API path: `/financial/v2/transactions`

#### ManageFinancialData

| Field | Description |
| --- | --- |
| `label_id` | Identifier of the label to attach. |

Operations: create, remove.

API path: `/financial/v2/transactions/{transaction_id}/labels`

#### ManageTransactionLabel

| Field | Description |
| --- | --- |
| `id` |  |

Operations: remove.

API path: `/financial/v2/transactions/labels/{label_id}`

#### RetrieveFinancialData

| Field | Description |
| --- | --- |
| `account_category` | The account category that the account type is a part of Currently supported values include (but are not necessarily limited to): - `credit` - `depository` - `insurance` - `investment` - `loan` - `savings` - `other` |
| `account_id` | A unique id associated with the account. |
| `account_name` | The name associated with the account, this is often a friendly name assigned to the account to make it easier to refer to. |
| `account_type` | The type of account. |
| `balances` | The balances ingested for the account. |
| `booked` | The current balance of the account |
| `closed_at` | The datetime the account was closed in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `credit_lines` | The latest credit limit available for the account. |
| `currency` | The three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) of the monetary amount. |
| `data` |  |
| `date` | The date of the given balance in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `details` | Provides more details for the authorised payment. |
| `first_transaction_date` | The date of the first transaction that bud has stored for this account. |
| `frequency` | The frequency of the authorised payment. |
| `holder` | The account holder. |
| `holders` | The account holder(s). |
| `id` | The id for the authorised payment. |
| `identifiers` | These are wellknown fields which uniquely identify the account. |
| `last_transaction_date` | The date of the last transaction that bud has stored for this account. |
| `metadata` |  |
| `name` | The name of the payment. |
| `opening_date_time` | The datetime the account was opened in the format [RFC 3339] |
| `operation_id` |  |
| `pending` | The balance of the account if all pending transactions have settled. |
| `provider` | The account's provider, this is usually the bank or building society the account is held with. |
| `provider_display_name` | The display name for the account's provider. |
| `provider_logo` | A link to an image for the accounts provider's logo. |
| `reference` | The reference associated with the authorised payment. |
| `restriction` | The restriction on the account, if any. |
| `status` | The status of the account. |
| `suggested_name` | The name Bud suggests client apps show for the account. |
| `transaction_windows` | The transaction windows indicate for which periods we have full coverage of transactions ingested for the account. |
| `type` | The type of the authorised payment. |
| `usage_type` | The intended usage of the account. |

Operations: list, load.

API path: `/financial/v3/accounts`

#### Similar

| Field | Description |
| --- | --- |
| `data` |  |
| `id` |  |
| `metadata` |  |
| `operation_id` |  |

Operations: load.

API path: `/corrections/v2/categories/similar/{transaction_id}`



## Entities


### CorrectFinancialData

Create an instance: `const correct_financial_data = client.CorrectFinancialData()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Date that the custom merchant was created, compliant with RFC3339. |
| `custom_merchant_id` | `string` | UUID representing the custom merchant. |
| `data` | `Array` | List of merchant that matched the query. |
| `frequency` | `string` | The frequency to assign. |
| `include_similar` | `boolean` | Apply the correction to the whole group of transactions similar to the specified transaction, rather than just the transaction itself |
| `logo_feedback` | `string` | The type of feedback for the merchant. |
| `metadata` | `Object` | Metadata associated with the response schema |
| `name` | `string` | The display name of the custom merchant to be created |
| `online_or_billing_only` | `boolean` | This is used to indicate if the custom merchant being created is an exclusively online or billing merchant. |
| `operation_id` | `string` | A unique identifier/reference associated with a given endpoint/operation |
| `reference_transaction_id` | `string` | The transaction whose group the specified transactions should join or leave. |
| `rule_definition` | `Object` | The definition of the rule, mirroring the request body that created it. |
| `rule_type` | `string` | The type of correction rule |
| `similar` | `boolean` | Whether the specified transactions are similar to each other (or to `reference_transaction_id`, if supplied) |
| `suggested_logo` | `string` |  |
| `suggested_url` | `string` | A suggested URL for the merchant, to help bud identify it and expand our merchant database. |
| `transaction_id` | `string` | The unique identifier for the transaction |
| `transaction_ids` | `Array` | The transactions the rule applies to |

#### Example: Load

```ts
const correct_financial_data = await client.CorrectFinancialData().load({ rule_id: 'rule_id' })
```

#### Example: List

```ts
const correct_financial_datas = await client.CorrectFinancialData().list()
```

#### Example: Create

```ts
const correct_financial_data = await client.CorrectFinancialData().create({
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


### CustomerMerchantCorrection

Create an instance: `const customer_merchant_correction = client.CustomerMerchantCorrection()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `Array` |  |
| `metadata` | `Object` | Metadata associated with the merchant correction response schema |
| `operation_id` | `string` |  |

#### Example: Create

```ts
const customer_merchant_correction = await client.CustomerMerchantCorrection().create({
  data: [],
  metadata: {},
  operation_id: 'example_operation_id',
})
```


### Label

Create an instance: `const label = client.Label()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | RFC3339 timestamp at which the label was created. |
| `id` | `string` | Unique identifier of the label. |
| `name` | `string` | Display name for the new label. |
| `updated_at` | `string` | RFC3339 timestamp at which the label was last updated. |

#### Example: Create

```ts
const label = await client.Label().create({
  created_at: 'example_created_at',
  id: 'example_id',
  name: 'example_name',
  updated_at: 'example_updated_at',
})
```


### ListLabel

Create an instance: `const list_label = client.ListLabel()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | RFC3339 timestamp at which the label was created. |
| `id` | `string` | Unique identifier of the label. |
| `name` | `string` | Display name of the label. |
| `updated_at` | `string` | RFC3339 timestamp at which the label was last updated. |

#### Example: List

```ts
const list_labels = await client.ListLabel().list()
```


### ListTransaction

Create an instance: `const list_transaction = client.ListTransaction()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` | Identifier for the account associated with the transaction. |
| `amount` | `Object` | The monetary amount. |
| `client_attributes` | `Object` | An optional map of key-value string pairs that allows clients to attach arbitrary metadata to the transaction. |
| `counterparty` | `Object` | An object containing details of the counterparty in this transaction |
| `credit_debit_indicator` | `string` | Credit/Debit Indicator |
| `date_time` | `string` | Date that the transaction occured compliant with RFC3339. |
| `description` | `string` | Description of the transaction. |
| `enrichments` | `Object` | Contextual enrichments associated with a Transaction |
| `labels` | `Array` | Customer-defined labels currently attached to this transaction. |
| `merchant_category_code` | `string` | Merchant category code conforming to ISO 18245, related to the type of services or goods the merchant provides for the transaction. |
| `posted_date_time` | `string` | Date the assets involved in the transaction transferred compliant with RFC3339. |
| `provider` | `string` | Name of the transaction source provider. |
| `running_balance` | `Object` | The running balance for the account that the transaction takes place against |
| `running_balance_credit_debit_indicator` | `string` | Credit/Debit Indicator for the running balance field |
| `status` | `string` | Status of the transaction. |
| `suggested_description` | `string` | The description Bud suggests client apps show for a given transaction. |
| `suggested_logo` | `string` | The logo Bud suggests client apps show for a given transaction. |
| `tags` | `Array` | A list of potential tags associated with the transaction after contextual enrichment, which can be used for filtering. |
| `transaction_id` | `string` | Unique identifier for the transaction. |
| `transaction_type` | `Object` | The code and a description of the transaction type. |
| `value_date_time` | `string` | Date the assets involved in the transaction transferred compliant with RFC3339. |

#### Example: List

```ts
const list_transactions = await client.ListTransaction().list()
```


### ManageFinancialData

Create an instance: `const manage_financial_data = client.ManageFinancialData()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `label_id` | `string` | Identifier of the label to attach. |

#### Example: Create

```ts
const manage_financial_data = await client.ManageFinancialData().create({
  transaction_id: 'example_transaction_id',
  label_id: 'example_label_id',
})
```


### ManageTransactionLabel

Create an instance: `const manage_transaction_label = client.ManageTransactionLabel()`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |


### RetrieveFinancialData

Create an instance: `const retrieve_financial_data = client.RetrieveFinancialData()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_category` | `string` | The account category that the account type is a part of Currently supported values include (but are not necessarily limited to): - `credit` - `depository` - `insurance` - `investment` - `loan` - `savings` - `other` |
| `account_id` | `string` | A unique id associated with the account. |
| `account_name` | `string` | The name associated with the account, this is often a friendly name assigned to the account to make it easier to refer to. |
| `account_type` | `string` | The type of account. |
| `balances` | `Object` | The balances ingested for the account. |
| `booked` | `*` | The current balance of the account |
| `closed_at` | `string` | The datetime the account was closed in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `credit_lines` | `Object` | The latest credit limit available for the account. |
| `currency` | `string` | The three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) of the monetary amount. |
| `data` | `Object` |  |
| `date` | `string` | The date of the given balance in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `details` | `Array` | Provides more details for the authorised payment. |
| `first_transaction_date` | `string` | The date of the first transaction that bud has stored for this account. |
| `frequency` | `string` | The frequency of the authorised payment. |
| `holder` | `Object` | The account holder. |
| `holders` | `Array` | The account holder(s). |
| `id` | `string` | The id for the authorised payment. |
| `identifiers` | `Object` | These are wellknown fields which uniquely identify the account. |
| `last_transaction_date` | `string` | The date of the last transaction that bud has stored for this account. |
| `metadata` | `Object` |  |
| `name` | `string` | The name of the payment. |
| `opening_date_time` | `string` | The datetime the account was opened in the format [RFC 3339] |
| `operation_id` | `string` |  |
| `pending` | `*` | The balance of the account if all pending transactions have settled. |
| `provider` | `string` | The account's provider, this is usually the bank or building society the account is held with. |
| `provider_display_name` | `string` | The display name for the account's provider. |
| `provider_logo` | `string` | A link to an image for the accounts provider's logo. |
| `reference` | `string` | The reference associated with the authorised payment. |
| `restriction` | `string` | The restriction on the account, if any. |
| `status` | `string` | The status of the account. |
| `suggested_name` | `string` | The name Bud suggests client apps show for the account. |
| `transaction_windows` | `Array` | The transaction windows indicate for which periods we have full coverage of transactions ingested for the account. |
| `type` | `string` | The type of the authorised payment. |
| `usage_type` | `string` | The intended usage of the account. |

#### Example: Load

```ts
const retrieve_financial_data = await client.RetrieveFinancialData().load({ account_id: 'account_id' })
```

#### Example: List

```ts
const retrieve_financial_datas = await client.RetrieveFinancialData().list()
```


### Similar

Create an instance: `const similar = client.Similar()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `Array` |  |
| `id` | `string` |  |
| `metadata` | `Object` |  |
| `operation_id` | `string` |  |

#### Example: Load

```ts
const similar = await client.Similar().load({ id: 'similar_id' })
```

## Features

This SDK ships 8 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`debug`](#debug) | Debug capture |
| [`idempotency`](#idempotency) | Idempotency |
| [`metrics`](#metrics) | Metrics |
| [`paging`](#paging) | Paging |
| [`ratelimit`](#ratelimit) | Rate limiting |
| [`retry`](#retry) | Retry |
| [`test`](#test) | Test transport |
| [`timeout`](#timeout) | Timeout |

> **Order matters for `ratelimit`, `retry`, `timeout`.** These wrap the
> transport, so each one wraps whatever is already installed: the order you
> activate them in IS the nesting order. Activating them as an ordered list
> rather than a map is what fixes that order.

### debug

Debug capture.

| Option | Default |
|---|---|
| `active` | `false` |
| `max` | `100` |
| `redact` | `['authorization', 'cookie', 'set-cookie', 'api-key', 'apikey', 'x-api-key', 'idempotency-key']` |

Set `feature.debug.active` to enable it, then override any of the options above.

### idempotency

Idempotency.

| Option | Default |
|---|---|
| `active` | `false` |
| `header` | `'Idempotency-Key'` |
| `methods` | `['POST', 'PUT', 'PATCH', 'DELETE']` |
| `ops` | `['create', 'update', 'remove']` |

Set `feature.idempotency.active` to enable it, then override any of the options above.

### metrics

Metrics.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.metrics.active` to enable it, then override any of the options above.

### paging

Paging.

| Option | Default |
|---|---|
| `active` | `false` |
| `afterVar` | `'after'` |
| `cursorParam` | `'cursor'` |
| `firstVar` | `'first'` |
| `limitParam` | `'limit'` |
| `pageParam` | `'page'` |
| `startPage` | `1` |

Set `feature.paging.active` to enable it, then override any of the options above.

### ratelimit

Rate limiting.

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

Set `feature.ratelimit.active` to enable it, then override any of the options above.

`ratelimit` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### retry

Retry.

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

Set `feature.retry.active` to enable it, then override any of the options above.

`retry` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### test

Test transport.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.

### timeout

Timeout.

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

Set `feature.timeout.active` to enable it, then override any of the options above.

`timeout` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **DebugFeature**: Debug capture
- **IdempotencyFeature**: Idempotency
- **MetricsFeature**: Metrics
- **PagingFeature**: Paging
- **RatelimitFeature**: Rate limiting
- **RetryFeature**: Retry
- **TestFeature**: Test transport
- **TimeoutFeature**: Timeout

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
bud-financial-data/
├── src/
│   ├── BudFinancialDataSDK.js        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
└── test/                   # Test suites
```

Import the SDK from the package root:

```js
const { BudFinancialDataSDK } = require('@voxgig-sdk/bud-financial-data-sdk-js')
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const listtransaction = client.ListTransaction()
await listtransaction.list()

// listtransaction.data() now returns the listtransaction data from the last `list`
// listtransaction.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
