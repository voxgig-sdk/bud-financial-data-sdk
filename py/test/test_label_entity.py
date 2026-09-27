# Label entity test

import json
import os
import time

import pytest

from budfinancialdata_sdk.utility.voxgig_struct import voxgig_struct as vs
from budfinancialdata_sdk import BudFinancialDataSDK
from budfinancialdata_sdk.core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestLabelEntity:

    def test_should_create_instance(self):
        testsdk = BudFinancialDataSDK.test(None, None)
        ent = testsdk.Label(None)
        assert ent is not None

    def test_should_run_basic_flow(self):
        setup = _label_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create", "update"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "label." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set BUD_FINANCIAL_DATA_TEST_LABEL_ENTID JSON to run live")
        client = setup["client"]

        # CREATE
        label_ref01_ent = client.Label(None)
        label_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.label"), "label_ref01"))

        label_ref01_data = helpers.to_map(runner.entity_data(label_ref01_ent.create(label_ref01_data, None)))
        assert label_ref01_data is not None
        assert label_ref01_data["id"] is not None

        # UPDATE
        label_ref01_data_up0_up = {
            "id": label_ref01_data["id"],
        }

        label_ref01_markdef_up0_name = "created_at"
        label_ref01_markdef_up0_value = "Mark01-label_ref01_" + str(setup["now"])
        label_ref01_data_up0_up[label_ref01_markdef_up0_name] = label_ref01_markdef_up0_value

        label_ref01_resdata_up0 = helpers.to_map(runner.entity_data(label_ref01_ent.update(label_ref01_data_up0_up, None)))
        assert label_ref01_resdata_up0 is not None
        assert label_ref01_resdata_up0["id"] == label_ref01_data_up0_up["id"]
        assert label_ref01_resdata_up0[label_ref01_markdef_up0_name] == label_ref01_markdef_up0_value



def _label_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/label/LabelTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = BudFinancialDataSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["label01", "label02", "label03"],
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
        "BUD_FINANCIAL_DATA_TEST_LABEL_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "BUD_FINANCIAL_DATA_TEST_LABEL_ENTID": idmap,
        "BUD_FINANCIAL_DATA_TEST_LIVE": "FALSE",
        "BUD_FINANCIAL_DATA_TEST_EXPLAIN": "FALSE",
        "BUD_FINANCIAL_DATA_APIKEY": "",
    })

    idmap_resolved = helpers.to_map(
        env.get("BUD_FINANCIAL_DATA_TEST_LABEL_ENTID"))
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
