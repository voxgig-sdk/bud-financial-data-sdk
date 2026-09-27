import { BudFinancialDataEntityBase } from '../BudFinancialDataEntityBase';
import type { BudFinancialDataSDK } from '../BudFinancialDataSDK';
import type { Control } from '../types';
import type { Similar, SimilarLoadMatch } from '../BudFinancialDataTypes';
declare class SimilarEntity extends BudFinancialDataEntityBase<Similar> {
    constructor(client: BudFinancialDataSDK, entopts: any);
    make(this: SimilarEntity): SimilarEntity;
    load(this: any, reqmatch?: SimilarLoadMatch, ctrl?: Control): Promise<SimilarEntity>;
}
export { SimilarEntity };
