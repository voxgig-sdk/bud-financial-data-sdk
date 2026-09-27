# CorrectFinancialData entity test

import json
import os
import time

import pytest

from budfinancialdata_sdk.utility.voxgig_struct import voxgig_struct as vs
from budfinancialdata_sdk import BudFinancialDataSDK
from budfinancialdata_sdk.core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestCorrectFinancialDataEntity:

    def test_should_create_instance(self):
        testsdk = BudFinancialDataSDK.test(None, None)
        ent = testsdk.CorrectFinancialData(None)
        assert ent is not None

    def test_should_stream(self):
        # Feature #4: the entity stream(action, ...) method runs the op
        # pipeline and yields result items. With the streaming feature active
        # it yields the feature's incremental output; otherwise it falls back
        # to the materialised list so stream always yields.
        seed = {
            "entity": {
                "correct_financial_data": {
                    "s1": {"id": "s1"},
                    "s2": {"id": "s2"},
                    "s3": {"id": "s3"},
                }
            }
        }

        # Fallback: streaming inactive -> yields the materialised list items.
        base = BudFinancialDataSDK.test(seed, None)
        seen = list(base.CorrectFinancialData(None).stream("list", None, None))
        assert len(seen) == 3

        # Inbound: streaming active -> yields each item from the feature.
        from budfinancialdata_sdk.config import shared_config
        cfg = shared_config()
        if isinstance(cfg.get("feature"), dict) and "streaming" in cfg["feature"]:
            sdk = BudFinancialDataSDK.test(
                seed, {"feature": {"streaming": {"active": True}}})
            got = []
            for item in sdk.CorrectFinancialData(None).stream("list", None, None):
                if isinstance(item, list):
                    got.extend(item)
                else:
                    got.append(item)
            assert len(got) == 3

    def test_should_run_basic_flow(self):
        setup = _correct_financial_data_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create", "list", "load", "remove"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "correct_financial_data." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set BUD_FINANCIAL_DATA_TEST_CORRECT_FINANCIAL_DATA_ENTID JSON to run live")
        client = setup["client"]

        # CREATE
        correct_financial_data_ref01_ent = client.CorrectFinancialData(None)
        correct_financial_data_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.correct_financial_data"), "correct_financial_data_ref01"))
        correct_financial_data_ref01_data["merchant_id"] = setup["idmap"]["merchant01"]

        correct_financial_data_ref01_data = helpers.to_map(runner.entity_data(correct_financial_data_ref01_ent.create(correct_financial_data_ref01_data, None)))
        assert correct_financial_data_ref01_data is not None

        # LIST
        correct_financial_data_ref01_match = {}

        correct_financial_data_ref01_list_result = correct_financial_data_ref01_ent.list(correct_financial_data_ref01_match, None)
        assert isinstance(correct_financial_data_ref01_list_result, list)

        # LOAD
        correct_financial_data_ref01_match_dt0 = {}
        correct_financial_data_ref01_data_dt0_loaded = correct_financial_data_ref01_ent.load(correct_financial_data_ref01_match_dt0, None)
        assert correct_financial_data_ref01_data_dt0_loaded is not None


        # LIST
        correct_financial_data_ref01_match_rt0 = {}

        correct_financial_data_ref01_list_rt0_result = correct_financial_data_ref01_ent.list(correct_financial_data_ref01_match_rt0, None)
        assert isinstance(correct_financial_data_ref01_list_rt0_result, list)



def _correct_financial_data_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/correct_financial_data/CorrectFinancialDataTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = BudFinancialDataSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["correct_financial_data01", "correct_financial_data02", "correct_financial_data03", "merchant01"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Detect ENTID env override before envOverride consumes it. When live
    # mode is on without a real override, the basic test runs against synthetic
    # IDs from the fixture and 4xx's. We surface this so the test can skip.
    _entid_env_raw = os.environ.get(
        "BUD_FINANCIAL_DATA_TEST_CORRECT_FINANCIAL_DATA_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "BUD_FINANCIAL_DATA_TEST_CORRECT_FINANCIAL_DATA_ENTID": idmap,
        "BUD_FINANCIAL_DATA_TEST_LIVE": "FALSE",
        "BUD_FINANCIAL_DATA_TEST_EXPLAIN": "FALSE",
        "BUD_FINANCIAL_DATA_APIKEY": "",
    })

    idmap_resolved = helpers.to_map(
        env.get("BUD_FINANCIAL_DATA_TEST_CORRECT_FINANCIAL_DATA_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

    if env.get("BUD_FINANCIAL_DATA_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            # FIRST, so the generated fields below win: sdk-test-control.json's
            # test.client.options adds to the live client, it does not
            # redirect it.
            runner.live_client_options(),
            {
                "apikey": env.get("BUD_FINANCIAL_DATA_APIKEY"),
            },
            extra or {},
        ])
        client = BudFinancialDataSDK(helpers.to_map(merged_opts))

    _live = env.get("BUD_FINANCIAL_DATA_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("BUD_FINANCIAL_DATA_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
