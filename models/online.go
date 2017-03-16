package models

import "flume-client/components/setting"

type OnlineModel struct {
    ProductModel
    Time string
    Num  int32      `ini:"NUM"`
}

var Online OnlineModel

func (o *OnlineModel) Init() error {
    Online = OnlineModel{
        ProductModel: Product,
        Time: IpTime.Time,
    }

    err := setting.Cfg.Section("models.online").MapTo(&Online)
    if err != nil {
        return err
    }
    return nil
}

func (OnlineModel) GetType() string {
    return "online"
}