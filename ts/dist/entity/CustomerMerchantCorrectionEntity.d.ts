import { BudFinancialDataEntityBase } from '../BudFinancialDataEntityBase';
import type { BudFinancialDataSDK } from '../BudFinancialDataSDK';
import type { Control } from '../types';
import type { CustomerMerchantCorrection, CustomerMerchantCorrectionCreateData } from '../BudFinancialDataTypes';
declare class CustomerMerchantCorrectionEntity extends BudFinancialDataEntityBase<CustomerMerchantCorrection> {
    constructor(client: BudFinancialDataSDK, entopts: any);
    make(this: CustomerMerchantCorrectionEntity): CustomerMerchantCorrectionEntity;
    create(this: any, reqdata?: CustomerMerchantCorrectionCreateData, ctrl?: Control): Promise<CustomerMerchantCorrectionEntity>;
}
export { CustomerMerchantCorrectionEntity };
