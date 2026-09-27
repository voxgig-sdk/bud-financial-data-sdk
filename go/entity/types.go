// Typed models for the BudFinancialData SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/bud-financial-data-sdk/go/core"
)

// CorrectFinancialData is the typed data model for the correct_financial_data entity.
type CorrectFinancialData struct {
}

// CorrectFinancialDataLoadMatch is the typed request payload for CorrectFinancialData.LoadTyped.
type CorrectFinancialDataLoadMatch struct {
	RuleId string `json:"rule_id"`
}

// CorrectFinancialDataListMatch is the typed request payload for CorrectFinancialData.ListTyped.
type CorrectFinancialDataListMatch struct {
	CreatedAt *string `json:"created_at,omitempty"`
	CustomMerchantId *string `json:"custom_merchant_id,omitempty"`
	Data *[]any `json:"data,omitempty"`
	Frequency *string `json:"frequency,omitempty"`
	IncludeSimilar *bool `json:"include_similar,omitempty"`
	LogoFeedback *string `json:"logo_feedback,omitempty"`
	Metadata *map[string]any `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
	OnlineOrBillingOnly *bool `json:"online_or_billing_only,omitempty"`
	OperationId *string `json:"operation_id,omitempty"`
	ReferenceTransactionId *string `json:"reference_transaction_id,omitempty"`
	RuleDefinition *map[string]any `json:"rule_definition,omitempty"`
	RuleType *string `json:"rule_type,omitempty"`
	Similar *bool `json:"similar,omitempty"`
	SuggestedLogo *string `json:"suggested_logo,omitempty"`
	SuggestedUrl *string `json:"suggested_url,omitempty"`
	TransactionId *string `json:"transaction_id,omitempty"`
	TransactionIds *[]any `json:"transaction_ids,omitempty"`
}

// CorrectFinancialDataCreateData is the typed request payload for CorrectFinancialData.CreateTyped.
type CorrectFinancialDataCreateData struct {
	MerchantId string `json:"merchant_id"`
	CreatedAt string `json:"created_at"`
	CustomMerchantId string `json:"custom_merchant_id"`
	Data []any `json:"data"`
	Frequency string `json:"frequency"`
	IncludeSimilar *bool `json:"include_similar,omitempty"`
	LogoFeedback string `json:"logo_feedback"`
	Metadata map[string]any `json:"metadata"`
	Name string `json:"name"`
	OnlineOrBillingOnly *bool `json:"online_or_billing_only,omitempty"`
	OperationId string `json:"operation_id"`
	ReferenceTransactionId *string `json:"reference_transaction_id,omitempty"`
	RuleDefinition map[string]any `json:"rule_definition"`
	RuleType string `json:"rule_type"`
	Similar *bool `json:"similar,omitempty"`
	SuggestedLogo *string `json:"suggested_logo,omitempty"`
	SuggestedUrl *string `json:"suggested_url,omitempty"`
	TransactionId string `json:"transaction_id"`
	TransactionIds []any `json:"transaction_ids"`
}

// CorrectFinancialDataRemoveMatch is the typed request payload for CorrectFinancialData.RemoveTyped.
type CorrectFinancialDataRemoveMatch struct {
	RuleId string `json:"rule_id"`
}

// CustomerMerchantCorrection is the typed data model for the customer_merchant_correction entity.
type CustomerMerchantCorrection struct {
}

// CustomerMerchantCorrectionCreateData is the typed request payload for CustomerMerchantCorrection.CreateTyped.
type CustomerMerchantCorrectionCreateData struct {
	Data []any `json:"data"`
	Metadata map[string]any `json:"metadata"`
	OperationId string `json:"operation_id"`
}

// Label is the typed data model for the label entity.
type Label struct {
}

// LabelCreateData is the typed request payload for Label.CreateTyped.
type LabelCreateData struct {
	CreatedAt string `json:"created_at"`
	Id string `json:"id"`
	Name string `json:"name"`
	UpdatedAt string `json:"updated_at"`
}

// LabelUpdateData is the typed request payload for Label.UpdateTyped.
type LabelUpdateData struct {
	Id string `json:"id"`
	CreatedAt *string `json:"created_at,omitempty"`
	Name *string `json:"name,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// ListLabel is the typed data model for the list_label entity.
type ListLabel struct {
}

// ListLabelListMatch is the typed request payload for ListLabel.ListTyped.
type ListLabelListMatch struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// ListTransaction is the typed data model for the list_transaction entity.
type ListTransaction struct {
}

// ListTransactionListMatch is the typed request payload for ListTransaction.ListTyped.
type ListTransactionListMatch struct {
	AccountId *string `json:"account_id,omitempty"`
	CategoryL1 *string `json:"category_l1,omitempty"`
	CategoryL2 *string `json:"category_l2,omitempty"`
	CreditDebitIndicator *string `json:"credit_debit_indicator,omitempty"`
	DateFrom *string `json:"date_from,omitempty"`
	DateTo *string `json:"date_to,omitempty"`
	ExcludeTag *[]any `json:"exclude_tag,omitempty"`
	IncludeLabelId *[]any `json:"include_label_id,omitempty"`
	IncludeTag *[]any `json:"include_tag,omitempty"`
	Merchant *string `json:"merchant,omitempty"`
	PageSize *int `json:"page_size,omitempty"`
	PageToken *string `json:"page_token,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAfter *string `json:"updated_after,omitempty"`
}

// ManageFinancialData is the typed data model for the manage_financial_data entity.
type ManageFinancialData struct {
}

// ManageFinancialDataCreateData is the typed request payload for ManageFinancialData.CreateTyped.
type ManageFinancialDataCreateData struct {
	TransactionId string `json:"transaction_id"`
	LabelId string `json:"label_id"`
}

// ManageFinancialDataRemoveMatch is the typed request payload for ManageFinancialData.RemoveTyped.
type ManageFinancialDataRemoveMatch struct {
	Provider string `json:"provider"`
}

// ManageTransactionLabel is the typed data model for the manage_transaction_label entity.
type ManageTransactionLabel struct {
}

// ManageTransactionLabelRemoveMatch is the typed request payload for ManageTransactionLabel.RemoveTyped.
type ManageTransactionLabelRemoveMatch struct {
	Id string `json:"id"`
}

// RetrieveFinancialData is the typed data model for the retrieve_financial_data entity.
type RetrieveFinancialData struct {
}

// RetrieveFinancialDataLoadMatch is the typed request payload for RetrieveFinancialData.LoadTyped.
type RetrieveFinancialDataLoadMatch struct {
	AccountId string `json:"account_id"`
}

// RetrieveFinancialDataListMatch is the typed request payload for RetrieveFinancialData.ListTyped.
type RetrieveFinancialDataListMatch struct {
	AccountType *string `json:"account_type,omitempty"`
	Currency *string `json:"currency,omitempty"`
	ExcludeRestricted *bool `json:"exclude_restricted,omitempty"`
	HolderRelationshipType *string `json:"holder_relationship_type,omitempty"`
	IncludeAllBalance *bool `json:"include_all_balance,omitempty"`
	Provider *string `json:"provider,omitempty"`
	UsageType *string `json:"usage_type,omitempty"`
}

// Similar is the typed data model for the similar entity.
type Similar struct {
}

// SimilarLoadMatch is the typed request payload for Similar.LoadTyped.
type SimilarLoadMatch struct {
	Id string `json:"id"`
	ExcludeSource *bool `json:"exclude_source,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
