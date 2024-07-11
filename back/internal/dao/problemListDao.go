package dao

import (
	"back/api/index"

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
