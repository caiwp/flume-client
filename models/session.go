package models

import (
	"flume-client/components/setting"
	"math/rand"
	"time"
	"fmt"
)

type SessionModel struct {
	ProductModel
	IpTimeModel
	AccountModel

	Session_id string `ini:"SESSION_ID"`
}

var Session SessionModel

func (SessionModel) Init() error {
	Session = SessionModel{
		ProductModel: Product,
		IpTimeModel:  IpTime,
		AccountModel: Account,
	}

	err := setting.Cfg.Section("models.session").MapTo(&Session)
	if err != nil {
		return err
	}
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation(TIME_FORMAT, Session.Time, loc)
	if err != nil {
		return err
	}
	dt := time.Unix(t.Unix() + rand.Int63n(1000), 0)
	Session.Time = dt.Format(TIME_FORMAT)

	Session.Session_id = fmt.Sprintf("%s-%d", Session.Session_id, time.Now().Unix())
	return nil
}

func (SessionModel) GetType() string {
	return "session"
}
