

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { BudFinancialDataSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


loadEnvLocal(__dirname + '/../../../.env.local')


describe('CustomerMerchantCorrectionEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when BUD_FINANCIAL_DATA_TEST_LIVE=TRUE.
  afterEach(liveDelay('BUD_FINANCIAL_DATA_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = BudFinancialDataSDK.test()
    const ent = testsdk.CustomerMerchantCorrection()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.BUD_FINANCIAL_DATA_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'customer_merchant_correction.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"data":{"a":true,"h":"Data","n":"data","r":true,"t":"`$ARRAY`","key$":"data","index$":0},"metadata":{"a":true,"h":"Metadata","n":"metadata","r":true,"sh":"Metadata associated with the merchant correction response schema","t":"`$OBJECT`","key$":"metadata","index$":1},"operation_id":{"a":true,"h":"Operation Id","n":"operation_id","r":true,"t":"`$STRING`","key$":"operation_id","index$":2}},"name":"customer_merchant_correction","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /corrections/v2/merchants","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","k":"header","n":"x_customer_id","or":"x_customer_id","r":true,"t":"`$STRING`","index$":1},{"a":true,"k":"header","n":"x_customer_idempotent_identifier","or":"x_customer_idempotent_identifier","r":false,"t":"`$STRING`","index$":2},{"a":true,"k":"header","n":"x_customer_secret","or":"x_customer_secret","r":false,"t":"`$STRING`","index$":3}]},"k":"http","m":"POST","o":"/corrections/v2/merchants","q":{"exist":["x_client_id","x_customer_id","x_customer_idempotent_identifier","x_customer_secret"]},"r":{},"s":[{"lit":"corrections"},{"lit":"v2"},{"lit":"merchants"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"customer_merchant_correction","name__orig":"customer_merchant_correction","Name":"CustomerMerchantCorrection","name_":"customer_merchant_correction","name-":"customer-merchant-correction","NAME":"CUSTOMER_MERCHANT_CORRECTION","index$":1}, {"active":true,"entity":"customer_merchant_correction","key$":"BasicCustomerMerchantCorrectionFlow","kind":"basic","name":"BasicCustomerMerchantCorrectionFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"customer_merchant_correction_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0}]}, 'CustomerMerchantCorrection', {"POST /corrections/v2/merchants":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"title":"Customer Merchant Corrections Request Payload","type":"array","description":"Expected request structure for the merchant corrections endpoint","items":{"title":"Customer Merchant Correction","type":"object","description":"Expected request structure for the merchant corrections endpoint","required":["transactions"],"properties":{"transactions":{"type":"array","minItems":1,"items":{"type":"string"},"description":"A list of transaction identifiers to be corrected"},"merchant_id":{"type":"string","description":"The identifier of the merchant that the transactions should be assigned\n\nTo remove a merchant instead of replacing it, leave this value blank\n"},"include_similar":{"type":"boolean","description":"Correct the merchant of other similar transactions to the specified transactions","default":false}}},"index$":1},"examples":{"Single Transaction Correction":{"value":[{"transactions":["7affc586-5600-4434-ac5f-1a845260d7d9"],"merchant_id":"f2b21ff7-836d-468c-8928-e81060d49988","include_similar":false}]},"Similar Transactions Correction":{"value":[{"transactions":["7affc586-5600-4434-ac5f-1a845260d7d9"],"merchant_id":"f2b21ff7-836d-468c-8928-e81060d49988","include_similar":true}]},"Remove Merchant From Single Transaction":{"value":[{"transactions":["7affc586-5600-4434-ac5f-1a845260d7d9"],"merchant_id":"","include_similar":false}]}}}}},"parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/0","index$":0},{"in":"header","name":"X-Customer-Id","schema":{"type":"string","format":"uuid"},"required":true,"description":"A unique identifier for a Customer, as registered on Bud's platform.","example":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/1","index$":1},{"in":"header","name":"X-Customer-Idempotent-Identifier","schema":{"type":"string"},"description":"Use the internal client identifier, provided in the `client_metadata` object when creating the customer in __Create Customer V3__, in place of an `X-Customer-Id` header.\n","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/2","index$":2},{"in":"header","name":"X-Customer-Secret","schema":{"type":"string"},"required":false,"description":"The Bud Customer secret used to encrypt customer data. This is required only if the customer secret is not already stored with Bud.","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/3","index$":3}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const customer_merchant_correction_ref01_ent = client.CustomerMerchantCorrection()
    let customer_merchant_correction_ref01_data = setup.data.new.customer_merchant_correction['customer_merchant_correction_ref01']

    customer_merchant_correction_ref01_data = (await customer_merchant_correction_ref01_ent.create(customer_merchant_correction_ref01_data)).data()
    assert(null != customer_merchant_correction_ref01_data)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/customer_merchant_correction/CustomerMerchantCorrectionTestData.json')

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
    ['customer_merchant_correction01','customer_merchant_correction02','customer_merchant_correction03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'BUD_FINANCIAL_DATA_TEST_CUSTOMER_MERCHANT_CORRECTION_ENTID': idmap,
    'BUD_FINANCIAL_DATA_TEST_LIVE': 'FALSE',
    'BUD_FINANCIAL_DATA_TEST_EXPLAIN': 'FALSE',
    'BUD_FINANCIAL_DATA_APIKEY': '',
  })

  idmap = env['BUD_FINANCIAL_DATA_TEST_CUSTOMER_MERCHANT_CORRECTION_ENTID']

  const live = 'TRUE' === env.BUD_FINANCIAL_DATA_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['BUD_FINANCIAL_DATA_TEST_CUSTOMER_MERCHANT_CORRECTION_ENTID']
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
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
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
  
