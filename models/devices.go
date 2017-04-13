package models

import (
	"flume-client/components/setting"
)

type DevicesModel struct {
	DeviceModel

	ProductName  string `ini:"PRODUCT_NAME"`
	PlatformName string `ini:"PLATFORM_NAME"`
	ChannelName  string `ini:"CHANNEL_NAME"`

	IpTimeModel
}

var Devices DevicesModel

func (DevicesModel) Init() error {
	Devices = DevicesModel{
		DeviceModel: Device,
		IpTimeModel: IpTime,
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
