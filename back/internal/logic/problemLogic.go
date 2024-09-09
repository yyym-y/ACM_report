package logic

import (
	"archive/zip"
	"back/api"
	"back/api/problem"
	"back/internal/crawler"
	"back/internal/dao"
	"bytes"
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/gogf/gf/v2/net/ghttp"
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

func (p *ProblemLogic) ExportProblem(nums []interface{}, r *ghttp.Request) []byte {
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)

	contents := make(chan []byte, len(nums))
	defer close(contents)
	for _, pr := range nums {
		num, _ := strconv.Atoi(fmt.Sprintf("%v", pr))
		go p.readProblem(num, &contents)
	}
	for i := 1; i <= len(nums); i++ {
		select {
		case ffile := <-contents:
			numStr := fmt.Sprintf("%v", i)
			f, _ := zw.Create(numStr + "pro.md")
			f.Write(ffile)
		}
	}
	zw.Close()
	return buf.Bytes()
}

func (p *ProblemLogic) readProblem(problemId int, ch *chan []byte) {
	detail := p.ReadProblemDetail(problemId)
	fileContent := detail.Description
	fileContent += "\n\n\n\n\n\n\n---\n\n\n## 题解\n"
	fileContent += detail.Solution
	*ch <- []byte(fileContent)
}
