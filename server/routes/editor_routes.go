package routes

import (
	"github.com/a9jain/coral-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterCompilerRoutes(r *gin.Engine) {
	compiler := r.Group("/editor")
	{
		compiler.POST("/compile", controllers.TestCode)
		compiler.POST("/create_new_file", controllers.CreateNewFile)
		compiler.POST("/save_new_file", controllers.SaveFile)
		compiler.POST("/push_to_redis", controllers.PushToRedis)
		compiler.POST("/undo", controllers.Undo)
		compiler.POST("/redo", controllers.Redo)
		// Make sure to add in some frontend feature to allow user to select language
		compiler.POST("/change_language", controllers.ChangeLanguage)
		// make sure to add in some frontend feature to allow users to join in on sessions
		compiler.POST("/join_session", controllers.ChangeLanguage)
	}
}
