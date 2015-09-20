package model

import (
	"github.com/Unknwon/goconfig"
	"github.com/jesusslim/slimgo"
	"github.com/jesusslim/slimmysql"
)

func init() {
	conf, _ := goconfig.LoadConfigFile("./conf/db.ini")
	conf_sq := "local"
	err := slimmysql.InitSqlDefault(conf.MustValue(conf_sq, "user"), conf.MustValue(conf_sq, "pass"), conf.MustValue(conf_sq, "ip"), conf.MustValue(conf_sq, "port"), conf.MustValue(conf_sq, "db"), conf.MustValue(conf_sq, "prefix"), false)
	if err != nil {
		slimgo.Error(err.Error())
	}
}
