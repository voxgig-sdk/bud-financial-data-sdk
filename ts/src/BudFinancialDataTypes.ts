// Typed models for the BudFinancialData SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface CorrectFinancialData {
  created_at: string
  custom_merchant_id: string
  data: any[]
  frequency: string
  include_similar?: boolean
  logo_feedback: string
  metadata: Record<string, any>
  name: string
  online_or_billing_only?: boolean
  operation_id: string
  reference_transaction_id?: string
  rule_definition: Record<string, any>
  rule_type: string
  similar?: boolean
  suggested_logo?: string
  suggested_url?: string
  transaction_id: string
  transaction_ids: any[]
}

export interface CorrectFinancialDataLoadMatch {
  rule_id: string
}

export interface CorrectFinancialDataListMatch {
  created_at?: string
  custom_merchant_id?: string
  data?: any[]
  frequency?: string
  include_similar?: boolean
  logo_feedback?: string
  metadata?: Record<string, any>
  name?: string
  online_or_billing_only?: boolean
  operation_id?: string
  reference_transaction_id?: string
  rule_definition?: Record<string, any>
  rule_type?: string
  similar?: boolean
  suggested_logo?: string
  suggested_url?: string
  transaction_id?: string
  transaction_ids?: any[]
}

export interface CorrectFinancialDataCreateData {
  merchant_id: string
  created_at: string
  custom_merchant_id: string
  data: any[]
  frequency: string
  include_similar?: boolean
  logo_feedback: string
  metadata: Record<string, any>
  name: string
  online_or_billing_only?: boolean
  operation_id: string
  reference_transaction_id?: string
  rule_definition: Record<string, any>
  rule_type: string
  similar?: boolean
  suggested_logo?: string
  suggested_url?: string
  transaction_id: string
  transaction_ids: any[]
}

export interface CorrectFinancialDataRemoveMatch {
  rule_id: string
}

export interface CustomerMerchantCorrection {
  data: any[]
  metadata: Record<string, any>
  operation_id: string
}

export interface CustomerMerchantCorrectionCreateData {
  data: any[]
  metadata: Record<string, any>
  operation_id: string
}

export interface Label {
  created_at: string
  id: string
  name: string
  updated_at: string
}

export interface LabelCreateData {
  created_at: string
  id: string
  name: string
  updated_at: string
}

export interface LabelUpdateData {
  id: string
  created_at?: string
  name?: string
  updated_at?: string
}

export interface ListLabel {
  created_at: string
  id: string
  name: string
  updated_at: string
}

export interface ListLabelListMatch {
  created_at?: string
  id?: string
  name?: string
  updated_at?: string
}

export interface ListTransaction {
  account_id: string
  amount: Record<string, any>
  client_attributes?: Record<string, any>
  counterparty?: Record<string, any>
  credit_debit_indicator: string
  date_time: string
  description: string
  enrichments?: Record<string, any>
  labels?: any[]
  merchant_category_code?: string
  posted_date_time?: string
  provider?: string
  running_balance: Record<string, any>
  running_balance_credit_debit_indicator?: string
  status: string
  suggested_description: string
  suggested_logo?: string
  tags?: any[]
  transaction_id: string
  transaction_type?: Record<string, any>
  value_date_time?: string
}

export interface ListTransactionListMatch {
  account_id?: string
  category_l1?: string
  category_l2?: string
  credit_debit_indicator?: string
  date_from?: string
  date_to?: string
  exclude_tag?: any[]
  include_label_id?: any[]
  include_tag?: any[]
  merchant?: string
  page_size?: number
  page_token?: string
  status?: string
  updated_after?: string
}

export interface ManageFinancialData {
  label_id: string
}

export interface ManageFinancialDataCreateData {
  transaction_id: string
  label_id: string
}

export interface ManageFinancialDataRemoveMatch {
  provider: string
}

export interface ManageTransactionLabel {
  id?: string
}

export interface ManageTransactionLabelRemoveMatch {
  id: string
}

export interface RetrieveFinancialData {
  account_category?: string
  account_id: string
  account_name?: string
  account_type?: string
  balances?: Record<string, any>
  booked: any
  closed_at?: string
  credit_lines?: Record<string, any>
  currency: string
  data: Record<string, any>
  date: string
  details?: any[]
  first_transaction_date: string
  frequency?: string
  holder?: Record<string, any>
  holders?: any[]
  id: string
  identifiers?: Record<string, any>
  last_transaction_date: string
  metadata: Record<string, any>
  name: string
  opening_date_time?: string
  operation_id: string
  pending: any
  provider?: string
  provider_display_name?: string
  provider_logo?: string
  reference: string
  restriction?: string
  status?: string
  suggested_name: string
  transaction_windows?: any[]
  type: string
  usage_type?: string
}

export interface RetrieveFinancialDataLoadMatch {
  account_id: string
}

export interface RetrieveFinancialDataListMatch {
  account_type?: string
  currency?: string
  exclude_restricted?: boolean
  holder_relationship_type?: string
  include_all_balance?: boolean
  provider?: string
  usage_type?: string
}

export interface Similar {
  data: any[]
  id?: string
  metadata: Record<string, any>
  operation_id: string
}

export interface SimilarLoadMatch {
  id: string
  exclude_source?: boolean
}

