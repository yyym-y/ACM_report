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

func (*DescripeListDao) CheckIfExist(problemId int, desType string) bool {
	des, _ := g.Model("descripe_list").Where("problem_id", problemId).Where("descripe_type", desType).All()
	return len(des) != 0
}

func (*DescripeListDao) UpdataDescription(problemId int, desType string, text string) bool {
	data := g.Map{
		"problem_id":    problemId,
		"descripe_type": desType,
		"description":   text,
	}
	res, err := g.Model("descripe_list").Where("problem_id", problemId).Update(data)
	if err != nil || res == nil {
		return false
	}
	return true
}

func (*DescripeListDao) InsertDescription(problemId int, desType string, text string) bool {
	data := g.Map{
		"problem_id":    problemId,
		"descripe_type": desType,
		"description":   text,
	}
	res, err := g.Model("descripe_list").Insert(data)
	if res == nil || err != nil {
		return false
	}
	return true
}
