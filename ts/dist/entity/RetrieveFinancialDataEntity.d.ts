import { BudFinancialDataEntityBase } from '../BudFinancialDataEntityBase';
import type { BudFinancialDataSDK } from '../BudFinancialDataSDK';
import type { Control } from '../types';
import type { RetrieveFinancialData, RetrieveFinancialDataLoadMatch, RetrieveFinancialDataListMatch } from '../BudFinancialDataTypes';
declare class RetrieveFinancialDataEntity extends BudFinancialDataEntityBase<RetrieveFinancialData> {
    constructor(client: BudFinancialDataSDK, entopts: any);
    make(this: RetrieveFinancialDataEntity): RetrieveFinancialDataEntity;
    load(this: any, reqmatch?: RetrieveFinancialDataLoadMatch, ctrl?: Control): Promise<RetrieveFinancialDataEntity>;
    list(this: any, reqmatch?: RetrieveFinancialDataListMatch, ctrl?: Control): Promise<RetrieveFinancialDataEntity[]>;
}
export { RetrieveFinancialDataEntity };
