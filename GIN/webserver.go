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
	name:=context.DefaultQuery("name","Guest")
	var data = map[string]interface{}{
	"message":"Hello "+name+" from ping",
	}
	//gin.H{"message":"Hello from ping",}

	context.JSON(http.StatusOK,data)
	})
	r.POST("/ping",func(context *gin.Context){
		address:=context.DefaultPostForm("addr","Viet Nam")
		age:=context.DefaultPostForm("age","18")

		context.JSON(http.StatusOK,gin.H{
			"message ":"Hello from" +address+" ping 1",
			"message 1":"Hello from"+ age +" POST ping",

		})
	})
	//curl -X POST "http://localhost:3333/ping" -H "Content-Type: application/x-www-form-urlencoded" -d "addr=TP.HCM&age=41"

	r.GET("/detail/:id",getDetail)

	r.Run(":3333")
}
func getDetail(context *gin.Context){
	id:=context.Param("id")
	context.JSON(http.StatusOK,gin.H{
		"id":id,
	})
}