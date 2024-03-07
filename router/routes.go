package router

import "github.com/gin-gonic/gin"

func InitializeRoutes(router *gin.Engine) {

	v1 := router.Group("api/v1")
	{
		//Show Oppenings
		v1.GET("/opening", func(ctx *gin.Context){
			ctx.JSON(200, gin.H{
				"message": "Get opening",
			})
		})

		v1.POST("/opening", func(ctx *gin.Context){
			ctx.JSON(200, gin.H{
				"message": "Post opening",
			})
		})
		v1.DELETE("/opening", func(ctx *gin.Context){
			ctx.JSON(200, gin.H{
				"message": "Delete opening",
			})
		})
		v1.PUT("/opening", func(ctx *gin.Context){
			ctx.JSON(200, gin.H{
				"message": "Put opening",
			})
		})
		v1.GET("/openings", func(ctx *gin.Context){
			ctx.JSON(200, gin.H{
				"message": "get openings",
			})
		})
	}

}
