-- BudFinancialData SDK error

local BudFinancialDataError = {}
BudFinancialDataError.__index = BudFinancialDataError


function BudFinancialDataError.new(code, msg, ctx)
  local self = setmetatable({}, BudFinancialDataError)
  self.is_sdk_error = true
  self.sdk = "BudFinancialData"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function BudFinancialDataError:error()
  return self.msg
end


function BudFinancialDataError:__tostring()
  return self.msg
end


return BudFinancialDataError
