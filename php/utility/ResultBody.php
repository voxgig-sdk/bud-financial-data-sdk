<?php
declare(strict_types=1);

// BudFinancialData SDK utility: result_body

class BudFinancialDataResultBody
{
    public static function call(BudFinancialDataContext $ctx): ?BudFinancialDataResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
