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
var difficultyDao = dao.NewDifficultyDao()
var typeListDao = dao.NewTypeListDao()

func (*IndexTableLogic) Serve() (res []index.IndexTableInfo) {
	res = problemListDao.ReadAllInfo()
	for i, v := range res {
		v.Diff = difficultyDao.ReadDifficulty(v.Diff_id)
		v.Type = typeListDao.ReadType(v.Type_id)
		res[i] = v
	}
	return
}
