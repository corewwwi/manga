package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
)

// StatusesController operations for Statuses
type StatusesController struct {
	beego.Controller
}

// URLMapping ...
func (c *StatusesController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title Get All
// @Description get Statuses
// @Success 200 {object} []models.Statuses
// @Failure 403
// @router / [get]
func (c *StatusesController) GetAll() {

	l, err := models.GetAllStatuses()
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}

	} else {
		var temp []interface{}
		temp = append(temp, l)
		c.Data["json"] = Response{
			Success:  true,
			Response: temp,
			Count:    1,
		}

	}
	c.ServeJSON()
}
