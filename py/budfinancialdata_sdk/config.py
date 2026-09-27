# BudFinancialData SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "BudFinancialData",
            "slug": "bud-financial-data",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "debug": {
        "options": {
          "active": False,
          "max": 100,
          "redact": [
            "authorization",
            "cookie",
            "set-cookie",
            "api-key",
            "apikey",
            "x-api-key",
            "idempotency-key",
          ],
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "onEntry": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "idempotency": {
        "options": {
          "active": False,
          "header": "Idempotency-Key",
          "methods": [
            "POST",
            "PUT",
            "PATCH",
            "DELETE",
          ],
          "ops": [
            "create",
            "update",
            "remove",
          ],
        },
        "optspec": {
          "keygen": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "metrics": {
        "options": {
          "active": False,
        },
        "optspec": {
          "now": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "paging": {
        "options": {
          "active": False,
          "afterVar": "after",
          "cursorParam": "cursor",
          "firstVar": "first",
          "limitParam": "limit",
          "pageParam": "page",
          "startPage": 1,
        },
        "optspec": {
          "limit": "`$NUMBER`",
          "ops": "`$LIST`",
        },
        "strict": False,
        "transport": "none",
      },
            "ratelimit": {
        "options": {
          "active": False,
          "burst": 5,
          "rate": 5,
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "retry": {
        "options": {
          "active": False,
          "factor": 2,
          "maxDelay": 2000,
          "minDelay": 50,
          "retries": 2,
          "statuses": [
            408,
            425,
            429,
            500,
            502,
            503,
            504,
          ],
        },
        "optspec": {
          "jitter": "`$BOOLEAN`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "test": {
        "options": {
          "active": False,
        },
        "optspec": {
          "entity": "`$MAP`",
          "net": "`$MAP`",
        },
        "strict": False,
        "transport": "base",
      },
            "timeout": {
        "options": {
          "active": False,
          "ms": 30000,
        },
        "optspec": {
          "clearTimer": "`$FUNCTION`",
          "setTimer": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
        },
        "options": {
            "base": "https://api-sandbox.thisisbud.com",
            "auth": {
                "prefix": "Bearer",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "correct_financial_data": {},
                "customer_merchant_correction": {},
                "label": {},
                "list_label": {},
                "list_transaction": {},
                "manage_financial_data": {},
                "manage_transaction_label": {},
                "retrieve_financial_data": {},
                "similar": {},
            },
        },
        "entity": {
      "correct_financial_data": {
        "fields": [
          {
            "name": "created_at",
            "title": "Created At",
            "type": "`$STRING`",
            "req": True,
            "short": "Date that the custom merchant was created, compliant with RFC3339.",
          },
          {
            "name": "custom_merchant_id",
            "title": "Custom Merchant Id",
            "type": "`$STRING`",
            "req": True,
            "short": "UUID representing the custom merchant.",
            "format": "uuid",
          },
          {
            "name": "data",
            "title": "Data",
            "type": "`$ARRAY`",
            "req": True,
            "short": "List of merchant that matched the query.",
          },
          {
            "name": "frequency",
            "title": "Frequency",
            "type": "`$STRING`",
            "req": True,
            "short": "The frequency to assign.",
          },
          {
            "name": "include_similar",
            "title": "Include Similar",
            "type": "`$BOOLEAN`",
            "short": "Apply the correction to the whole group of transactions similar to the specified transaction, rather than just the transaction itself",
          },
          {
            "name": "logo_feedback",
            "title": "Logo Feedback",
            "type": "`$STRING`",
            "req": True,
            "short": "The type of feedback for the merchant.",
          },
          {
            "name": "metadata",
            "title": "Metadata",
            "type": "`$OBJECT`",
            "req": True,
            "short": "Metadata associated with the response schema",
          },
          {
            "name": "name",
            "title": "Name",
            "type": "`$STRING`",
            "req": True,
            "short": "The display name of the custom merchant to be created",
          },
          {
            "name": "online_or_billing_only",
            "title": "Online Or Billing Only",
            "type": "`$BOOLEAN`",
            "op": {
              "list": {
                "req": True,
                "type": "`$BOOLEAN`",
              },
            },
            "short": "This is used to indicate if the custom merchant being created is an exclusively online or billing merchant.",
          },
          {
            "name": "operation_id",
            "title": "Operation Id",
            "type": "`$STRING`",
            "req": True,
            "short": "A unique identifier/reference associated with a given endpoint/operation",
          },
          {
            "name": "reference_transaction_id",
            "title": "Reference Transaction Id",
            "type": "`$STRING`",
            "short": "The transaction whose group the specified transactions should join or leave.",
          },
          {
            "name": "rule_definition",
            "title": "Rule Definition",
            "type": "`$OBJECT`",
            "req": True,
            "short": "The definition of the rule, mirroring the request body that created it.",
          },
          {
            "name": "rule_type",
            "title": "Rule Type",
            "type": "`$STRING`",
            "req": True,
            "short": "The type of correction rule",
          },
          {
            "name": "similar",
            "title": "Similar",
            "type": "`$BOOLEAN`",
            "short": "Whether the specified transactions are similar to each other (or to `reference_transaction_id`, if supplied)",
          },
          {
            "name": "suggested_logo",
            "title": "Suggested Logo",
            "type": "`$STRING`",
          },
          {
            "name": "suggested_url",
            "title": "Suggested Url",
            "type": "`$STRING`",
            "op": {
              "list": {
                "req": True,
                "type": "`$STRING`",
              },
            },
            "short": "A suggested URL for the merchant, to help bud identify it and expand our merchant database.",
          },
          {
            "name": "transaction_id",
            "title": "Transaction Id",
            "type": "`$STRING`",
            "req": True,
            "short": "The unique identifier for the transaction",
          },
          {
            "name": "transaction_ids",
            "title": "Transaction Ids",
            "type": "`$ARRAY`",
            "req": True,
            "short": "The transactions the rule applies to",
          },
        ],
        "name": "correct_financial_data",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "kind": "http",
                "method": "POST",
                "orig": "/corrections/v2/categories",
                "segments": [
                  {
                    "lit": "corrections",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "categories",
                  },
                ],
                "parts": [
                  "corrections",
                  "v2",
                  "categories",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "POST",
                "orig": "/corrections/v2/custom-merchants",
                "segments": [
                  {
                    "lit": "corrections",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "custom-merchants",
                  },
                ],
                "parts": [
                  "corrections",
                  "v2",
                  "custom-merchants",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "POST",
                "orig": "/corrections/v3/regularity",
                "segments": [
                  {
                    "lit": "corrections",
                  },
                  {
                    "lit": "v3",
                  },
                  {
                    "lit": "regularity",
                  },
                ],
                "parts": [
                  "corrections",
                  "v3",
                  "regularity",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "POST",
                "orig": "/corrections/v3/regularity/end",
                "segments": [
                  {
                    "lit": "corrections",
                  },
                  {
                    "lit": "v3",
                  },
                  {
                    "lit": "regularity",
                  },
                  {
                    "lit": "end",
                  },
                ],
                "parts": [
                  "corrections",
                  "v3",
                  "regularity",
                  "end",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "POST",
                "orig": "/corrections/v3/similar",
                "segments": [
                  {
                    "lit": "corrections",
                  },
                  {
                    "lit": "v3",
                  },
                  {
                    "lit": "similar",
                  },
                ],
                "parts": [
                  "corrections",
                  "v3",
                  "similar",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "POST",
                "orig": "/corrections/v2/merchant-feedback/merchant/{merchant_id}",
                "segments": [
                  {
                    "lit": "corrections",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "merchant-feedback",
                  },
                  {
                    "lit": "merchant",
                  },
                  {
                    "var": "merchant_id",
                  },
                ],
                "parts": [
                  "corrections",
                  "v2",
                  "merchant-feedback",
                  "merchant",
                  "{merchant_id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                  "params": [
                    {
                      "name": "merchant_id",
                      "orig": "merchant_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "merchant_id",
                    "x_client_id",
                  ],
                },
              },
            ],
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/corrections/v2/custom-merchants",
                "segments": [
                  {
                    "lit": "corrections",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "custom-merchants",
                  },
                ],
                "parts": [
                  "corrections",
                  "v2",
                  "custom-merchants",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                  ],
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/corrections/v2/merchants/search/{merchant_query}",
                "segments": [
                  {
                    "lit": "corrections",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "merchants",
                  },
                  {
                    "lit": "search",
                  },
                  {
                    "var": "merchant_query",
                  },
                ],
                "parts": [
                  "corrections",
                  "v2",
                  "merchants",
                  "search",
                  "{merchant_query}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "merchant_query",
                      "orig": "merchant_query",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "merchant_query",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/corrections/v3/rules/{rule_id}",
                "segments": [
                  {
                    "lit": "corrections",
                  },
                  {
                    "lit": "v3",
                  },
                  {
                    "lit": "rules",
                  },
                  {
                    "var": "rule_id",
                  },
                ],
                "parts": [
                  "corrections",
                  "v3",
                  "rules",
                  "{rule_id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "rule_id",
                      "orig": "rule_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "rule_id",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                  ],
                },
              },
            ],
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "kind": "http",
                "method": "DELETE",
                "orig": "/corrections/v3/rules/{rule_id}",
                "segments": [
                  {
                    "lit": "corrections",
                  },
                  {
                    "lit": "v3",
                  },
                  {
                    "lit": "rules",
                  },
                  {
                    "var": "rule_id",
                  },
                ],
                "parts": [
                  "corrections",
                  "v3",
                  "rules",
                  "{rule_id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "rule_id",
                      "orig": "rule_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "rule_id",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "customer_merchant_correction": {
        "fields": [
          {
            "name": "data",
            "title": "Data",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "metadata",
            "title": "Metadata",
            "type": "`$OBJECT`",
            "req": True,
            "short": "Metadata associated with the merchant correction response schema",
          },
          {
            "name": "operation_id",
            "title": "Operation Id",
            "type": "`$STRING`",
            "req": True,
          },
        ],
        "name": "customer_merchant_correction",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "kind": "http",
                "method": "POST",
                "orig": "/corrections/v2/merchants",
                "segments": [
                  {
                    "lit": "corrections",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "merchants",
                  },
                ],
                "parts": [
                  "corrections",
                  "v2",
                  "merchants",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "label": {
        "fields": [
          {
            "name": "created_at",
            "title": "Created At",
            "type": "`$STRING`",
            "req": True,
            "short": "RFC3339 timestamp at which the label was created.",
            "format": "date-time",
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
            "req": True,
            "short": "Unique identifier of the label.",
            "format": "uuid",
          },
          {
            "name": "name",
            "title": "Name",
            "type": "`$STRING`",
            "req": True,
            "short": "Display name for the new label.",
          },
          {
            "name": "updated_at",
            "title": "Updated At",
            "type": "`$STRING`",
            "req": True,
            "short": "RFC3339 timestamp at which the label was last updated.",
            "format": "date-time",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "label",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "kind": "http",
                "method": "POST",
                "orig": "/financial/v2/transactions/labels",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "transactions",
                  },
                  {
                    "lit": "labels",
                  },
                ],
                "parts": [
                  "financial",
                  "v2",
                  "transactions",
                  "labels",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
            ],
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "kind": "http",
                "method": "PATCH",
                "orig": "/financial/v2/transactions/labels/{label_id}",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "transactions",
                  },
                  {
                    "lit": "labels",
                  },
                  {
                    "var": "id",
                  },
                ],
                "parts": [
                  "financial",
                  "v2",
                  "transactions",
                  "labels",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "label_id": "id",
                  },
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "id",
                      "orig": "label_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "list_label": {
        "fields": [
          {
            "name": "created_at",
            "title": "Created At",
            "type": "`$STRING`",
            "req": True,
            "short": "RFC3339 timestamp at which the label was created.",
            "format": "date-time",
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
            "req": True,
            "short": "Unique identifier of the label.",
            "format": "uuid",
          },
          {
            "name": "name",
            "title": "Name",
            "type": "`$STRING`",
            "req": True,
            "short": "Display name of the label.",
          },
          {
            "name": "updated_at",
            "title": "Updated At",
            "type": "`$STRING`",
            "req": True,
            "short": "RFC3339 timestamp at which the label was last updated.",
            "format": "date-time",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "list_label",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/financial/v2/transactions/labels",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "transactions",
                  },
                  {
                    "lit": "labels",
                  },
                ],
                "parts": [
                  "financial",
                  "v2",
                  "transactions",
                  "labels",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "list_transaction": {
        "fields": [
          {
            "name": "account_id",
            "title": "Account Id",
            "type": "`$STRING`",
            "req": True,
            "short": "Identifier for the account associated with the transaction.",
          },
          {
            "name": "amount",
            "title": "Amount",
            "type": "`$OBJECT`",
            "req": True,
            "short": "The monetary amount.",
          },
          {
            "name": "client_attributes",
            "title": "Client Attributes",
            "type": "`$OBJECT`",
            "short": "An optional map of key-value string pairs that allows clients to attach arbitrary metadata to the transaction.",
          },
          {
            "name": "counterparty",
            "title": "Counterparty",
            "type": "`$OBJECT`",
            "short": "An object containing details of the counterparty in this transaction",
          },
          {
            "name": "credit_debit_indicator",
            "title": "Credit Debit Indicator",
            "type": "`$STRING`",
            "req": True,
            "short": "Credit/Debit Indicator",
          },
          {
            "name": "date_time",
            "title": "Date Time",
            "type": "`$STRING`",
            "req": True,
            "short": "Date that the transaction occured compliant with RFC3339.",
          },
          {
            "name": "description",
            "title": "Description",
            "type": "`$STRING`",
            "req": True,
            "short": "Description of the transaction.",
          },
          {
            "name": "enrichments",
            "title": "Enrichments",
            "type": "`$OBJECT`",
            "short": "Contextual enrichments associated with a Transaction",
          },
          {
            "name": "labels",
            "title": "Labels",
            "type": "`$ARRAY`",
            "short": "Customer-defined labels currently attached to this transaction.",
          },
          {
            "name": "merchant_category_code",
            "title": "Merchant Category Code",
            "type": "`$STRING`",
            "short": "Merchant category code conforming to ISO 18245, related to the type of services or goods the merchant provides for the transaction.",
          },
          {
            "name": "posted_date_time",
            "title": "Posted Date Time",
            "type": "`$STRING`",
            "short": "Date the assets involved in the transaction transferred compliant with RFC3339.",
            "deprecated": True,
          },
          {
            "name": "provider",
            "title": "Provider",
            "type": "`$STRING`",
            "short": "Name of the transaction source provider.",
          },
          {
            "name": "running_balance",
            "title": "Running Balance",
            "type": "`$OBJECT`",
            "req": True,
            "short": "The running balance for the account that the transaction takes place against",
          },
          {
            "name": "running_balance_credit_debit_indicator",
            "title": "Running Balance Credit Debit Indicator",
            "type": "`$STRING`",
            "short": "Credit/Debit Indicator for the running balance field",
          },
          {
            "name": "status",
            "title": "Status",
            "type": "`$STRING`",
            "req": True,
            "short": "Status of the transaction.",
          },
          {
            "name": "suggested_description",
            "title": "Suggested Description",
            "type": "`$STRING`",
            "req": True,
            "short": "The description Bud suggests client apps show for a given transaction.",
          },
          {
            "name": "suggested_logo",
            "title": "Suggested Logo",
            "type": "`$STRING`",
            "short": "The logo Bud suggests client apps show for a given transaction.",
          },
          {
            "name": "tags",
            "title": "Tags",
            "type": "`$ARRAY`",
            "short": "A list of potential tags associated with the transaction after contextual enrichment, which can be used for filtering.",
          },
          {
            "name": "transaction_id",
            "title": "Transaction Id",
            "type": "`$STRING`",
            "req": True,
            "short": "Unique identifier for the transaction.",
          },
          {
            "name": "transaction_type",
            "title": "Transaction Type",
            "type": "`$OBJECT`",
            "short": "The code and a description of the transaction type.",
          },
          {
            "name": "value_date_time",
            "title": "Value Date Time",
            "type": "`$STRING`",
            "short": "Date the assets involved in the transaction transferred compliant with RFC3339.",
          },
        ],
        "name": "list_transaction",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/financial/v2/transactions",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "transactions",
                  },
                ],
                "parts": [
                  "financial",
                  "v2",
                  "transactions",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "account_id",
                      "orig": "account_id",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "5bd9ecaf31fc2d5172fd89cb61b89fc6",
                    },
                    {
                      "name": "category_l1",
                      "orig": "category_l1",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "bills",
                    },
                    {
                      "name": "category_l2",
                      "orig": "category_l2",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "tv_and_broadband",
                    },
                    {
                      "name": "credit_debit_indicator",
                      "orig": "credit_debit_indicator",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "credit",
                    },
                    {
                      "name": "date_from",
                      "orig": "date_from",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "2022-05-28T00:00:00Z",
                    },
                    {
                      "name": "date_to",
                      "orig": "date_to",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "2022-05-29T00:00:00Z",
                    },
                    {
                      "name": "exclude_tag",
                      "orig": "exclude_tag",
                      "type": "`$ARRAY`",
                      "kind": "query",
                      "example": [
                        "benefit",
                      ],
                    },
                    {
                      "name": "include_label_id",
                      "orig": "include_label_id",
                      "type": "`$ARRAY`",
                      "kind": "query",
                      "example": [
                        "1f10e6de-0f5f-4d8c-987a-9635a9d1b8c6",
                      ],
                    },
                    {
                      "name": "include_tag",
                      "orig": "include_tag",
                      "type": "`$ARRAY`",
                      "kind": "query",
                      "example": [
                        "regular-transaction",
                      ],
                    },
                    {
                      "name": "merchant",
                      "orig": "merchant",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "netflix",
                    },
                    {
                      "name": "page_size",
                      "orig": "page_size",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 100,
                    },
                    {
                      "name": "page_token",
                      "orig": "page_token",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "eyJvZmZzZXQiOjEwMH0",
                    },
                    {
                      "name": "status",
                      "orig": "status",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "booked",
                    },
                    {
                      "name": "updated_after",
                      "orig": "updated_after",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "2024-06-01T11:00:00Z",
                    },
                  ],
                },
                "select": {
                  "exist": [
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
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "manage_financial_data": {
        "fields": [
          {
            "name": "label_id",
            "title": "Label Id",
            "type": "`$STRING`",
            "req": True,
            "short": "Identifier of the label to attach.",
            "format": "uuid",
          },
        ],
        "name": "manage_financial_data",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "kind": "http",
                "method": "POST",
                "orig": "/financial/v2/transactions/{transaction_id}/labels",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "transactions",
                  },
                  {
                    "var": "transaction_id",
                  },
                  {
                    "lit": "labels",
                  },
                ],
                "parts": [
                  "financial",
                  "v2",
                  "transactions",
                  "{transaction_id}",
                  "labels",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "transaction_id",
                      "orig": "transaction_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "transaction_id",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
            ],
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "kind": "http",
                "method": "DELETE",
                "orig": "/financial/v2/transactions/{transaction_id}/labels/{label_id}",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "transactions",
                  },
                  {
                    "var": "transaction_id",
                  },
                  {
                    "lit": "labels",
                  },
                  {
                    "var": "label_id",
                  },
                ],
                "parts": [
                  "financial",
                  "v2",
                  "transactions",
                  "{transaction_id}",
                  "labels",
                  "{label_id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "label_id",
                      "orig": "label_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                    {
                      "name": "transaction_id",
                      "orig": "transaction_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "label_id",
                    "transaction_id",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "DELETE",
                "orig": "/financial/v2/accounts/{account_id}",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "accounts",
                  },
                  {
                    "var": "account_id",
                  },
                ],
                "parts": [
                  "financial",
                  "v2",
                  "accounts",
                  "{account_id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "account_id",
                      "orig": "account_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "account_id",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "DELETE",
                "orig": "/v1/provider/{provider}",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "provider",
                  },
                  {
                    "var": "provider",
                  },
                ],
                "parts": [
                  "v1",
                  "provider",
                  "{provider}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "provider",
                      "orig": "provider",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "provider",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [
            [
              "$.main.kit.entity.label",
            ],
          ],
        },
      },
      "manage_transaction_label": {
        "fields": [
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "manage_transaction_label",
        "op": {
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "kind": "http",
                "method": "DELETE",
                "orig": "/financial/v2/transactions/labels/{label_id}",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "transactions",
                  },
                  {
                    "lit": "labels",
                  },
                  {
                    "var": "id",
                  },
                ],
                "parts": [
                  "financial",
                  "v2",
                  "transactions",
                  "labels",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "label_id": "id",
                  },
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "id",
                      "orig": "label_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "retrieve_financial_data": {
        "fields": [
          {
            "name": "account_category",
            "title": "Account Category",
            "type": "`$STRING`",
            "short": "The account category that the account type is a part of Currently supported values include (but are not necessarily limited to): - `credit` - `depository` - `insurance` - `investment` - `loan` - `savings` - `other`",
          },
          {
            "name": "account_id",
            "title": "Account Id",
            "type": "`$STRING`",
            "req": True,
            "op": {
              "list": {
                "type": "`$STRING`",
              },
            },
            "short": "A unique id associated with the account.",
          },
          {
            "name": "account_name",
            "title": "Account Name",
            "type": "`$STRING`",
            "short": "The name associated with the account, this is often a friendly name assigned to the account to make it easier to refer to.",
          },
          {
            "name": "account_type",
            "title": "Account Type",
            "type": "`$STRING`",
            "short": "The type of account.",
          },
          {
            "name": "balances",
            "title": "Balances",
            "type": "`$OBJECT`",
            "op": {
              "list": {
                "req": True,
                "type": "`$ARRAY`",
              },
            },
            "short": "The balances ingested for the account.",
          },
          {
            "name": "booked",
            "title": "Booked",
            "type": "`$ANY`",
            "req": True,
            "short": "The current balance of the account",
          },
          {
            "name": "closed_at",
            "title": "Closed At",
            "type": "`$STRING`",
            "short": "The datetime the account was closed in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339).",
            "format": "date-time",
          },
          {
            "name": "credit_lines",
            "title": "Credit Lines",
            "type": "`$OBJECT`",
            "short": "The latest credit limit available for the account.",
          },
          {
            "name": "currency",
            "title": "Currency",
            "type": "`$STRING`",
            "req": True,
            "short": "The three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) of the monetary amount.",
          },
          {
            "name": "data",
            "title": "Data",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "date",
            "title": "Date",
            "type": "`$STRING`",
            "req": True,
            "short": "The date of the given balance in the format [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339).",
          },
          {
            "name": "details",
            "title": "Details",
            "type": "`$ARRAY`",
            "short": "Provides more details for the authorised payment.",
          },
          {
            "name": "first_transaction_date",
            "title": "First Transaction Date",
            "type": "`$STRING`",
            "req": True,
            "short": "The date of the first transaction that bud has stored for this account.",
            "format": "date-time",
          },
          {
            "name": "frequency",
            "title": "Frequency",
            "type": "`$STRING`",
            "short": "The frequency of the authorised payment.",
          },
          {
            "name": "holder",
            "title": "Holder",
            "type": "`$OBJECT`",
            "short": "The account holder.",
          },
          {
            "name": "holders",
            "title": "Holders",
            "type": "`$ARRAY`",
            "short": "The account holder(s).",
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
            "req": True,
            "short": "The id for the authorised payment.",
          },
          {
            "name": "identifiers",
            "title": "Identifiers",
            "type": "`$OBJECT`",
            "short": "These are wellknown fields which uniquely identify the account.",
          },
          {
            "name": "last_transaction_date",
            "title": "Last Transaction Date",
            "type": "`$STRING`",
            "req": True,
            "short": "The date of the last transaction that bud has stored for this account.",
            "format": "date-time",
          },
          {
            "name": "metadata",
            "title": "Metadata",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "name",
            "title": "Name",
            "type": "`$STRING`",
            "req": True,
            "short": "The name of the payment.",
          },
          {
            "name": "opening_date_time",
            "title": "Opening Date Time",
            "type": "`$STRING`",
            "short": "The datetime the account was opened in the format [RFC 3339]",
            "format": "date-time",
          },
          {
            "name": "operation_id",
            "title": "Operation Id",
            "type": "`$STRING`",
            "req": True,
          },
          {
            "name": "pending",
            "title": "Pending",
            "type": "`$ANY`",
            "req": True,
            "short": "The balance of the account if all pending transactions have settled.",
          },
          {
            "name": "provider",
            "title": "Provider",
            "type": "`$STRING`",
            "short": "The account's provider, this is usually the bank or building society the account is held with.",
          },
          {
            "name": "provider_display_name",
            "title": "Provider Display Name",
            "type": "`$STRING`",
            "short": "The display name for the account's provider.",
          },
          {
            "name": "provider_logo",
            "title": "Provider Logo",
            "type": "`$STRING`",
            "short": "A link to an image for the accounts provider's logo.",
          },
          {
            "name": "reference",
            "title": "Reference",
            "type": "`$STRING`",
            "req": True,
            "short": "The reference associated with the authorised payment.",
          },
          {
            "name": "restriction",
            "title": "Restriction",
            "type": "`$STRING`",
            "short": "The restriction on the account, if any.",
          },
          {
            "name": "status",
            "title": "Status",
            "type": "`$STRING`",
            "op": {
              "list": {
                "req": True,
                "type": "`$STRING`",
              },
            },
            "short": "The status of the account.",
          },
          {
            "name": "suggested_name",
            "title": "Suggested Name",
            "type": "`$STRING`",
            "req": True,
            "short": "The name Bud suggests client apps show for the account.",
          },
          {
            "name": "transaction_windows",
            "title": "Transaction Windows",
            "type": "`$ARRAY`",
            "short": "The transaction windows indicate for which periods we have full coverage of transactions ingested for the account.",
          },
          {
            "name": "type",
            "title": "Type",
            "type": "`$STRING`",
            "req": True,
            "short": "The type of the authorised payment.",
          },
          {
            "name": "usage_type",
            "title": "Usage Type",
            "type": "`$STRING`",
            "short": "The intended usage of the account.",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "retrieve_financial_data",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/financial/v3/accounts",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v3",
                  },
                  {
                    "lit": "accounts",
                  },
                ],
                "parts": [
                  "financial",
                  "v3",
                  "accounts",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "account_type",
                      "orig": "account_type",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "current_account",
                    },
                    {
                      "name": "currency",
                      "orig": "currency",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "GBP",
                    },
                    {
                      "name": "exclude_restricted",
                      "orig": "exclude_restricted",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                      "example": True,
                    },
                    {
                      "name": "holder_relationship_type",
                      "orig": "holder_relationship_type",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "joint",
                    },
                    {
                      "name": "include_all_balance",
                      "orig": "include_all_balance",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                      "example": True,
                    },
                    {
                      "name": "provider",
                      "orig": "provider",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "BankOfBud",
                    },
                    {
                      "name": "usage_type",
                      "orig": "usage_type",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "personal",
                    },
                  ],
                },
                "select": {
                  "exist": [
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
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/financial/v2/accounts",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "accounts",
                  },
                ],
                "parts": [
                  "financial",
                  "v2",
                  "accounts",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "account_type",
                      "orig": "account_type",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "current_account",
                    },
                    {
                      "name": "currency",
                      "orig": "currency",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "GBP",
                    },
                    {
                      "name": "exclude_restricted",
                      "orig": "exclude_restricted",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                      "example": True,
                    },
                    {
                      "name": "holder_relationship_type",
                      "orig": "holder_relationship_type",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "joint",
                    },
                    {
                      "name": "provider",
                      "orig": "provider",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "BankOfBud",
                    },
                    {
                      "name": "usage_type",
                      "orig": "usage_type",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "personal",
                    },
                  ],
                },
                "select": {
                  "exist": [
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
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/financial/v2/accounts/{account_id}/balances",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "accounts",
                  },
                  {
                    "var": "account_id",
                  },
                  {
                    "lit": "balances",
                  },
                ],
                "parts": [
                  "financial",
                  "v2",
                  "accounts",
                  "{account_id}",
                  "balances",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "account_id",
                      "orig": "account_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                  "query": [
                    {
                      "name": "date_field",
                      "orig": "date_field",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "date_time",
                    },
                    {
                      "name": "from",
                      "orig": "from",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "2019-02-29T12:23:00Z",
                    },
                    {
                      "name": "granularity",
                      "orig": "granularity",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "daily",
                    },
                    {
                      "name": "to",
                      "orig": "to",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "2019-02-29T12:23:00Z",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "account_id",
                    "date_field",
                    "from",
                    "granularity",
                    "to",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/financial/v3/accounts/{account_id}/balances",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v3",
                  },
                  {
                    "lit": "accounts",
                  },
                  {
                    "var": "account_id",
                  },
                  {
                    "lit": "balances",
                  },
                ],
                "parts": [
                  "financial",
                  "v3",
                  "accounts",
                  "{account_id}",
                  "balances",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "account_id",
                      "orig": "account_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                  "query": [
                    {
                      "name": "date_field",
                      "orig": "date_field",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "date_time",
                    },
                    {
                      "name": "from",
                      "orig": "from",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "2019-02-29T12:23:00Z",
                    },
                    {
                      "name": "granularity",
                      "orig": "granularity",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "daily",
                    },
                    {
                      "name": "to",
                      "orig": "to",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "2019-02-29T12:23:00Z",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "account_id",
                    "date_field",
                    "from",
                    "granularity",
                    "to",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/financial/v2/authorised-payments",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "authorised-payments",
                  },
                ],
                "parts": [
                  "financial",
                  "v2",
                  "authorised-payments",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "account_id",
                      "orig": "account_id",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "page_size",
                      "orig": "page_size",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 100,
                    },
                    {
                      "name": "page_token",
                      "orig": "page_token",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "eyJvZmZzZXQiOjEwMH0",
                    },
                    {
                      "name": "status",
                      "orig": "status",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "active",
                    },
                    {
                      "name": "type",
                      "orig": "type",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "direct_debit",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "account_id",
                    "page_size",
                    "page_token",
                    "status",
                    "type",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/financial/v2/balances",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "balances",
                  },
                ],
                "parts": [
                  "financial",
                  "v2",
                  "balances",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "date_field",
                      "orig": "date_field",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "date_time",
                    },
                    {
                      "name": "from",
                      "orig": "from",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "2019-02-29T12:23:00Z",
                    },
                    {
                      "name": "granularity",
                      "orig": "granularity",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "daily",
                    },
                    {
                      "name": "to",
                      "orig": "to",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "2019-02-29T12:23:00Z",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "date_field",
                    "from",
                    "granularity",
                    "to",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/financial/v3/balances",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v3",
                  },
                  {
                    "lit": "balances",
                  },
                ],
                "parts": [
                  "financial",
                  "v3",
                  "balances",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "date_field",
                      "orig": "date_field",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "date_time",
                    },
                    {
                      "name": "from",
                      "orig": "from",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "2019-02-29T12:23:00Z",
                    },
                    {
                      "name": "granularity",
                      "orig": "granularity",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "daily",
                    },
                    {
                      "name": "to",
                      "orig": "to",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "2019-02-29T12:23:00Z",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "date_field",
                    "from",
                    "granularity",
                    "to",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/financial/v3/accounts/transaction-dates",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v3",
                  },
                  {
                    "lit": "accounts",
                  },
                  {
                    "lit": "transaction-dates",
                  },
                ],
                "parts": [
                  "financial",
                  "v3",
                  "accounts",
                  "transaction-dates",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/financial/v2/accounts/{account_id}",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "accounts",
                  },
                  {
                    "var": "account_id",
                  },
                ],
                "parts": [
                  "financial",
                  "v2",
                  "accounts",
                  "{account_id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "account_id",
                      "orig": "account_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "account_id",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/financial/v3/accounts/{account_id}",
                "segments": [
                  {
                    "lit": "financial",
                  },
                  {
                    "lit": "v3",
                  },
                  {
                    "lit": "accounts",
                  },
                  {
                    "var": "account_id",
                  },
                ],
                "parts": [
                  "financial",
                  "v3",
                  "accounts",
                  "{account_id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "account_id",
                      "orig": "account_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "account_id",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "similar": {
        "fields": [
          {
            "name": "data",
            "title": "Data",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
          },
          {
            "name": "metadata",
            "title": "Metadata",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "operation_id",
            "title": "Operation Id",
            "type": "`$STRING`",
            "req": True,
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "similar",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/corrections/v2/categories/similar/{transaction_id}",
                "segments": [
                  {
                    "lit": "corrections",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "categories",
                  },
                  {
                    "lit": "similar",
                  },
                  {
                    "var": "id",
                  },
                ],
                "parts": [
                  "corrections",
                  "v2",
                  "categories",
                  "similar",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "transaction_id": "id",
                  },
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "id",
                      "orig": "transaction_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                  "query": [
                    {
                      "name": "exclude_source",
                      "orig": "exclude_source",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "exclude_source",
                    "id",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/corrections/v2/merchants/similar/{transaction_id}",
                "segments": [
                  {
                    "lit": "corrections",
                  },
                  {
                    "lit": "v2",
                  },
                  {
                    "lit": "merchants",
                  },
                  {
                    "lit": "similar",
                  },
                  {
                    "var": "id",
                  },
                ],
                "parts": [
                  "corrections",
                  "v2",
                  "merchants",
                  "similar",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "transaction_id": "id",
                  },
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_customer_id",
                      "orig": "x_customer_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
                    },
                    {
                      "name": "x_customer_idempotent_identifier",
                      "orig": "x_customer_idempotent_identifier",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                    {
                      "name": "x_customer_secret",
                      "orig": "x_customer_secret",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "params": [
                    {
                      "name": "id",
                      "orig": "transaction_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                  "query": [
                    {
                      "name": "exclude_source",
                      "orig": "exclude_source",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "exclude_source",
                    "id",
                    "x_client_id",
                    "x_customer_id",
                    "x_customer_idempotent_identifier",
                    "x_customer_secret",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
