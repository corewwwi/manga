package controllers

import (
	"github.com/astaxie/beego"
)

// LogoutController operations for Logout
type LogoutController struct {
	beego.Controller
}

// URLMapping ...
func (c *LogoutController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Post)
}

// Post ...
// @Title Logout
// @Description Logout
// @Param	body		body 	models.Logout	true		"body for Logout content"
// @Success 201 {object} models.Logout
// @Failure 403 body is empty
// @router / [post]
func (c *LogoutController) Post() {

}