package controller

import (
	"back/api"
	"back/api/problem"
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

func (c *IndexController) InsertProblem(r *ghttp.Request) {
	r.Response.CORSDefault()
	data := problem.NewProblemInfo()
	r.GetForm("data").Scan(&data)
	r.Response.WriteJson(indexTableLogic.InsertProblem(*data))
}

func (c *IndexController) UpdataProblem(r *ghttp.Request) {
	r.Response.CORSDefault()
	problemId := r.GetForm("Problem_id").Int()
	data := problem.NewProblemInfo()
	r.GetForm("data").Scan(&data)
	r.Response.WriteJson(indexTableLogic.UpdataProblem(problemId, *data))
}

func (c *IndexController) DeleteProblem(r *ghttp.Request) {
	r.Response.CORSDefault()
	problemId := r.GetForm("Problem_id").Int()
	r.Response.WriteJson(indexTableLogic.DeleteProblem(problemId))
}

func (c *IndexController) QueryAllType(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Response.WriteJson(indexTableLogic.QueryAllType())
}

func (c *IndexController) QueryAllDiffical(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Response.WriteJson(indexTableLogic.QueryAllDiffical())
}
