

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


describe('SimilarEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when BUD_FINANCIAL_DATA_TEST_LIVE=TRUE.
  afterEach(liveDelay('BUD_FINANCIAL_DATA_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = BudFinancialDataSDK.test()
    const ent = testsdk.Similar()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.BUD_FINANCIAL_DATA_TEST_LIVE
    for (const op of ['load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'similar.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"data":{"a":true,"h":"Data","n":"data","r":true,"t":"`$ARRAY`","key$":"data","index$":0},"id":{"a":true,"h":"Id","n":"id","r":false,"t":"`$STRING`","key$":"id","index$":1},"metadata":{"a":true,"h":"Metadata","n":"metadata","r":true,"t":"`$OBJECT`","key$":"metadata","index$":2},"operation_id":{"a":true,"h":"Operation Id","n":"operation_id","r":true,"t":"`$STRING`","key$":"operation_id","index$":3}},"id":{"field":"id","name":"id"},"name":"similar","op":{"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /corrections/v2/categories/similar/{transaction_id}","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","k":"header","n":"x_customer_id","or":"x_customer_id","r":true,"t":"`$STRING`","index$":1},{"a":true,"k":"header","n":"x_customer_idempotent_identifier","or":"x_customer_idempotent_identifier","r":false,"t":"`$STRING`","index$":2},{"a":true,"k":"header","n":"x_customer_secret","or":"x_customer_secret","r":false,"t":"`$STRING`","index$":3}],"params":[{"a":true,"k":"param","n":"id","or":"transaction_id","r":true,"t":"`$STRING`","index$":0}],"query":[{"a":true,"k":"query","n":"exclude_source","or":"exclude_source","r":false,"t":"`$BOOLEAN`","index$":0}]},"k":"http","m":"GET","o":"/corrections/v2/categories/similar/{transaction_id}","q":{"exist":["exclude_source","id","x_client_id","x_customer_id","x_customer_idempotent_identifier","x_customer_secret"]},"r":{"param":{"transaction_id":"id"}},"s":[{"lit":"corrections"},{"lit":"v2"},{"lit":"categories"},{"lit":"similar"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"co":{"id":"GET /corrections/v2/merchants/similar/{transaction_id}","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","k":"header","n":"x_customer_id","or":"x_customer_id","r":true,"t":"`$STRING`","index$":1},{"a":true,"k":"header","n":"x_customer_idempotent_identifier","or":"x_customer_idempotent_identifier","r":false,"t":"`$STRING`","index$":2},{"a":true,"k":"header","n":"x_customer_secret","or":"x_customer_secret","r":false,"t":"`$STRING`","index$":3}],"params":[{"a":true,"k":"param","n":"id","or":"transaction_id","r":true,"t":"`$STRING`","index$":0}],"query":[{"a":true,"k":"query","n":"exclude_source","or":"exclude_source","r":false,"t":"`$BOOLEAN`","index$":0}]},"k":"http","m":"GET","o":"/corrections/v2/merchants/similar/{transaction_id}","q":{"exist":["exclude_source","id","x_client_id","x_customer_id","x_customer_idempotent_identifier","x_customer_secret"]},"r":{"param":{"transaction_id":"id"}},"s":[{"lit":"corrections"},{"lit":"v2"},{"lit":"merchants"},{"lit":"similar"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"similar","name__orig":"similar","Name":"Similar","name_":"similar","name-":"similar","NAME":"SIMILAR","index$":8}, {"active":true,"entity":"similar","key$":"BasicSimilarFlow","kind":"basic","name":"BasicSimilarFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"similar_ref01","srcdatavar":"similar_ref01_data","suffix":"_dt0"},"m":{"id":"similar01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-similar_ref01"}}],"index$":0}]}, 'Similar', {"GET /corrections/v2/categories/similar/{transaction_id}":{"protocol":"http","parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/0","index$":0},{"in":"header","name":"X-Customer-Id","schema":{"type":"string","format":"uuid"},"required":true,"description":"A unique identifier for a Customer, as registered on Bud's platform.","example":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/1","index$":1},{"in":"header","name":"X-Customer-Idempotent-Identifier","schema":{"type":"string"},"description":"Use the internal client identifier, provided in the `client_metadata` object when creating the customer in __Create Customer V3__, in place of an `X-Customer-Id` header.\n","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/2","index$":2},{"in":"header","name":"X-Customer-Secret","schema":{"type":"string"},"required":false,"description":"The Bud Customer secret used to encrypt customer data. This is required only if the customer secret is not already stored with Bud.","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/3","index$":3},{"in":"query","name":"exclude_source","schema":{"type":"boolean"},"required":false,"description":"Toggle whether the endpoint returns the transaction used as part of the request","x-ref":"#/paths/~1corrections~1v2~1merchants~1similar~1{transaction_id}/get/parameters/4","index$":4},{"in":"path","name":"transaction_id","schema":{"type":"string","format":"uuid"},"required":true,"description":"A transaction identifier","x-ref":"#/paths/~1corrections~1v2~1merchants~1similar~1{transaction_id}/get/parameters/5","index$":5}]},"GET /corrections/v2/merchants/similar/{transaction_id}":{"protocol":"http","parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/0","index$":0},{"in":"header","name":"X-Customer-Id","schema":{"type":"string","format":"uuid"},"required":true,"description":"A unique identifier for a Customer, as registered on Bud's platform.","example":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/1","index$":1},{"in":"header","name":"X-Customer-Idempotent-Identifier","schema":{"type":"string"},"description":"Use the internal client identifier, provided in the `client_metadata` object when creating the customer in __Create Customer V3__, in place of an `X-Customer-Id` header.\n","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/2","index$":2},{"in":"header","name":"X-Customer-Secret","schema":{"type":"string"},"required":false,"description":"The Bud Customer secret used to encrypt customer data. This is required only if the customer secret is not already stored with Bud.","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/3","index$":3},{"in":"query","name":"exclude_source","schema":{"type":"boolean"},"required":false,"description":"Toggle whether the endpoint returns the transaction used as part of the request","index$":4},{"in":"path","name":"transaction_id","schema":{"type":"string","format":"uuid"},"required":true,"description":"A transaction identifier","index$":5}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let similar_ref01_data = Object.values(setup.data.existing.similar)[0] as any

    // LOAD
    const similar_ref01_ent = client.Similar()
    const similar_ref01_match_dt0: any = {}
    similar_ref01_match_dt0.id = similar_ref01_data.id
    const similar_ref01_data_dt0 = (await similar_ref01_ent.load(similar_ref01_match_dt0)).data()
    assert(similar_ref01_data_dt0.id === similar_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/similar/SimilarTestData.json')

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
    ['similar01','similar02','similar03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'BUD_FINANCIAL_DATA_TEST_SIMILAR_ENTID': idmap,
    'BUD_FINANCIAL_DATA_TEST_LIVE': 'FALSE',
    'BUD_FINANCIAL_DATA_TEST_EXPLAIN': 'FALSE',
    'BUD_FINANCIAL_DATA_APIKEY': '',
  })

  idmap = env['BUD_FINANCIAL_DATA_TEST_SIMILAR_ENTID']

  const live = 'TRUE' === env.BUD_FINANCIAL_DATA_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['BUD_FINANCIAL_DATA_TEST_SIMILAR_ENTID']
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
  
