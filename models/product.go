package models

type ProductModel struct {
    ProductName  string
    PlatformName string
    ChannelName  string
    GameserverNo int32
}

var Product ProductModel

func init() {
    Product = ProductModel{
        "xcqy",
        "plat1",
        "test1",
        1,
    }
}
