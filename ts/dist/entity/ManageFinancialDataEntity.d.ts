import { BudFinancialDataEntityBase } from '../BudFinancialDataEntityBase';
import type { BudFinancialDataSDK } from '../BudFinancialDataSDK';
import type { Control } from '../types';
import type { ManageFinancialData, ManageFinancialDataCreateData, ManageFinancialDataRemoveMatch } from '../BudFinancialDataTypes';
declare class ManageFinancialDataEntity extends BudFinancialDataEntityBase<ManageFinancialData> {
    constructor(client: BudFinancialDataSDK, entopts: any);
    make(this: ManageFinancialDataEntity): ManageFinancialDataEntity;
    create(this: any, reqdata?: ManageFinancialDataCreateData, ctrl?: Control): Promise<ManageFinancialDataEntity>;
    remove(this: any, reqmatch?: ManageFinancialDataRemoveMatch, ctrl?: Control): Promise<ManageFinancialDataEntity>;
}
export { ManageFinancialDataEntity };
