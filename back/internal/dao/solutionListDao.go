package dao

import (
	"back/api/problem"

	"github.com/gogf/gf/v2/frame/g"
)

type SolutioonListDao struct{}

func NewSolutionListDao() *SolutioonListDao {
	return &SolutioonListDao{}
}

func (*SolutioonListDao) ReadSolution(problemId int) (res problem.ProblemSolution) {
	g.Model("solution_list").Where("problem_id", problemId).Scan(&res)
	return
}
