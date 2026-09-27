# BudFinancialData Lua SDK



The Lua SDK for the BudFinancialData API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:CorrectFinancialData()` — each with the same small set of operations (`list`, `load`, `create`, `update`, `remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/bud-financial-data-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("bud-financial-data_sdk")

local client = sdk.new({
  apikey = os.getenv("BUD_FINANCIAL_DATA_APIKEY"),
})
```

### 2. List correctfinancialdata records

Entity operations return `(value, err)`. For `list`, `value` is the
array of records itself — iterate it directly (there is no wrapper).

```lua
local correctfinancialdatas, err = client:CorrectFinancialData():list()
if err then error(err) end

for _, item in ipairs(correctfinancialdatas) do
  print(item)
end
```

### 3. Load a correctfinancialdata

```lua
local correctfinancialdata, err = client:CorrectFinancialData():load({ rule_id = "example_rule_id" })
if err then error(err) end
print(correctfinancialdata)
```

### 4. Create, update, and remove

```lua
-- Create
local created, err = client:CorrectFinancialData():create({ merchant_id = "example_merchant_id", created_at = "example_created_at", custom_merchant_id = "example_custom_merchant_id", data = {}, frequency = "example_frequency", logo_feedback = "example_logo_feedback", metadata = {}, name = "example_name", operation_id = "example_operation_id", rule_definition = {}, rule_type = "example_rule_type", transaction_id = "example_transaction_id", transaction_ids = {} })
if err then error(err) end

-- Remove
client:CorrectFinancialData():remove({ rule_id = "example_rule_id" })
```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local listtransactions, err = client:ListTransaction():list()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:ListTransaction():list()
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
  },
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
cd lua && busted test/
```


## Reference

### BudFinancialDataSDK

```lua
local sdk = require("bud-financial-data_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### BudFinancialDataSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
| `CorrectFinancialData` | `(data) -> CorrectFinancialDataEntity` | Create a CorrectFinancialData entity instance. |
| `CustomerMerchantCorrection` | `(data) -> CustomerMerchantCorrectionEntity` | Create a CustomerMerchantCorrection entity instance. |
| `Label` | `(data) -> LabelEntity` | Create a Label entity instance. |
| `ListLabel` | `(data) -> ListLabelEntity` | Create a ListLabel entity instance. |
| `ListTransaction` | `(data) -> ListTransactionEntity` | Create a ListTransaction entity instance. |
| `ManageFinancialData` | `(data) -> ManageFinancialDataEntity` | Create a ManageFinancialData entity instance. |
| `ManageTransactionLabel` | `(data) -> ManageTransactionLabelEntity` | Create a ManageTransactionLabel entity instance. |
| `RetrieveFinancialData` | `(data) -> RetrieveFinancialDataEntity` | Create a RetrieveFinancialData entity instance. |
| `Similar` | `(data) -> SimilarEntity` | Create a Similar entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `create` | `(reqdata, ctrl) -> any, err` | Create a new entity. |
| `update` | `(reqdata, ctrl) -> any, err` | Update an existing entity. |
| `remove` | `(reqmatch, ctrl) -> any, err` | Remove an entity. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` / `create` / `update` / `remove` | the entity record (a `table`) |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local correct_financial_data, err = client:CorrectFinancialData():load()
    if err then error(err) end
    -- correct_financial_data is the loaded record

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

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

Operations: Create, List, Load, Remove.

API path: `/corrections/v2/categories`

#### CustomerMerchantCorrection

| Field | Description |
| --- | --- |
| `data` |  |
| `metadata` | Metadata associated with the merchant correction response schema |
| `operation_id` |  |

Operations: Create.

API path: `/corrections/v2/merchants`

#### Label

| Field | Description |
| --- | --- |
| `created_at` | RFC3339 timestamp at which the label was created. |
| `id` | Unique identifier of the label. |
| `name` | Display name for the new label. |
| `updated_at` | RFC3339 timestamp at which the label was last updated. |

Operations: Create, Update.

API path: `/financial/v2/transactions/labels`

#### ListLabel

| Field | Description |
| --- | --- |
| `created_at` | RFC3339 timestamp at which the label was created. |
| `id` | Unique identifier of the label. |
| `name` | Display name of the label. |
| `updated_at` | RFC3339 timestamp at which the label was last updated. |

Operations: List.

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

Operations: List.

API path: `/financial/v2/transactions`

#### ManageFinancialData

