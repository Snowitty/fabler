package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/snowitty/fabler/conf"
	"github.com/snowitty/fabler/internal/controller/rest"
	"github.com/snowitty/fabler/internal/controller/web"
	"github.com/snowitty/fabler/internal/middleware"
	"github.com/snowitty/fabler/internal/view"
)

func SetupAPIRouter(r *gin.Engine) {

	r.Use(middleware.Logger())

	GroupAPI := r.Group("/api")
	{
		//group web
		GroupWeb := GroupAPI.Group("/web")
		{
			GroupWeb.POST("/security/sign", web.SecuritySign)
			GroupWeb.POST("/security/login", web.SecurityLogout)
			GroupWeb.POST("/security/logout", web.SecurityLogout)

			GroupWeb.POST("/question", middleware.Authorizer(), web.CreateQuestion)

			GroupWeb.POST("/answer", middleware.Authorizer(), web.CreateAnswer)
			GroupWeb.POST("/answer/:id/supporters", middleware.Administartor(), web.CreateSupporter)
			GroupWeb.DELETE("/answer/:id/supporters", middleware.Authorizer(), web.DeleteSupporter)

			GroupWeb.POST("/comment", middleware.Authorizer(), web.CreateComment)

			GroupWeb.POST("/reply", middleware.Authorizer(), web.CreateReply)

			GroupWeb.POST("/file/avatar", middleware.Authorizer(), web.FileAvatarResolve)

		}

		//group rest
		GroupRest := GroupAPI.Group("/rest")
		GroupRest.Use(middleware.Authorizer())
		{

			GroupRest.GET("/user/:id", rest.GetUser)
			GroupRest.GET("/users", rest.GetUsers)
			GroupRest.GET("/users/counts", rest.GetUsersCounts)
			GroupRest.PUT("/user/:id", rest.UpdateUser)
			GroupRest.DELETE("/user/:id", rest.DeleteUser)

			GroupRest.GET("/profile/:id", rest.GetProfile)
			GroupRest.GET("/profiles", rest.GetProfiles)
			GroupRest.GET("/profiles/counts", rest.GetProfilesCounts)
			GroupRest.PUT("/profile/:id", rest.UpdateProfile)
			GroupRest.DELETE("/profile/:id", rest.DeleteProfile)

			GroupRest.GET("/question/:id", rest.GetQuestion)
			GroupRest.GET("/questions", rest.GetQuestions)
			GroupRest.GET("/questions/counts", rest.GetQuestionsCounts)
			GroupRest.PUT("/question/:id", rest.UpdateQuestion)
			GroupRest.DELETE("/question/:id", rest.DeleteQuestion)

			GroupRest.GET("/answer/:id", rest.GetAnswer)
			GroupRest.GET("/answers", rest.GetAnswers)
			GroupRest.GET("/answers/counts", rest.GetAnswersCounts)
			GroupRest.PUT("/answer/:id", rest.UpdateAnswer)
			GroupRest.DELETE("/answer/:id", rest.DeleteAnswer)

			GroupRest.GET("/comment/:id", rest.GetComment)
			GroupRest.GET("/comments", rest.GetComments)
			GroupRest.GET("/comments/counts", rest.GetCommentsCounts)
			GroupRest.PUT("/comment/:id", rest.UpdateComment)
			GroupRest.DELETE("/comment/:id", rest.DeleteComment)

			GroupRest.GET("/reply/:id", rest.GetReply)
			GroupRest.GET("/replies", rest.GetReplies)
			GroupRest.GET("/replies/counts", rest.GetRepliesCounts)
			GroupRest.PUT("/reply/:id", rest.UpdateReply)
			GroupRest.DELETE("/reply/:id", rest.DeleteReply)

		}
	}

}

func SetupViewRouter(r *gin.Engine) {
	{
		r.Delims("\"/{{", "}}/\"")
		r.LoadHTMLGlob("web/*.html")
		r.Static("/static", "web/staitc")

		r.GET("/", middleware.Authorizer(), view.Index)
		r.GET("/profile", middleware.Authorizer(), view.Profile)
		r.GET("/answer", middleware.Authorizer(), view.Answer)
		r.GET("/admin", middleware.Authorizer(), middleware.Administartor(), view.Admin)
		r.GET("/login", view.Login)
		r.GET("/error", view.Error)
	}
}

func main() {
	var shouldInit = flag.Bool("init", false, "initialize all")
	flag.Parse()

	if *shouldInit {
		initAll(conf.Config())
	}

	r := gin.Default()
	SetupAPIRouter(r)
	SetupViewRouter(r)

	r.Run(conf.Config().Address)
}
