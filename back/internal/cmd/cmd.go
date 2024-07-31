package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"back/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			indexController := controller.NewIndexController()
			problemController := controller.NewProblemController()
			s.Group("/index", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.POST("/insertProblem", indexController.InsertProblem)
				group.POST("/updataProblem", indexController.UpdataProblem)
				group.POST("/deleteProblem", indexController.DeleteProblem)
				group.POST("/insertDifficulty", indexController.InsertDifficulty)
				group.GET("/queryAllType", indexController.QueryAllType)
				group.GET("/queryAllDiffical", indexController.QueryAllDiffical)
				group.GET("/table", indexController.Table)
			})
			s.Group("/problem", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.POST("/changeDescription", problemController.ChangeDescription)
				group.POST("/changeSolution", problemController.ChangeSolution)
				group.GET("/queryDetail", problemController.QueryDetail)
			})
			s.Run()
			return nil
		},
	}
)
