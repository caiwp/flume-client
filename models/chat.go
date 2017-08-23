package models

import "flume-client/components/setting"

type ChatModel struct {
	ProductModel
	IPTimeModel
	AccountModel

	Type    string `json:"type" ini:"TYPE"`
	Content string `json:"content" ini:"CONTENT"`
	Target  string `json:"target" ini:"TARGET"`
	Keyword string `json:"keyword" ini:"KEYWORD"`
}

var Chat ChatModel

func (ChatModel) Init() error {
	Chat = ChatModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
	}

	err := setting.Cfg.Section("models.chat").MapTo(&Chat)
	if err != nil {
		return err
	}
	return nil
}

func (ChatModel) GetType() string {
	return "chat"
}
