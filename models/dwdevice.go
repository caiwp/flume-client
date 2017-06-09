package models

import (
	"flume-client/components/setting"
	"fmt"
	"time"
)

// DeviceModel 设备
type DeviceModel struct {
	DeviceID    string `json:"device_id" ini:"DEVICE_ID"`
	DeviceType  string `json:"device_type" ini:"DEVICE_TYPE"`
	Os          string `json:"os" ini:"OS"`
	Carrier     string `json:"carrier" ini:"CARRIER"`
	NetworkType string `json:"network_type" ini:"NETWORK_TYPE"`
	Resolution  string `json:"resolution" ini:"RESOLUTION"`
}

// Device 设备
var Device DeviceModel

func init() {
	err := setting.Cfg.Section("models.device").MapTo(&Device)
	if err != nil {
		panic(err)
	}
	Device.DeviceID = fmt.Sprintf("%s-%d", Device.DeviceID, time.Now().Unix())
}
