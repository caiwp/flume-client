package models

import (
	"flume-client/components/setting"
)

type EntryModel struct {
	ProductName  string `json:"product_name" ini:"PRODUCT_NAME"`
	PlatformName string `json:"platform_name" ini:"PLATFORM_NAME"`
	ChannelName  string `json:"channel_name" ini:"CHANNEL_NAME"`

	IPTimeModel
	VersionModel

	AccountID   string `json:"account_id" ini:"ACCOUNT_ID"`
	AccountName string `json:"account_name" ini:"ACCOUNT_NAME"`
	Step        int32  `json:"step" ini:"STEP"`
}

var Entry EntryModel

func (EntryModel) Init() error {
	Entry = EntryModel{
		IPTimeModel:  IPTime,
		VersionModel: Version,
	}

	var err error
	err = setting.Cfg.Section("models.product").MapTo(&Entry)
	if err != nil {
		return err
	}
	err = setting.Cfg.Section("models.entry").MapTo(&Entry)
	if err != nil {
		return err
	}
	err = setting.Cfg.Section("models.account").MapTo(&Entry)
	if err != nil {
		return err
	}
	return nil
}

func (EntryModel) GetType() string {
	return "entry"
}
