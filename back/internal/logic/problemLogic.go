package logic

import (
	"back/api/problem"
	"back/internal/dao"
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
