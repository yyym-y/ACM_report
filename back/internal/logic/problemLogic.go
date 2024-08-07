package logic

import (
	"back/api"
	"back/api/problem"
	"back/internal/crawler"
	"back/internal/dao"

	"github.com/gocolly/colly"
)

type ProblemLogic struct{}

func NewProblemLogic() *ProblemLogic {
	return &ProblemLogic{}
}

var descripeListDao = dao.NewDescripeListDao()
var solutioonListDao = dao.NewSolutionListDao()

func (*ProblemLogic) ReadProblemDetail(problemId int) problem.ProblemDetail {
	descrip := descripeListDao.ReadDescription(problemId)
	solution := solutioonListDao.ReadSolution(problemId)
	return *problem.NewP_DetailWithData(
		problemId, descrip, solution,
	)
}

func (*ProblemLogic) ChangeDescription(problemId int, text string) api.ApiRes {
	type_ := "codeforce"
	success := false
	if descripeListDao.CheckIfExist(problemId, type_) {
		success = descripeListDao.UpdataDescription(problemId, type_, text)
	} else {
		success = descripeListDao.InsertDescription(problemId, type_, text)
	}
	if !success {
		return *api.ErrorRes("change description error")
	}
	return *api.SuccessResNone()
}

func (*ProblemLogic) ChangeSolution(problemId int, text string) api.ApiRes {
	success := false
	if solutioonListDao.CheckIfExist(problemId) {
		success = solutioonListDao.UpdataSolution(problemId, text)
	} else {
		success = solutioonListDao.InsertSolution(problemId, text)
	}
	if !success {
		return *api.ErrorRes("change solution error")
	}
	return *api.SuccessResNone()
}

func (*ProblemLogic) ProblemCrawler(Type string, url string) api.ApiRes {
	c := colly.NewCollector()
	if Type == "codeforces" {
		return *api.SuccessRes(crawler.CfCrawler(c, url))
	}
	return *api.ErrorRes("crowler not found")
}
