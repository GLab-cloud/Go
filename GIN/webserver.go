package main

import (

	"net/http"

	"github.com/gin-gonic/gin"
)
func main(){
 	r:=gin.Default()
	r.GET("/ping",func (context *gin.Context){
	// 		context.String(http.StatusOK,"Ping")
	//  })
	var data =map[string]interface{}{
	"message":"Hello from ping",
	}
	//gin.H{"message":"Hello from ping",}

	context.JSON(http.StatusOK,data)
	})
	r.POST("/ping",func(context *gin.Context){
		context.JSON(http.StatusOK,gin.H{
			"message":"Hello from POST ping",
			"message 1":"Hello from POST ping 1",

		})
	})
	r.GET("/detail/:id",getDetail)

	r.Run(":3333")
}
func getDetail(context *gin.Context){
	id:=context.Param("id")
	context.JSON(http.StatusOK,gin.H{
		"id":id,
	})
}