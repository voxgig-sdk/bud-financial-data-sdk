import { BudFinancialDataEntityBase } from '../BudFinancialDataEntityBase';
import type { BudFinancialDataSDK } from '../BudFinancialDataSDK';
import type { Control } from '../types';
import type { ManageTransactionLabel, ManageTransactionLabelRemoveMatch } from '../BudFinancialDataTypes';
declare class ManageTransactionLabelEntity extends BudFinancialDataEntityBase<ManageTransactionLabel> {
    constructor(client: BudFinancialDataSDK, entopts: any);
    make(this: ManageTransactionLabelEntity): ManageTransactionLabelEntity;
    remove(this: any, reqmatch?: ManageTransactionLabelRemoveMatch, ctrl?: Control): Promise<ManageTransactionLabelEntity>;
}
export { ManageTransactionLabelEntity };
