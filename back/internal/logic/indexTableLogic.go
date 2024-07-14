package logic

import (
	"back/api"
	"back/api/index"
	"back/api/problem"
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

func (*IndexTableLogic) InsertProblem(data problem.ProblemInfo) api.ApiRes {
	flag := problemListDao.InsertProblem(data)
	return *api.CheckSuccessNone(flag, "insert error")
}

func (*IndexTableLogic) UpdataProblem(problemId int, data problem.ProblemInfo) api.ApiRes {
	flag := problemListDao.UpdataProblem(problemId, data)
	return *api.CheckSuccessNone(flag, "updata error")
}

func (*IndexTableLogic) DeleteProblem(problemId int) api.ApiRes {
	flag := problemListDao.DeleteProblem(problemId)
	return *api.CheckSuccessNone(flag, "delete error")
}

func (c *IndexTableLogic) QueryAllType() api.ApiRes {
	return *api.SuccessRes(typeListDao.ReadAllType())
}

func (c *IndexTableLogic) QueryAllDiffical() api.ApiRes {
	return *api.SuccessRes(difficultyDao.ReadAllDifficulty())
}
