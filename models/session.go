package models

import "flume-client/components/setting"

type SessionModel struct {
	ProductModel
	IpTimeModel
	AccountModel

	Session_id string `ini:"SESSION_ID"`
}

var Session SessionModel

func (SessionModel) Init() error {
	Session = SessionModel{
		ProductModel: Product,
		IpTimeModel:  IpTime,
		AccountModel: Account,
	}

	err := setting.Cfg.Section("models.session").MapTo(&Session)
	if err != nil {
		return err
	}
	return nil
}

func (SessionModel) GetType() string {
	return "session"
}
