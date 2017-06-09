package models

import (
	"flume-client/components/setting"
	"fmt"
	"time"
)

type TaskLogModel struct {
	ProductModel
	IPTimeModel
	AccountModel

	Operate   int32  `json:"operate" ini:"OPERATE"`
	TaskID    int32  `json:"task_id" ini:"TASK_ID"`
	Name      string `json:"name" ini:"NAME"`
	TypeID    int32  `json:"type_id" ini:"TYPE_ID"`
	Type      string `json:"type" ini:"TYPE"`
	Result    int32  `json:"result" ini:"RESULT"`
	Reason    string `json:"reason" ini:"REASON"`
	SessionID string `json:"session_id" ini:"SESSION_ID"`
}

var TaskLog TaskLogModel

func (TaskLogModel) Init() error {
	TaskLog = TaskLogModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
	}

	err := setting.Cfg.Section("models.task_log").MapTo(&TaskLog)
	if err != nil {
		return err
	}
	TaskLog.SessionID = fmt.Sprintf("%s-%d", TaskLog.SessionID, time.Now().Unix())
	return nil
}

func (TaskLogModel) GetType() string {
	return "task-log"
}
