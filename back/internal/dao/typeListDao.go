package dao

import "github.com/gogf/gf/v2/frame/g"

type TypeListDao struct{}

func NewTypeListDao() *TypeListDao {
	return &TypeListDao{}
}

func (*TypeListDao) ReadType(type_id int) string {
	ff, _ := g.Model("type_list").Where("type_id", type_id).One()
	return ff["type_name"].String()
}
