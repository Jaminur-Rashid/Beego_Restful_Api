package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type TestController struct {
	beego.Controller
}
type LIKE struct {
	Food   string
	Watch  string
	Listen string
}

type JSONS struct {
	//Must write
	Code string
	Msg  string
	User []string `json:"user_info"` //Key is renamed, the outermost is the trick
	Like LIKE
}

func (c *TestController) Get() {
	data := &JSONS{"100", "Get successful",
		[]string{"maple", "18"}, LIKE{"cake", "Movie", "music"}}
	c.Data["json"] = data
	c.ServeJSON()
}
