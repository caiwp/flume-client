package models

import "flume-client/components/setting"

// ProductModel 游戏信息
type ProductModel struct {
	ProductName  string `json:"product_name" ini:"PRODUCT_NAME"`
	PlatformName string `json:"platform_name" ini:"PLATFORM_NAME"`
	ChannelName  string `json:"channel_name" ini:"CHANNEL_NAME"`
	GameserverNo int32  `json:"gameserver_no" ini:"GAMESERVER_NO"`
}

// Product 游戏实例
var Product ProductModel

func init() {
	err := setting.Cfg.Section("models.product").MapTo(&Product)
	if err != nil {
		panic(err)
	}
}
