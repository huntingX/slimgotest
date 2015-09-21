package controller

import (
	"github.com/jesusslim/slimgo"
	"github.com/jesusslim/slimmysql"
	"math/rand"
	"strconv"
)

type TestController struct {
	slimgo.Controller
}

func (this *TestController) Index() {
	funcs := []string{
		"test",
		"list",
		"condition",
		"insert?nickname=Test" + strconv.Itoa(rand.Intn(100)),
		"cookie",
		"jsonp?callback=test",
		"xml",
		"stoptimetask",
		"restarttimetask",
	}
	html := ""
	for _, v := range funcs {
		html += "<a href='/test/" + v + "' target='_blank'>The example link of:" + v + "</a><br>"
	}
	this.Context.WriteString(html)
}

func (this *TestController) Test() {
	this.Data["json"] = slimgo.AppName
	this.ServeJson()
}

func (this *TestController) List() {
	sm := new(slimmysql.Sql)
	stds, err := sm.Table("students").Select()
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = stds
	}
	this.ServeJson()
}

func (this *TestController) Condition() {
	sm := new(slimmysql.Sql)
	condition := map[string]interface{}{
		"nickname__like": "sl",
	}
	stds, err := sm.Table("students").Where(condition).GetField("id,nickname")
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = map[string]interface{}{
			"students": stds,
		}
	}
	this.ServeJson()
}

func (this *TestController) Insert() {
	sm := new(slimmysql.Sql)
	nickname := this.Context.Request.FormValue("nickname")
	id, err := sm.Table("students").Add(map[string]interface{}{
		"nickname": nickname,
	})
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = "insert success,id:" + strconv.Itoa(int(id))
	}
	this.ServeJson()
}

func (this *TestController) Cookie() {
	test1 := this.Context.GetCookie("test2")
	if test1 != "" {
		this.Data["json"] = map[string]string{
			"Result": "find cookie",
			"Cookie": test1,
		}
	} else {
		this.Context.SetCookie("test2", "ooooooook", 200)
		this.Data["json"] = map[string]string{
			"Result": "not found,set it",
		}
	}
	this.ServeJson()
}

func (this *TestController) Jsonp() {
	this.ServeJsonp(map[string]string{
		"nickname": "jsonp",
		"age":      "27",
	})
}

func (this *TestController) Xml() {
	this.ServeXml("testxml")
}

func (this *TestController) StopTimeTask() {
	err := slimgo.ShutDownTimeTask("showtime")
	if err != nil {
		this.ServeJson(err.Error())
	} else {
		this.ServeJson("shut down success")
	}
}

func (this *TestController) RestartTimeTask() {
	err := slimgo.RestartTimeTask("showtime")
	if err != nil {
		this.ServeJson(err.Error())
	} else {
		this.ServeJson("restart success")
	}
}
