<?php
declare(strict_types=1);

// Typed models for the BudFinancialData SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** CorrectFinancialData entity data model. */
class CorrectFinancialData
{
    public string $created_at;
    public string $custom_merchant_id;
    public array $data;
    public string $frequency;
    public ?bool $include_similar = null;
    public string $logo_feedback;
    public array $metadata;
    public string $name;
    public ?bool $online_or_billing_only = null;
    public string $operation_id;
    public ?string $reference_transaction_id = null;
    public array $rule_definition;
    public string $rule_type;
    public ?bool $similar = null;
    public ?string $suggested_logo = null;
    public ?string $suggested_url = null;
    public string $transaction_id;
    public array $transaction_ids;
}

/** Request payload for CorrectFinancialData#load. */
class CorrectFinancialDataLoadMatch
{
    public string $rule_id;
}

/** Request payload for CorrectFinancialData#list. */
class CorrectFinancialDataListMatch
{
    public ?string $created_at = null;
    public ?string $custom_merchant_id = null;
    public ?array $data = null;
    public ?string $frequency = null;
    public ?bool $include_similar = null;
    public ?string $logo_feedback = null;
    public ?array $metadata = null;
    public ?string $name = null;
    public ?bool $online_or_billing_only = null;
    public ?string $operation_id = null;
    public ?string $reference_transaction_id = null;
    public ?array $rule_definition = null;
    public ?string $rule_type = null;
    public ?bool $similar = null;
    public ?string $suggested_logo = null;
    public ?string $suggested_url = null;
    public ?string $transaction_id = null;
    public ?array $transaction_ids = null;
}

/** Request payload for CorrectFinancialData#create. */
class CorrectFinancialDataCreateData
{
    public string $merchant_id;
    public string $created_at;
    public string $custom_merchant_id;
    public array $data;
    public string $frequency;
    public ?bool $include_similar = null;
    public string $logo_feedback;
    public array $metadata;
    public string $name;
    public ?bool $online_or_billing_only = null;
    public string $operation_id;
    public ?string $reference_transaction_id = null;
    public array $rule_definition;
    public string $rule_type;
    public ?bool $similar = null;
    public ?string $suggested_logo = null;
    public ?string $suggested_url = null;
    public string $transaction_id;
    public array $transaction_ids;
}

/** Request payload for CorrectFinancialData#remove. */
class CorrectFinancialDataRemoveMatch
{
    public string $rule_id;
}

/** CustomerMerchantCorrection entity data model. */
class CustomerMerchantCorrection
{
    public array $data;
    public array $metadata;
    public string $operation_id;
}

/** Request payload for CustomerMerchantCorrection#create. */
class CustomerMerchantCorrectionCreateData
{
    public array $data;
    public array $metadata;
    public string $operation_id;
}

/** Label entity data model. */
class Label
{
    public string $created_at;
    public string $id;
    public string $name;
    public string $updated_at;
}

/** Request payload for Label#create. */
class LabelCreateData
{
    public string $created_at;
    public string $id;
    public string $name;
    public string $updated_at;
}

/** Request payload for Label#update. */
class LabelUpdateData
{
    public string $id;
    public ?string $created_at = null;
    public ?string $name = null;
    public ?string $updated_at = null;
}

/** ListLabel entity data model. */
class ListLabel
{
    public string $created_at;
    public string $id;
    public string $name;
    public string $updated_at;
}

/** Request payload for ListLabel#list. */
class ListLabelListMatch
{
    public ?string $created_at = null;
    public ?string $id = null;
    public ?string $name = null;
    public ?string $updated_at = null;
}

/** ListTransaction entity data model. */
class ListTransaction
{
    public string $account_id;
    public array $amount;
    public ?array $client_attributes = null;
    public ?array $counterparty = null;
    public string $credit_debit_indicator;
    public string $date_time;
    public string $description;
    public ?array $enrichments = null;
    public ?array $labels = null;
    public ?string $merchant_category_code = null;
    public ?string $posted_date_time = null;
    public ?string $provider = null;
    public array $running_balance;
    public ?string $running_balance_credit_debit_indicator = null;
    public string $status;
    public string $suggested_description;
    public ?string $suggested_logo = null;
    public ?array $tags = null;
    public string $transaction_id;
    public ?array $transaction_type = null;
    public ?string $value_date_time = null;
}

/** Request payload for ListTransaction#list. */
class ListTransactionListMatch
{
    public ?string $account_id = null;
    public ?string $category_l1 = null;
    public ?string $category_l2 = null;
    public ?string $credit_debit_indicator = null;
    public ?string $date_from = null;
    public ?string $date_to = null;
    public ?array $exclude_tag = null;
    public ?array $include_label_id = null;
    public ?array $include_tag = null;
    public ?string $merchant = null;
    public ?int $page_size = null;
    public ?string $page_token = null;
    public ?string $status = null;
    public ?string $updated_after = null;
}

/** ManageFinancialData entity data model. */
class ManageFinancialData
{
    public string $label_id;
}

/** Request payload for ManageFinancialData#create. */
class ManageFinancialDataCreateData
{
    public string $transaction_id;
    public string $label_id;
}

/** Request payload for ManageFinancialData#remove. */
class ManageFinancialDataRemoveMatch
{
    public string $provider;
}

/** ManageTransactionLabel entity data model. */
class ManageTransactionLabel
{
    public ?string $id = null;
}

/** Request payload for ManageTransactionLabel#remove. */
class ManageTransactionLabelRemoveMatch
{
    public string $id;
}

/** RetrieveFinancialData entity data model. */
class RetrieveFinancialData
{
    public ?string $account_category = null;
    public string $account_id;
    public ?string $account_name = null;
    public ?string $account_type = null;
    public ?array $balances = null;
    public mixed $booked;
    public ?string $closed_at = null;
    public ?array $credit_lines = null;
    public string $currency;
    public array $data;
    public string $date;
    public ?array $details = null;
    public string $first_transaction_date;
    public ?string $frequency = null;
    public ?array $holder = null;
    public ?array $holders = null;
    public string $id;
    public ?array $identifiers = null;
    public string $last_transaction_date;
    public array $metadata;
    public string $name;
    public ?string $opening_date_time = null;
    public string $operation_id;
    public mixed $pending;
    public ?string $provider = null;
    public ?string $provider_display_name = null;
    public ?string $provider_logo = null;
    public string $reference;
    public ?string $restriction = null;
    public ?string $status = null;
    public string $suggested_name;
    public ?array $transaction_windows = null;
    public string $type;
    public ?string $usage_type = null;
}

/** Request payload for RetrieveFinancialData#load. */
class RetrieveFinancialDataLoadMatch
{
    public string $account_id;
}

/** Request payload for RetrieveFinancialData#list. */
class RetrieveFinancialDataListMatch
{
    public ?string $account_type = null;
    public ?string $currency = null;
    public ?bool $exclude_restricted = null;
    public ?string $holder_relationship_type = null;
    public ?bool $include_all_balance = null;
    public ?string $provider = null;
    public ?string $usage_type = null;
}

/** Similar entity data model. */
class Similar
{
    public array $data;
    public ?string $id = null;
    public array $metadata;
    public string $operation_id;
}

/** Request payload for Similar#load. */
class SimilarLoadMatch
{
    public string $id;
    public ?bool $exclude_source = null;
}

