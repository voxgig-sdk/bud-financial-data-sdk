import { BudFinancialDataEntityBase } from '../BudFinancialDataEntityBase';
import type { BudFinancialDataSDK } from '../BudFinancialDataSDK';
import type { Control } from '../types';
import type { Label, LabelCreateData, LabelUpdateData } from '../BudFinancialDataTypes';
declare class LabelEntity extends BudFinancialDataEntityBase<Label> {
    constructor(client: BudFinancialDataSDK, entopts: any);
    make(this: LabelEntity): LabelEntity;
    create(this: any, reqdata?: LabelCreateData, ctrl?: Control): Promise<LabelEntity>;
    update(this: any, reqdata?: LabelUpdateData, ctrl?: Control): Promise<LabelEntity>;
}
export { LabelEntity };
