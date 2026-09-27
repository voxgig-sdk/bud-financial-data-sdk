<?php
declare(strict_types=1);

// BudFinancialData SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class BudFinancialDataMakeContext
{
    public static function call(array $ctxmap, ?BudFinancialDataContext $basectx): BudFinancialDataContext
    {
        return new BudFinancialDataContext($ctxmap, $basectx);
    }
}
