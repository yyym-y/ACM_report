package controller

import (
	"back/api"
	"back/internal/logic"

	"github.com/gogf/gf/v2/net/ghttp"
)

type IndexController struct{}

func NewIndexController() *IndexController {
	return &IndexController{}
}

var indexTableLogic = logic.NewIndexTableLogic()

func (*IndexController) Table(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Response.WriteJson(api.SuccessRes(indexTableLogic.Serve()))
}
