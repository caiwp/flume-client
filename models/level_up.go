package models

import "flume-client/components/setting"

type LevelUpModel struct {
	ProductModel
	IPTimeModel
	AccountModel
	CurrentlyExp int32 `json:"currently_exp" ini:"CURRENTLY_EXP"`
	Exp          int32 `json:"exp" ini:"EXP"`
}

var LevelUp LevelUpModel

func (LevelUpModel) Init() error {
	LevelUp = LevelUpModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
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
