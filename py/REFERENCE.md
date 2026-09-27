# BudFinancialData Python SDK Reference

Complete API reference for the BudFinancialData Python SDK.


## BudFinancialDataSDK

### Constructor

```python
from budfinancialdata_sdk import BudFinancialDataSDK

client = BudFinancialDataSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `BudFinancialDataSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = BudFinancialDataSDK.test()
```


### Instance Methods

#### `CorrectFinancialData(data=None)`

Create a new `CorrectFinancialDataEntity` instance. Pass `None` for no initial data.

#### `CustomerMerchantCorrection(data=None)`

Create a new `CustomerMerchantCorrectionEntity` instance. Pass `None` for no initial data.

#### `Label(data=None)`

Create a new `LabelEntity` instance. Pass `None` for no initial data.

#### `ListLabel(data=None)`

Create a new `ListLabelEntity` instance. Pass `None` for no initial data.

#### `ListTransaction(data=None)`

Create a new `ListTransactionEntity` instance. Pass `None` for no initial data.

#### `ManageFinancialData(data=None)`

Create a new `ManageFinancialDataEntity` instance. Pass `None` for no initial data.

#### `ManageTransactionLabel(data=None)`

Create a new `ManageTransactionLabelEntity` instance. Pass `None` for no initial data.

#### `RetrieveFinancialData(data=None)`

Create a new `RetrieveFinancialDataEntity` instance. Pass `None` for no initial data.

#### `Similar(data=None)`

Create a new `SimilarEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## CorrectFinancialDataEntity

