package dao

import "github.com/gogf/gf/v2/frame/g"

type DifficultyDao struct{}

func NewDifficultyDao() *DifficultyDao {
	return &DifficultyDao{}
}

func (*DifficultyDao) ReadDifficulty(diff_id int) string {
	ff, _ := g.Model("difficulty").Where("diff_id", diff_id).One()
	return ff["diff_name"].String()
}
