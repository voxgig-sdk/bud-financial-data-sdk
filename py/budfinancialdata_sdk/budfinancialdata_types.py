# Typed models for the BudFinancialData SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class CorrectFinancialDataRequired(TypedDict):
    created_at: str
    custom_merchant_id: str
    data: list
    frequency: str
    logo_feedback: str
    metadata: dict
    name: str
    operation_id: str
    rule_definition: dict
    rule_type: str
    transaction_id: str
    transaction_ids: list


class CorrectFinancialData(CorrectFinancialDataRequired, total=False):
    include_similar: bool
    online_or_billing_only: bool
    reference_transaction_id: str
    similar: bool
    suggested_logo: str
    suggested_url: str


class CorrectFinancialDataLoadMatch(TypedDict):
    rule_id: str


class CorrectFinancialDataListMatch(TypedDict, total=False):
    created_at: str
    custom_merchant_id: str
    data: list
    frequency: str
    include_similar: bool
    logo_feedback: str
    metadata: dict
    name: str
    online_or_billing_only: bool
    operation_id: str
    reference_transaction_id: str
    rule_definition: dict
    rule_type: str
    similar: bool
    suggested_logo: str
    suggested_url: str
    transaction_id: str
    transaction_ids: list


class CorrectFinancialDataCreateDataRequired(TypedDict):
    merchant_id: str
    created_at: str
    custom_merchant_id: str
    data: list
    frequency: str
    logo_feedback: str
    metadata: dict
    name: str
    operation_id: str
    rule_definition: dict
    rule_type: str
    transaction_id: str
    transaction_ids: list


class CorrectFinancialDataCreateData(CorrectFinancialDataCreateDataRequired, total=False):
    include_similar: bool
    online_or_billing_only: bool
    reference_transaction_id: str
    similar: bool
    suggested_logo: str
    suggested_url: str


class CorrectFinancialDataRemoveMatch(TypedDict):
    rule_id: str


class CustomerMerchantCorrection(TypedDict):
    data: list
    metadata: dict
    operation_id: str


class CustomerMerchantCorrectionCreateData(TypedDict):
    data: list
    metadata: dict
    operation_id: str


class Label(TypedDict):
    created_at: str
    id: str
    name: str
    updated_at: str


class LabelCreateData(TypedDict):
    created_at: str
    id: str
    name: str
    updated_at: str


class LabelUpdateDataRequired(TypedDict):
    id: str


class LabelUpdateData(LabelUpdateDataRequired, total=False):
    created_at: str
    name: str
    updated_at: str


class ListLabel(TypedDict):
    created_at: str
    id: str
    name: str
    updated_at: str


class ListLabelListMatch(TypedDict, total=False):
    created_at: str
    id: str
    name: str
    updated_at: str


class ListTransactionRequired(TypedDict):
    account_id: str
    amount: dict
    credit_debit_indicator: str
    date_time: str
    description: str
    running_balance: dict
    status: str
    suggested_description: str
    transaction_id: str


class ListTransaction(ListTransactionRequired, total=False):
    client_attributes: dict
    counterparty: dict
    enrichments: dict
    labels: list
    merchant_category_code: str
    posted_date_time: str
    provider: str
    running_balance_credit_debit_indicator: str
    suggested_logo: str
    tags: list
    transaction_type: dict
    value_date_time: str


class ListTransactionListMatch(TypedDict, total=False):
    account_id: str
    category_l1: str
    category_l2: str
    credit_debit_indicator: str
    date_from: str
    date_to: str
    exclude_tag: list
    include_label_id: list
    include_tag: list
    merchant: str
    page_size: int
    page_token: str
    status: str
    updated_after: str


class ManageFinancialData(TypedDict):
    label_id: str


class ManageFinancialDataCreateData(TypedDict):
    transaction_id: str
    label_id: str


class ManageFinancialDataRemoveMatch(TypedDict):
    provider: str


class ManageTransactionLabel(TypedDict, total=False):
    id: str


class ManageTransactionLabelRemoveMatch(TypedDict):
    id: str


class RetrieveFinancialDataRequired(TypedDict):
    account_id: str
    booked: Any
    currency: str
    data: dict
    date: str
    first_transaction_date: str
    id: str
    last_transaction_date: str
    metadata: dict
    name: str
    operation_id: str
    pending: Any
    reference: str
    suggested_name: str
    type: str


class RetrieveFinancialData(RetrieveFinancialDataRequired, total=False):
    account_category: str
    account_name: str
    account_type: str
    balances: dict
    closed_at: str
    credit_lines: dict
    details: list
    frequency: str
    holder: dict
    holders: list
    identifiers: dict
    opening_date_time: str
    provider: str
    provider_display_name: str
    provider_logo: str
    restriction: str
    status: str
    transaction_windows: list
    usage_type: str


class RetrieveFinancialDataLoadMatch(TypedDict):
    account_id: str


class RetrieveFinancialDataListMatch(TypedDict, total=False):
    account_type: str
    currency: str
    exclude_restricted: bool
    holder_relationship_type: str
    include_all_balance: bool
    provider: str
    usage_type: str


class SimilarRequired(TypedDict):
    data: list
    metadata: dict
    operation_id: str


class Similar(SimilarRequired, total=False):
    id: str


class SimilarLoadMatchRequired(TypedDict):
    id: str


class SimilarLoadMatch(SimilarLoadMatchRequired, total=False):
    exclude_source: bool
