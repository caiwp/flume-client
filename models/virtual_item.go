package models

type VirtualItemModel struct {
    ProductModel
    IpTimeModel
    AccountModel
    ItemId   string
    ItemName string
    TypeName string
    OpCount  int32
}

var VirtualItem VirtualItemModel

func (v *VirtualItemModel) Init() {
    VirtualItem = VirtualItemModel{
        ProductModel: Product,
        IpTimeModel: IpTime,
        AccountModel: Account,
        ItemId: "123123",
        ItemName: "牛皮副具",
        TypeName: "商场购买",
        OpCount: 20,
    }
}

func (VirtualItemModel) GetType() string {
    return "virtual-item"
}