# Financial Data API

These endpoints handle the aggregation and retrieval of financial data. They are designed to be generic across a range of ingestion sources and regions.

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 9 entities and 32 HTTP routes. There are 6 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### [CorrectFinancialData](docs/api/correct_financial_data.html)

Results: The request was successfully processed; Created; Accepted; OK; The rule was successfully deleted.

SDK operations: `create`, `list`, `load`, `remove`.

Key fields to recognise:

- `created_at`: Date that the custom merchant was created, compliant with RFC3339.
- `custom_merchant_id`: UUID representing the custom merchant.
- `data`: List of merchant that matched the query. Best match first
- `frequency`: The frequency to assign.
- `include_similar`: Correct the category of other similar transactions to the specified transaction

### [CustomerMerchantCorrection](docs/api/customer_merchant_correction.html)

Results: The request was successfully processed.

SDK operations: `create`.

Key fields to recognise:

- `metadata`: Metadata associated with the merchant correction response schema
- `operation_id`: A unique identifier/reference associated with a given endpoint/operation

### [Label](docs/api/label.html)

Results: The label was created successfully.; The label was renamed successfully.

SDK operations: `create`, `update`.

Key fields to recognise:

- `created_at`: RFC3339 timestamp at which the label was created.
- `id`: Unique identifier of the label.
- `name`: Display name of the label.
- `updated_at`: RFC3339 timestamp at which the label was last updated.

### [ListLabel](docs/api/list_label.html)

Results: The request was successfully processed.

SDK operations: `list`.

Key fields to recognise:

- `created_at`: RFC3339 timestamp at which the label was created.
- `id`: Unique identifier of the label.
- `name`: Display name of the label.
- `updated_at`: RFC3339 timestamp at which the label was last updated.

### [ListTransaction](docs/api/list_transaction.html)

Results: The request was successfully processed.

SDK operations: `list`.

Key fields to recognise:

- `account_id`: Identifier for the account associated with the transaction.
- `amount`: The monetary amount.
- `client_attributes`: An optional map of key-value string pairs that allows clients to attach arbitrary metadata to the transaction. This field can be used to store any relevant information specific to the client&#39;s needs, such as: These attributes can be added using the [Ingest Transactions](#operation/v2_ingest_transactions_post) endpoint.
- `counterparty`: An object containing details of the counterparty in this transaction
- `credit_debit_indicator`: Credit/Debit Indicator

### [ManageFinancialData](docs/api/manage_financial_data.html)

Results: The label was attached to the transaction.; The label was removed from the transaction.; No Content - The account data has been successfully deleted. Note: This response is also given if the account did not exist.; The request was successfully processed.

SDK operations: `create`, `remove`.

Key fields to recognise:

- `label_id`: Identifier of the label to attach.

### [ManageTransactionLabel](docs/api/manage_transaction_label.html)

Results: The label was deleted successfully.

SDK operations: `remove`.

### [RetrieveFinancialData](docs/api/retrieve_financial_data.html)

Results: OK.

SDK operations: `list`, `load`.

Key fields to recognise:

- `account_category`: The account category that the account type is a part of Currently supported values include (but are not necessarily limited to): - `credit` - `depository` - `insurance` - `investment` - `loan` - `savings` - `other`
- `account_id`: A unique id associated with the account. Example: `sdb4eRMaad8XVEvdIoISOQ`
- `account_name`: The name associated with the account, this is often a friendly name assigned to the account to make it easier to refer to. Example: `Household Savings Account` Not present if empty.
- `account_type`: The type of account. Currently supported values include (but are not necessarily limited to): - `auto_loan` - `boat_loan` - `brokerage` - `business_loan` - `certificate_of_deposit` - `charge_card` - `checking_account` - `credit_card` - `current_account` - `debit_card` - `disability_insurance` - `e_money` - `health_insurance` - `home_equity_loan` - `insurance` - `investment` - `ira` - `liability_insurance` - `life_insurance` - `line_of_credit` - `loan` - `money_market` - `mortgage` - `other` - `personal_loan` - `pre_paid_card` - `property_insurance` - `roth` - `rv_loan` - `safe_deposit_box` - `savings` - `student_loan` - `travel_insurance` - `vehicle_insurance` - `wealth_account` Not present if empty.
- `balances`: The balances ingested for the account. &gt; 📘 Note &gt; &gt; Normalised balances are derived from ingested data, therefore we&#39;re unable to guarantee that either a `booked` or `pending` balance will always be present. Where possible we will look to return both.

### [Similar](docs/api/similar.html)

Results: The request was successfully processed.

SDK operations: `load`.

