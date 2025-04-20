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

r.Run(":3333")
}