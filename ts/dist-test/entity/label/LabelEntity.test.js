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
(0, node_test_1.describe)('LabelEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when BUD_FINANCIAL_DATA_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('BUD_FINANCIAL_DATA_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.BudFinancialDataSDK.test();
        const ent = testsdk.Label();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.BUD_FINANCIAL_DATA_TEST_LIVE;
        for (const op of ['create', 'update']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'label.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "created_at": { "a": true, "fo": "date-time", "h": "Created At", "n": "created_at", "r": true, "sh": "RFC3339 timestamp at which the label was created.", "t": "`$STRING`", "key$": "created_at", "index$": 0 }, "id": { "a": true, "fo": "uuid", "h": "Id", "n": "id", "r": true, "sh": "Unique identifier of the label.", "t": "`$STRING`", "key$": "id", "index$": 1 }, "name": { "a": true, "h": "Name", "n": "name", "r": true, "sh": "Display name for the new label.", "t": "`$STRING`", "key$": "name", "index$": 2 }, "updated_at": { "a": true, "fo": "date-time", "h": "Updated At", "n": "updated_at", "r": true, "sh": "RFC3339 timestamp at which the label was last updated.", "t": "`$STRING`", "key$": "updated_at", "index$": 3 } }, "id": { "field": "id", "name": "id" }, "name": "label", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "co": { "id": "POST /financial/v2/transactions/labels", "source": "openapi3", "version": 2 }, "g": { "header": [{ "a": true, "ex": "8711bbef-b357-4c2d-97ca-0d9df4206e9a", "k": "header", "n": "x_client_id", "or": "x_client_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "c3339af9-2426-44d4-9c4d-ddb8fd281e23", "k": "header", "n": "x_customer_id", "or": "x_customer_id", "r": true, "t": "`$STRING`", "index$": 1 }, { "a": true, "k": "header", "n": "x_customer_idempotent_identifier", "or": "x_customer_idempotent_identifier", "r": false, "t": "`$STRING`", "index$": 2 }, { "a": true, "k": "header", "n": "x_customer_secret", "or": "x_customer_secret", "r": false, "t": "`$STRING`", "index$": 3 }] }, "k": "http", "m": "POST", "o": "/financial/v2/transactions/labels", "q": { "exist": ["x_client_id", "x_customer_id", "x_customer_idempotent_identifier", "x_customer_secret"] }, "r": {}, "s": [{ "lit": "financial" }, { "lit": "v2" }, { "lit": "transactions" }, { "lit": "labels" }], "t": { "req": "`reqdata`", "res": "`body.data`" }, "index$": 0 }], "key$": "create" }, "update": { "input": "data", "name": "update", "points": [{ "a": true, "co": { "id": "PATCH /financial/v2/transactions/labels/{label_id}", "source": "openapi3", "version": 2 }, "g": { "header": [{ "a": true, "ex": "8711bbef-b357-4c2d-97ca-0d9df4206e9a", "k": "header", "n": "x_client_id", "or": "x_client_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "c3339af9-2426-44d4-9c4d-ddb8fd281e23", "k": "header", "n": "x_customer_id", "or": "x_customer_id", "r": true, "t": "`$STRING`", "index$": 1 }, { "a": true, "k": "header", "n": "x_customer_idempotent_identifier", "or": "x_customer_idempotent_identifier", "r": false, "t": "`$STRING`", "index$": 2 }, { "a": true, "k": "header", "n": "x_customer_secret", "or": "x_customer_secret", "r": false, "t": "`$STRING`", "index$": 3 }], "params": [{ "a": true, "k": "param", "n": "id", "or": "label_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "PATCH", "o": "/financial/v2/transactions/labels/{label_id}", "q": { "exist": ["id", "x_client_id", "x_customer_id", "x_customer_idempotent_identifier", "x_customer_secret"] }, "r": { "param": { "label_id": "id" } }, "s": [{ "lit": "financial" }, { "lit": "v2" }, { "lit": "transactions" }, { "lit": "labels" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body.data`" }, "index$": 0 }], "key$": "update" } }, "relations": { "ancestors": [] }, "key$": "label", "name__orig": "label", "Name": "Label", "name_": "label", "name-": "label", "NAME": "LABEL", "index$": 2 }, { "active": true, "entity": "label", "key$": "BasicLabelFlow", "kind": "basic", "name": "BasicLabelFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "label_ref01" }, "m": {}, "o": "create", "s": [], "v": [], "index$": 0 }, { "a": true, "d": {}, "i": { "ref": "label_ref01", "srcdatavar": "label_ref01_data", "suffix": "_up0", "textfield": "created_at" }, "m": {}, "o": "update", "s": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-label_ref01" } }], "v": [], "index$": 1 }] }, 'Label', { "POST /financial/v2/transactions/labels": { "protocol": "http", "requestBody": { "required": true, "content": { "application/json": { "schema": { "title": "Create Label Request", "type": "object", "required": ["name"], "properties": { "name": { "type": "string", "minLength": 1, "maxLength": 50, "description": "Display name for the new label. Must be non-empty and no more than 50 characters.", "key$": "name" } }, "x-ref": "#/components/schemas/CreateLabelRequest", "index$": 1 }, "example": { "name": "Trip to Lisbon" } } } }, "parameters": [{ "in": "header", "name": "X-Client-Id", "schema": { "type": "string" }, "required": true, "description": "The API Client Identifier (Service Application Identifier).", "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a", "x-ref": "#/paths/~1financial~1v3~1accounts/get/parameters/0", "index$": 0 }, { "in": "header", "name": "X-Customer-Id", "schema": { "type": "string", "format": "uuid" }, "required": true, "description": "A unique identifier for a Customer, as registered on Bud's platform.", "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23", "x-ref": "#/paths/~1financial~1v3~1accounts/get/parameters/1", "index$": 1 }, { "in": "header", "name": "X-Customer-Idempotent-Identifier", "schema": { "type": "string" }, "description": "Use the internal client identifier, provided in the `client_metadata` object when creating the customer in __Create Customer V3__, in place of an `X-Customer-Id` header.\n", "x-ref": "#/paths/~1financial~1v3~1accounts/get/parameters/2", "index$": 2 }, { "in": "header", "name": "X-Customer-Secret", "schema": { "type": "string" }, "required": false, "description": "The Bud Customer secret used to encrypt customer data. This is required only if the customer secret is not already stored with Bud.", "x-ref": "#/paths/~1financial~1v3~1accounts/get/parameters/3", "index$": 3 }] }, "PATCH /financial/v2/transactions/labels/{label_id}": { "protocol": "http", "requestBody": { "required": true, "content": { "application/json": { "schema": { "title": "Rename Label Request", "type": "object", "required": ["name"], "properties": { "name": { "type": "string", "minLength": 1, "maxLength": 50, "description": "New display name for the label. Must be non-empty and no more than 50 characters.", "key$": "name" } }, "x-ref": "#/components/schemas/RenameLabelRequest", "index$": 1 }, "example": { "name": "Home renovation" } } } }, "parameters": [{ "in": "header", "name": "X-Client-Id", "schema": { "type": "string" }, "required": true, "description": "The API Client Identifier (Service Application Identifier).", "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a", "x-ref": "#/paths/~1financial~1v3~1accounts/get/parameters/0", "index$": 0 }, { "in": "header", "name": "X-Customer-Id", "schema": { "type": "string", "format": "uuid" }, "required": true, "description": "A unique identifier for a Customer, as registered on Bud's platform.", "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23", "x-ref": "#/paths/~1financial~1v3~1accounts/get/parameters/1", "index$": 1 }, { "in": "header", "name": "X-Customer-Idempotent-Identifier", "schema": { "type": "string" }, "description": "Use the internal client identifier, provided in the `client_metadata` object when creating the customer in __Create Customer V3__, in place of an `X-Customer-Id` header.\n", "x-ref": "#/paths/~1financial~1v3~1accounts/get/parameters/2", "index$": 2 }, { "in": "header", "name": "X-Customer-Secret", "schema": { "type": "string" }, "required": false, "description": "The Bud Customer secret used to encrypt customer data. This is required only if the customer secret is not already stored with Bud.", "x-ref": "#/paths/~1financial~1v3~1accounts/get/parameters/3", "index$": 3 }, { "in": "path", "name": "label_id", "required": true, "schema": { "type": "string", "format": "uuid" }, "description": "Identifier of the label to rename.", "index$": 4 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const label_ref01_ent = client.Label();
        let label_ref01_data = setup.data.new.label['label_ref01'];
        label_ref01_data = (await label_ref01_ent.create(label_ref01_data)).data();
        (0, node_assert_1.default)(null != label_ref01_data.id);
        // UPDATE
        const label_ref01_data_up0 = {};
        label_ref01_data_up0.id = label_ref01_data.id;
        const label_ref01_markdef_up0 = { name: 'created_at', value: 'Mark01-label_ref01_' + setup.now };
        label_ref01_data_up0[label_ref01_markdef_up0.name] = label_ref01_markdef_up0.value;
        const label_ref01_resdata_up0 = (await label_ref01_ent.update(label_ref01_data_up0)).data();
        (0, node_assert_1.default)(label_ref01_resdata_up0.id === label_ref01_data_up0.id);
        (0, node_assert_1.default)(label_ref01_resdata_up0[label_ref01_markdef_up0.name] === label_ref01_markdef_up0.value);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/label/LabelTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.BudFinancialDataSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['label01', 'label02', 'label03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'BUD_FINANCIAL_DATA_TEST_LABEL_ENTID': idmap,
        'BUD_FINANCIAL_DATA_TEST_LIVE': 'FALSE',
        'BUD_FINANCIAL_DATA_TEST_EXPLAIN': 'FALSE',
        'BUD_FINANCIAL_DATA_APIKEY': '',
    });
    idmap = env['BUD_FINANCIAL_DATA_TEST_LABEL_ENTID'];
    const live = 'TRUE' === env.BUD_FINANCIAL_DATA_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['BUD_FINANCIAL_DATA_TEST_LABEL_ENTID'];
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
//# sourceMappingURL=LabelEntity.test.js.map