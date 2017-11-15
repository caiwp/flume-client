package models

import "flume-client/components/setting"

type KafkaModel struct {
	ProductModel
	IPTimeModel
	AccountModel
	VersionModel

	Type    string `json:"type" ini:"TYPE"`
	Content string `json:"content" ini:"CONTENT"`
	Target  string `json:"target" ini:"TARGET"`
	Keyword string `json:"keyword" ini:"KEYWORD"`

	Key int `json:"key"`
}

var Kafka KafkaModel

func (KafkaModel) Init() error {
	Kafka = KafkaModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
		VersionModel: Version,
	}

	Kafka.Key = 2

	err := setting.Cfg.Section("models.chat").MapTo(&Kafka)
	if err != nil {
		return err
	}
	return nil
}

func (KafkaModel) GetType() string {
	return "kafka"
}
