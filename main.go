package main

import (
    "flume-client/components/client"
    "flume-client/models"
    "flume-client/components/log"
    "time"
    "flume-client/components/setting"
)

var (
    l = log.NewLogger("default")

    host string
    port int
)

func main() {
    t0 := time.Now()
    defer func() {
        l.Info("End time duration: %.4fs", time.Since(t0).Seconds())
    }()
    if err := loadConfig(); err != nil {
        panic(err)
    }

    client := client.NewFlumeClient(host, port)

    done := make(chan bool, 1)
    ms := []models.Model{
        &models.VirtualCurrency,
        &models.VirtualItem,
        &models.LevelUp,
    }

    go func() {
        for _, v := range ms {
            err := sendData(client, v)
            if err != nil {
                l.Error("send data failed: %s", err)
            }
        }
        done <- true
    }()

    <-done
    return
}

func loadConfig() error {
    var err error
    sec := setting.Cfg.Section("client")
    host = sec.Key("DIST_HOST").String()
    port, err = sec.Key("DIST_PORT").Int()

    if err != nil {
        return err
    }
    return nil
}

func sendData(cl *client.FlumeClient, m models.Model) error {
    m.Init()
    evt := models.GetEvent(m)
    return cl.Append(evt)
}
