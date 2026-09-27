"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('ListTransactionEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when BUD_FINANCIAL_DATA_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('BUD_FINANCIAL_DATA_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.BudFinancialDataSDK.test();
        const ent = testsdk.ListTransaction();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.BUD_FINANCIAL_DATA_TEST_LIVE;
        for (const op of ['list']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'list_transaction.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "account_id": { "a": true, "h": "Account Id", "n": "account_id", "r": true, "sh": "Identifier for the account associated with the transaction.", "t": "`$STRING`", "key$": "account_id", "index$": 0 }, "amount": { "a": true, "h": "Amount", "n": "amount", "r": true, "sh": "The monetary amount.", "t": "`$OBJECT`", "key$": "amount", "index$": 1 }, "client_attributes": { "a": true, "h": "Client Attributes", "n": "client_attributes", "r": false, "sh": "An optional map of key-value string pairs that allows clients to attach arbitrary metadata to the transaction.", "t": "`$OBJECT`", "key$": "client_attributes", "index$": 2 }, "counterparty": { "a": true, "h": "Counterparty", "n": "counterparty", "r": false, "sh": "An object containing details of the counterparty in this transaction", "t": "`$OBJECT`", "key$": "counterparty", "index$": 3 }, "credit_debit_indicator": { "a": true, "h": "Credit Debit Indicator", "n": "credit_debit_indicator", "r": true, "sh": "Credit/Debit Indicator", "t": "`$STRING`", "key$": "credit_debit_indicator", "index$": 4 }, "date_time": { "a": true, "h": "Date Time", "n": "date_time", "r": true, "sh": "Date that the transaction occured compliant with RFC3339.", "t": "`$STRING`", "key$": "date_time", "index$": 5 }, "description": { "a": true, "h": "Description", "n": "description", "r": true, "sh": "Description of the transaction.", "t": "`$STRING`", "key$": "description", "index$": 6 }, "enrichments": { "a": true, "h": "Enrichments", "n": "enrichments", "r": false, "sh": "Contextual enrichments associated with a Transaction", "t": "`$OBJECT`", "key$": "enrichments", "index$": 7 }, "labels": { "a": true, "h": "Labels", "n": "labels", "r": false, "sh": "Customer-defined labels currently attached to this transaction.", "t": "`$ARRAY`", "key$": "labels", "index$": 8 }, "merchant_category_code": { "a": true, "h": "Merchant Category Code", "n": "merchant_category_code", "r": false, "sh": "Merchant category code conforming to ISO 18245, related to the type of services or goods the merchant provides for the transaction.", "t": "`$STRING`", "key$": "merchant_category_code", "index$": 9 }, "posted_date_time": { "a": true, "de": true, "h": "Posted Date Time", "n": "posted_date_time", "r": false, "sh": "Date the assets involved in the transaction transferred compliant with RFC3339.", "t": "`$STRING`", "key$": "posted_date_time", "index$": 10 }, "provider": { "a": true, "h": "Provider", "n": "provider", "r": false, "sh": "Name of the transaction source provider.", "t": "`$STRING`", "key$": "provider", "index$": 11 }, "running_balance": { "a": true, "h": "Running Balance", "n": "running_balance", "r": true, "sh": "The running balance for the account that the transaction takes place against", "t": "`$OBJECT`", "key$": "running_balance", "index$": 12 }, "running_balance_credit_debit_indicator": { "a": true, "h": "Running Balance Credit Debit Indicator", "n": "running_balance_credit_debit_indicator", "r": false, "sh": "Credit/Debit Indicator for the running balance field", "t": "`$STRING`", "key$": "running_balance_credit_debit_indicator", "index$": 13 }, "status": { "a": true, "h": "Status", "n": "status", "r": true, "sh": "Status of the transaction.", "t": "`$STRING`", "key$": "status", "index$": 14 }, "suggested_description": { "a": true, "h": "Suggested Description", "n": "suggested_description", "r": true, "sh": "The description Bud suggests client apps show for a given transaction.", "t": "`$STRING`", "key$": "suggested_description", "index$": 15 }, "suggested_logo": { "a": true, "h": "Suggested Logo", "n": "suggested_logo", "r": false, "sh": "The logo Bud suggests client apps show for a given transaction.", "t": "`$STRING`", "key$": "suggested_logo", "index$": 16 }, "tags": { "a": true, "h": "Tags", "n": "tags", "r": false, "sh": "A list of potential tags associated with the transaction after contextual enrichment, which can be used for filtering.", "t": "`$ARRAY`", "key$": "tags", "index$": 17 }, "transaction_id": { "a": true, "h": "Transaction Id", "n": "transaction_id", "r": true, "sh": "Unique identifier for the transaction.", "t": "`$STRING`", "key$": "transaction_id", "index$": 18 }, "transaction_type": { "a": true, "h": "Transaction Type", "n": "transaction_type", "r": false, "sh": "The code and a description of the transaction type.", "t": "`$OBJECT`", "key$": "transaction_type", "index$": 19 }, "value_date_time": { "a": true, "h": "Value Date Time", "n": "value_date_time", "r": false, "sh": "Date the assets involved in the transaction transferred compliant with RFC3339.", "t": "`$STRING`", "key$": "value_date_time", "index$": 20 } }, "name": "list_transaction", "op": { "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /financial/v2/transactions", "source": "openapi3", "version": 2 }, "g": { "header": [{ "a": true, "ex": "8711bbef-b357-4c2d-97ca-0d9df4206e9a", "k": "header", "n": "x_client_id", "or": "x_client_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "c3339af9-2426-44d4-9c4d-ddb8fd281e23", "k": "header", "n": "x_customer_id", "or": "x_customer_id", "r": true, "t": "`$STRING`", "index$": 1 }, { "a": true, "k": "header", "n": "x_customer_idempotent_identifier", "or": "x_customer_idempotent_identifier", "r": false, "t": "`$STRING`", "index$": 2 }, { "a": true, "k": "header", "n": "x_customer_secret", "or": "x_customer_secret", "r": false, "t": "`$STRING`", "index$": 3 }], "query": [{ "a": true, "ex": "5bd9ecaf31fc2d5172fd89cb61b89fc6", "k": "query", "n": "account_id", "or": "account_id", "r": false, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "bills", "k": "query", "n": "category_l1", "or": "category_l1", "r": false, "t": "`$STRING`", "index$": 1 }, { "a": true, "ex": "tv_and_broadband", "k": "query", "n": "category_l2", "or": "category_l2", "r": false, "t": "`$STRING`", "index$": 2 }, { "a": true, "ex": "credit", "k": "query", "n": "credit_debit_indicator", "or": "credit_debit_indicator", "r": false, "t": "`$STRING`", "index$": 3 }, { "a": true, "ex": "2022-05-28T00:00:00Z", "k": "query", "n": "date_from", "or": "date_from", "r": false, "t": "`$STRING`", "index$": 4 }, { "a": true, "ex": "2022-05-29T00:00:00Z", "k": "query", "n": "date_to", "or": "date_to", "r": false, "t": "`$STRING`", "index$": 5 }, { "a": true, "ex": ["benefit"], "k": "query", "n": "exclude_tag", "or": "exclude_tag", "r": false, "t": "`$ARRAY`", "index$": 6 }, { "a": true, "ex": ["1f10e6de-0f5f-4d8c-987a-9635a9d1b8c6"], "k": "query", "n": "include_label_id", "or": "include_label_id", "r": false, "t": "`$ARRAY`", "index$": 7 }, { "a": true, "ex": ["regular-transaction"], "k": "query", "n": "include_tag", "or": "include_tag", "r": false, "t": "`$ARRAY`", "index$": 8 }, { "a": true, "ex": "netflix", "k": "query", "n": "merchant", "or": "merchant", "r": false, "t": "`$STRING`", "index$": 9 }, { "a": true, "ex": 100, "k": "query", "n": "page_size", "or": "page_size", "r": false, "t": "`$INTEGER`", "index$": 10 }, { "a": true, "ex": "eyJvZmZzZXQiOjEwMH0", "k": "query", "n": "page_token", "or": "page_token", "r": false, "t": "`$STRING`", "index$": 11 }, { "a": true, "ex": "booked", "k": "query", "n": "status", "or": "status", "r": false, "t": "`$STRING`", "index$": 12 }, { "a": true, "ex": "2024-06-01T11:00:00Z", "k": "query", "n": "updated_after", "or": "updated_after", "r": false, "t": "`$STRING`", "index$": 13 }] }, "k": "http", "m": "GET", "o": "/financial/v2/transactions", "q": { "exist": ["account_id", "category_l1", "category_l2", "credit_debit_indicator", "date_from", "date_to", "exclude_tag", "include_label_id", "include_tag", "merchant", "page_size", "page_token", "status", "updated_after", "x_client_id", "x_customer_id", "x_customer_idempotent_identifier", "x_customer_secret"] }, "r": {}, "s": [{ "lit": "financial" }, { "lit": "v2" }, { "lit": "transactions" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" } }, "relations": { "ancestors": [] }, "key$": "list_transaction", "name__orig": "list_transaction", "Name": "ListTransaction", "name_": "list_transaction", "name-": "list-transaction", "NAME": "LIST_TRANSACTION", "index$": 4 }, { "active": true, "entity": "list_transaction", "key$": "BasicListTransactionFlow", "kind": "basic", "name": "BasicListTransactionFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": {}, "m": {}, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "list_transaction_ref01" } }], "index$": 0 }] }, 'ListTransaction', { "GET /financial/v2/transactions": { "protocol": "http", "parameters": [{ "in": "header", "name": "X-Client-Id", "schema": { "type": "string" }, "required": true, "description": "The API Client Identifier (Service Application Identifier).", "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a", "x-ref": "#/paths/~1financial~1v3~1accounts/get/parameters/0", "index$": 0 }, { "in": "header", "name": "X-Customer-Id", "schema": { "type": "string", "format": "uuid" }, "required": true, "description": "A unique identifier for a Customer, as registered on Bud's platform.", "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23", "x-ref": "#/paths/~1financial~1v3~1accounts/get/parameters/1", "index$": 1 }, { "in": "header", "name": "X-Customer-Idempotent-Identifier", "schema": { "type": "string" }, "description": "Use the internal client identifier, provided in the `client_metadata` object when creating the customer in __Create Customer V3__, in place of an `X-Customer-Id` header.\n", "x-ref": "#/paths/~1financial~1v3~1accounts/get/parameters/2", "index$": 2 }, { "in": "header", "name": "X-Customer-Secret", "schema": { "type": "string" }, "required": false, "description": "The Bud Customer secret used to encrypt customer data. This is required only if the customer secret is not already stored with Bud.", "x-ref": "#/paths/~1financial~1v3~1accounts/get/parameters/3", "index$": 3 }, { "in": "query", "name": "date_from", "schema": { "type": "string" }, "example": "2022-05-28T00:00:00Z", "description": "Date (RFC3339) from which the transactions should be returned from. To maintain backwards compatibility,  dates in the format (YYYY-MM-DD) can also be used and will be assumed to be UTC.\n", "index$": 4 }, { "in": "query", "name": "date_to", "schema": { "type": "string" }, "example": "2022-05-29T00:00:00Z", "description": "Date (RFC3339) from which the transactions should be returned to. To maintain backwards compatibility,  dates in the format (YYYY-MM-DD) can also be used and will be assumed to be UTC. Uses the current date by default.\n", "index$": 5 }, { "in": "query", "name": "page_size", "schema": { "type": "integer" }, "example": 100, "description": "Maximum number of results to be returned. Defaults to 100, maximum is 200. \nNote: When using the `exclude_tags` filter, the response may contain more transactions than the specified  `page_size` to ensure complete filtering. Always use the presence of `next_page_token` to determine if  additional pages are available.\n", "index$": 6 }, { "in": "query", "name": "page_token", "schema": { "type": "string" }, "example": "eyJvZmZzZXQiOjEwMH0", "description": "The token required to fetch a specific page of results. Provided by the `next_page_token` field in the previous request.", "index$": 7 }, { "in": "query", "name": "account_id", "schema": { "type": "string" }, "example": "5bd9ecaf31fc2d5172fd89cb61b89fc6", "description": "Identifier of the accounts from which transactions should be returned.", "index$": 8 }, { "in": "query", "name": "merchant", "schema": { "type": "string" }, "description": "Human readable identifier of the merchants for which associated transactions should be returned.", "example": "netflix", "index$": 9 }, { "in": "query", "name": "category_l1", "schema": { "type": "string" }, "example": "bills", "description": "Identifier of the Category Level 1 for which associated transactions should be returned.", "index$": 10 }, { "in": "query", "name": "category_l2", "schema": { "type": "string" }, "example": "tv_and_broadband", "description": "Identifier of the Category Level 2 for which associated transactions should be returned.", "index$": 11 }, { "in": "query", "name": "include_tags", "schema": { "type": "array", "items": { "type": "string" } }, "example": ["regular-transaction"], "description": "Tags for which associated transactions should be returned. See response body example for an exhaustive list of possible tags.", "index$": 12 }, { "in": "query", "name": "exclude_tags", "schema": { "type": "array", "items": { "type": "string" } }, "example": ["benefit"], "description": "Tags for which associated transactions should not be returned. See response body example for an exhaustive list of possible tags.\nWhen this filter is applied, the response may include more transactions than the `page_size` to ensure accurate filtering. Always use the presence of `next_page_token` to determine if additional pages are available.\n", "index$": 13 }, { "in": "query", "name": "include_label_ids", "schema": { "type": "array", "items": { "type": "string", "format": "uuid" } }, "example": ["1f10e6de-0f5f-4d8c-987a-9635a9d1b8c6"], "description": "Customer-defined `label_id` values to filter on. Only transactions that have all of the supplied `label_id` values attached are returned. Repeat the parameter to supply multiple values.\nLabels are managed via the __Manage Transaction Labels__ endpoints.\n", "index$": 14 }, { "in": "query", "name": "credit_debit_indicator", "schema": { "type": "string", "enum": ["credit", "debit"] }, "example": "credit", "description": "Credit debit indicator for which associated transactions should be returned.", "index$": 15 }, { "in": "query", "name": "updated_after", "schema": { "type": "string" }, "example": "2024-06-01T11:00:00Z", "description": "The `updated_after` parameter lets you retrieve transactions that have been modified since  a specific date and time. This format should follow the RFC3339 standard (e.g., \"2024-06-01T11:00:00Z\").\nThis parameter is particularly useful for keeping local copies of transactions up-to-date. By storing the  timestamp of your last successful fetch, you can request only transactions updated after that point during  subsequent retrievals. \nThis approach ensures you only fetch the latest changes, making the process more streamlined.\n", "index$": 16 }, { "in": "query", "name": "status", "schema": { "type": "string", "enum": ["booked", "pending", "declined"] }, "example": "booked", "description": "Statuses of transactions to be returned. This parameter can be supplied multiple times to fetch transactions with a combination of statuses.\n\nBy default, declined transactions are excluded from the results.\nNote: Transactions ingested via Open Banking will only come as `booked` or `pending`.\nThe `declined` status is only applicable to transactions ingested through our `Ingest First Party Data` endpoint.\n", "index$": 17 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let list_transaction_ref01_data = Object.values(setup.data.existing.list_transaction)[0];
        // LIST
        const list_transaction_ref01_ent = client.ListTransaction();
        const list_transaction_ref01_match = {};
        const list_transaction_ref01_list = (await list_transaction_ref01_ent.list(list_transaction_ref01_match)).map((e) => e.data());
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/list_transaction/ListTransactionTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.BudFinancialDataSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['list_transaction01', 'list_transaction02', 'list_transaction03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'BUD_FINANCIAL_DATA_TEST_LIST_TRANSACTION_ENTID': idmap,
        'BUD_FINANCIAL_DATA_TEST_LIVE': 'FALSE',
        'BUD_FINANCIAL_DATA_TEST_EXPLAIN': 'FALSE',
        'BUD_FINANCIAL_DATA_APIKEY': '',
    });
    idmap = env['BUD_FINANCIAL_DATA_TEST_LIST_TRANSACTION_ENTID'];
    const live = 'TRUE' === env.BUD_FINANCIAL_DATA_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['BUD_FINANCIAL_DATA_TEST_LIST_TRANSACTION_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.BudFinancialDataSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {
                apikey: env.BUD_FINANCIAL_DATA_APIKEY,
            },
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
    }
    const setup = {
        idmap,
        env,
        options,
        client,
        struct,
        data: entityData,
        explain: 'TRUE' === env.BUD_FINANCIAL_DATA_TEST_EXPLAIN,
        live,
        transport,
        now: Date.now(),
    };
    return setup;
}
//# sourceMappingURL=ListTransactionEntity.test.js.map