Key fields to recognise:

- `id`: UUID corresponding to the merchant, if available. Subject to version control.
- `operation_id`: A unique identifier/reference associated with a given endpoint/operation

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| [CorrectFinancialData](docs/api/correct_financial_data.html) | `create` | `POST /corrections/v2/categories` | Required |
| [CorrectFinancialData](docs/api/correct_financial_data.html) | `create` | `POST /corrections/v2/custom-merchants` | Required |
| [CorrectFinancialData](docs/api/correct_financial_data.html) | `create` | `POST /corrections/v3/regularity` | Required |
| [CorrectFinancialData](docs/api/correct_financial_data.html) | `create` | `POST /corrections/v3/regularity/end` | Required |
| [CorrectFinancialData](docs/api/correct_financial_data.html) | `create` | `POST /corrections/v3/similar` | Required |
| [CorrectFinancialData](docs/api/correct_financial_data.html) | `create` | `POST /corrections/v2/merchant-feedback/merchant/{merchant_id}` | Required |
| [CorrectFinancialData](docs/api/correct_financial_data.html) | `list` | `GET /corrections/v2/custom-merchants` | Required |
| [CorrectFinancialData](docs/api/correct_financial_data.html) | `load` | `GET /corrections/v2/merchants/search/{merchant_query}` | Required |
| [CorrectFinancialData](docs/api/correct_financial_data.html) | `load` | `GET /corrections/v3/rules/{rule_id}` | Required |
| [CorrectFinancialData](docs/api/correct_financial_data.html) | `remove` | `DELETE /corrections/v3/rules/{rule_id}` | Required |
| [CustomerMerchantCorrection](docs/api/customer_merchant_correction.html) | `create` | `POST /corrections/v2/merchants` | Required |
| [Label](docs/api/label.html) | `create` | `POST /financial/v2/transactions/labels` | Required |
| [Label](docs/api/label.html) | `update` | `PATCH /financial/v2/transactions/labels/{label_id}` | Required |
| [ListLabel](docs/api/list_label.html) | `list` | `GET /financial/v2/transactions/labels` | Required |
| [ListTransaction](docs/api/list_transaction.html) | `list` | `GET /financial/v2/transactions` | Required |
| [ManageFinancialData](docs/api/manage_financial_data.html) | `create` | `POST /financial/v2/transactions/{transaction_id}/labels` | Required |
| [ManageFinancialData](docs/api/manage_financial_data.html) | `remove` | `DELETE /financial/v2/transactions/{transaction_id}/labels/{label_id}` | Required |
| [ManageFinancialData](docs/api/manage_financial_data.html) | `remove` | `DELETE /financial/v2/accounts/{account_id}` | Required |
| [ManageFinancialData](docs/api/manage_financial_data.html) | `remove` | `DELETE /v1/provider/{provider}` | Required |
| [ManageTransactionLabel](docs/api/manage_transaction_label.html) | `remove` | `DELETE /financial/v2/transactions/labels/{label_id}` | Required |
| [RetrieveFinancialData](docs/api/retrieve_financial_data.html) | `list` | `GET /financial/v3/accounts` | Required |
| [RetrieveFinancialData](docs/api/retrieve_financial_data.html) | `list` | `GET /financial/v2/accounts` | Required |
| [RetrieveFinancialData](docs/api/retrieve_financial_data.html) | `list` | `GET /financial/v2/accounts/{account_id}/balances` | Required |
| [RetrieveFinancialData](docs/api/retrieve_financial_data.html) | `list` | `GET /financial/v3/accounts/{account_id}/balances` | Required |
| [RetrieveFinancialData](docs/api/retrieve_financial_data.html) | `list` | `GET /financial/v2/authorised-payments` | Required |
| [RetrieveFinancialData](docs/api/retrieve_financial_data.html) | `list` | `GET /financial/v2/balances` | Required |
| [RetrieveFinancialData](docs/api/retrieve_financial_data.html) | `list` | `GET /financial/v3/balances` | Required |
| [RetrieveFinancialData](docs/api/retrieve_financial_data.html) | `list` | `GET /financial/v3/accounts/transaction-dates` | Required |
| [RetrieveFinancialData](docs/api/retrieve_financial_data.html) | `load` | `GET /financial/v2/accounts/{account_id}` | Required |
| [RetrieveFinancialData](docs/api/retrieve_financial_data.html) | `load` | `GET /financial/v3/accounts/{account_id}` | Required |
| [Similar](docs/api/similar.html) | `load` | `GET /corrections/v2/categories/similar/{transaction_id}` | Required |
| [Similar](docs/api/similar.html) | `load` | `GET /corrections/v2/merchants/similar/{transaction_id}` | Required |

