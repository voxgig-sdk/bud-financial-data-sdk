import { BudFinancialDataEntityBase } from '../BudFinancialDataEntityBase';
import type { BudFinancialDataSDK } from '../BudFinancialDataSDK';
import type { Control } from '../types';
import type { ListTransaction, ListTransactionListMatch } from '../BudFinancialDataTypes';
declare class ListTransactionEntity extends BudFinancialDataEntityBase<ListTransaction> {
    constructor(client: BudFinancialDataSDK, entopts: any);
    make(this: ListTransactionEntity): ListTransactionEntity;
    list(this: any, reqmatch?: ListTransactionListMatch, ctrl?: Control): Promise<ListTransactionEntity[]>;
}
export { ListTransactionEntity };
