package models

import "flume-client/components/setting"

type LoginModel struct {
	ProductModel
	IpTimeModel
	AccountModel
	DeviceModel

	Session_id string `ini:"SESSION_ID"`
}

var Login LoginModel

func (LoginModel) Init() error {
	Login = LoginModel{
		ProductModel: Product,
		IpTimeModel:  IpTime,
		AccountModel: Account,
		DeviceModel:  Device,
	}

	err := setting.Cfg.Section("models.login").MapTo(&Login)
	if err != nil {
		return err
	}
	return nil
}

func (LoginModel) GetType() string {
	return "login"
}
