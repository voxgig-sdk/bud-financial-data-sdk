

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


describe('ListLabelEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when BUD_FINANCIAL_DATA_TEST_LIVE=TRUE.
  afterEach(liveDelay('BUD_FINANCIAL_DATA_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = BudFinancialDataSDK.test()
    const ent = testsdk.ListLabel()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.BUD_FINANCIAL_DATA_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'list_label.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"created_at":{"a":true,"fo":"date-time","h":"Created At","n":"created_at","r":true,"sh":"RFC3339 timestamp at which the label was created.","t":"`$STRING`","key$":"created_at","index$":0},"id":{"a":true,"fo":"uuid","h":"Id","n":"id","r":true,"sh":"Unique identifier of the label.","t":"`$STRING`","key$":"id","index$":1},"name":{"a":true,"h":"Name","n":"name","r":true,"sh":"Display name of the label.","t":"`$STRING`","key$":"name","index$":2},"updated_at":{"a":true,"fo":"date-time","h":"Updated At","n":"updated_at","r":true,"sh":"RFC3339 timestamp at which the label was last updated.","t":"`$STRING`","key$":"updated_at","index$":3}},"id":{"field":"id","name":"id"},"name":"list_label","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /financial/v2/transactions/labels","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","k":"header","n":"x_customer_id","or":"x_customer_id","r":true,"t":"`$STRING`","index$":1},{"a":true,"k":"header","n":"x_customer_idempotent_identifier","or":"x_customer_idempotent_identifier","r":false,"t":"`$STRING`","index$":2},{"a":true,"k":"header","n":"x_customer_secret","or":"x_customer_secret","r":false,"t":"`$STRING`","index$":3}]},"k":"http","m":"GET","o":"/financial/v2/transactions/labels","q":{"exist":["x_client_id","x_customer_id","x_customer_idempotent_identifier","x_customer_secret"]},"r":{},"s":[{"lit":"financial"},{"lit":"v2"},{"lit":"transactions"},{"lit":"labels"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"list_label","name__orig":"list_label","Name":"ListLabel","name_":"list_label","name-":"list-label","NAME":"LIST_LABEL","index$":3}, {"active":true,"entity":"list_label","key$":"BasicListLabelFlow","kind":"basic","name":"BasicListLabelFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"list_label_ref01"}}],"index$":0}]}, 'ListLabel', {"GET /financial/v2/transactions/labels":{"protocol":"http","parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/0","index$":0},{"in":"header","name":"X-Customer-Id","schema":{"type":"string","format":"uuid"},"required":true,"description":"A unique identifier for a Customer, as registered on Bud's platform.","example":"c3339af9-2426-44d4-9c4d-ddb8fd281e23","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/1","index$":1},{"in":"header","name":"X-Customer-Idempotent-Identifier","schema":{"type":"string"},"description":"Use the internal client identifier, provided in the `client_metadata` object when creating the customer in __Create Customer V3__, in place of an `X-Customer-Id` header.\n","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/2","index$":2},{"in":"header","name":"X-Customer-Secret","schema":{"type":"string"},"required":false,"description":"The Bud Customer secret used to encrypt customer data. This is required only if the customer secret is not already stored with Bud.","x-ref":"#/paths/~1financial~1v3~1accounts/get/parameters/3","index$":3}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let list_label_ref01_data = Object.values(setup.data.existing.list_label)[0] as any

    // LIST
    const list_label_ref01_ent = client.ListLabel()
    const list_label_ref01_match: any = {}

    const list_label_ref01_list = (await list_label_ref01_ent.list(list_label_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/list_label/ListLabelTestData.json')

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
    ['list_label01','list_label02','list_label03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'BUD_FINANCIAL_DATA_TEST_LIST_LABEL_ENTID': idmap,
    'BUD_FINANCIAL_DATA_TEST_LIVE': 'FALSE',
    'BUD_FINANCIAL_DATA_TEST_EXPLAIN': 'FALSE',
    'BUD_FINANCIAL_DATA_APIKEY': '',
  })

  idmap = env['BUD_FINANCIAL_DATA_TEST_LIST_LABEL_ENTID']

  const live = 'TRUE' === env.BUD_FINANCIAL_DATA_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['BUD_FINANCIAL_DATA_TEST_LIST_LABEL_ENTID']
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
  
