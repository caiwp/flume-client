package models

import "flume-client/components/setting"

type PetModel struct {
	ProductModel
	IPTimeModel
	AccountModel
	VersionModel

	PetID     int32  `json:"pet_id" ini:"PET_ID"`
	PetName   string `json:"pet_name" ini:"PET_NAME"`
	EventType string `json:"event_type" ini:"EVENT_TYPE"`
	Event     string `json:"event" ini:"EVENT"`
	Info      string `json:"info" ini:"INFO"`
}

var Pet PetModel

func (PetModel) Init() error {
	Pet = PetModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
		VersionModel: Version,
	}

	err := setting.Cfg.Section("models.pet").MapTo(&Pet)
	if err != nil {
		return err
	}
	return nil
}

func (PetModel) GetType() string {
	return "pet"
}
