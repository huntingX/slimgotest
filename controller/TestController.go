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
		"test", "list", "condition", "insert?nickname=Test" + strconv.Itoa(rand.Intn(100)),
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
