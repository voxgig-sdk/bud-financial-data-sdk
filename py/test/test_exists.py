# BudFinancialData SDK exists test

import pytest
from budfinancialdata_sdk import BudFinancialDataSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = BudFinancialDataSDK.test(None, None)
        assert testsdk is not None
