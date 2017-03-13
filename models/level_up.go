package models

type LevelUpModel struct {
    ProductModel
    IpTimeModel
    AccountModel
    CurrentlyExp int32
    Exp          int32
}

var LevelUp LevelUpModel

func (l *LevelUpModel) Init() {
    LevelUp = LevelUpModel{
        ProductModel: Product,
        IpTimeModel: IpTime,
        AccountModel: Account,
        CurrentlyExp: 2002,
        Exp: 222,
    }
}

func (LevelUpModel) GetType() string {
    return "level-up"
}
