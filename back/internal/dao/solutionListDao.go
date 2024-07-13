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

func (*SolutioonListDao) CheckIfExist(problemId int) bool {
	sol, _ := g.Model("solution_list").Where("problem_id", problemId).All()
	return len(sol) != 0
}

func (*SolutioonListDao) UpdataSolution(problemId int, text string) bool {
	data := g.Map{
		"problem_id": problemId,
		"solution":   text,
	}
	res, err := g.Model("solution_list").Where("problem_id", problemId).Update(data)
	if err != nil || res == nil {
		return false
	}
	return true
}

func (*SolutioonListDao) InsertSolution(problemId int, text string) bool {
	data := g.Map{
		"problem_id": problemId,
		"solution":   text,
	}
	res, err := g.Model("solution_list").Insert(data)
	if res == nil || err != nil {
		return false
	}
	return true
}
