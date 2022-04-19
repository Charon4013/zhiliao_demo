package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"zhiliao_mvc_demo/controller"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	// middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(Cors)

	common := app.Party("/")
	{
		common.Options("*", func(ctx iris.Context) {
			ctx.Next()
		})
	}

	mvc.New(app.Party("/user")).Handle(controller.NewUserController())
	mvc.New(app.Party("/user/{uid}")).Handle(controller.NewUserInfoController())
	mvc.New(app.Party("/main")).Handle(controller.NewMainPageController())
	mvc.New(app.Party("/question")).Handle(controller.NewQuestionController())
	mvc.New(app.Party("/question/{qid}")).Handle(controller.NewAnswerController())
	mvc.New(app.Party("/question/{qid}/answer/{aid}")).Handle(controller.NewAnswerSupportController())
	mvc.New(app.Party("/question/{qid}/answer/{aid}/comment")).Handle(controller.NewCommentController())
	//mvc.New(app.Party("/auth")).Handle(controller.NewAuthController())
	mvc.New(app.Party("/manage")).Handle(controller.NewManageController())

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

// Cors : Solve CORS problem
func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization, X-Requested-With, Token")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
