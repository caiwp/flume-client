package models

import (
	"flume-client/components/setting"
	"fmt"
	"time"
)

type LoginModel struct {
	ProductModel
	IPTimeModel
	AccountModel
	DeviceModel
	VersionModel

	SessionID string `json:"session_id" ini:"SESSION_ID"`
}

var Login LoginModel

func (LoginModel) Init() error {
	Login = LoginModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
		DeviceModel:  Device,
		VersionModel: Version,
	}

	err := setting.Cfg.Section("models.login").MapTo(&Login)
	if err != nil {
		return err
	}
	Login.SessionID = fmt.Sprintf("%s-%d", Login.SessionID, time.Now().Unix())
	return nil
}

func (LoginModel) GetType() string {
	return "login"
}
