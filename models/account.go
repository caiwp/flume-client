package models

import (
	"math/rand"

	"flume-client/components/setting"
)

type AccountModel struct {
	AccountId   string `ini:"ACCOUNT_ID"`
	AccountName string `ini:"ACCOUNT_NAME"`
	ChrId       string `ini:"CHR_ID"`
	ChrName     string `ini:"CHR_NAME"`
	ChrLvl      int32  `ini:"CHR_LVL"`
	ChrLvlVip   int32  `ini:"CHR_LVL_VIP"`
	Career      string `ini:"CAREER"`
	Gender      string `ini:"GENDER"`
}

var Account AccountModel

func init() {
	err := setting.Cfg.Section("models.account").MapTo(&Account)
	if err != nil {
		panic(err)
	}
	Account.ChrLvl = rand.Int31n(100)
	Account.ChrLvlVip = rand.Int31n(5)
}
