package models

type VirtualCurrencyModel struct {
    ProductModel
    IpTimeModel
    AccountModel
    Category       string
    Project        string
    Info           string
    OpCount        int32
    QuantityBefore int32
}

var VirtualCurrency VirtualCurrencyModel

func (v *VirtualCurrencyModel) Init() {
    VirtualCurrency = VirtualCurrencyModel{
        ProductModel: Product,
        IpTimeModel: IpTime,
        AccountModel: Account,
        Category: "元宝",
        Project: "装备出售",
        Info: "",
        OpCount: 20,
        QuantityBefore: 2000,
    }
}

func (VirtualCurrencyModel) GetType() string {
    return "virtual-currency"
}