# BudFinancialData Python SDK



The Python SDK for the BudFinancialData API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.CorrectFinancialData()` — each
carrying a small, uniform set of operations (`list`, `load`, `create`, `update`, `remove`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/bud-financial-data-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from budfinancialdata_sdk import BudFinancialDataSDK

client = BudFinancialDataSDK({
    "apikey": os.environ.get("BUD_FINANCIAL_DATA_APIKEY"),
})
```

### 2. List correctfinancialdata records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    correctfinancialdatas = client.CorrectFinancialData().list()
    for correctfinancialdata in correctfinancialdatas:
        print(correctfinancialdata)
except Exception as err:
    print(f"list failed: {err}")
```

### 3. Load a correctfinancialdata

`load()` returns the ENTITY — call data_get() for the record — and raises on error.

```python
try:
    correctfinancialdata = client.CorrectFinancialData().load({"rule_id": "example_rule_id"})
    print(correctfinancialdata)
except Exception as err:
    print(f"load failed: {err}")
```

### 4. Create, update, and remove

```python
# Create — returns the ENTITY (call data_get() for the record)
created = client.CorrectFinancialData().create({"merchant_id": "example_merchant_id", "created_at": "example_created_at", "custom_merchant_id": "example_custom_merchant_id", "data": [], "frequency": "example_frequency", "logo_feedback": "example_logo_feedback", "metadata": {}, "name": "example_name", "operation_id": "example_operation_id", "rule_definition": {}, "rule_type": "example_rule_type", "transaction_id": "example_transaction_id", "transaction_ids": []})

# Remove
client.CorrectFinancialData().remove({"rule_id": "example_rule_id"})
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    listtransactions = client.ListTransaction().list()
    print(listtransactions)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = BudFinancialDataSDK.test()

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
listtransaction = client.ListTransaction().list()
# listtransaction contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = BudFinancialDataSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### BudFinancialDataSDK

```python
from budfinancialdata_sdk import BudFinancialDataSDK

client = BudFinancialDataSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = BudFinancialDataSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### BudFinancialDataSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
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
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `correct_financial_data = client.CorrectFinancialData()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | Date that the custom merchant was created, compliant with RFC3339. |
| `custom_merchant_id` | `str` | UUID representing the custom merchant. |
| `data` | `list` | List of merchant that matched the query. |
| `frequency` | `str` | The frequency to assign. |
| `include_similar` | `bool` | Apply the correction to the whole group of transactions similar to the specified transaction, rather than just the transaction itself |
| `logo_feedback` | `str` | The type of feedback for the merchant. |
| `metadata` | `dict` | Metadata associated with the response schema |
| `name` | `str` | The display name of the custom merchant to be created |
| `online_or_billing_only` | `bool` | This is used to indicate if the custom merchant being created is an exclusively online or billing merchant. |
| `operation_id` | `str` | A unique identifier/reference associated with a given endpoint/operation |
| `reference_transaction_id` | `str` | The transaction whose group the specified transactions should join or leave. |
| `rule_definition` | `dict` | The definition of the rule, mirroring the request body that created it. |
| `rule_type` | `str` | The type of correction rule |
| `similar` | `bool` | Whether the specified transactions are similar to each other (or to `reference_transaction_id`, if supplied) |
| `suggested_logo` | `str` |  |
| `suggested_url` | `str` | A suggested URL for the merchant, to help bud identify it and expand our merchant database. |
| `transaction_id` | `str` | The unique identifier for the transaction |
| `transaction_ids` | `list` | The transactions the rule applies to |

#### Example: Load

```python
correct_financial_data = client.CorrectFinancialData().load({"rule_id": "rule_id"})
```

#### Example: List

```python
correct_financial_datas = client.CorrectFinancialData().list()
```

#### Example: Create

```python
correct_financial_data = client.CorrectFinancialData().create({
    "merchant_id": "example_merchant_id",  # str
    "created_at": "example_created_at",  # str
    "custom_merchant_id": "example_custom_merchant_id",  # str
    "data": [],  # list
    "frequency": "example_frequency",  # str
    "logo_feedback": "example_logo_feedback",  # str
    "metadata": {},  # dict
    "name": "example_name",  # str
    "operation_id": "example_operation_id",  # str
    "rule_definition": {},  # dict
    "rule_type": "example_rule_type",  # str
    "transaction_id": "example_transaction_id",  # str
    "transaction_ids": [],  # list
})
```


### CustomerMerchantCorrection

Create an instance: `customer_merchant_correction = client.CustomerMerchantCorrection()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `list` |  |
| `metadata` | `dict` | Metadata associated with the merchant correction response schema |
| `operation_id` | `str` |  |

#### Example: Create

```python
customer_merchant_correction = client.CustomerMerchantCorrection().create({
    "data": [],  # list
    "metadata": {},  # dict
    "operation_id": "example_operation_id",  # str
})
```


### Label

Create an instance: `label = client.Label()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | RFC3339 timestamp at which the label was created. |
| `id` | `str` | Unique identifier of the label. |
| `name` | `str` | Display name for the new label. |
| `updated_at` | `str` | RFC3339 timestamp at which the label was last updated. |

#### Example: Create

```python
label = client.Label().create({
    "created_at": "example_created_at",  # str
    "id": "example_id",  # str
    "name": "example_name",  # str
    "updated_at": "example_updated_at",  # str
})
```


### ListLabel

