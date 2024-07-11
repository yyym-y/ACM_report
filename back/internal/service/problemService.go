package service

import (
	"back/api/problem"
	"back/internal/dao"
)

type ProblemService interface {
	Serve() problem.ProblemDetail
}
