package models

import "flume-client/components/setting"

type ShopModel struct {
	ProductModel
	IPTimeModel
	AccountModel
	VersionModel

	Type          string `json:"type" ini:"TYPE"`
	ItemID        int32  `json:"item_id" ini:"ITEM_ID"`
	ItemName      string `json:"item_name" ini:"ITEM_NAME"`
	ItemCount     int32  `json:"item_count" ini:"ITEM_COUNT"`
	Category      string `json:"category" ini:"CATEGORY"`
	CategoryCount int32  `json:"category_count" ini:"CATEGORY_COUNT"`
	Info          string `json:"info" ini:"INFO"`
}

var Shop ShopModel

func (ShopModel) Init() error {
	Shop = ShopModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
		VersionModel: Version,
	}

	err := setting.Cfg.Section("models.shop").MapTo(&Shop)
	if err != nil {
		return err
	}
	return nil
}

func (ShopModel) GetType() string {
	return "shop"
}
