package logic

import (
	"back/api/index"
	"back/internal/dao"
)

type IndexTableLogic struct{}

func NewIndexTableLogic() *IndexTableLogic {
	return &IndexTableLogic{}
}

var problemListDao = dao.NewProblemListDao()

func (*IndexTableLogic) Serve() (res []index.IndexTableInfo) {
	res = problemListDao.ReadAllInfo()
	return
}
