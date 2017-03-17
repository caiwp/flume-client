package models

import (
	"flume-client/components/setting"
	"time"
)

type IpTimeModel struct {
	Ip   string `ini:"IP"`
	Time string `ini:"TIME"`
}

var IpTime IpTimeModel

func init() {
	err := setting.Cfg.Section("models.ip_time").MapTo(&IpTime)
	if err != nil {
		panic(err)
	}
	if IpTime.Time == "" {
		IpTime.Time = time.Now().Format("2006-01-02 15:04:05")
	}
}
