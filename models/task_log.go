package models

import "flume-client/components/setting"

type TaskLogModel struct {
    ProductModel
    IpTimeModel
    AccountModel

    Operate    int32    `ini:"OPERATE"`
    Name       string   `ini:"NAME"`
    Type       string   `ini:"TYPE"`
    Result     int32    `ini:"RESULT"`
    Reason     string   `ini:"REASON"`
    Session_id string   `ini:"SESSION_ID"`
}

var TaskLog TaskLogModel

func (l *TaskLogModel) Init() error {
    TaskLog = TaskLogModel{
        ProductModel: Product,
        IpTimeModel: IpTime,
        AccountModel: Account,
    }

    err := setting.Cfg.Section("models.task_log").MapTo(&TaskLog)
    if err != nil {
        return err
    }
    return nil
}

func (TaskLogModel) GetType() string {
    return "task-log"
}
