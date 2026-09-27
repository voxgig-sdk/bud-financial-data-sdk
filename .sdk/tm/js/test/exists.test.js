
const { test, describe } = require('node:test')
const { equal } = require('node:assert')


const { BudFinancialDataSDK } = require('..')


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await BudFinancialDataSDK.test()
    equal(null !== testsdk, true)
  })

})
