package models

import (
	"math/rand"

	"flume-client/components/setting"
	"time"
)

// AccountModel 角色信息
type AccountModel struct {
	AccountID   string `json:"account_id" ini:"ACCOUNT_ID"`
	AccountName string `json:"account_name" ini:"ACCOUNT_NAME"`
	ChrID       string `json:"chr_id" ini:"CHR_ID"`
	ChrName     string `json:"chr_name" ini:"CHR_NAME"`
	ChrLevel    int32  `json:"chr_level" ini:"CHR_LEVEL"`
	ChrLevelVip int32  `json:"chr_level_vip" ini:"CHR_LEVEL_VIP"`
	Career      string `json:"career" ini:"CAREER"`
	Gender      string `json:"gender" ini:"GENDER"`
}

// Account 角色
var Account AccountModel

func init() {
	err := setting.Cfg.Section("models.account").MapTo(&Account)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().Unix())
	if Account.ChrLevel == 0 {
		Account.ChrLevel = rand.Int31n(100)
	}
	if Account.ChrLevelVip == 0 {
		Account.ChrLevelVip = rand.Int31n(5)
	}
}