## Connect to the API

- Bud sandbox: `https://api-sandbox.thisisbud.com`

The default credential is sent in the `Authorization` header with the `Bearer` prefix.

Authentication flow: 1. Perform OAuth2 Client Credentials authentication using API Credentials (`client_id`,`client_secret`) to obtain an `access_token` against `/v1/oauth/token` endpoint, 2. Use `access_token` as Bearer Authorisation for every other API request, 3. Include `X-Client-Id` (=client_id) within the header of every API request, 4. Note that some of the requests may also require `X-Customer-Id` to be provided within the request header. ### Examples Obtain OAuth2 `access_token` and `refresh_token` using `grant_type=client_credentials` and HTTP Basic auth header ``` curl --basic --user &#123;&#123;client_id&#125;&#125;:&#123;&#123;client_secret&#125;&#125; \ -X POST https://api-sandbox.thisisbud.com/v1/oauth/token \ -H &#39;Content-Type: application/x-www-form-urlencoded&#39; \ -d grant_type=client_credentials ``` Successful response: ``` &#123; &quot;operation_id&quot;: &quot;oauth_token_post&quot;, &quot;data&quot;: &#123; &quot;access_token&quot;: &quot;dd0c17e3fd6d2ce94aa091257a3ea393b4f9b5cf3d3e998f07dc9826da86ff15&quot;, &quot;token_type&quot;: &quot;bearer&quot;, &quot;expires_in&quot;: 3600, &quot;refresh_token&quot;: &quot;fac32cca7559d9f6e8f1dfe9a99c71fa1dcfeb482bedf287d7934d2667ae54b3&quot; &#125; &#125; ``` Refresh `access_token` token using `refresh_token` against `/v1/oauth/token` endpoint with `grant_type=refresh_token` ``` curl -X POST \ https://api-sandbox.thisisbud.com/v1/oauth/token \ -H &#39;Content-Type: application/x-www-form-urlencoded&#39; \ -H &#39;X-Client-Id: &#123;&#123;client_id&#125;&#125;&#39; \ -d &#39;grant_type=refresh_token&amp;refresh_token=&#123;&#123;refresh_token&#125;&#125;&#39; ``` Successful response: ``` &#123; &quot;operation_id&quot;: &quot;oauth_token_post&quot;, &quot;data&quot;: &#123; &quot;access_token&quot;: &quot;cc0c17e3fd6d2ce94aa091257a3ea393b4f9b5cf3d3e998f07dc9826da86ff94&quot;, &quot;token_type&quot;: &quot;bearer&quot;, &quot;expires_in&quot;: 3600, &quot;refresh_token&quot;: &quot;ffc30cca7559d9f6e8f1dfe9a99c71fa1dcfeb482bedf287d7934d2667ae54b3&quot; &#125; &#125; ```

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| [Golang](docs/sdks/go.html) | `go/` | Build from source |
| [JavaScript](docs/sdks/js.html) | `js/` | Build from source |
| [Lua](docs/sdks/lua.html) | `lua/` | Build from source |
| [PHP](docs/sdks/php.html) | `php/` | Build from source |
| [Python](docs/sdks/py.html) | `py/` | Build from source |
| [TypeScript](docs/sdks/ts.html) | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### [Go CLI](docs/tools/go-cli.html)

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### [Go MCP server](docs/tools/go-mcp.html)

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `bud-financial-data_list`: List records for an entity. Supported entities: `correct_financial_data`, `list_label`, `list_transaction`, `retrieve_financial_data`.
- `bud-financial-data_load`: Load one record for an entity. Supported entities: `correct_financial_data`, `retrieve_financial_data`, `similar`.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- [`debug`](docs/features/debug.html): Request/response capture ring buffer for debugging
- [`idempotency`](docs/features/idempotency.html): Idempotency keys for safe retries of mutating operations
- [`metrics`](docs/features/metrics.html): Statistics capture: per-operation counters and latency
- [`paging`](docs/features/paging.html): Pagination signals for list operations
- [`ratelimit`](docs/features/ratelimit.html): Client-side rate limiting via a token bucket
- [`retry`](docs/features/retry.html): Automatic retry of transient failures with exponential backoff
- [`test`](docs/features/test.html): In-memory mock transport for testing without a live server
- [`timeout`](docs/features/timeout.html): Per-request timeout with transport abort

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the [first-call guide](docs/guides/first-call.html) for the setup sequence.
- Read the [authentication guide](docs/guides/authentication.html) before using protected routes.
- Use the [API reference](docs/api/index.html) for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

