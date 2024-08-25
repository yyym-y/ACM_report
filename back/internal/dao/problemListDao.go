package dao

import (
	"back/api/index"
	"back/api/problem"

	"github.com/gogf/gf/v2/frame/g"
)

type ProblemListDao struct{}

func NewProblemListDao() *ProblemListDao {
	return &ProblemListDao{}
}

func (*ProblemListDao) ReadAllInfo() (res []index.IndexTableInfo) {
	g.Model("problem_list").Scan(&res)
	return
}

func (*ProblemListDao) InsertProblem(data problem.ProblemInfo) bool {
	res, err := g.Model("problem_list").Insert(data)
	if res == nil || err != nil {
		println(err)
		return false
	}
	return true
}

func (*ProblemListDao) DeleteProblem(problemId int) bool {
	res, err := g.Model("problem_list").Where("problem_id", problemId).Delete()
	if res == nil || err != nil {
		println(err)
		return false
	}
	return true
}

func (*ProblemListDao) UpdataProblem(problemId int, data problem.ProblemInfo) bool {
	res, err := g.Model("problem_list").Where("problem_id", problemId).Update(data)
	if res == nil || err != nil {
		println(err)
		return false
	}
	return true
}

func (*ProblemListDao) QueryOneProblem(problemId int) (res index.IndexTableInfo) {
	g.Model("problem_list").Where("problem_id", problemId).Scan(&res)
	return
}
