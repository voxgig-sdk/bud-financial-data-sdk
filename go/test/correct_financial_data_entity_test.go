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

func TestCorrectFinancialDataEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.CorrectFinancialData(nil)
		if ent == nil {
			t.Fatal("expected non-nil CorrectFinancialDataEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"correct_financial_data": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.CorrectFinancialData(nil).Stream("list", nil, nil) {
			seen = append(seen, item)
		}
		if len(seen) != 3 {
			t.Fatalf("expected 3 streamed items, got %d", len(seen))
		}

		// Inbound: streaming active -> yields each item from the feature iterator.
		hasStreaming := false
		if fm, ok := core.SharedConfig()["feature"].(map[string]any); ok {
			_, hasStreaming = fm["streaming"]
		}
		if hasStreaming {
			streamSdk := sdk.TestSDK(seed, map[string]any{
				"feature": map[string]any{"streaming": map[string]any{"active": true}},
			})
			var got []any
			for item := range streamSdk.CorrectFinancialData(nil).Stream("list", nil, nil) {
				if sub, ok := item.([]any); ok {
					got = append(got, sub...)
				} else {
					got = append(got, item)
				}
			}
			if len(got) != 3 {
				t.Fatalf("expected 3 items via streaming feature, got %d", len(got))
			}
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := correct_financial_dataBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "correct_financial_data." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set BUD_FINANCIAL_DATA_TEST_CORRECT_FINANCIAL_DATA_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		correctFinancialDataRef01Ent := client.CorrectFinancialData(nil)
		correctFinancialDataRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "correct_financial_data"}), "correct_financial_data_ref01"))
		correctFinancialDataRef01Data["merchant_id"] = setup.idmap["merchant01"]

		correctFinancialDataRef01DataResult, err := correctFinancialDataRef01Ent.Create(correctFinancialDataRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		correctFinancialDataRef01Data = core.ToMapAny(entityData(correctFinancialDataRef01DataResult))
		if correctFinancialDataRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

		// LIST
		correctFinancialDataRef01Match := map[string]any{}

		correctFinancialDataRef01ListResult, err := correctFinancialDataRef01Ent.List(correctFinancialDataRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, correctFinancialDataRef01ListOk := correctFinancialDataRef01ListResult.([]any)
		if !correctFinancialDataRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", correctFinancialDataRef01ListResult)
		}

		// LOAD
		correctFinancialDataRef01MatchDt0 := map[string]any{}
		correctFinancialDataRef01DataDt0Loaded, err := correctFinancialDataRef01Ent.Load(correctFinancialDataRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if correctFinancialDataRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}


		// LIST
		correctFinancialDataRef01MatchRt0 := map[string]any{}

		correctFinancialDataRef01ListRt0Result, err := correctFinancialDataRef01Ent.List(correctFinancialDataRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, correctFinancialDataRef01ListRt0Ok := correctFinancialDataRef01ListRt0Result.([]any)
		if !correctFinancialDataRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", correctFinancialDataRef01ListRt0Result)
		}

	})
}

func correct_financial_dataBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "correct_financial_data", "CorrectFinancialDataTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read correct_financial_data test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse correct_financial_data test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"correct_financial_data01", "correct_financial_data02", "correct_financial_data03", "merchant01"},
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
	entidEnvRaw := os.Getenv("BUD_FINANCIAL_DATA_TEST_CORRECT_FINANCIAL_DATA_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"BUD_FINANCIAL_DATA_TEST_CORRECT_FINANCIAL_DATA_ENTID": idmap,
		"BUD_FINANCIAL_DATA_TEST_LIVE":      "FALSE",
		"BUD_FINANCIAL_DATA_TEST_EXPLAIN":   "FALSE",
		"BUD_FINANCIAL_DATA_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["BUD_FINANCIAL_DATA_TEST_CORRECT_FINANCIAL_DATA_ENTID"])
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
