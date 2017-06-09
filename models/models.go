package models

import (
	"encoding/json"

	"flume-client/components/flume"

	"github.com/Sirupsen/logrus"
	"gopkg.in/op/go-logging.v1"
)

var l logging.Logger

type Model interface {
	Init() error
	GetType() string
}

func getJson(l interface{}) ([]byte, error) {
	res, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetEvent(m Model) *flume.ThriftFlumeEvent {
	event := flume.NewThriftFlumeEvent()
	event.Headers = map[string]string{
		"type": m.GetType(),
	}
	var err error
	event.Body, err = getJson(m)
	logrus.Info(m.GetType() + ":" + string(event.Body))

	if err != nil {
		panic(err)
	}
	return event
}
