package models

import (
	"flume-client/components/setting"
	"fmt"
	"time"
)

type DeviceModel struct {
	DeviceId    string `ini:"DEVICE_ID"`
	Model       string `ini:"MODEL"`
	Os          string `ini:"OS"`
	Carrier     string `ini:"CARRIER"`
	NetworkType string `ini:"NETWORK_TYPE"`
	Resolution  string `ini:"RESOLUTION"`
}

var Device DeviceModel

func init() {
	err := setting.Cfg.Section("models.device").MapTo(&Device)
	if err != nil {
		panic(err)
	}
	Device.DeviceId = fmt.Sprintf("%s-%d", Device.DeviceId, time.Now().Unix())
}
