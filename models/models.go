package models

import (
    "gopkg.in/op/go-logging.v1"
    "flume-client/components/flume"
    "encoding/json"
    "reflect"
)

var l logging.Logger

type Model interface {
    Init() error
    GetType() string
}

func getData(l interface{}) ([]byte, error) {
    var sl []interface{}

    val := reflect.ValueOf(l)
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }

    for i := 0; i < val.NumField(); i ++ {
        f := val.Field(i)
        if f.Kind() == reflect.Struct {
            ej := reflect.ValueOf(f.Interface())
            for j := 0; j < ej.NumField(); j ++ {
                fj := ej.Field(j)
                sl = append(sl, fj.Interface())
            }
        } else if (f.Kind() == reflect.Ptr) {
            continue
        } else {
            sl = append(sl, f.Interface())
        }
    }
    return json.Marshal(sl)
}

func GetEvent(m Model) *flume.ThriftFlumeEvent {
    event := flume.NewThriftFlumeEvent()
    event.Headers = map[string]string{
        "type" : m.GetType(),
    }
    var err error
    event.Body, err = getData(m)

    l.Info("Body: %s", string(event.GetBody()))
    if err != nil {
        panic(err)
    }
    return event
}

