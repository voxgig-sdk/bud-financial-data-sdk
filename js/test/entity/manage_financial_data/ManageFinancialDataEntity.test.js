
const envlocal = __dirname + '/../../../.env.local'
require('../../utility').loadEnvLocal(envlocal)

const Path = require('node:path')
const Fs = require('node:fs')

const { test, describe, afterEach } = require('node:test')
const assert = require('node:assert')
const { createLiveTransport } = require('../../live-runner')
const { runLiveEntity } = require('../../live-entity')


const { BudFinancialDataSDK, BaseFeature, stdutil, config } = require('../../..')

const {
  envOverride,
  liveClientOptions,
  liveDelay,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
} = require('../../utility')


describe('ManageFinancialDataEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when BUD_FINANCIAL_DATA_TEST_LIVE=TRUE.
  afterEach(liveDelay('BUD_FINANCIAL_DATA_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = BudFinancialDataSDK.test()
    const ent = testsdk.ManageFinancialData()
    assert(null != ent)
  })


  test('basic', async (t) => {

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"label_id":{"a":true,"fo":"uuid","h":"Label Id","n":"label_id","r":true,"sh":"Identifier of the label to attach.","t":"`$STRING`","key$":"label_id","index$":0}},"name":"manage_financial_data","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /financial/v2/transactions/{transaction_id}/labels","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","k":"header","n":"x_customer_id","or":"x_customer_id","r":true,"t":"`$STRING`","index$":1},{"a":true,"k":"header","n":"x_customer_idempotent_identifier","or":"x_customer_idempotent_identifier","r":false,"t":"`$STRING`","index$":2},{"a":true,"k":"header","n":"x_customer_secret","or":"x_customer_secret","r":false,"t":"`$STRING`","index$":3}],"params":[{"a":true,"k":"param","n":"transaction_id","or":"transaction_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/financial/v2/transactions/{transaction_id}/labels","q":{"exist":["transaction_id","x_client_id","x_customer_id","x_customer_idempotent_identifier","x_customer_secret"]},"r":{},"s":[{"lit":"financial"},{"lit":"v2"},{"lit":"transactions"},{"var":"transaction_id"},{"lit":"labels"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"},"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /financial/v2/transactions/{transaction_id}/labels/{label_id}","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","k":"header","n":"x_customer_id","or":"x_customer_id","r":true,"t":"`$STRING`","index$":1},{"a":true,"k":"header","n":"x_customer_idempotent_identifier","or":"x_customer_idempotent_identifier","r":false,"t":"`$STRING`","index$":2},{"a":true,"k":"header","n":"x_customer_secret","or":"x_customer_secret","r":false,"t":"`$STRING`","index$":3}],"params":[{"a":true,"k":"param","n":"label_id","or":"label_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"param","n":"transaction_id","or":"transaction_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"DELETE","o":"/financial/v2/transactions/{transaction_id}/labels/{label_id}","q":{"exist":["label_id","transaction_id","x_client_id","x_customer_id","x_customer_idempotent_identifier","x_customer_secret"]},"r":{},"s":[{"lit":"financial"},{"lit":"v2"},{"lit":"transactions"},{"var":"transaction_id"},{"lit":"labels"},{"var":"label_id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"co":{"id":"DELETE /financial/v2/accounts/{account_id}","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","k":"header","n":"x_customer_id","or":"x_customer_id","r":true,"t":"`$STRING`","index$":1},{"a":true,"k":"header","n":"x_customer_idempotent_identifier","or":"x_customer_idempotent_identifier","r":false,"t":"`$STRING`","index$":2},{"a":true,"k":"header","n":"x_customer_secret","or":"x_customer_secret","r":false,"t":"`$STRING`","index$":3}],"params":[{"a":true,"k":"param","n":"account_id","or":"account_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"DELETE","o":"/financial/v2/accounts/{account_id}","q":{"exist":["account_id","x_client_id","x_customer_id","x_customer_idempotent_identifier","x_customer_secret"]},"r":{},"s":[{"lit":"financial"},{"lit":"v2"},{"lit":"accounts"},{"var":"account_id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1},{"a":true,"co":{"id":"DELETE /v1/provider/{provider}","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","k":"header","n":"x_customer_id","or":"x_customer_id","r":true,"t":"`$STRING`","index$":1},{"a":true,"k":"header","n":"x_customer_idempotent_identifier","or":"x_customer_idempotent_identifier","r":false,"t":"`$STRING`","index$":2},{"a":true,"k":"header","n":"x_customer_secret","or":"x_customer_secret","r":false,"t":"`$STRING`","index$":3}],"params":[{"a":true,"k":"param","n":"provider","or":"provider","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"DELETE","o":"/v1/provider/{provider}","q":{"exist":["provider","x_client_id","x_customer_id","x_customer_idempotent_identifier","x_customer_secret"]},"r":{},"s":[{"lit":"v1"},{"lit":"provider"},{"var":"provider"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"}},"relations":{"ancestors":[["$.main.kit.entity.label"]]},"key$":"manage_financial_data","name__orig":"manage_financial_data","Name":"ManageFinancialData","name_":"manage_financial_data","name-":"manage-financial-data","NAME":"MANAGE_FINANCIAL_DATA","index$":5}, {"active":true,"entity":"manage_financial_data","key$":"BasicManageFinancialDataFlow","kind":"basic","name":"BasicManageFinancialDataFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"manage_financial_data_ref01"},"m":{"transaction_id":"transaction01"},"o":"create","s":[],"v":[],"index$":0},{"a":true,"d":{},"i":{"ref":"manage_financial_data_ref01","suffix":"_rm0"},"m":{"id":"manage_financial_data01"},"o":"remove","s":[],"v":[],"index$":1}]}, 'ManageFinancialData', {"POST /financial/v2/transactions/{transaction_id}/labels":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"title":"Attach Transaction Label Request","type":"object","required":["label_id"],"properties":{"label_id":{"type":"string","format":"uuid","description":"Identifier of the label to attach. The label must already exist for the authenticated customer.","key$":"label_id"}},"x-ref":"#/components/schemas/AssignTransactionLabelRequest","index$":1},"example":{"label_id":"1f10e6de-0f5f-4d8c-987a-9635a9d1b8c6"}}}},"parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/0","index$":0},{"in":"header","name":"X-Customer-Id","schema":{"type":"string","format":"uuid"},"required":true,"description":"A unique identifier for a Customer, as registered on Bud's platform.","example":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/1","index$":1},{"in":"header","name":"X-Customer-Idempotent-Identifier","schema":{"type":"string"},"description":"Use the internal client identifier, provided in the `client_metadata` object when creating the customer in __Create Customer V3__, in place of an `X-Customer-Id` header.\n","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/2","index$":2},{"in":"header","name":"X-Customer-Secret","schema":{"type":"string"},"required":false,"description":"The Bud Customer secret used to encrypt customer data. This is required only if the customer secret is not already stored with Bud.","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/3","index$":3},{"in":"path","name":"transaction_id","schema":{"type":"string","format":"uuid"},"required":true,"description":"A transaction identifier","x-ref":"#/paths/~1corrections~1v2~1merchants~1similar~1{transaction_id}/get/parameters/5","index$":4}]},"DELETE /financial/v2/transactions/{transaction_id}/labels/{label_id}":{"protocol":"http","parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/0","index$":0},{"in":"header","name":"X-Customer-Id","schema":{"type":"string","format":"uuid"},"required":true,"description":"A unique identifier for a Customer, as registered on Bud's platform.","example":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/1","index$":1},{"in":"header","name":"X-Customer-Idempotent-Identifier","schema":{"type":"string"},"description":"Use the internal client identifier, provided in the `client_metadata` object when creating the customer in __Create Customer V3__, in place of an `X-Customer-Id` header.\n","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/2","index$":2},{"in":"header","name":"X-Customer-Secret","schema":{"type":"string"},"required":false,"description":"The Bud Customer secret used to encrypt customer data. This is required only if the customer secret is not already stored with Bud.","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/3","index$":3},{"in":"path","name":"transaction_id","schema":{"type":"string","format":"uuid"},"required":true,"description":"A transaction identifier","x-ref":"#/paths/~1corrections~1v2~1merchants~1similar~1{transaction_id}/get/parameters/5","index$":4},{"in":"path","name":"label_id","required":true,"schema":{"type":"string","format":"uuid"},"description":"Identifier of the label to remove from the transaction.","index$":5}]},"DELETE /financial/v2/accounts/{account_id}":{"protocol":"http","parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/0","index$":0},{"in":"header","name":"X-Customer-Id","schema":{"type":"string","format":"uuid"},"required":true,"description":"A unique identifier for a Customer, as registered on Bud's platform.","example":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/1","index$":1},{"in":"header","name":"X-Customer-Idempotent-Identifier","schema":{"type":"string"},"description":"Use the internal client identifier, provided in the `client_metadata` object when creating the customer in __Create Customer V3__, in place of an `X-Customer-Id` header.\n","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/2","index$":2},{"in":"header","name":"X-Customer-Secret","schema":{"type":"string"},"required":false,"description":"The Bud Customer secret used to encrypt customer data. This is required only if the customer secret is not already stored with Bud.","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/3","index$":3},{"in":"path","name":"account_id","required":true,"schema":{"type":"string"},"description":"The unique identifier of the account to be deleted.","index$":4}]},"DELETE /v1/provider/{provider}":{"protocol":"http","parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/0","index$":0},{"in":"header","name":"X-Customer-Id","schema":{"type":"string","format":"uuid"},"required":true,"description":"A unique identifier for a Customer, as registered on Bud's platform.","example":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/1","index$":1},{"in":"header","name":"X-Customer-Idempotent-Identifier","schema":{"type":"string"},"description":"Use the internal client identifier, provided in the `client_metadata` object when creating the customer in __Create Customer V3__, in place of an `X-Customer-Id` header.\n","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/2","index$":2},{"in":"header","name":"X-Customer-Secret","schema":{"type":"string"},"required":false,"description":"The Bud Customer secret used to encrypt customer data. This is required only if the customer secret is not already stored with Bud.","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/3","index$":3},{"in":"path","name":"provider","schema":{"type":"string"},"required":true,"description":"Name of the provider (banking institution)","index$":4}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const manage_financial_data_ref01_ent = client.ManageFinancialData()
    let manage_financial_data_ref01_data = setup.data.new.manage_financial_data['manage_financial_data_ref01']
    manage_financial_data_ref01_data['transaction_id'] = setup.idmap['transaction01']

    manage_financial_data_ref01_data = (await manage_financial_data_ref01_ent.create(manage_financial_data_ref01_data)).data()
    assert(null != manage_financial_data_ref01_data)



  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/manage_financial_data/ManageFinancialDataTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = BudFinancialDataSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['manage_financial_data01','manage_financial_data02','manage_financial_data03','label01','label02','label03','transaction01'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'BUD_FINANCIAL_DATA_TEST_MANAGE_FINANCIAL_DATA_ENTID': idmap,
    'BUD_FINANCIAL_DATA_TEST_LIVE': 'FALSE',
    'BUD_FINANCIAL_DATA_TEST_EXPLAIN': 'FALSE',
    'BUD_FINANCIAL_DATA_APIKEY': '',
  })

  idmap = env['BUD_FINANCIAL_DATA_TEST_MANAGE_FINANCIAL_DATA_ENTID']

  const live = 'TRUE' === env.BUD_FINANCIAL_DATA_TEST_LIVE
  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['BUD_FINANCIAL_DATA_TEST_MANAGE_FINANCIAL_DATA_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new BudFinancialDataSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
        apikey: env.BUD_FINANCIAL_DATA_APIKEY,
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when
      // the last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey and
      // server values above and handed the SDK undefined.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
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
  }

  return setup
}
  
