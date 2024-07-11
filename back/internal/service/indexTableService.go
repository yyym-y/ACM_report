package service

import "back/api/index"

type IndexTableService interface {
	Serve() []index.IndexTableInfo
}
