-- BudFinancialData SDK exists test

local sdk = require("bud-financial-data_sdk")

describe("BudFinancialDataSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
