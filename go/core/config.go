package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "BudFinancialData",
			"slug": "bud-financial-data",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"debug": map[string]any{
				"options": map[string]any{
					"active": false,
					"max": 100,
					"redact": []any{
						"authorization",
						"cookie",
						"set-cookie",
						"api-key",
						"apikey",
						"x-api-key",
						"idempotency-key",
					},
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"onEntry": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"idempotency": map[string]any{
				"options": map[string]any{
					"active": false,
					"header": "Idempotency-Key",
					"methods": []any{
						"POST",
						"PUT",
						"PATCH",
						"DELETE",
					},
					"ops": []any{
						"create",
						"update",
						"remove",
					},
				},
				"optspec": map[string]any{
					"keygen": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"metrics": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"paging": map[string]any{
				"options": map[string]any{
					"active": false,
					"afterVar": "after",
					"cursorParam": "cursor",
					"firstVar": "first",
					"limitParam": "limit",
					"pageParam": "page",
					"startPage": 1,
				},
				"optspec": map[string]any{
					"limit": "`$NUMBER`",
					"ops": "`$LIST`",
				},
				"strict": false,
				"transport": "none",
			},
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://api-sandbox.thisisbud.com",
			"auth": map[string]any{
				"prefix": "Bearer",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"correct_financial_data": map[string]any{},
				"customer_merchant_correction": map[string]any{},
				"label": map[string]any{},
				"list_label": map[string]any{},
				"list_transaction": map[string]any{},
				"manage_financial_data": map[string]any{},
				"manage_transaction_label": map[string]any{},
				"retrieve_financial_data": map[string]any{},
				"similar": map[string]any{},
			},
		},
		"entity": map[string]any{
			"correct_financial_data": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"short": "Date that the custom merchant was created, compliant with RFC3339.",
					},
					map[string]any{
						"name": "custom_merchant_id",
						"title": "Custom Merchant Id",
						"type": "`$STRING`",
						"req": true,
						"short": "UUID representing the custom merchant.",
						"format": "uuid",
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$ARRAY`",
						"req": true,
						"short": "List of merchant that matched the query.",
					},
					map[string]any{
						"name": "frequency",
						"title": "Frequency",
						"type": "`$STRING`",
						"req": true,
						"short": "The frequency to assign.",
					},
					map[string]any{
						"name": "include_similar",
						"title": "Include Similar",
						"type": "`$BOOLEAN`",
						"short": "Apply the correction to the whole group of transactions similar to the specified transaction, rather than just the transaction itself",
					},
					map[string]any{
						"name": "logo_feedback",
						"title": "Logo Feedback",
						"type": "`$STRING`",
						"req": true,
						"short": "The type of feedback for the merchant.",
					},
					map[string]any{
						"name": "metadata",
						"title": "Metadata",
						"type": "`$OBJECT`",
						"req": true,
						"short": "Metadata associated with the response schema",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"short": "The display name of the custom merchant to be created",
					},
					map[string]any{
						"name": "online_or_billing_only",
						"title": "Online Or Billing Only",
						"type": "`$BOOLEAN`",
						"op": map[string]any{
							"list": map[string]any{
								"req": true,
								"type": "`$BOOLEAN`",
							},
						},
						"short": "This is used to indicate if the custom merchant being created is an exclusively online or billing merchant.",
					},
					map[string]any{
						"name": "operation_id",
						"title": "Operation Id",
						"type": "`$STRING`",
						"req": true,
						"short": "A unique identifier/reference associated with a given endpoint/operation",
					},
					map[string]any{
						"name": "reference_transaction_id",
						"title": "Reference Transaction Id",
						"type": "`$STRING`",
						"short": "The transaction whose group the specified transactions should join or leave.",
					},
					map[string]any{
						"name": "rule_definition",
						"title": "Rule Definition",
						"type": "`$OBJECT`",
						"req": true,
						"short": "The definition of the rule, mirroring the request body that created it.",
					},
					map[string]any{
						"name": "rule_type",
						"title": "Rule Type",
						"type": "`$STRING`",
						"req": true,
						"short": "The type of correction rule",
					},
					map[string]any{
						"name": "similar",
						"title": "Similar",
						"type": "`$BOOLEAN`",
						"short": "Whether the specified transactions are similar to each other (or to `reference_transaction_id`, if supplied)",
					},
					map[string]any{
						"name": "suggested_logo",
						"title": "Suggested Logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "suggested_url",
						"title": "Suggested Url",
						"type": "`$STRING`",
						"op": map[string]any{
							"list": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "A suggested URL for the merchant, to help bud identify it and expand our merchant database.",
					},
					map[string]any{
						"name": "transaction_id",
						"title": "Transaction Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The unique identifier for the transaction",
					},
					map[string]any{
						"name": "transaction_ids",
						"title": "Transaction Ids",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The transactions the rule applies to",
					},
				},
				"name": "correct_financial_data",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/corrections/v2/categories",
								"segments": []any{
									map[string]any{
										"lit": "corrections",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "categories",
									},
								},
								"parts": []any{
									"corrections",
									"v2",
									"categories",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/corrections/v2/custom-merchants",
								"segments": []any{
									map[string]any{
										"lit": "corrections",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "custom-merchants",
									},
								},
								"parts": []any{
									"corrections",
									"v2",
									"custom-merchants",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/corrections/v3/regularity",
								"segments": []any{
									map[string]any{
										"lit": "corrections",
									},
									map[string]any{
										"lit": "v3",
									},
									map[string]any{
										"lit": "regularity",
									},
								},
								"parts": []any{
									"corrections",
									"v3",
									"regularity",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/corrections/v3/regularity/end",
								"segments": []any{
									map[string]any{
										"lit": "corrections",
									},
									map[string]any{
										"lit": "v3",
									},
									map[string]any{
										"lit": "regularity",
									},
									map[string]any{
										"lit": "end",
									},
								},
								"parts": []any{
									"corrections",
									"v3",
									"regularity",
									"end",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/corrections/v3/similar",
								"segments": []any{
									map[string]any{
										"lit": "corrections",
									},
									map[string]any{
										"lit": "v3",
									},
									map[string]any{
										"lit": "similar",
									},
								},
								"parts": []any{
									"corrections",
									"v3",
									"similar",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/corrections/v2/merchant-feedback/merchant/{merchant_id}",
								"segments": []any{
									map[string]any{
										"lit": "corrections",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "merchant-feedback",
									},
									map[string]any{
										"lit": "merchant",
									},
									map[string]any{
										"var": "merchant_id",
									},
								},
								"parts": []any{
									"corrections",
									"v2",
									"merchant-feedback",
									"merchant",
									"{merchant_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"merchant_id",
										"x_client_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/corrections/v2/custom-merchants",
								"segments": []any{
									map[string]any{
										"lit": "corrections",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "custom-merchants",
									},
								},
								"parts": []any{
									"corrections",
									"v2",
									"custom-merchants",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/corrections/v2/merchants/search/{merchant_query}",
								"segments": []any{
									map[string]any{
										"lit": "corrections",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "merchants",
									},
									map[string]any{
										"lit": "search",
									},
									map[string]any{
										"var": "merchant_query",
									},
								},
								"parts": []any{
									"corrections",
									"v2",
									"merchants",
									"search",
									"{merchant_query}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "merchant_query",
											"orig": "merchant_query",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"merchant_query",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/corrections/v3/rules/{rule_id}",
								"segments": []any{
									map[string]any{
										"lit": "corrections",
									},
									map[string]any{
										"lit": "v3",
									},
									map[string]any{
										"lit": "rules",
									},
									map[string]any{
										"var": "rule_id",
									},
								},
								"parts": []any{
									"corrections",
									"v3",
									"rules",
									"{rule_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "rule_id",
											"orig": "rule_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"rule_id",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/corrections/v3/rules/{rule_id}",
								"segments": []any{
									map[string]any{
										"lit": "corrections",
									},
									map[string]any{
										"lit": "v3",
									},
									map[string]any{
										"lit": "rules",
									},
									map[string]any{
										"var": "rule_id",
									},
								},
								"parts": []any{
									"corrections",
									"v3",
									"rules",
									"{rule_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "rule_id",
											"orig": "rule_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"rule_id",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"customer_merchant_correction": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "metadata",
						"title": "Metadata",
						"type": "`$OBJECT`",
						"req": true,
						"short": "Metadata associated with the merchant correction response schema",
					},
					map[string]any{
						"name": "operation_id",
						"title": "Operation Id",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "customer_merchant_correction",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/corrections/v2/merchants",
								"segments": []any{
									map[string]any{
										"lit": "corrections",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "merchants",
									},
								},
								"parts": []any{
									"corrections",
									"v2",
									"merchants",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"label": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"short": "RFC3339 timestamp at which the label was created.",
						"format": "date-time",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "Unique identifier of the label.",
						"format": "uuid",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"short": "Display name for the new label.",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"req": true,
						"short": "RFC3339 timestamp at which the label was last updated.",
						"format": "date-time",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "label",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/financial/v2/transactions/labels",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "transactions",
									},
									map[string]any{
										"lit": "labels",
									},
								},
								"parts": []any{
									"financial",
									"v2",
									"transactions",
									"labels",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/financial/v2/transactions/labels/{label_id}",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "transactions",
									},
									map[string]any{
										"lit": "labels",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"financial",
									"v2",
									"transactions",
									"labels",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"label_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "label_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"list_label": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"short": "RFC3339 timestamp at which the label was created.",
						"format": "date-time",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "Unique identifier of the label.",
						"format": "uuid",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"short": "Display name of the label.",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"req": true,
						"short": "RFC3339 timestamp at which the label was last updated.",
						"format": "date-time",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "list_label",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/financial/v2/transactions/labels",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "transactions",
									},
									map[string]any{
										"lit": "labels",
									},
								},
								"parts": []any{
									"financial",
									"v2",
									"transactions",
									"labels",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"list_transaction": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "account_id",
						"title": "Account Id",
						"type": "`$STRING`",
						"req": true,
						"short": "Identifier for the account associated with the transaction.",
					},
					map[string]any{
						"name": "amount",
						"title": "Amount",
						"type": "`$OBJECT`",
						"req": true,
						"short": "The monetary amount.",
					},
					map[string]any{
						"name": "client_attributes",
						"title": "Client Attributes",
						"type": "`$OBJECT`",
						"short": "An optional map of key-value string pairs that allows clients to attach arbitrary metadata to the transaction.",
					},
					map[string]any{
						"name": "counterparty",
						"title": "Counterparty",
						"type": "`$OBJECT`",
						"short": "An object containing details of the counterparty in this transaction",
					},
					map[string]any{
						"name": "credit_debit_indicator",
						"title": "Credit Debit Indicator",
						"type": "`$STRING`",
						"req": true,
						"short": "Credit/Debit Indicator",
					},
					map[string]any{
						"name": "date_time",
						"title": "Date Time",
						"type": "`$STRING`",
						"req": true,
						"short": "Date that the transaction occured compliant with RFC3339.",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
						"req": true,
						"short": "Description of the transaction.",
					},
					map[string]any{
						"name": "enrichments",
						"title": "Enrichments",
						"type": "`$OBJECT`",
						"short": "Contextual enrichments associated with a Transaction",
					},
					map[string]any{
						"name": "labels",
						"title": "Labels",
						"type": "`$ARRAY`",
						"short": "Customer-defined labels currently attached to this transaction.",
					},
					map[string]any{
						"name": "merchant_category_code",
						"title": "Merchant Category Code",
						"type": "`$STRING`",
						"short": "Merchant category code conforming to ISO 18245, related to the type of services or goods the merchant provides for the transaction.",
					},
					map[string]any{
						"name": "posted_date_time",
						"title": "Posted Date Time",
						"type": "`$STRING`",
						"short": "Date the assets involved in the transaction transferred compliant with RFC3339.",
						"deprecated": true,
					},
					map[string]any{
						"name": "provider",
						"title": "Provider",
						"type": "`$STRING`",
						"short": "Name of the transaction source provider.",
					},
					map[string]any{
						"name": "running_balance",
						"title": "Running Balance",
						"type": "`$OBJECT`",
						"req": true,
						"short": "The running balance for the account that the transaction takes place against",
					},
					map[string]any{
						"name": "running_balance_credit_debit_indicator",
						"title": "Running Balance Credit Debit Indicator",
						"type": "`$STRING`",
						"short": "Credit/Debit Indicator for the running balance field",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"req": true,
						"short": "Status of the transaction.",
					},
					map[string]any{
						"name": "suggested_description",
						"title": "Suggested Description",
						"type": "`$STRING`",
						"req": true,
						"short": "The description Bud suggests client apps show for a given transaction.",
					},
					map[string]any{
						"name": "suggested_logo",
						"title": "Suggested Logo",
						"type": "`$STRING`",
						"short": "The logo Bud suggests client apps show for a given transaction.",
					},
					map[string]any{
						"name": "tags",
						"title": "Tags",
						"type": "`$ARRAY`",
						"short": "A list of potential tags associated with the transaction after contextual enrichment, which can be used for filtering.",
					},
					map[string]any{
						"name": "transaction_id",
						"title": "Transaction Id",
						"type": "`$STRING`",
						"req": true,
						"short": "Unique identifier for the transaction.",
					},
					map[string]any{
						"name": "transaction_type",
						"title": "Transaction Type",
						"type": "`$OBJECT`",
						"short": "The code and a description of the transaction type.",
					},
					map[string]any{
						"name": "value_date_time",
						"title": "Value Date Time",
						"type": "`$STRING`",
						"short": "Date the assets involved in the transaction transferred compliant with RFC3339.",
					},
				},
				"name": "list_transaction",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/financial/v2/transactions",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "transactions",
									},
								},
								"parts": []any{
									"financial",
									"v2",
									"transactions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "query",
											"example": "5bd9ecaf31fc2d5172fd89cb61b89fc6",
										},
										map[string]any{
											"name": "category_l1",
											"orig": "category_l1",
											"type": "`$STRING`",
											"kind": "query",
											"example": "bills",
										},
										map[string]any{
											"name": "category_l2",
											"orig": "category_l2",
											"type": "`$STRING`",
											"kind": "query",
											"example": "tv_and_broadband",
										},
										map[string]any{
											"name": "credit_debit_indicator",
											"orig": "credit_debit_indicator",
											"type": "`$STRING`",
											"kind": "query",
											"example": "credit",
										},
										map[string]any{
											"name": "date_from",
											"orig": "date_from",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2022-05-28T00:00:00Z",
										},
										map[string]any{
											"name": "date_to",
											"orig": "date_to",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2022-05-29T00:00:00Z",
										},
										map[string]any{
											"name": "exclude_tag",
											"orig": "exclude_tag",
											"type": "`$ARRAY`",
											"kind": "query",
											"example": []any{
												"benefit",
											},
										},
										map[string]any{
											"name": "include_label_id",
											"orig": "include_label_id",
											"type": "`$ARRAY`",
											"kind": "query",
											"example": []any{
												"1f10e6de-0f5f-4d8c-987a-9635a9d1b8c6",
											},
										},
										map[string]any{
											"name": "include_tag",
											"orig": "include_tag",
											"type": "`$ARRAY`",
											"kind": "query",
											"example": []any{
												"regular-transaction",
											},
										},
										map[string]any{
											"name": "merchant",
											"orig": "merchant",
											"type": "`$STRING`",
											"kind": "query",
											"example": "netflix",
										},
										map[string]any{
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 100,
										},
										map[string]any{
											"name": "page_token",
											"orig": "page_token",
											"type": "`$STRING`",
											"kind": "query",
											"example": "eyJvZmZzZXQiOjEwMH0",
										},
										map[string]any{
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
											"kind": "query",
											"example": "booked",
										},
										map[string]any{
											"name": "updated_after",
											"orig": "updated_after",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2024-06-01T11:00:00Z",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"category_l1",
										"category_l2",
										"credit_debit_indicator",
										"date_from",
										"date_to",
										"exclude_tag",
										"include_label_id",
										"include_tag",
										"merchant",
										"page_size",
										"page_token",
										"status",
										"updated_after",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"manage_financial_data": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "label_id",
						"title": "Label Id",
						"type": "`$STRING`",
						"req": true,
						"short": "Identifier of the label to attach.",
						"format": "uuid",
					},
				},
				"name": "manage_financial_data",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/financial/v2/transactions/{transaction_id}/labels",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "transactions",
									},
									map[string]any{
										"var": "transaction_id",
									},
									map[string]any{
										"lit": "labels",
									},
								},
								"parts": []any{
									"financial",
									"v2",
									"transactions",
									"{transaction_id}",
									"labels",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "transaction_id",
											"orig": "transaction_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"transaction_id",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/financial/v2/transactions/{transaction_id}/labels/{label_id}",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "transactions",
									},
									map[string]any{
										"var": "transaction_id",
									},
									map[string]any{
										"lit": "labels",
									},
									map[string]any{
										"var": "label_id",
									},
								},
								"parts": []any{
									"financial",
									"v2",
									"transactions",
									"{transaction_id}",
									"labels",
									"{label_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "label_id",
											"orig": "label_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "transaction_id",
											"orig": "transaction_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"label_id",
										"transaction_id",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/financial/v2/accounts/{account_id}",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
								},
								"parts": []any{
									"financial",
									"v2",
									"accounts",
									"{account_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/v1/provider/{provider}",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "provider",
									},
									map[string]any{
										"var": "provider",
									},
								},
								"parts": []any{
									"v1",
									"provider",
									"{provider}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "provider",
											"orig": "provider",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"provider",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.label",
						},
					},
				},
			},
			"manage_transaction_label": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "manage_transaction_label",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/financial/v2/transactions/labels/{label_id}",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "transactions",
									},
									map[string]any{
										"lit": "labels",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"financial",
									"v2",
									"transactions",
									"labels",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"label_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "label_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"retrieve_financial_data": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "account_category",
						"title": "Account Category",
						"type": "`$STRING`",
						"short": "The account category that the account type is a part of Currently supported values include (but are not necessarily limited to): - `credit` - `depository` - `insurance` - `investment` - `loan` - `savings` - `other`",
					},
					map[string]any{
						"name": "account_id",
						"title": "Account Id",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"list": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "A unique id associated with the account.",
					},
					map[string]any{
						"name": "account_name",
						"title": "Account Name",
						"type": "`$STRING`",
						"short": "The name associated with the account, this is often a friendly name assigned to the account to make it easier to refer to.",
					},
					map[string]any{
						"name": "account_type",
						"title": "Account Type",
						"type": "`$STRING`",
						"short": "The type of account.",
					},
					map[string]any{
						"name": "balances",
						"title": "Balances",
						"type": "`$OBJECT`",
						"op": map[string]any{
							"list": map[string]any{
								"req": true,
								"type": "`$ARRAY`",
							},
						},
						"short": "The balances ingested for the account.",
					},
					map[string]any{
						"name": "booked",
						"title": "Booked",
						"type": "`$ANY`",
						"req": true,
						"short": "The current balance of the account",
					},
					map[string]any{
						"name": "closed_at",
						"title": "Closed At",
						"type": "`$STRING`",
						"short": "The datetime the account was closed in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339).",
						"format": "date-time",
					},
					map[string]any{
						"name": "credit_lines",
						"title": "Credit Lines",
						"type": "`$OBJECT`",
						"short": "The latest credit limit available for the account.",
					},
					map[string]any{
						"name": "currency",
						"title": "Currency",
						"type": "`$STRING`",
						"req": true,
						"short": "The three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) of the monetary amount.",
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "date",
						"title": "Date",
						"type": "`$STRING`",
						"req": true,
						"short": "The date of the given balance in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339).",
					},
					map[string]any{
						"name": "details",
						"title": "Details",
						"type": "`$ARRAY`",
						"short": "Provides more details for the authorised payment.",
					},
					map[string]any{
						"name": "first_transaction_date",
						"title": "First Transaction Date",
						"type": "`$STRING`",
						"req": true,
						"short": "The date of the first transaction that bud has stored for this account.",
						"format": "date-time",
					},
					map[string]any{
						"name": "frequency",
						"title": "Frequency",
						"type": "`$STRING`",
						"short": "The frequency of the authorised payment.",
					},
					map[string]any{
						"name": "holder",
						"title": "Holder",
						"type": "`$OBJECT`",
						"short": "The account holder.",
					},
					map[string]any{
						"name": "holders",
						"title": "Holders",
						"type": "`$ARRAY`",
						"short": "The account holder(s).",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The id for the authorised payment.",
					},
					map[string]any{
						"name": "identifiers",
						"title": "Identifiers",
						"type": "`$OBJECT`",
						"short": "These are wellknown fields which uniquely identify the account.",
					},
					map[string]any{
						"name": "last_transaction_date",
						"title": "Last Transaction Date",
						"type": "`$STRING`",
						"req": true,
						"short": "The date of the last transaction that bud has stored for this account.",
						"format": "date-time",
					},
					map[string]any{
						"name": "metadata",
						"title": "Metadata",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"short": "The name of the payment.",
					},
					map[string]any{
						"name": "opening_date_time",
						"title": "Opening Date Time",
						"type": "`$STRING`",
						"short": "The datetime the account was opened in the format [RFC 3339]",
						"format": "date-time",
					},
					map[string]any{
						"name": "operation_id",
						"title": "Operation Id",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "pending",
						"title": "Pending",
						"type": "`$ANY`",
						"req": true,
						"short": "The balance of the account if all pending transactions have settled.",
					},
					map[string]any{
						"name": "provider",
						"title": "Provider",
						"type": "`$STRING`",
						"short": "The account's provider, this is usually the bank or building society the account is held with.",
					},
					map[string]any{
						"name": "provider_display_name",
						"title": "Provider Display Name",
						"type": "`$STRING`",
						"short": "The display name for the account's provider.",
					},
					map[string]any{
						"name": "provider_logo",
						"title": "Provider Logo",
						"type": "`$STRING`",
						"short": "A link to an image for the accounts provider's logo.",
					},
					map[string]any{
						"name": "reference",
						"title": "Reference",
						"type": "`$STRING`",
						"req": true,
						"short": "The reference associated with the authorised payment.",
					},
					map[string]any{
						"name": "restriction",
						"title": "Restriction",
						"type": "`$STRING`",
						"short": "The restriction on the account, if any.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"op": map[string]any{
							"list": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "The status of the account.",
					},
					map[string]any{
						"name": "suggested_name",
						"title": "Suggested Name",
						"type": "`$STRING`",
						"req": true,
						"short": "The name Bud suggests client apps show for the account.",
					},
					map[string]any{
						"name": "transaction_windows",
						"title": "Transaction Windows",
						"type": "`$ARRAY`",
						"short": "The transaction windows indicate for which periods we have full coverage of transactions ingested for the account.",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$STRING`",
						"req": true,
						"short": "The type of the authorised payment.",
					},
					map[string]any{
						"name": "usage_type",
						"title": "Usage Type",
						"type": "`$STRING`",
						"short": "The intended usage of the account.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "retrieve_financial_data",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/financial/v3/accounts",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v3",
									},
									map[string]any{
										"lit": "accounts",
									},
								},
								"parts": []any{
									"financial",
									"v3",
									"accounts",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "account_type",
											"orig": "account_type",
											"type": "`$STRING`",
											"kind": "query",
											"example": "current_account",
										},
										map[string]any{
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
											"kind": "query",
											"example": "GBP",
										},
										map[string]any{
											"name": "exclude_restricted",
											"orig": "exclude_restricted",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": true,
										},
										map[string]any{
											"name": "holder_relationship_type",
											"orig": "holder_relationship_type",
											"type": "`$STRING`",
											"kind": "query",
											"example": "joint",
										},
										map[string]any{
											"name": "include_all_balance",
											"orig": "include_all_balance",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": true,
										},
										map[string]any{
											"name": "provider",
											"orig": "provider",
											"type": "`$STRING`",
											"kind": "query",
											"example": "BankOfBud",
										},
										map[string]any{
											"name": "usage_type",
											"orig": "usage_type",
											"type": "`$STRING`",
											"kind": "query",
											"example": "personal",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_type",
										"currency",
										"exclude_restricted",
										"holder_relationship_type",
										"include_all_balance",
										"provider",
										"usage_type",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/financial/v2/accounts",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "accounts",
									},
								},
								"parts": []any{
									"financial",
									"v2",
									"accounts",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "account_type",
											"orig": "account_type",
											"type": "`$STRING`",
											"kind": "query",
											"example": "current_account",
										},
										map[string]any{
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
											"kind": "query",
											"example": "GBP",
										},
										map[string]any{
											"name": "exclude_restricted",
											"orig": "exclude_restricted",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": true,
										},
										map[string]any{
											"name": "holder_relationship_type",
											"orig": "holder_relationship_type",
											"type": "`$STRING`",
											"kind": "query",
											"example": "joint",
										},
										map[string]any{
											"name": "provider",
											"orig": "provider",
											"type": "`$STRING`",
											"kind": "query",
											"example": "BankOfBud",
										},
										map[string]any{
											"name": "usage_type",
											"orig": "usage_type",
											"type": "`$STRING`",
											"kind": "query",
											"example": "personal",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_type",
										"currency",
										"exclude_restricted",
										"holder_relationship_type",
										"provider",
										"usage_type",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/financial/v2/accounts/{account_id}/balances",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "balances",
									},
								},
								"parts": []any{
									"financial",
									"v2",
									"accounts",
									"{account_id}",
									"balances",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "date_field",
											"orig": "date_field",
											"type": "`$STRING`",
											"kind": "query",
											"example": "date_time",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2019-02-29T12:23:00Z",
										},
										map[string]any{
											"name": "granularity",
											"orig": "granularity",
											"type": "`$STRING`",
											"kind": "query",
											"example": "daily",
										},
										map[string]any{
											"name": "to",
											"orig": "to",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2019-02-29T12:23:00Z",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"date_field",
										"from",
										"granularity",
										"to",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/financial/v3/accounts/{account_id}/balances",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v3",
									},
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
									map[string]any{
										"lit": "balances",
									},
								},
								"parts": []any{
									"financial",
									"v3",
									"accounts",
									"{account_id}",
									"balances",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "date_field",
											"orig": "date_field",
											"type": "`$STRING`",
											"kind": "query",
											"example": "date_time",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2019-02-29T12:23:00Z",
										},
										map[string]any{
											"name": "granularity",
											"orig": "granularity",
											"type": "`$STRING`",
											"kind": "query",
											"example": "daily",
										},
										map[string]any{
											"name": "to",
											"orig": "to",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2019-02-29T12:23:00Z",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"date_field",
										"from",
										"granularity",
										"to",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/financial/v2/authorised-payments",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "authorised-payments",
									},
								},
								"parts": []any{
									"financial",
									"v2",
									"authorised-payments",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 100,
										},
										map[string]any{
											"name": "page_token",
											"orig": "page_token",
											"type": "`$STRING`",
											"kind": "query",
											"example": "eyJvZmZzZXQiOjEwMH0",
										},
										map[string]any{
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
											"kind": "query",
											"example": "active",
										},
										map[string]any{
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
											"kind": "query",
											"example": "direct_debit",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"page_size",
										"page_token",
										"status",
										"type",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/financial/v2/balances",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "balances",
									},
								},
								"parts": []any{
									"financial",
									"v2",
									"balances",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "date_field",
											"orig": "date_field",
											"type": "`$STRING`",
											"kind": "query",
											"example": "date_time",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2019-02-29T12:23:00Z",
										},
										map[string]any{
											"name": "granularity",
											"orig": "granularity",
											"type": "`$STRING`",
											"kind": "query",
											"example": "daily",
										},
										map[string]any{
											"name": "to",
											"orig": "to",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2019-02-29T12:23:00Z",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"date_field",
										"from",
										"granularity",
										"to",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/financial/v3/balances",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v3",
									},
									map[string]any{
										"lit": "balances",
									},
								},
								"parts": []any{
									"financial",
									"v3",
									"balances",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "date_field",
											"orig": "date_field",
											"type": "`$STRING`",
											"kind": "query",
											"example": "date_time",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2019-02-29T12:23:00Z",
										},
										map[string]any{
											"name": "granularity",
											"orig": "granularity",
											"type": "`$STRING`",
											"kind": "query",
											"example": "daily",
										},
										map[string]any{
											"name": "to",
											"orig": "to",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2019-02-29T12:23:00Z",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"date_field",
										"from",
										"granularity",
										"to",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/financial/v3/accounts/transaction-dates",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v3",
									},
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"lit": "transaction-dates",
									},
								},
								"parts": []any{
									"financial",
									"v3",
									"accounts",
									"transaction-dates",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/financial/v2/accounts/{account_id}",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
								},
								"parts": []any{
									"financial",
									"v2",
									"accounts",
									"{account_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/financial/v3/accounts/{account_id}",
								"segments": []any{
									map[string]any{
										"lit": "financial",
									},
									map[string]any{
										"lit": "v3",
									},
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"var": "account_id",
									},
								},
								"parts": []any{
									"financial",
									"v3",
									"accounts",
									"{account_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "account_id",
											"orig": "account_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"similar": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metadata",
						"title": "Metadata",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "operation_id",
						"title": "Operation Id",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "similar",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/corrections/v2/categories/similar/{transaction_id}",
								"segments": []any{
									map[string]any{
										"lit": "corrections",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "categories",
									},
									map[string]any{
										"lit": "similar",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"corrections",
									"v2",
									"categories",
									"similar",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"transaction_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "transaction_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "exclude_source",
											"orig": "exclude_source",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"exclude_source",
										"id",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/corrections/v2/merchants/similar/{transaction_id}",
								"segments": []any{
									map[string]any{
										"lit": "corrections",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "merchants",
									},
									map[string]any{
										"lit": "similar",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"corrections",
									"v2",
									"merchants",
									"similar",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"transaction_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
										map[string]any{
											"name": "x_customer_idempotent_identifier",
											"orig": "x_customer_idempotent_identifier",
											"type": "`$STRING`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_customer_secret",
											"orig": "x_customer_secret",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "transaction_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "exclude_source",
											"orig": "exclude_source",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"exclude_source",
										"id",
										"x_client_id",
										"x_customer_id",
										"x_customer_idempotent_identifier",
										"x_customer_secret",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "debug":
		if NewDebugFeatureFunc != nil {
			return NewDebugFeatureFunc()
		}
	case "idempotency":
		if NewIdempotencyFeatureFunc != nil {
			return NewIdempotencyFeatureFunc()
		}
	case "metrics":
		if NewMetricsFeatureFunc != nil {
			return NewMetricsFeatureFunc()
		}
	case "paging":
		if NewPagingFeatureFunc != nil {
			return NewPagingFeatureFunc()
		}
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
