package models

import "flume-client/components/setting"

type MailModel struct {
	ProductModel
	IPTimeModel
	AccountModel
	VersionModel

	Title   string `json:"title" ini:"TITLE"`
	Content string `json:"content" ini:"CONTENT"`
	Options string `json:"options" ini:"OPTIONS"`
	OpType  string `json:"op_type" ini:"OP_TYPE"`
	Info    string `json:"info" ini:"INFO"`
}

var Mail MailModel

func (MailModel) Init() error {
	Mail = MailModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
		VersionModel: Version,
	}

	err := setting.Cfg.Section("models.mail").MapTo(&Mail)
	if err != nil {
		return err
	}
	return nil
}

func (MailModel) GetType() string {
	return "mail"
}
