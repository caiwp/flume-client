package models

import "flume-client/components/setting"

type VirtualCurrencyModel struct {
	ProductModel
	IPTimeModel
	AccountModel
	VersionModel

	Category       string `json:"category" ini:"CATEGORY"`
	Project        string `json:"project" ini:"PROJECT"`
	Info           string `json:"info" ini:"INFO"`
	OpCount        int32  `json:"op_count" ini:"OP_COUNT"`
	QuantityBefore int32  `json:"quantity_before" ini:"QUANTITY_BEFORE"`
}

var VirtualCurrency VirtualCurrencyModel

func (VirtualCurrencyModel) Init() error {
	VirtualCurrency = VirtualCurrencyModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
		VersionModel: Version,
	}

	err := setting.Cfg.Section("models.virtual_currency").MapTo(&VirtualCurrency)
	if err != nil {
		return err
	}
	return nil
}

func (VirtualCurrencyModel) GetType() string {
	return "virtual-currencies"
}
