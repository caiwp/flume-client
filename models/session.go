package models

import (
	"flume-client/components/setting"
	"fmt"
	"math/rand"
	"time"
)

type SessionModel struct {
	ProductModel
	IPTimeModel
	AccountModel
	VersionModel

	SessionID string `json:"session_id" ini:"SESSION_ID"`
}

var Session SessionModel

func (SessionModel) Init() error {
	Session = SessionModel{
		ProductModel: Product,
		IPTimeModel:  IPTime,
		AccountModel: Account,
		VersionModel: Version,
	}

	err := setting.Cfg.Section("models.session").MapTo(&Session)
	if err != nil {
		return err
	}
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation(TimeFormat, Session.DateTime, loc)
	if err != nil {
		return err
	}
	dt := time.Unix(t.Unix()+rand.Int63n(1000), 0)
	Session.DateTime = dt.Format(TimeFormat)

	Session.SessionID = fmt.Sprintf("%s-%d", Session.SessionID, time.Now().Unix())
	return nil
}

func (SessionModel) GetType() string {
	return "session"
}
