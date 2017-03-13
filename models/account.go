package models

type AccountModel struct {
    AccountId   string
    AccountName string
    ChrId       string
    ChrName     string
    ChrLvl      int32
    ChrLvlVip   int32
}

var Account AccountModel

func init() {
    Account = AccountModel{
        "123123",
        "caicai",
        "123",
        "hehe",
        4,
        2,
    }
}
