package models

import (
	"flume-client/components/setting"
	"time"
)

// TimeFormat 自定义时间格式
const TimeFormat string = "2006-01-02 15:04:05"

// IPTimeModel Ip 和 时间 信息
type IPTimeModel struct {
	IP       string `json:"ip" ini:"IP"`
	DateTime string `json:"date_time" ini:"DATE_TIME"`
}

// IPTime 实例
var IPTime IPTimeModel

func init() {
	err := setting.Cfg.Section("models.ip_time").MapTo(&IPTime)
	if err != nil {
		panic(err)
	}
	if IPTime.DateTime == "" {
		IPTime.DateTime = time.Now().Format(TimeFormat)
	}
}