| Field | Description |
| --- | --- |
| `label_id` | Identifier of the label to attach. |

Operations: Create, Remove.

API path: `/financial/v2/transactions/{transaction_id}/labels`

#### ManageTransactionLabel

| Field | Description |
| --- | --- |
| `id` |  |

Operations: Remove.

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

Operations: List, Load.

API path: `/financial/v3/accounts`

#### Similar

| Field | Description |
| --- | --- |
| `data` |  |
| `id` |  |
| `metadata` |  |
| `operation_id` |  |

Operations: Load.

API path: `/corrections/v2/categories/similar/{transaction_id}`



## Entities


### CorrectFinancialData

Create an instance: `local correct_financial_data = client:CorrectFinancialData(nil)`

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
| `data` | `table` | List of merchant that matched the query. |
| `frequency` | `string` | The frequency to assign. |
| `include_similar` | `boolean` | Apply the correction to the whole group of transactions similar to the specified transaction, rather than just the transaction itself |
| `logo_feedback` | `string` | The type of feedback for the merchant. |
| `metadata` | `table` | Metadata associated with the response schema |
| `name` | `string` | The display name of the custom merchant to be created |
| `online_or_billing_only` | `boolean` | This is used to indicate if the custom merchant being created is an exclusively online or billing merchant. |
| `operation_id` | `string` | A unique identifier/reference associated with a given endpoint/operation |
| `reference_transaction_id` | `string` | The transaction whose group the specified transactions should join or leave. |
| `rule_definition` | `table` | The definition of the rule, mirroring the request body that created it. |
| `rule_type` | `string` | The type of correction rule |
| `similar` | `boolean` | Whether the specified transactions are similar to each other (or to `reference_transaction_id`, if supplied) |
| `suggested_logo` | `string` |  |
| `suggested_url` | `string` | A suggested URL for the merchant, to help bud identify it and expand our merchant database. |
| `transaction_id` | `string` | The unique identifier for the transaction |
| `transaction_ids` | `table` | The transactions the rule applies to |

#### Example: Load

```lua
local correct_financial_data, err = client:CorrectFinancialData():load({ rule_id = "rule_id" })
```

#### Example: List

```lua
local correct_financial_datas, err = client:CorrectFinancialData():list()
```

#### Example: Create

```lua
local correct_financial_data, err = client:CorrectFinancialData():create({
  merchant_id = "example_merchant_id", -- string
  created_at = "example_created_at", -- string
  custom_merchant_id = "example_custom_merchant_id", -- string
  data = {}, -- table
  frequency = "example_frequency", -- string
  logo_feedback = "example_logo_feedback", -- string
  metadata = {}, -- table
  name = "example_name", -- string
  operation_id = "example_operation_id", -- string
  rule_definition = {}, -- table
  rule_type = "example_rule_type", -- string
  transaction_id = "example_transaction_id", -- string
  transaction_ids = {}, -- table
})
```


### CustomerMerchantCorrection

