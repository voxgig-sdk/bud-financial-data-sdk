package sdktest

import (
	"encoding/json"
	"fmt"
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

func TestLabelEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Label(nil)
		if ent == nil {
			t.Fatal("expected non-nil LabelEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := labelBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "update"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "label." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set BUD_FINANCIAL_DATA_TEST_LABEL_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		labelRef01Ent := client.Label(nil)
		labelRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "label"}), "label_ref01"))

		labelRef01DataResult, err := labelRef01Ent.Create(labelRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		labelRef01Data = core.ToMapAny(entityData(labelRef01DataResult))
		if labelRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if labelRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// UPDATE
		labelRef01DataUp0Up := map[string]any{
			"id": labelRef01Data["id"],
		}

		labelRef01MarkdefUp0Name := "created_at"
		labelRef01MarkdefUp0Value := fmt.Sprintf("Mark01-label_ref01_%d", setup.now)
		labelRef01DataUp0Up[labelRef01MarkdefUp0Name] = labelRef01MarkdefUp0Value

		labelRef01ResdataUp0Result, err := labelRef01Ent.Update(labelRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		labelRef01ResdataUp0 := core.ToMapAny(entityData(labelRef01ResdataUp0Result))
		if labelRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if labelRef01ResdataUp0["id"] != labelRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if labelRef01ResdataUp0[labelRef01MarkdefUp0Name] != labelRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", labelRef01MarkdefUp0Name, labelRef01ResdataUp0[labelRef01MarkdefUp0Name])
		}

	})
}

func labelBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "label", "LabelTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read label test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse label test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"label01", "label02", "label03"},
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
	entidEnvRaw := os.Getenv("BUD_FINANCIAL_DATA_TEST_LABEL_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"BUD_FINANCIAL_DATA_TEST_LABEL_ENTID": idmap,
		"BUD_FINANCIAL_DATA_TEST_LIVE":      "FALSE",
		"BUD_FINANCIAL_DATA_TEST_EXPLAIN":   "FALSE",
		"BUD_FINANCIAL_DATA_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["BUD_FINANCIAL_DATA_TEST_LABEL_ENTID"])
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
