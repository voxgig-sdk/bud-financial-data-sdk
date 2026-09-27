import { BudFinancialDataEntityBase } from '../BudFinancialDataEntityBase';
import type { BudFinancialDataSDK } from '../BudFinancialDataSDK';
import type { Control } from '../types';
import type { CorrectFinancialData, CorrectFinancialDataLoadMatch, CorrectFinancialDataListMatch, CorrectFinancialDataCreateData, CorrectFinancialDataRemoveMatch } from '../BudFinancialDataTypes';
declare class CorrectFinancialDataEntity extends BudFinancialDataEntityBase<CorrectFinancialData> {
    constructor(client: BudFinancialDataSDK, entopts: any);
    make(this: CorrectFinancialDataEntity): CorrectFinancialDataEntity;
    load(this: any, reqmatch?: CorrectFinancialDataLoadMatch, ctrl?: Control): Promise<CorrectFinancialDataEntity>;
    list(this: any, reqmatch?: CorrectFinancialDataListMatch, ctrl?: Control): Promise<CorrectFinancialDataEntity[]>;
    create(this: any, reqdata?: CorrectFinancialDataCreateData, ctrl?: Control): Promise<CorrectFinancialDataEntity>;
    remove(this: any, reqmatch?: CorrectFinancialDataRemoveMatch, ctrl?: Control): Promise<CorrectFinancialDataEntity>;
}
export { CorrectFinancialDataEntity };
