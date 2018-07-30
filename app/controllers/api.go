package controllers

import (
	"ccptsgo/app/models"
	"github.com/revel/revel"
	//"errors"
)

type Api struct {
	*revel.Controller
}

func (c Api) Index() revel.Result {
	arcanas := []models.Arcana{}
	DB.Limit(20).Order("id desc").Find(&arcanas)
	return c.RenderJSON(arcanas)
}
