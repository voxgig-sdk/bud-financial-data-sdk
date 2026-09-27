import { Context } from './Context';
declare class BudFinancialDataError extends Error {
    isBudFinancialDataError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { BudFinancialDataError };