Create an instance: `list_label = client.ListLabel()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | RFC3339 timestamp at which the label was created. |
| `id` | `str` | Unique identifier of the label. |
| `name` | `str` | Display name of the label. |
| `updated_at` | `str` | RFC3339 timestamp at which the label was last updated. |

#### Example: List

```python
list_labels = client.ListLabel().list()
```


### ListTransaction

Create an instance: `list_transaction = client.ListTransaction()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `str` | Identifier for the account associated with the transaction. |
| `amount` | `dict` | The monetary amount. |
| `client_attributes` | `dict` | An optional map of key-value string pairs that allows clients to attach arbitrary metadata to the transaction. |
| `counterparty` | `dict` | An object containing details of the counterparty in this transaction |
| `credit_debit_indicator` | `str` | Credit/Debit Indicator |
| `date_time` | `str` | Date that the transaction occured compliant with RFC3339. |
| `description` | `str` | Description of the transaction. |
| `enrichments` | `dict` | Contextual enrichments associated with a Transaction |
| `labels` | `list` | Customer-defined labels currently attached to this transaction. |
| `merchant_category_code` | `str` | Merchant category code conforming to ISO 18245, related to the type of services or goods the merchant provides for the transaction. |
| `posted_date_time` | `str` | Date the assets involved in the transaction transferred compliant with RFC3339. |
| `provider` | `str` | Name of the transaction source provider. |
| `running_balance` | `dict` | The running balance for the account that the transaction takes place against |
| `running_balance_credit_debit_indicator` | `str` | Credit/Debit Indicator for the running balance field |
| `status` | `str` | Status of the transaction. |
| `suggested_description` | `str` | The description Bud suggests client apps show for a given transaction. |
| `suggested_logo` | `str` | The logo Bud suggests client apps show for a given transaction. |
| `tags` | `list` | A list of potential tags associated with the transaction after contextual enrichment, which can be used for filtering. |
| `transaction_id` | `str` | Unique identifier for the transaction. |
| `transaction_type` | `dict` | The code and a description of the transaction type. |
| `value_date_time` | `str` | Date the assets involved in the transaction transferred compliant with RFC3339. |

#### Example: List

```python
list_transactions = client.ListTransaction().list()
```


### ManageFinancialData

Create an instance: `manage_financial_data = client.ManageFinancialData()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `label_id` | `str` | Identifier of the label to attach. |

#### Example: Create

```python
manage_financial_data = client.ManageFinancialData().create({
    "transaction_id": "example_transaction_id",  # str
    "label_id": "example_label_id",  # str
})
```


### ManageTransactionLabel

Create an instance: `manage_transaction_label = client.ManageTransactionLabel()`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `str` |  |


### RetrieveFinancialData

Create an instance: `retrieve_financial_data = client.RetrieveFinancialData()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_category` | `str` | The account category that the account type is a part of Currently supported values include (but are not necessarily limited to): - `credit` - `depository` - `insurance` - `investment` - `loan` - `savings` - `other` |
| `account_id` | `str` | A unique id associated with the account. |
| `account_name` | `str` | The name associated with the account, this is often a friendly name assigned to the account to make it easier to refer to. |
| `account_type` | `str` | The type of account. |
| `balances` | `dict` | The balances ingested for the account. |
| `booked` | `Any` | The current balance of the account |
| `closed_at` | `str` | The datetime the account was closed in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `credit_lines` | `dict` | The latest credit limit available for the account. |
| `currency` | `str` | The three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) of the monetary amount. |
| `data` | `dict` |  |
| `date` | `str` | The date of the given balance in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `details` | `list` | Provides more details for the authorised payment. |
| `first_transaction_date` | `str` | The date of the first transaction that bud has stored for this account. |
| `frequency` | `str` | The frequency of the authorised payment. |
| `holder` | `dict` | The account holder. |
| `holders` | `list` | The account holder(s). |
| `id` | `str` | The id for the authorised payment. |
| `identifiers` | `dict` | These are wellknown fields which uniquely identify the account. |
| `last_transaction_date` | `str` | The date of the last transaction that bud has stored for this account. |
| `metadata` | `dict` |  |
| `name` | `str` | The name of the payment. |
| `opening_date_time` | `str` | The datetime the account was opened in the format [RFC 3339] |
| `operation_id` | `str` |  |
| `pending` | `Any` | The balance of the account if all pending transactions have settled. |
| `provider` | `str` | The account's provider, this is usually the bank or building society the account is held with. |
| `provider_display_name` | `str` | The display name for the account's provider. |
| `provider_logo` | `str` | A link to an image for the accounts provider's logo. |
| `reference` | `str` | The reference associated with the authorised payment. |
| `restriction` | `str` | The restriction on the account, if any. |
| `status` | `str` | The status of the account. |
| `suggested_name` | `str` | The name Bud suggests client apps show for the account. |
| `transaction_windows` | `list` | The transaction windows indicate for which periods we have full coverage of transactions ingested for the account. |
| `type` | `str` | The type of the authorised payment. |
| `usage_type` | `str` | The intended usage of the account. |

#### Example: Load

```python
retrieve_financial_data = client.RetrieveFinancialData().load({"account_id": "account_id"})
```

#### Example: List

```python
retrieve_financial_datas = client.RetrieveFinancialData().list()
```


### Similar

Create an instance: `similar = client.Similar()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `list` |  |
| `id` | `str` |  |
| `metadata` | `dict` |  |
| `operation_id` | `str` |  |

#### Example: Load

```python
similar = client.Similar().load({"id": "similar_id"})
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

Features are the extension mechanism. A feature is a Python class
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

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── budfinancialdata_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── schema.py                    -- Generated option + entity specs
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`budfinancialdata_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
listtransaction = client.ListTransaction()
listtransaction.list()

# listtransaction.data_get() now returns the listtransaction data from the last list
# listtransaction.match_get() returns the last match criteria
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
