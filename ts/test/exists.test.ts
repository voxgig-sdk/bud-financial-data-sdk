
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { BudFinancialDataSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = BudFinancialDataSDK.test()
    equal(testsdk instanceof BudFinancialDataSDK, true,
      'BudFinancialDataSDK.test() must return a client synchronously')
  })

})
