package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}

	routerGroupApp.SettingsRouter()
	routerGroupApp.ImagesRouter()
	routerGroupApp.AdvertiseRouter()
	routerGroupApp.MenuRouter()
	routerGroupApp.UserRouter()
	routerGroupApp.TagRouter()
	routerGroupApp.MessageRouter()
	routerGroupApp.ArticleRouter()
	routerGroupApp.DiggRouter()
	routerGroupApp.CommentRouter()

	return router
}
