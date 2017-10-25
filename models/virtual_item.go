package models

import "flume-client/components/setting"

type VirtualItemModel struct {
	ProductModel
	IPTimeModel
	AccountModel
	VersionModel

	ItemID   int32  `json:"item_id" ini:"ITEM_ID"`
	ItemName string `json:"item_name" ini:"ITEM_NAME"`
	TypeName string `json:"type_name" ini:"TYPE_NAME"`
	OpCount  int32  `json:"op_count" ini:"OP_COUNT"`
}

var VirtualItem VirtualItemModel

func (VirtualItemModel) Init() error {
	VirtualItem = VirtualItemModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
		VersionModel: Version,
	}

	err := setting.Cfg.Section("models.virtual_item").MapTo(&VirtualItem)
	if err != nil {
		return err
	}
	return nil
}

func (VirtualItemModel) GetType() string {
	return "virtual-items"
}
