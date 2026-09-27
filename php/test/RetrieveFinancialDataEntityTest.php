<?php
declare(strict_types=1);

// RetrieveFinancialData entity test

require_once __DIR__ . '/../budfinancialdata_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class RetrieveFinancialDataEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = BudFinancialDataSDK::test(null, null);
        $ent = $testsdk->RetrieveFinancialData(null);
        $this->assertNotNull($ent);
    }

    // Feature #4: the entity stream(action, ...) method runs the op pipeline
    // and yields result items. With the streaming feature active it yields the
    // feature's incremental output; otherwise it falls back to the materialised
    // list so stream always yields.
    public function test_stream(): void
    {
        $seed = [
            "entity" => [
                "retrieve_financial_data" => [
                    "s1" => ["id" => "s1"],
                    "s2" => ["id" => "s2"],
                    "s3" => ["id" => "s3"],
                ],
            ],
        ];

        // Fallback: streaming inactive -> yields the materialised list items.
        $base = BudFinancialDataSDK::test($seed, null);
        $seen = iterator_to_array($base->RetrieveFinancialData(null)->stream("list", null, null), false);
        $this->assertCount(3, $seen);

        // Inbound: streaming active -> yields each item from the feature.
        $cfg = BudFinancialDataConfig::shared_config();
        if (isset($cfg["feature"]) && is_array($cfg["feature"]) && isset($cfg["feature"]["streaming"])) {
            $sdk = BudFinancialDataSDK::test($seed, ["feature" => ["streaming" => ["active" => true]]]);
            $got = [];
            foreach ($sdk->RetrieveFinancialData(null)->stream("list", null, null) as $item) {
                if (is_array($item) && array_is_list($item)) {
                    foreach ($item as $sub) {
                        $got[] = $sub;
                    }
                } else {
                    $got[] = $item;
                }
            }
            $this->assertCount(3, $got);
        }
    }

    public function test_basic_flow(): void
    {
        $setup = retrieve_financial_data_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["list", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "retrieve_financial_data." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set BUD_FINANCIAL_DATA_TEST_RETRIEVE_FINANCIAL_DATA_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $retrieve_financial_data_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.retrieve_financial_data")));
        $retrieve_financial_data_ref01_data = null;
        if (count($retrieve_financial_data_ref01_data_raw) > 0) {
            $retrieve_financial_data_ref01_data = Helpers::to_map($retrieve_financial_data_ref01_data_raw[0][1]);
        }

        // LIST
        $retrieve_financial_data_ref01_ent = $client->RetrieveFinancialData(null);
        $retrieve_financial_data_ref01_match = [];

        $retrieve_financial_data_ref01_list_result = $retrieve_financial_data_ref01_ent->list($retrieve_financial_data_ref01_match, null);
        $this->assertIsArray($retrieve_financial_data_ref01_list_result);

        // LOAD
        $retrieve_financial_data_ref01_match_dt0 = [
            "id" => $retrieve_financial_data_ref01_data["id"],
        ];
        $retrieve_financial_data_ref01_data_dt0_loaded = $retrieve_financial_data_ref01_ent->load($retrieve_financial_data_ref01_match_dt0, null);
        $retrieve_financial_data_ref01_data_dt0_load_result = Helpers::to_map(is_object($retrieve_financial_data_ref01_data_dt0_loaded) && method_exists($retrieve_financial_data_ref01_data_dt0_loaded, 'data_get') ? $retrieve_financial_data_ref01_data_dt0_loaded->data_get() : $retrieve_financial_data_ref01_data_dt0_loaded);
        $this->assertNotNull($retrieve_financial_data_ref01_data_dt0_load_result);
        $this->assertEquals($retrieve_financial_data_ref01_data_dt0_load_result["id"], $retrieve_financial_data_ref01_data["id"]);

    }
}

function retrieve_financial_data_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/retrieve_financial_data/RetrieveFinancialDataTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = BudFinancialDataSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["retrieve_financial_data01", "retrieve_financial_data02", "retrieve_financial_data03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("BUD_FINANCIAL_DATA_TEST_RETRIEVE_FINANCIAL_DATA_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "BUD_FINANCIAL_DATA_TEST_RETRIEVE_FINANCIAL_DATA_ENTID" => $idmap,
        "BUD_FINANCIAL_DATA_TEST_LIVE" => "FALSE",
        "BUD_FINANCIAL_DATA_TEST_EXPLAIN" => "FALSE",
        "BUD_FINANCIAL_DATA_APIKEY" => "",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["BUD_FINANCIAL_DATA_TEST_RETRIEVE_FINANCIAL_DATA_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["BUD_FINANCIAL_DATA_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            Runner::live_client_options(),
            [
                "apikey" => $env["BUD_FINANCIAL_DATA_APIKEY"],
            ],
            // ismap, not a plain "?? []" default: an empty PHP array is a
            // LIST, and a non-map later entry REPLACES the accumulated map in
            // merge - so the no-extras call discarded live_client_options()
            // and the apikey/server map above it.
            Vs::ismap($extra) ? $extra : new \stdClass(),
        ]);
        // "?? []" because merge legitimately answers with a stdClass when every
        // contributing entry is an EMPTY map - an SDK with no apikey and no
        // server variables generates an empty middle entry, so that is the
        // common case, not the edge one. to_map returns null for a non-array by
        // design, and the constructor takes a non-nullable array, so without the
        // fallback every such SDK died on "must be of type array, null given"
        // the moment live mode was switched on. Offline mode never reaches this
        // branch, which is why the offline suite stayed green.
        $client = new BudFinancialDataSDK(Helpers::to_map($merged_opts) ?? []);
    }

    $live = $env["BUD_FINANCIAL_DATA_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["BUD_FINANCIAL_DATA_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
