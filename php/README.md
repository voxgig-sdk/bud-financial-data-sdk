# BudFinancialData PHP SDK



The PHP SDK for the BudFinancialData API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->CorrectFinancialData()` — with named operations (`list`/`load`/`create`/`update`/`remove`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/bud-financial-data-sdk/releases](https://github.com/voxgig-sdk/bud-financial-data-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'budfinancialdata_sdk.php';

$client = new BudFinancialDataSDK([
    "apikey" => getenv("BUD_FINANCIAL_DATA_APIKEY"),
]);
```

### 2. List correctfinancialdata records

```php
try {
    // list() returns entity instances; data_get() reads each record.
    $correctfinancialdatas = $client->CorrectFinancialData()->list();
    foreach ($correctfinancialdatas as $record) {
        $item = $record->data_get();
        echo $item["created_at"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 3. Load a correctfinancialdata

```php
try {
    // load() returns the ENTITY — call data_get() for the CorrectFinancialData record (throws on error).
    $correctfinancialdata = $client->CorrectFinancialData()->load(["rule_id" => "example_rule_id"]);
    print_r($correctfinancialdata->data_get());
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 4. Create, update, and remove

```php
// create() returns the ENTITY — call data_get() for the created CorrectFinancialData record.
$created = $client->CorrectFinancialData()->create(["merchant_id" => "example_merchant_id", "created_at" => "example_created_at", "custom_merchant_id" => "example_custom_merchant_id", "data" => [], "frequency" => "example_frequency", "logo_feedback" => "example_logo_feedback", "metadata" => [], "name" => "example_name", "operation_id" => "example_operation_id", "rule_definition" => [], "rule_type" => "example_rule_type", "transaction_id" => "example_transaction_id", "transaction_ids" => []]);

// Remove
$client->CorrectFinancialData()->remove(["rule_id" => "example_rule_id"]);
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $listtransactions = $client->ListTransaction()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```php
$client = BudFinancialDataSDK::test([
    "entity" => ["similar" => ["test01" => ["id" => "test01"]]],
]);

// Entity ops return the ENTITY (throws on error);
// call data_get() for the mock record.
$similar = $client->Similar()->load(["id" => "test01"]);
print_r($similar->data_get());
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new BudFinancialDataSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
BUD_FINANCIAL_DATA_TEST_LIVE=TRUE
BUD_FINANCIAL_DATA_APIKEY=<your-key>
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### BudFinancialDataSDK

```php
require_once 'budfinancialdata_sdk.php';
$client = new BudFinancialDataSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = BudFinancialDataSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### BudFinancialDataSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `CorrectFinancialData` | `($data): CorrectFinancialDataEntity` | Create a CorrectFinancialData entity instance. |
| `CustomerMerchantCorrection` | `($data): CustomerMerchantCorrectionEntity` | Create a CustomerMerchantCorrection entity instance. |
| `Label` | `($data): LabelEntity` | Create a Label entity instance. |
| `ListLabel` | `($data): ListLabelEntity` | Create a ListLabel entity instance. |
| `ListTransaction` | `($data): ListTransactionEntity` | Create a ListTransaction entity instance. |
| `ManageFinancialData` | `($data): ManageFinancialDataEntity` | Create a ManageFinancialData entity instance. |
| `ManageTransactionLabel` | `($data): ManageTransactionLabelEntity` | Create a ManageTransactionLabel entity instance. |
| `RetrieveFinancialData` | `($data): RetrieveFinancialDataEntity` | Create a RetrieveFinancialData entity instance. |
| `Similar` | `($data): SimilarEntity` | Create a Similar entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `update` | `($reqdata, $ctrl): array` | Update an existing entity. |
| `remove` | `($reqmatch, $ctrl): array` | Remove an entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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

Create an instance: `$correct_financial_data = $client->CorrectFinancialData();`

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
| `data` | `array` | List of merchant that matched the query. |
| `frequency` | `string` | The frequency to assign. |
| `include_similar` | `bool` | Apply the correction to the whole group of transactions similar to the specified transaction, rather than just the transaction itself |
| `logo_feedback` | `string` | The type of feedback for the merchant. |
| `metadata` | `array` | Metadata associated with the response schema |
| `name` | `string` | The display name of the custom merchant to be created |
| `online_or_billing_only` | `bool` | This is used to indicate if the custom merchant being created is an exclusively online or billing merchant. |
| `operation_id` | `string` | A unique identifier/reference associated with a given endpoint/operation |
| `reference_transaction_id` | `string` | The transaction whose group the specified transactions should join or leave. |
| `rule_definition` | `array` | The definition of the rule, mirroring the request body that created it. |
| `rule_type` | `string` | The type of correction rule |
| `similar` | `bool` | Whether the specified transactions are similar to each other (or to `reference_transaction_id`, if supplied) |
| `suggested_logo` | `string` |  |
| `suggested_url` | `string` | A suggested URL for the merchant, to help bud identify it and expand our merchant database. |
| `transaction_id` | `string` | The unique identifier for the transaction |
| `transaction_ids` | `array` | The transactions the rule applies to |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the CorrectFinancialData record (throws on error).
$correct_financial_data = $client->CorrectFinancialData()->load(["rule_id" => "rule_id"]);
```

#### Example: List

```php
// list() returns an array of CorrectFinancialData records (throws on error).
$correct_financial_datas = $client->CorrectFinancialData()->list();
```

#### Example: Create

```php
$correct_financial_data = $client->CorrectFinancialData()->create([
    "merchant_id" => null, // string
    "created_at" => null, // string
    "custom_merchant_id" => null, // string
    "data" => null, // array
    "frequency" => null, // string
    "logo_feedback" => null, // string
    "metadata" => null, // array
    "name" => null, // string
    "operation_id" => null, // string
    "rule_definition" => null, // array
    "rule_type" => null, // string
    "transaction_id" => null, // string
    "transaction_ids" => null, // array
]);
```


### CustomerMerchantCorrection

Create an instance: `$customer_merchant_correction = $client->CustomerMerchantCorrection();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `array` |  |
| `metadata` | `array` | Metadata associated with the merchant correction response schema |
| `operation_id` | `string` |  |

#### Example: Create

```php
$customer_merchant_correction = $client->CustomerMerchantCorrection()->create([
    "data" => null, // array
    "metadata" => null, // array
    "operation_id" => null, // string
]);
```


### Label

Create an instance: `$label = $client->Label();`

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

```php
$label = $client->Label()->create([
    "created_at" => null, // string
    "id" => null, // string
    "name" => null, // string
    "updated_at" => null, // string
]);
```


### ListLabel

Create an instance: `$list_label = $client->ListLabel();`

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

```php
// list() returns an array of ListLabel records (throws on error).
$list_labels = $client->ListLabel()->list();
```


### ListTransaction

Create an instance: `$list_transaction = $client->ListTransaction();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` | Identifier for the account associated with the transaction. |
| `amount` | `array` | The monetary amount. |
| `client_attributes` | `array` | An optional map of key-value string pairs that allows clients to attach arbitrary metadata to the transaction. |
| `counterparty` | `array` | An object containing details of the counterparty in this transaction |
| `credit_debit_indicator` | `string` | Credit/Debit Indicator |
| `date_time` | `string` | Date that the transaction occured compliant with RFC3339. |
| `description` | `string` | Description of the transaction. |
| `enrichments` | `array` | Contextual enrichments associated with a Transaction |
| `labels` | `array` | Customer-defined labels currently attached to this transaction. |
| `merchant_category_code` | `string` | Merchant category code conforming to ISO 18245, related to the type of services or goods the merchant provides for the transaction. |
| `posted_date_time` | `string` | Date the assets involved in the transaction transferred compliant with RFC3339. |
| `provider` | `string` | Name of the transaction source provider. |
| `running_balance` | `array` | The running balance for the account that the transaction takes place against |
| `running_balance_credit_debit_indicator` | `string` | Credit/Debit Indicator for the running balance field |
| `status` | `string` | Status of the transaction. |
| `suggested_description` | `string` | The description Bud suggests client apps show for a given transaction. |
| `suggested_logo` | `string` | The logo Bud suggests client apps show for a given transaction. |
| `tags` | `array` | A list of potential tags associated with the transaction after contextual enrichment, which can be used for filtering. |
| `transaction_id` | `string` | Unique identifier for the transaction. |
| `transaction_type` | `array` | The code and a description of the transaction type. |
| `value_date_time` | `string` | Date the assets involved in the transaction transferred compliant with RFC3339. |

#### Example: List

```php
// list() returns an array of ListTransaction records (throws on error).
$list_transactions = $client->ListTransaction()->list();
```


### ManageFinancialData

Create an instance: `$manage_financial_data = $client->ManageFinancialData();`

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

```php
$manage_financial_data = $client->ManageFinancialData()->create([
    "transaction_id" => null, // string
    "label_id" => null, // string
]);
```


### ManageTransactionLabel

Create an instance: `$manage_transaction_label = $client->ManageTransactionLabel();`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |


### RetrieveFinancialData

Create an instance: `$retrieve_financial_data = $client->RetrieveFinancialData();`

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
| `balances` | `array` | The balances ingested for the account. |
| `booked` | `mixed` | The current balance of the account |
| `closed_at` | `string` | The datetime the account was closed in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `credit_lines` | `array` | The latest credit limit available for the account. |
| `currency` | `string` | The three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) of the monetary amount. |
| `data` | `array` |  |
| `date` | `string` | The date of the given balance in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `details` | `array` | Provides more details for the authorised payment. |
| `first_transaction_date` | `string` | The date of the first transaction that bud has stored for this account. |
| `frequency` | `string` | The frequency of the authorised payment. |
| `holder` | `array` | The account holder. |
| `holders` | `array` | The account holder(s). |
| `id` | `string` | The id for the authorised payment. |
| `identifiers` | `array` | These are wellknown fields which uniquely identify the account. |
| `last_transaction_date` | `string` | The date of the last transaction that bud has stored for this account. |
| `metadata` | `array` |  |
| `name` | `string` | The name of the payment. |
| `opening_date_time` | `string` | The datetime the account was opened in the format [RFC 3339] |
| `operation_id` | `string` |  |
| `pending` | `mixed` | The balance of the account if all pending transactions have settled. |
| `provider` | `string` | The account's provider, this is usually the bank or building society the account is held with. |
| `provider_display_name` | `string` | The display name for the account's provider. |
| `provider_logo` | `string` | A link to an image for the accounts provider's logo. |
| `reference` | `string` | The reference associated with the authorised payment. |
| `restriction` | `string` | The restriction on the account, if any. |
| `status` | `string` | The status of the account. |
| `suggested_name` | `string` | The name Bud suggests client apps show for the account. |
| `transaction_windows` | `array` | The transaction windows indicate for which periods we have full coverage of transactions ingested for the account. |
| `type` | `string` | The type of the authorised payment. |
| `usage_type` | `string` | The intended usage of the account. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the RetrieveFinancialData record (throws on error).
$retrieve_financial_data = $client->RetrieveFinancialData()->load(["account_id" => "account_id"]);
```

#### Example: List

```php
// list() returns an array of RetrieveFinancialData records (throws on error).
$retrieve_financial_datas = $client->RetrieveFinancialData()->list();
```


### Similar

Create an instance: `$similar = $client->Similar();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `array` |  |
| `id` | `string` |  |
| `metadata` | `array` |  |
| `operation_id` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Similar record (throws on error).
$similar = $client->Similar()->load(["id" => "similar_id"]);
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

Features are the extension mechanism. A feature is a PHP class
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

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── budfinancialdata_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── schema.php                     -- Generated option + entity specs
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`budfinancialdata_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$listtransaction = $client->ListTransaction();
$listtransaction->list();

// $listtransaction->data_get() now returns the listtransaction data from the last list
// $listtransaction->match_get() returns the last match criteria
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
