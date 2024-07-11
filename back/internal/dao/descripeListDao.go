package dao

import (
	"back/api/problem"

	"github.com/gogf/gf/v2/frame/g"
)

type DescripeListDao struct{}

func NewDescripeListDao() *DescripeListDao {
	return &DescripeListDao{}
}

func (*DescripeListDao) ReadDescription(problemId int) (res problem.ProblemDescrip) {
	g.Model("descripe_list").Where("problem_id", problemId).Scan(&res)
	return
}
