# BudFinancialData SDK utility: make_context

from budfinancialdata_sdk.core.context import BudFinancialDataContext


def make_context_util(ctxmap, basectx):
    return BudFinancialDataContext(ctxmap, basectx)
