<?php
declare(strict_types=1);

// BudFinancialData SDK utility: result_headers

class BudFinancialDataResultHeaders
{
    public static function call(BudFinancialDataContext $ctx): ?BudFinancialDataResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
