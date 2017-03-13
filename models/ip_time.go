package models

import "time"

type IpTimeModel struct {
    Ip   string
    Time string
}

var IpTime IpTimeModel

func init() {
    t := time.Now().Format("2006-01-02 15:04:05")
    IpTime = IpTimeModel{
        "192.168.1.1",
        t,
    }
}
