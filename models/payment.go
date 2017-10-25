package models

import (
	"fmt"
	"time"

	"flume-client/components/setting"
)

type PaymentModel struct {
	ProductModel
	IPTimeModel
	AccountModel
	DeviceModel
	VersionModel

	OrderID        string `json:"order_id" ini:"ORDER_ID"`
	CurrencyType   string `json:"currency_type" ini:"CURRENCY_TYPE"`
	CurrencyAmount int    `json:"currency_amount" ini:"CURRENCY_AMOUNT"`
	PayCategory    int    `json:"pay_category" ini:"PAY_CATEGORY"`
	PackageName    string `json:"package_name" ini:"PACKAGE_NAME"`
	PayChannel     string `json:"pay_channel" ini:"PAY_CHANNEL"`
	VirtualAmount  int    `json:"virtual_amount" ini:"VIRTUAL_AMOUNT"`
}

var Payment PaymentModel

func (PaymentModel) Init() error {
	Payment = PaymentModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
		DeviceModel:  Device,
		VersionModel: Version,
	}

	err := setting.Cfg.Section("models.payment").MapTo(&Payment)
	if err != nil {
		return err
	}

	Payment.OrderID = fmt.Sprintf("%s-%d", Payment.OrderID, time.Now().Unix())
	return nil
}

func (PaymentModel) GetType() string {
	return "payment"
}
