package main

import (
	//"fmt"
	"net/http"
	"github.com/labstack/echo"
)

// func helloWorldPage(w http.ResponseWriter, r *http.Request ){
// 	fmt.Fprint(w,"Hello World")
// }

func main() {
// 	http.HandleFunc("/",helloWorldPage)
// 	http.ListenAndServe(":8080",nil)
e:=echo.New()
e.GET("/",func(c echo.Context) error{
	return c.String(http.StatusOK,"Hello, World")
})
e.Logger.Fatal(e.Start(":1323"))
}
