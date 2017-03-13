package models

import "flume-client/components/setting"

type VirtualItemModel struct {
    ProductModel
    IpTimeModel
    AccountModel
    ItemId   string     `ini:"ITEM_ID"`
    ItemName string     `ini:"ITEM_NAME"`
    TypeName string     `ini:"TYPE_NAME"`
    OpCount  int32      `ini:"OP_COUNT"`
}

var VirtualItem VirtualItemModel

func (v *VirtualItemModel) Init() error {
    VirtualItem = VirtualItemModel{
        ProductModel: Product,
        IpTimeModel: IpTime,
        AccountModel: Account,
    }

    err := setting.Cfg.Section("models.virtual_item").MapTo(&VirtualItem)
    if err != nil {
        return err
    }
    return nil
}

func (VirtualItemModel) GetType() string {
    return "virtual-item"
}