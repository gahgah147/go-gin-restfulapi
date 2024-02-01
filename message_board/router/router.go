package router

import (
	"message/controller"
	"message/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Swagger 註解

// @title Gin swagger Restful API 留言板範例
// @version 1.0
// @description Gin Restful API 留言板範例 swagger 文件

// @contact.name nalson
// @contact.url https://gahgah147.github.io/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8082
// schemes http
func SetRouter() *gin.Engine {
	//顯示 debug 模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	v1 := r.Group("api/v1")
	{
		//新增留言
		v1.POST("/message", controller.Create)
		//查詢全部留言
		v1.GET("/message", controller.GetAll)
		//查詢 {id} 留言
		v1.GET("/message/:id", controller.Get)
		//修改 {id} 留言
		v1.PATCH("/message/:id", controller.Update)
		//刪除 {id} 留言
		v1.DELETE("/message/:id", controller.Delete)
	}
	// 設定swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
