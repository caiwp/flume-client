package main

import (
    "flume-client/components/client"
    "flume-client/models"
    "flume-client/components/log"
    "time"
)

var l = log.NewLogger("default")

func main() {
    t0 := time.Now()
    defer func(t0 time.Time) {
        t1 := time.Now()
        l.Info("End time duration: %.4fs", t1.Sub(t0).Seconds())
    }(t0)

    client := client.NewFlumeClient("192.168.1.110", 60000)
    client.Connect()
    defer client.Destroy()

    done := make(chan bool, 1)
    ms := []models.Model{
        &models.VirtualCurrency,
        &models.VirtualItem,
        &models.LevelUp,
    }

    go func() {
        for _, v := range ms {
            sendData(client, v)
        }
        done <- true
    }()

    <-done
    return
}

func sendData(cl *client.FlumeClient, m models.Model) {
    m.Init()
    evt := models.GetEvent(m)
    cl.Append(evt)
}
