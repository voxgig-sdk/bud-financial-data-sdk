import { CorrectFinancialDataEntity } from './entity/CorrectFinancialDataEntity';
import { CustomerMerchantCorrectionEntity } from './entity/CustomerMerchantCorrectionEntity';
import { LabelEntity } from './entity/LabelEntity';
import { ListLabelEntity } from './entity/ListLabelEntity';
import { ListTransactionEntity } from './entity/ListTransactionEntity';
import { ManageFinancialDataEntity } from './entity/ManageFinancialDataEntity';
import { ManageTransactionLabelEntity } from './entity/ManageTransactionLabelEntity';
import { RetrieveFinancialDataEntity } from './entity/RetrieveFinancialDataEntity';
import { SimilarEntity } from './entity/SimilarEntity';
export type * from './BudFinancialDataTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { BudFinancialDataEntityBase } from './BudFinancialDataEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class BudFinancialDataSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    CorrectFinancialData(entopts?: Record<string, any>): CorrectFinancialDataEntity;
    CustomerMerchantCorrection(entopts?: Record<string, any>): CustomerMerchantCorrectionEntity;
    Label(entopts?: Record<string, any>): LabelEntity;
    ListLabel(entopts?: Record<string, any>): ListLabelEntity;
    ListTransaction(entopts?: Record<string, any>): ListTransactionEntity;
    ManageFinancialData(entopts?: Record<string, any>): ManageFinancialDataEntity;
    ManageTransactionLabel(entopts?: Record<string, any>): ManageTransactionLabelEntity;
    RetrieveFinancialData(entopts?: Record<string, any>): RetrieveFinancialDataEntity;
    Similar(entopts?: Record<string, any>): SimilarEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): BudFinancialDataSDK;
    tester(testopts?: any, sdkopts?: any): BudFinancialDataSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof BudFinancialDataSDK;
export { stdutil, config, BaseFeature, BudFinancialDataEntityBase, BudFinancialDataSDK, SDK, };
