package index

type IndexTableInfo struct {
	Problem_id   int
	Problem_name string
	Problem_url  string
	Type_id      int
	Solve_time   string
	Diff_id      int
	Type         string
	Diff         string
}

func NewIndexTableInfo() *IndexTableInfo {
	return &IndexTableInfo{}
}
