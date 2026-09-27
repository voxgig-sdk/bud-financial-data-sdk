package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/bud-financial-data-sdk/go"
	"github.com/voxgig-sdk/bud-financial-data-sdk/go/core"

	vs "github.com/voxgig-sdk/bud-financial-data-sdk/go/utility/struct"
)

func TestCustomerMerchantCorrectionEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.CustomerMerchantCorrection(nil)
		if ent == nil {
			t.Fatal("expected non-nil CustomerMerchantCorrectionEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := customer_merchant_correctionBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "customer_merchant_correction." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set BUD_FINANCIAL_DATA_TEST_CUSTOMER_MERCHANT_CORRECTION_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		customerMerchantCorrectionRef01Ent := client.CustomerMerchantCorrection(nil)
		customerMerchantCorrectionRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "customer_merchant_correction"}), "customer_merchant_correction_ref01"))

		customerMerchantCorrectionRef01DataResult, err := customerMerchantCorrectionRef01Ent.Create(customerMerchantCorrectionRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		customerMerchantCorrectionRef01Data = core.ToMapAny(entityData(customerMerchantCorrectionRef01DataResult))
		if customerMerchantCorrectionRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

	})
}

func customer_merchant_correctionBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "customer_merchant_correction", "CustomerMerchantCorrectionTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read customer_merchant_correction test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse customer_merchant_correction test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"customer_merchant_correction01", "customer_merchant_correction02", "customer_merchant_correction03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("BUD_FINANCIAL_DATA_TEST_CUSTOMER_MERCHANT_CORRECTION_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"BUD_FINANCIAL_DATA_TEST_CUSTOMER_MERCHANT_CORRECTION_ENTID": idmap,
		"BUD_FINANCIAL_DATA_TEST_LIVE":      "FALSE",
		"BUD_FINANCIAL_DATA_TEST_EXPLAIN":   "FALSE",
		"BUD_FINANCIAL_DATA_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["BUD_FINANCIAL_DATA_TEST_CUSTOMER_MERCHANT_CORRECTION_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["BUD_FINANCIAL_DATA_TEST_LIVE"] == "TRUE" {
		// An empty map, not a nil one: Merge returns nil when its last entry
		// is nil, and BasicSetup is normally called with no extras - so a
		// bare nil silently discarded the apikey and server values below.
		extraOpts := extra
		if extraOpts == nil {
			extraOpts = map[string]any{}
		}

		mergedOpts := vs.Merge([]any{
			// liveClientOptions() FIRST, so the generated fields below win:
			// sdk-test-control.json's test.client.options adds to the live
			// client, it does not redirect it.
			liveClientOptions(),
			map[string]any{
				"apikey": env["BUD_FINANCIAL_DATA_APIKEY"],
			},
			extraOpts,
		})
		client = sdk.NewBudFinancialDataSDK(core.ToMapAny(mergedOpts))
	}

	live := env["BUD_FINANCIAL_DATA_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["BUD_FINANCIAL_DATA_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