Create an instance: `local customer_merchant_correction = client:CustomerMerchantCorrection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `table` |  |
| `metadata` | `table` | Metadata associated with the merchant correction response schema |
| `operation_id` | `string` |  |

#### Example: Create

```lua
local customer_merchant_correction, err = client:CustomerMerchantCorrection():create({
  data = {}, -- table
  metadata = {}, -- table
  operation_id = "example_operation_id", -- string
})
```


### Label

Create an instance: `local label = client:Label(nil)`

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

```lua
local label, err = client:Label():create({
  created_at = "example_created_at", -- string
  id = "example_id", -- string
  name = "example_name", -- string
  updated_at = "example_updated_at", -- string
})
```


### ListLabel

Create an instance: `local list_label = client:ListLabel(nil)`

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

```lua
local list_labels, err = client:ListLabel():list()
```


### ListTransaction

Create an instance: `local list_transaction = client:ListTransaction(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` | Identifier for the account associated with the transaction. |
| `amount` | `table` | The monetary amount. |
| `client_attributes` | `table` | An optional map of key-value string pairs that allows clients to attach arbitrary metadata to the transaction. |
| `counterparty` | `table` | An object containing details of the counterparty in this transaction |
| `credit_debit_indicator` | `string` | Credit/Debit Indicator |
| `date_time` | `string` | Date that the transaction occured compliant with RFC3339. |
| `description` | `string` | Description of the transaction. |
| `enrichments` | `table` | Contextual enrichments associated with a Transaction |
| `labels` | `table` | Customer-defined labels currently attached to this transaction. |
| `merchant_category_code` | `string` | Merchant category code conforming to ISO 18245, related to the type of services or goods the merchant provides for the transaction. |
| `posted_date_time` | `string` | Date the assets involved in the transaction transferred compliant with RFC3339. |
| `provider` | `string` | Name of the transaction source provider. |
| `running_balance` | `table` | The running balance for the account that the transaction takes place against |
| `running_balance_credit_debit_indicator` | `string` | Credit/Debit Indicator for the running balance field |
| `status` | `string` | Status of the transaction. |
| `suggested_description` | `string` | The description Bud suggests client apps show for a given transaction. |
| `suggested_logo` | `string` | The logo Bud suggests client apps show for a given transaction. |
| `tags` | `table` | A list of potential tags associated with the transaction after contextual enrichment, which can be used for filtering. |
| `transaction_id` | `string` | Unique identifier for the transaction. |
| `transaction_type` | `table` | The code and a description of the transaction type. |
| `value_date_time` | `string` | Date the assets involved in the transaction transferred compliant with RFC3339. |

#### Example: List

```lua
local list_transactions, err = client:ListTransaction():list()
```


### ManageFinancialData

Create an instance: `local manage_financial_data = client:ManageFinancialData(nil)`

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

```lua
local manage_financial_data, err = client:ManageFinancialData():create({
  transaction_id = "example_transaction_id", -- string
  label_id = "example_label_id", -- string
})
```


### ManageTransactionLabel

Create an instance: `local manage_transaction_label = client:ManageTransactionLabel(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |


### RetrieveFinancialData

Create an instance: `local retrieve_financial_data = client:RetrieveFinancialData(nil)`

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
| `balances` | `table` | The balances ingested for the account. |
| `booked` | `any` | The current balance of the account |
| `closed_at` | `string` | The datetime the account was closed in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `credit_lines` | `table` | The latest credit limit available for the account. |
| `currency` | `string` | The three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) of the monetary amount. |
| `data` | `table` |  |
| `date` | `string` | The date of the given balance in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `details` | `table` | Provides more details for the authorised payment. |
| `first_transaction_date` | `string` | The date of the first transaction that bud has stored for this account. |
| `frequency` | `string` | The frequency of the authorised payment. |
| `holder` | `table` | The account holder. |
| `holders` | `table` | The account holder(s). |
| `id` | `string` | The id for the authorised payment. |
| `identifiers` | `table` | These are wellknown fields which uniquely identify the account. |
| `last_transaction_date` | `string` | The date of the last transaction that bud has stored for this account. |
| `metadata` | `table` |  |
| `name` | `string` | The name of the payment. |
| `opening_date_time` | `string` | The datetime the account was opened in the format [RFC 3339] |
| `operation_id` | `string` |  |
| `pending` | `any` | The balance of the account if all pending transactions have settled. |
| `provider` | `string` | The account's provider, this is usually the bank or building society the account is held with. |
| `provider_display_name` | `string` | The display name for the account's provider. |
| `provider_logo` | `string` | A link to an image for the accounts provider's logo. |
| `reference` | `string` | The reference associated with the authorised payment. |
| `restriction` | `string` | The restriction on the account, if any. |
| `status` | `string` | The status of the account. |
| `suggested_name` | `string` | The name Bud suggests client apps show for the account. |
| `transaction_windows` | `table` | The transaction windows indicate for which periods we have full coverage of transactions ingested for the account. |
| `type` | `string` | The type of the authorised payment. |
| `usage_type` | `string` | The intended usage of the account. |

#### Example: Load

```lua
local retrieve_financial_data, err = client:RetrieveFinancialData():load({ account_id = "account_id" })
```

#### Example: List

```lua
local retrieve_financial_datas, err = client:RetrieveFinancialData():list()
```


### Similar

Create an instance: `local similar = client:Similar(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `table` |  |
| `id` | `string` |  |
| `metadata` | `table` |  |
| `operation_id` | `string` |  |

#### Example: Load

```lua
local similar, err = client:Similar():load({ id = "similar_id" })
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

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

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

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── bud-financial-data_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── schema.lua               -- Generated option + entity specs
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`bud-financial-data_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```lua
local listtransaction = client:ListTransaction()
listtransaction:list()

-- listtransaction:data_get() now returns the listtransaction data from the last list
-- listtransaction:match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
