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
	problemId := r.GetQuery("problem_id").Int()
	r.Response.WriteJson(api.SuccessRes(problemLogic.ReadProblemDetail(problemId)))
}
