package models

import "flume-client/components/setting"

type LevelUpModel struct {
	ProductModel
	IpTimeModel
	AccountModel
	CurrentlyExp int32 `ini:"CURRENTLY_EXP"`
	Exp          int32 `ini:"EXP"`
}

var LevelUp LevelUpModel

func (LevelUpModel) Init() error {
	LevelUp = LevelUpModel{
		ProductModel: Product,
		IpTimeModel:  IpTime,
		AccountModel: Account,
	}

	err := setting.Cfg.Section("models.level_up").MapTo(&LevelUp)
	if err != nil {
		return err
	}
	return nil
}

func (LevelUpModel) GetType() string {
	return "level-up"
}
