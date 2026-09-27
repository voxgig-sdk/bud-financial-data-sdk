
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


describe('ManageTransactionLabelEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when BUD_FINANCIAL_DATA_TEST_LIVE=TRUE.
  afterEach(liveDelay('BUD_FINANCIAL_DATA_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = BudFinancialDataSDK.test()
    const ent = testsdk.ManageTransactionLabel()
    assert(null != ent)
  })


  test('basic', async (t) => {

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"id":{"a":true,"h":"Id","n":"id","r":false,"t":"`$STRING`","key$":"id","index$":0}},"id":{"field":"id","name":"id"},"name":"manage_transaction_label","op":{"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /financial/v2/transactions/labels/{label_id}","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","k":"header","n":"x_customer_id","or":"x_customer_id","r":true,"t":"`$STRING`","index$":1},{"a":true,"k":"header","n":"x_customer_idempotent_identifier","or":"x_customer_idempotent_identifier","r":false,"t":"`$STRING`","index$":2},{"a":true,"k":"header","n":"x_customer_secret","or":"x_customer_secret","r":false,"t":"`$STRING`","index$":3}],"params":[{"a":true,"k":"param","n":"id","or":"label_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"DELETE","o":"/financial/v2/transactions/labels/{label_id}","q":{"exist":["id","x_client_id","x_customer_id","x_customer_idempotent_identifier","x_customer_secret"]},"r":{"param":{"label_id":"id"}},"s":[{"lit":"financial"},{"lit":"v2"},{"lit":"transactions"},{"lit":"labels"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"}},"relations":{"ancestors":[]},"key$":"manage_transaction_label","name__orig":"manage_transaction_label","Name":"ManageTransactionLabel","name_":"manage_transaction_label","name-":"manage-transaction-label","NAME":"MANAGE_TRANSACTION_LABEL","index$":6}, {"active":true,"entity":"manage_transaction_label","key$":"BasicManageTransactionLabelFlow","kind":"basic","name":"BasicManageTransactionLabelFlow","param":{},"step":[]}, 'ManageTransactionLabel', {"DELETE /financial/v2/transactions/labels/{label_id}":{"protocol":"http","parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/0","index$":0},{"in":"header","name":"X-Customer-Id","schema":{"type":"string","format":"uuid"},"required":true,"description":"A unique identifier for a Customer, as registered on Bud's platform.","example":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/1","index$":1},{"in":"header","name":"X-Customer-Idempotent-Identifier","schema":{"type":"string"},"description":"Use the internal client identifier, provided in the `client_metadata` object when creating the customer in __Create Customer V3__, in place of an `X-Customer-Id` header.\n","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/2","index$":2},{"in":"header","name":"X-Customer-Secret","schema":{"type":"string"},"required":false,"description":"The Bud Customer secret used to encrypt customer data. This is required only if the customer secret is not already stored with Bud.","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/3","index$":3},{"in":"path","name":"label_id","required":true,"schema":{"type":"string","format":"uuid"},"description":"Identifier of the label to delete.","index$":4}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let manage_transaction_label_ref01_data = Object.values(setup.data.existing.manage_transaction_label)[0]

  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/manage_transaction_label/ManageTransactionLabelTestData.json')

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
    ['manage_transaction_label01','manage_transaction_label02','manage_transaction_label03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'BUD_FINANCIAL_DATA_TEST_MANAGE_TRANSACTION_LABEL_ENTID': idmap,
    'BUD_FINANCIAL_DATA_TEST_LIVE': 'FALSE',
    'BUD_FINANCIAL_DATA_TEST_EXPLAIN': 'FALSE',
    'BUD_FINANCIAL_DATA_APIKEY': '',
  })

  idmap = env['BUD_FINANCIAL_DATA_TEST_MANAGE_TRANSACTION_LABEL_ENTID']

  const live = 'TRUE' === env.BUD_FINANCIAL_DATA_TEST_LIVE
  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['BUD_FINANCIAL_DATA_TEST_MANAGE_TRANSACTION_LABEL_ENTID']
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
  
