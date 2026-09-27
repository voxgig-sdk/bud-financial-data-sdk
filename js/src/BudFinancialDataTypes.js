// Typed models for the BudFinancialData SDK (JSDoc typedefs).
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
// edit by hand.

/**
 * @typedef {Object} CorrectFinancialData
 * @property {string} created_at
 * @property {string} custom_merchant_id
 * @property {Array} data
 * @property {string} frequency
 * @property {boolean} [include_similar]
 * @property {string} logo_feedback
 * @property {Object} metadata
 * @property {string} name
 * @property {boolean} [online_or_billing_only]
 * @property {string} operation_id
 * @property {string} [reference_transaction_id]
 * @property {Object} rule_definition
 * @property {string} rule_type
 * @property {boolean} [similar]
 * @property {string} [suggested_logo]
 * @property {string} [suggested_url]
 * @property {string} transaction_id
 * @property {Array} transaction_ids
 */

/**
 * @typedef {Object} CorrectFinancialDataLoadMatch
 * @property {string} rule_id
 */

/**
 * @typedef {Object} CorrectFinancialDataListMatch
 * @property {string} [created_at]
 * @property {string} [custom_merchant_id]
 * @property {Array} [data]
 * @property {string} [frequency]
 * @property {boolean} [include_similar]
 * @property {string} [logo_feedback]
 * @property {Object} [metadata]
 * @property {string} [name]
 * @property {boolean} [online_or_billing_only]
 * @property {string} [operation_id]
 * @property {string} [reference_transaction_id]
 * @property {Object} [rule_definition]
 * @property {string} [rule_type]
 * @property {boolean} [similar]
 * @property {string} [suggested_logo]
 * @property {string} [suggested_url]
 * @property {string} [transaction_id]
 * @property {Array} [transaction_ids]
 */

/**
 * @typedef {Object} CorrectFinancialDataCreateData
 * @property {string} merchant_id
 * @property {string} created_at
 * @property {string} custom_merchant_id
 * @property {Array} data
 * @property {string} frequency
 * @property {boolean} [include_similar]
 * @property {string} logo_feedback
 * @property {Object} metadata
 * @property {string} name
 * @property {boolean} [online_or_billing_only]
 * @property {string} operation_id
 * @property {string} [reference_transaction_id]
 * @property {Object} rule_definition
 * @property {string} rule_type
 * @property {boolean} [similar]
 * @property {string} [suggested_logo]
 * @property {string} [suggested_url]
 * @property {string} transaction_id
 * @property {Array} transaction_ids
 */

/**
 * @typedef {Object} CorrectFinancialDataRemoveMatch
 * @property {string} rule_id
 */

/**
 * @typedef {Object} CustomerMerchantCorrection
 * @property {Array} data
 * @property {Object} metadata
 * @property {string} operation_id
 */

/**
 * @typedef {Object} CustomerMerchantCorrectionCreateData
 * @property {Array} data
 * @property {Object} metadata
 * @property {string} operation_id
 */

/**
 * @typedef {Object} Label
 * @property {string} created_at
 * @property {string} id
 * @property {string} name
 * @property {string} updated_at
 */

/**
 * @typedef {Object} LabelCreateData
 * @property {string} created_at
 * @property {string} id
 * @property {string} name
 * @property {string} updated_at
 */

/**
 * @typedef {Object} LabelUpdateData
 * @property {string} id
 * @property {string} [created_at]
 * @property {string} [name]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} ListLabel
 * @property {string} created_at
 * @property {string} id
 * @property {string} name
 * @property {string} updated_at
 */

/**
 * @typedef {Object} ListLabelListMatch
 * @property {string} [created_at]
 * @property {string} [id]
 * @property {string} [name]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} ListTransaction
 * @property {string} account_id
 * @property {Object} amount
 * @property {Object} [client_attributes]
 * @property {Object} [counterparty]
 * @property {string} credit_debit_indicator
 * @property {string} date_time
 * @property {string} description
 * @property {Object} [enrichments]
 * @property {Array} [labels]
 * @property {string} [merchant_category_code]
 * @property {string} [posted_date_time]
 * @property {string} [provider]
 * @property {Object} running_balance
 * @property {string} [running_balance_credit_debit_indicator]
 * @property {string} status
 * @property {string} suggested_description
 * @property {string} [suggested_logo]
 * @property {Array} [tags]
 * @property {string} transaction_id
 * @property {Object} [transaction_type]
 * @property {string} [value_date_time]
 */

/**
 * @typedef {Object} ListTransactionListMatch
 * @property {string} [account_id]
 * @property {string} [category_l1]
 * @property {string} [category_l2]
 * @property {string} [credit_debit_indicator]
 * @property {string} [date_from]
 * @property {string} [date_to]
 * @property {Array} [exclude_tag]
 * @property {Array} [include_label_id]
 * @property {Array} [include_tag]
 * @property {string} [merchant]
 * @property {number} [page_size]
 * @property {string} [page_token]
 * @property {string} [status]
 * @property {string} [updated_after]
 */

/**
 * @typedef {Object} ManageFinancialData
 * @property {string} label_id
 */

/**
 * @typedef {Object} ManageFinancialDataCreateData
 * @property {string} transaction_id
 * @property {string} label_id
 */

/**
 * @typedef {Object} ManageFinancialDataRemoveMatch
 * @property {string} provider
 */

/**
 * @typedef {Object} ManageTransactionLabel
 * @property {string} [id]
 */

/**
 * @typedef {Object} ManageTransactionLabelRemoveMatch
 * @property {string} id
 */

/**
 * @typedef {Object} RetrieveFinancialData
 * @property {string} [account_category]
 * @property {string} account_id
 * @property {string} [account_name]
 * @property {string} [account_type]
 * @property {Object} [balances]
 * @property {*} booked
 * @property {string} [closed_at]
 * @property {Object} [credit_lines]
 * @property {string} currency
 * @property {Object} data
 * @property {string} date
 * @property {Array} [details]
 * @property {string} first_transaction_date
 * @property {string} [frequency]
 * @property {Object} [holder]
 * @property {Array} [holders]
 * @property {string} id
 * @property {Object} [identifiers]
 * @property {string} last_transaction_date
 * @property {Object} metadata
 * @property {string} name
 * @property {string} [opening_date_time]
 * @property {string} operation_id
 * @property {*} pending
 * @property {string} [provider]
 * @property {string} [provider_display_name]
 * @property {string} [provider_logo]
 * @property {string} reference
 * @property {string} [restriction]
 * @property {string} [status]
 * @property {string} suggested_name
 * @property {Array} [transaction_windows]
 * @property {string} type
 * @property {string} [usage_type]
 */

/**
 * @typedef {Object} RetrieveFinancialDataLoadMatch
 * @property {string} account_id
 */

/**
 * @typedef {Object} RetrieveFinancialDataListMatch
 * @property {string} [account_type]
 * @property {string} [currency]
 * @property {boolean} [exclude_restricted]
 * @property {string} [holder_relationship_type]
 * @property {boolean} [include_all_balance]
 * @property {string} [provider]
 * @property {string} [usage_type]
 */

/**
 * @typedef {Object} Similar
 * @property {Array} data
 * @property {string} [id]
 * @property {Object} metadata
 * @property {string} operation_id
 */

/**
 * @typedef {Object} SimilarLoadMatch
 * @property {string} id
 * @property {boolean} [exclude_source]
 */

