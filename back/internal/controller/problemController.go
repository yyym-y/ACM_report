package controller

import (
	"back/api"
	"back/internal/logic"

	"github.com/gogf/gf/v2/net/ghttp"
)

type ProblemController struct{}

func NewProblemController() *ProblemController {
	return &ProblemController{}
}

var problemLogic = logic.NewProblemLogic()

func (c *ProblemController) QueryDetail(r *ghttp.Request) {
	r.Response.CORSDefault()
	problemId := r.GetQuery("Problem_id").Int()
	println(problemId)
	r.Response.WriteJson(api.SuccessRes(problemLogic.ReadProblemDetail(problemId)))
}

func (c *ProblemController) ChangeDescription(r *ghttp.Request) {
	r.Response.CORSDefault()
	problemId := r.GetForm("Problem_id").Int()
	text := r.GetForm("text").String()
	r.Response.WriteJson(problemLogic.ChangeDescription(problemId, text))
}

func (c *ProblemController) ChangeSolution(r *ghttp.Request) {
	r.Response.CORSDefault()
	problemId := r.GetForm("Problem_id").Int()
	text := r.GetForm("text").String()
	r.Response.WriteJson(problemLogic.ChangeSolution(problemId, text))
}
