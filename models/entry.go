package models

import (
	"flume-client/components/setting"
)

type EntryModel struct {
	ProductName  string `ini:"PRODUCT_NAME"`
	PlatformName string `ini:"PLATFORM_NAME"`
	ChannelName  string `ini:"CHANNEL_NAME"`

	IpTimeModel

	AccountId   string `ini:"ACCOUNT_ID"`
	AccountName string `ini:"ACCOUNT_NAME"`
	Step        int32 `ini:"STEP"`
}

var Entry EntryModel

func (EntryModel) Init() error {
	Entry = EntryModel{
		IpTimeModel: IpTime,
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
