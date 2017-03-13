package models

import "flume-client/components/setting"

type VirtualCurrencyModel struct {
    ProductModel
    IpTimeModel
    AccountModel
    Category       string   `ini:"CATEGORY"`
    Project        string   `ini:"PROJECT"`
    Info           string   `ini:"INFO"`
    OpCount        int32    `ini:"OP_COUNT"`
    QuantityBefore int32    `ini:"QUANTITY_BEFORE"`
}

var VirtualCurrency VirtualCurrencyModel

func (v *VirtualCurrencyModel) Init() error {
    VirtualCurrency = VirtualCurrencyModel{
        ProductModel: Product,
        IpTimeModel: IpTime,
        AccountModel: Account,
    }

    err := setting.Cfg.Section("models.virtual_currency").MapTo(&VirtualCurrency)
    if err != nil {
        return err
    }
    return nil
}

func (VirtualCurrencyModel) GetType() string {
    return "virtual-currency"
}