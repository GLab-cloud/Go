package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
func main(){
 	r:=gin.Default()
	r.Use(MyCustomMiddleware)
	r.GET("/",func (context *gin.Context){
		// 		context.String(http.StatusOK,"Ping")
		//  })
		log.Println("I'm in index handler")

		var data = map[string]interface{}{
		"message":"Hello from Golang web server",
		}
		//gin.H{"message":"Hello from ping",}

		context.JSON(http.StatusOK,data)
	})

	r.GET("/ping",func (context *gin.Context){
	// 		context.String(http.StatusOK,"Ping")
	//  })
	log.Println("I'm in get Ping handler")

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
	log.Println("I'm in get Detail handler")

	id:=context.Param("id")
	context.JSON(http.StatusOK,gin.H{
		"id":id,
	})
}
//Middleware
func MyCustomMiddleware(context *gin.Context){
	log.Println("I'm a global middleware")
	context.Next()
}