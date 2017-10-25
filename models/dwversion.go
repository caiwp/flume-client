package models

import "flume-client/components/setting"

type VersionModel struct {
	Version string `json:"version" ini:"VERSION"`
}

var Version VersionModel

func init() {
	err := setting.Cfg.Section("models.version").MapTo(&Version)
	if err != nil {
		panic(err)
	}
}
