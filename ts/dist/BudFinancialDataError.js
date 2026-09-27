"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.BudFinancialDataError = void 0;
class BudFinancialDataError extends Error {
    isBudFinancialDataError = true;
    sdk = 'BudFinancialData';
    code;
    ctx;
    status = -1;
    // `err.notFound` rather than a magic number at every call site.
    get notFound() { return 404 === this.status; }
    constructor(code, msg, ctx) {
        super(msg);
        this.code = code;
        this.ctx = ctx;
    }
}
exports.BudFinancialDataError = BudFinancialDataError;
//# sourceMappingURL=BudFinancialDataError.js.map