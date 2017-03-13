package models

import "flume-client/components/setting"

type ProductModel struct {
    ProductName  string     `ini:"PRODUCT_NAME"`
    PlatformName string     `ini:"PLATFORM_NAME"`
    ChannelName  string     `ini:"CHANNEL_NAME"`
    GameserverNo int32      `ini:"GAMESERVER_NO"`
}

var Product ProductModel

func init() {
    err := setting.Cfg.Section("models.product").MapTo(&Product)
    if err != nil {
        panic(err)
    }
}
