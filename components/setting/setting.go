package setting

import "gopkg.in/ini.v1"

var Cfg *ini.File

func init() {
    conf := "conf/app.ini"
    var err error
    Cfg, err = ini.Load(conf)
    if err != nil {
        panic(err)
    }
}