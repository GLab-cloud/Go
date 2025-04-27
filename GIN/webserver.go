package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
func main(){
 	r:=gin.Default()
	r.Static("/static_file","./assets")
	r.Use(CORSMiddleware())
	r.Use(MyCustomMiddleware)
	r.Use(NewCustomMiddleware())

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

	//Group API
	api:=r.Group("/api")
	{
		v1:=api.Group("/v1")
		{
			v1.Use(MyGroupV1Middleware())
			v1.GET("/ping",func(context *gin.Context){
				context.JSON(http.StatusOK,gin.H{
					"message ":"Hello from api/v1/ping",
				})
			})
			v1.GET("/pong",func(context *gin.Context){
				context.JSON(http.StatusOK,gin.H{
					"message ":"Hello from api/v1/pong",
				})
			})
		}
		v2:=api.Group("/v2")
		{
			v2.Use(MyGroupV2Middleware())
			v2.GET("/ping",func(context *gin.Context){
				context.JSON(http.StatusOK,gin.H{
					"message ":"Hello from api/v2/ping",
				})
			})
			v2.GET("/pong",func(context *gin.Context){
				context.JSON(http.StatusOK,gin.H{
					"message ":"Hello from api/v2/pong",
				})
			})
		}
	}
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
	if(true){ // check if authentication condition,...
		context.Next() // CustomMiddleware -> next handler
	}
}
func NewCustomMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		log.Println("I'm a new global middleware - HandlerFunc ")

		if(true){ // check if authentication condition,...
			ctx.Next() // CustomMiddleware -> next handler
		}
	}
}
//CORS Middleware
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		log.Println("I'm a CORS Middleware - HandlerFunc ")

        c.Writer.Header().Set("Content-Type", "application/json")
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(200)
        } else {
            c.Next()
        }
    }
}
//Group Middleware
func MyGroupV1Middleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		log.Println("I'm in a group v1 middleware")
		ctx.Next()
	}
}
func MyGroupV2Middleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		log.Println("I'm in a group v2 middleware")
		ctx.Next()
	}
}