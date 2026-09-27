# BudFinancialData SDK feature factory

from budfinancialdata_sdk.feature.base_feature import BudFinancialDataBaseFeature
from budfinancialdata_sdk.feature.debug_feature import BudFinancialDataDebugFeature
from budfinancialdata_sdk.feature.idempotency_feature import BudFinancialDataIdempotencyFeature
from budfinancialdata_sdk.feature.metrics_feature import BudFinancialDataMetricsFeature
from budfinancialdata_sdk.feature.paging_feature import BudFinancialDataPagingFeature
from budfinancialdata_sdk.feature.ratelimit_feature import BudFinancialDataRatelimitFeature
from budfinancialdata_sdk.feature.retry_feature import BudFinancialDataRetryFeature
from budfinancialdata_sdk.feature.test_feature import BudFinancialDataTestFeature
from budfinancialdata_sdk.feature.timeout_feature import BudFinancialDataTimeoutFeature


_FEATURES = {
    "base": lambda: BudFinancialDataBaseFeature(),
    "debug": lambda: BudFinancialDataDebugFeature(),
    "idempotency": lambda: BudFinancialDataIdempotencyFeature(),
    "metrics": lambda: BudFinancialDataMetricsFeature(),
    "paging": lambda: BudFinancialDataPagingFeature(),
    "ratelimit": lambda: BudFinancialDataRatelimitFeature(),
    "retry": lambda: BudFinancialDataRetryFeature(),
    "test": lambda: BudFinancialDataTestFeature(),
    "timeout": lambda: BudFinancialDataTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
