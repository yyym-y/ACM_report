package problem

type ProblemDetail struct {
	Problem_id    int
	Descripe_type string
	Description   string
	Solution      string
}

type ProblemDescrip struct {
	Problem_id    int
	Descripe_type string
	Description   string
}

type ProblemSolution struct {
	Problem_id int
	Solution   string
}

func NewProblemDetail() *ProblemDetail {
	return &ProblemDetail{}
}

func NewP_DetailWithData(id int, des ProblemDescrip, solution ProblemSolution) *ProblemDetail {
	return &ProblemDetail{
		Problem_id:    id,
		Descripe_type: des.Descripe_type,
		Description:   des.Description,
		Solution:      solution.Solution,
	}
}

func NewProblemDescrip() *ProblemDescrip {
	return &ProblemDescrip{}
}

func NewProblemSolution() *ProblemSolution {
	return &ProblemSolution{}
}
