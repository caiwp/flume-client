package models

import "flume-client/components/setting"

type CustomEventModel struct {
	ProductModel
	IPTimeModel
	AccountModel

	EventName string `json:"event_name" ini:"EVENT_NAME"`
	Comments  string `json:"comments" ini:"COMMENTS"`
	SessionID string `json:"session_id" ini:"SESSION_ID"`
}

var CustomEvent CustomEventModel

func (CustomEventModel) Init() error {
	CustomEvent = CustomEventModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
	}

	err := setting.Cfg.Section("models.custom_event").MapTo(&CustomEvent)
	if err != nil {
		return err
	}
	return nil
}

func (CustomEventModel) GetType() string {
	return "custom-event"
}
