package models

import (
	"flume-client/components/setting"
)

type DevicesModel struct {
	DeviceModel

	ProductName  string `json:"product_name" ini:"PRODUCT_NAME"`
	PlatformName string `json:"platform_name" ini:"PLATFORM_NAME"`
	ChannelName  string `json:"channel_name" ini:"CHANNEL_NAME"`

	IPTimeModel
}

var Devices DevicesModel

func (DevicesModel) Init() error {
	Devices = DevicesModel{
		DeviceModel: Device,
		IPTimeModel: IPTime,
	}

	var err error
	err = setting.Cfg.Section("models.product").MapTo(&Devices)
	if err != nil {
		return err
	}
	return nil
}

func (DevicesModel) GetType() string {
	return "device"
}
