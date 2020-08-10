package controllers

import (
	"beego-hw/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	// log pv
	models.LogPV("/user")

	// tpl data assign
	c.Data["mysqluser"] = beego.AppConfig.String("mysqluser")
	c.Data["mysqlpass"] = beego.AppConfig.String("mysqlpass")
	// tpl view
	c.TplName = "user.tpl"
}