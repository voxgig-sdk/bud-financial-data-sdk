import { BudFinancialDataEntityBase } from '../BudFinancialDataEntityBase';
import type { BudFinancialDataSDK } from '../BudFinancialDataSDK';
import type { Control } from '../types';
import type { ListLabel, ListLabelListMatch } from '../BudFinancialDataTypes';
declare class ListLabelEntity extends BudFinancialDataEntityBase<ListLabel> {
    constructor(client: BudFinancialDataSDK, entopts: any);
    make(this: ListLabelEntity): ListLabelEntity;
    list(this: any, reqmatch?: ListLabelListMatch, ctrl?: Control): Promise<ListLabelEntity[]>;
}
export { ListLabelEntity };
