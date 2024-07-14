package problem

type DifficultyInfo struct {
	Diff_id   int
	Diff_name string
}

func NewDifficultyInfo() *DifficultyInfo {
	return &DifficultyInfo{}
}
