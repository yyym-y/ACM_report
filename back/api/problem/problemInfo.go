package problem

type ProblemInfo struct {
	Problem_name string
	Problem_url  string
	Type_id      int
	Solve_time   string
	Diff_id      int
}

func NewProblemInfo() *ProblemInfo {
	return &ProblemInfo{}
}
