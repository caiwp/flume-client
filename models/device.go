package models

import "flume-client/components/setting"

type DeviceModel struct {
    DeviceId string     `ini:"DEVICE_ID"`
    Model    string     `ini:"MODEL"`
    Os       string     `ini:"OS"`
}

var Device DeviceModel

func init() {
    err := setting.Cfg.Section("models.device").MapTo(&Device)
    if err != nil {
        panic(err)
    }
}