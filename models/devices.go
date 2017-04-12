package models

import (
	"flume-client/components/setting"
	"time"
)

type DevicesModel struct {
	DeviceModel

	ProductName  string `ini:"PRODUCT_NAME"`
	PlatformName string `ini:"PLATFORM_NAME"`
	ChannelName  string `ini:"CHANNEL_NAME"`
	Time         string `ini:"TIME"`
}

var Devices DevicesModel

func (DevicesModel) Init() error {
	Devices = DevicesModel{
		DeviceModel: Device,
	}

	err := setting.Cfg.Section("models.product").MapTo(&Devices)
	if err != nil {
		return err
	}
	if Devices.Time == "" {
		Devices.Time = time.Now().Format("2006-01-02 15:04:05")
	}
	return nil
}

func (DevicesModel) GetType() string {
	return "device"
}
