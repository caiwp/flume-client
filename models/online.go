package models

import (
	"flume-client/components/setting"
	"math/rand"
)

type OnlineModel struct {
	ProductModel
	VersionModel

	DateTime string `json:"date_time"`
	Num      int32  `json:"num" ini:"NUM"`
}

var Online OnlineModel

func (o *OnlineModel) Init() error {
	Online = OnlineModel{
		ProductModel: Product,
		DateTime:     IPTime.DateTime,
		VersionModel: Version,
	}

	err := setting.Cfg.Section("models.online").MapTo(&Online)
	if err != nil {
		return err
	}
	Online.Num = rand.Int31n(1000)
	return nil
}

func (OnlineModel) GetType() string {
	return "online"
}