```python
correct_financial_data = client.CorrectFinancialData()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | Yes | Date that the custom merchant was created, compliant with RFC3339. |
| `custom_merchant_id` | `str` | Yes | UUID representing the custom merchant. |
| `data` | `list` | Yes | List of merchant that matched the query. |
| `frequency` | `str` | Yes | The frequency to assign. |
| `include_similar` | `bool` | No | Apply the correction to the whole group of transactions similar to the specified transaction, rather than just the transaction itself |
| `logo_feedback` | `str` | Yes | The type of feedback for the merchant. |
| `metadata` | `dict` | Yes | Metadata associated with the response schema |
| `name` | `str` | Yes | The display name of the custom merchant to be created |
| `online_or_billing_only` | `bool` | No | This is used to indicate if the custom merchant being created is an exclusively online or billing merchant. |
| `operation_id` | `str` | Yes | A unique identifier/reference associated with a given endpoint/operation |
| `reference_transaction_id` | `str` | No | The transaction whose group the specified transactions should join or leave. |
| `rule_definition` | `dict` | Yes | The definition of the rule, mirroring the request body that created it. |
| `rule_type` | `str` | Yes | The type of correction rule |
| `similar` | `bool` | No | Whether the specified transactions are similar to each other (or to `reference_transaction_id`, if supplied) |
| `suggested_logo` | `str` | No |  |
| `suggested_url` | `str` | No | A suggested URL for the merchant, to help bud identify it and expand our merchant database. |
| `transaction_id` | `str` | Yes | The unique identifier for the transaction |
| `transaction_ids` | `list` | Yes | The transactions the rule applies to |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CorrectFinancialData().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.CorrectFinancialData().list()
for correct_financial_data in results:
    print(correct_financial_data)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.CorrectFinancialData().load({"rule_id": "rule_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.CorrectFinancialData().remove({"rule_id": "rule_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CorrectFinancialDataEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CustomerMerchantCorrectionEntity

```python
customer_merchant_correction = client.CustomerMerchantCorrection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `list` | Yes |  |
| `metadata` | `dict` | Yes | Metadata associated with the merchant correction response schema |
| `operation_id` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CustomerMerchantCorrection().create({
    "data": [],  # list
    "metadata": {},  # dict
    "operation_id": "example_operation_id",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CustomerMerchantCorrectionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## LabelEntity

```python
label = client.Label()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | Yes | RFC3339 timestamp at which the label was created. |
| `id` | `str` | Yes | Unique identifier of the label. |
| `name` | `str` | Yes | Display name for the new label. |
| `updated_at` | `str` | Yes | RFC3339 timestamp at which the label was last updated. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Label().create({
    "created_at": "example_created_at",  # str
    "id": "example_id",  # str
    "name": "example_name",  # str
    "updated_at": "example_updated_at",  # str
})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Label().update({
    "id": "id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `LabelEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ListLabelEntity

```python
list_label = client.ListLabel()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | Yes | RFC3339 timestamp at which the label was created. |
| `id` | `str` | Yes | Unique identifier of the label. |
| `name` | `str` | Yes | Display name of the label. |
| `updated_at` | `str` | Yes | RFC3339 timestamp at which the label was last updated. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.ListLabel().list()
for list_label in results:
    print(list_label)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListLabelEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ListTransactionEntity

```python
list_transaction = client.ListTransaction()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `str` | Yes | Identifier for the account associated with the transaction. |
| `amount` | `dict` | Yes | The monetary amount. |
| `client_attributes` | `dict` | No | An optional map of key-value string pairs that allows clients to attach arbitrary metadata to the transaction. |
| `counterparty` | `dict` | No | An object containing details of the counterparty in this transaction |
| `credit_debit_indicator` | `str` | Yes | Credit/Debit Indicator |
| `date_time` | `str` | Yes | Date that the transaction occured compliant with RFC3339. |
| `description` | `str` | Yes | Description of the transaction. |
| `enrichments` | `dict` | No | Contextual enrichments associated with a Transaction |
| `labels` | `list` | No | Customer-defined labels currently attached to this transaction. |
| `merchant_category_code` | `str` | No | Merchant category code conforming to ISO 18245, related to the type of services or goods the merchant provides for the transaction. |
| `posted_date_time` | `str` | No | Date the assets involved in the transaction transferred compliant with RFC3339. |
| `provider` | `str` | No | Name of the transaction source provider. |
| `running_balance` | `dict` | Yes | The running balance for the account that the transaction takes place against |
| `running_balance_credit_debit_indicator` | `str` | No | Credit/Debit Indicator for the running balance field |
| `status` | `str` | Yes | Status of the transaction. |
| `suggested_description` | `str` | Yes | The description Bud suggests client apps show for a given transaction. |
| `suggested_logo` | `str` | No | The logo Bud suggests client apps show for a given transaction. |
| `tags` | `list` | No | A list of potential tags associated with the transaction after contextual enrichment, which can be used for filtering. |
| `transaction_id` | `str` | Yes | Unique identifier for the transaction. |
| `transaction_type` | `dict` | No | The code and a description of the transaction type. |
| `value_date_time` | `str` | No | Date the assets involved in the transaction transferred compliant with RFC3339. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.ListTransaction().list()
for list_transaction in results:
    print(list_transaction)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListTransactionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ManageFinancialDataEntity

```python
manage_financial_data = client.ManageFinancialData()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `label_id` | `str` | Yes | Identifier of the label to attach. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ManageFinancialData().create({
    "transaction_id": "example_transaction_id",  # str
    "label_id": "example_label_id",  # str
})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.ManageFinancialData().remove({"provider": "provider"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ManageFinancialDataEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ManageTransactionLabelEntity

```python
manage_transaction_label = client.ManageTransactionLabel()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No |  |

### Operations

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.ManageTransactionLabel().remove({"id": "id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ManageTransactionLabelEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RetrieveFinancialDataEntity

```python
retrieve_financial_data = client.RetrieveFinancialData()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_category` | `str` | No | The account category that the account type is a part of Currently supported values include (but are not necessarily limited to): - `credit` - `depository` - `insurance` - `investment` - `loan` - `savings` - `other` |
| `account_id` | `str` | Yes | A unique id associated with the account. |
| `account_name` | `str` | No | The name associated with the account, this is often a friendly name assigned to the account to make it easier to refer to. |
| `account_type` | `str` | No | The type of account. |
| `balances` | `dict` | No | The balances ingested for the account. |
| `booked` | `Any` | Yes | The current balance of the account |
| `closed_at` | `str` | No | The datetime the account was closed in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `credit_lines` | `dict` | No | The latest credit limit available for the account. |
| `currency` | `str` | Yes | The three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) of the monetary amount. |
| `data` | `dict` | Yes |  |
| `date` | `str` | Yes | The date of the given balance in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339). |
| `details` | `list` | No | Provides more details for the authorised payment. |
| `first_transaction_date` | `str` | Yes | The date of the first transaction that bud has stored for this account. |
| `frequency` | `str` | No | The frequency of the authorised payment. |
| `holder` | `dict` | No | The account holder. |
| `holders` | `list` | No | The account holder(s). |
| `id` | `str` | Yes | The id for the authorised payment. |
| `identifiers` | `dict` | No | These are wellknown fields which uniquely identify the account. |
| `last_transaction_date` | `str` | Yes | The date of the last transaction that bud has stored for this account. |
| `metadata` | `dict` | Yes |  |
| `name` | `str` | Yes | The name of the payment. |
| `opening_date_time` | `str` | No | The datetime the account was opened in the format [RFC 3339] |
| `operation_id` | `str` | Yes |  |
| `pending` | `Any` | Yes | The balance of the account if all pending transactions have settled. |
| `provider` | `str` | No | The account's provider, this is usually the bank or building society the account is held with. |
| `provider_display_name` | `str` | No | The display name for the account's provider. |
| `provider_logo` | `str` | No | A link to an image for the accounts provider's logo. |
| `reference` | `str` | Yes | The reference associated with the authorised payment. |
| `restriction` | `str` | No | The restriction on the account, if any. |
| `status` | `str` | No | The status of the account. |
| `suggested_name` | `str` | Yes | The name Bud suggests client apps show for the account. |
| `transaction_windows` | `list` | No | The transaction windows indicate for which periods we have full coverage of transactions ingested for the account. |
| `type` | `str` | Yes | The type of the authorised payment. |
| `usage_type` | `str` | No | The intended usage of the account. |

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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.RetrieveFinancialData().list()
for retrieve_financial_data in results:
    print(retrieve_financial_data)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.RetrieveFinancialData().load({"account_id": "account_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RetrieveFinancialDataEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SimilarEntity

```python
similar = client.Similar()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `list` | Yes |  |
| `id` | `str` | No |  |
| `metadata` | `dict` | Yes |  |
| `operation_id` | `str` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Similar().load({"id": "similar_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SimilarEntity` instance with the same options.

#### `get_name() -> str`

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

```python
client = BudFinancialDataSDK({
    "feature": {
        "debug": {"active": True},
        "idempotency": {"active": True},
        "metrics": {"active": True},
        "paging": {"active": True},
        "ratelimit": {"active": True},
        "retry": {"active": True},
        "test": {"active": True},
        "timeout": {"active": True},
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

