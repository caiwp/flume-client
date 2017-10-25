package main

import (
	"flag"
	"time"

	"flume-client/components/log"
	"flume-client/components/setting"
	"flume-client/models"

	"github.com/caiwp/utils"
)

var (
	l = log.NewLogger("default")

	host string
	port int

	done chan struct{}

	category string
	ms       map[string]models.Model
)

func Init() {
	done = make(chan struct{})
	flag.StringVar(&category, "category", "all", "a string var")
	ms = map[string]models.Model{
		"virtual-currency": &models.VirtualCurrency,
		"virtual-item":     &models.VirtualItem,
		"level-up":         &models.LevelUp,
		"task-log":         &models.TaskLog,
		"online":           &models.Online,
		"login":            &models.Login,
		"session":          &models.Session,
		"payment":          &models.Payment,
		"device":           &models.Devices,
		//"entry":            &models.Entry,
		"custom-event": &models.CustomEvent,
		"join-log":     &models.JoinLog,
		"rank-list":    &models.RankList,
		"chat":         &models.Chat,
		"shop":         &models.Shop,
	}
}

func main() {
	Init()

	t0 := time.Now()
	defer func() {
		l.Info("End time duration: %.4fs", time.Since(t0).Seconds())
	}()
	if err := loadConfig(); err != nil {
		panic(err)
	}

	flag.Parse()

	client := utils.NewFlumeClient(host, port)

	if category != "all" {
		m, ok := ms[category]
		if !ok {
			l.Error("category not found: %s", category)
			return
		}

		go func() {
			if err := sendData(client, m); err != nil {
				l.Error("send data failed: %s", err)
			}
			stop()
		}()
	} else {
		go func() {
			for _, v := range ms {
				if err := sendData(client, v); err != nil {
					l.Error("send data failed: %s", err)
					stop()
				}
			}
			stop()
		}()
	}

	waitForStop()
	return
}

func stop() {
	close(done)
}

func waitForStop() {
	<-done
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

func sendData(cl *utils.FlumeClient, m models.Model) error {
	m.Init()
	evt := models.GetEvent(m)
	return cl.Append(evt)
}
