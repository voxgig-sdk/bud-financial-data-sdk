-- Typed models for the BudFinancialData SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
-- params (op.<name>.points[].g.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class CorrectFinancialData
---@field created_at string
---@field custom_merchant_id string
---@field data table
---@field frequency string
---@field include_similar? boolean
---@field logo_feedback string
---@field metadata table
---@field name string
---@field online_or_billing_only? boolean
---@field operation_id string
---@field reference_transaction_id? string
---@field rule_definition table
---@field rule_type string
---@field similar? boolean
---@field suggested_logo? string
---@field suggested_url? string
---@field transaction_id string
---@field transaction_ids table

---@class CorrectFinancialDataLoadMatch
---@field rule_id string

---@class CorrectFinancialDataListMatch
---@field created_at? string
---@field custom_merchant_id? string
---@field data? table
---@field frequency? string
---@field include_similar? boolean
---@field logo_feedback? string
---@field metadata? table
---@field name? string
---@field online_or_billing_only? boolean
---@field operation_id? string
---@field reference_transaction_id? string
---@field rule_definition? table
---@field rule_type? string
---@field similar? boolean
---@field suggested_logo? string
---@field suggested_url? string
---@field transaction_id? string
---@field transaction_ids? table

---@class CorrectFinancialDataCreateData
---@field merchant_id string
---@field created_at string
---@field custom_merchant_id string
---@field data table
---@field frequency string
---@field include_similar? boolean
---@field logo_feedback string
---@field metadata table
---@field name string
---@field online_or_billing_only? boolean
---@field operation_id string
---@field reference_transaction_id? string
---@field rule_definition table
---@field rule_type string
---@field similar? boolean
---@field suggested_logo? string
---@field suggested_url? string
---@field transaction_id string
---@field transaction_ids table

---@class CorrectFinancialDataRemoveMatch
---@field rule_id string

---@class CustomerMerchantCorrection
---@field data table
---@field metadata table
---@field operation_id string

---@class CustomerMerchantCorrectionCreateData
---@field data table
---@field metadata table
---@field operation_id string

---@class Label
---@field created_at string
---@field id string
---@field name string
---@field updated_at string

---@class LabelCreateData
---@field created_at string
---@field id string
---@field name string
---@field updated_at string

---@class LabelUpdateData
---@field id string
---@field created_at? string
---@field name? string
---@field updated_at? string

---@class ListLabel
---@field created_at string
---@field id string
---@field name string
---@field updated_at string

---@class ListLabelListMatch
---@field created_at? string
---@field id? string
---@field name? string
---@field updated_at? string

---@class ListTransaction
---@field account_id string
---@field amount table
---@field client_attributes? table
---@field counterparty? table
---@field credit_debit_indicator string
---@field date_time string
---@field description string
---@field enrichments? table
---@field labels? table
---@field merchant_category_code? string
---@field posted_date_time? string
---@field provider? string
---@field running_balance table
---@field running_balance_credit_debit_indicator? string
---@field status string
---@field suggested_description string
---@field suggested_logo? string
---@field tags? table
---@field transaction_id string
---@field transaction_type? table
---@field value_date_time? string

---@class ListTransactionListMatch
---@field account_id? string
---@field category_l1? string
---@field category_l2? string
---@field credit_debit_indicator? string
---@field date_from? string
---@field date_to? string
---@field exclude_tag? table
---@field include_label_id? table
---@field include_tag? table
---@field merchant? string
---@field page_size? number
---@field page_token? string
---@field status? string
---@field updated_after? string

---@class ManageFinancialData
---@field label_id string

---@class ManageFinancialDataCreateData
---@field transaction_id string
---@field label_id string

---@class ManageFinancialDataRemoveMatch
---@field provider string

---@class ManageTransactionLabel
---@field id? string

---@class ManageTransactionLabelRemoveMatch
---@field id string

---@class RetrieveFinancialData
---@field account_category? string
---@field account_id string
---@field account_name? string
---@field account_type? string
---@field balances? table
---@field booked any
---@field closed_at? string
---@field credit_lines? table
---@field currency string
---@field data table
---@field date string
---@field details? table
---@field first_transaction_date string
---@field frequency? string
---@field holder? table
---@field holders? table
---@field id string
---@field identifiers? table
---@field last_transaction_date string
---@field metadata table
---@field name string
---@field opening_date_time? string
---@field operation_id string
---@field pending any
---@field provider? string
---@field provider_display_name? string
---@field provider_logo? string
---@field reference string
---@field restriction? string
---@field status? string
---@field suggested_name string
---@field transaction_windows? table
---@field type string
---@field usage_type? string

---@class RetrieveFinancialDataLoadMatch
---@field account_id string

---@class RetrieveFinancialDataListMatch
---@field account_type? string
---@field currency? string
---@field exclude_restricted? boolean
---@field holder_relationship_type? string
---@field include_all_balance? boolean
---@field provider? string
---@field usage_type? string

---@class Similar
---@field data table
---@field id? string
---@field metadata table
---@field operation_id string

---@class SimilarLoadMatch
---@field id string
---@field exclude_source? boolean

local M = {}

return M
