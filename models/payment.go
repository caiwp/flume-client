package models

import (
	"fmt"
	"time"

	"flume-client/components/setting"
)

type PaymentModel struct {
	ProductModel
	IpTimeModel
	AccountModel
	DeviceModel

	OrderId        string `ini:"ORDER_ID"`
	CurrencyType   string `ini:"CURRENCY_TYPE"`
	CurrencyAmount int    `ini:"CURRENCY_AMOUNT"`
	PayCategory    int    `ini:"PAY_CATEGORY"`
	PackageName    string `ini:"PACKAGE_NAME"`
	PayChannel     string `ini:"PAY_CHANNEL"`
	VirtualAmount  int    `ini:"VIRTUAL_AMOUNT"`
}

var Payment PaymentModel

func (PaymentModel) Init() error {
	Payment = PaymentModel{
		ProductModel: Product,
		IpTimeModel:  IpTime,
		AccountModel: Account,
		DeviceModel:  Device,
	}

	err := setting.Cfg.Section("models.payment").MapTo(&Payment)
	if err != nil {
		return err
	}

	Payment.OrderId = fmt.Sprintf("%s-%d", Payment.OrderId, time.Now().Unix())
	return nil
}

func (PaymentModel) GetType() string {
	return "payment"
}
