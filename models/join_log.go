package models

import "flume-client/components/setting"

type JoinLogModel struct {
	ProductModel
	IPTimeModel
	AccountModel
	VersionModel

	Type      string `json:"type" ini:"TYPE"`
	Name      string `json:"name" ini:"NAME"`
	Result    int32  `json:"result" ini:"RESULT"`
	SessionID string `json:"session_id" ini:"SESSION_ID"`
}

var JoinLog JoinLogModel

func (JoinLogModel) Init() error {
	JoinLog = JoinLogModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
		VersionModel: Version,
	}

	err := setting.Cfg.Section("models.join_log").MapTo(&JoinLog)
	if err != nil {
		return err
	}
	return nil
}

func (JoinLogModel) GetType() string {
	return "join-log"
}
