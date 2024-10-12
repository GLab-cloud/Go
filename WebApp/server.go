package main

import (
	//"fmt"
	"github-trend-BE/handler"
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
//Init(e)
e.GET("/",welcome)
//users APIs
e.GET("/user/sign-in",handler.HandleSignIn)
e.GET("/user/sign-up",handler.HandleSignUp)

e.GET("/users/:id", getUser)
//e.PUT("/users/:id", updateUser)
//e.DELETE("/users/:id", deleteUser)
e.GET("/show", show)
e.Logger.Fatal(e.Start(":1000"))
//e.POST("/users", saveUser)

}
func welcome (c echo.Context) error{
	return c.String(http.StatusOK,"welcome to My App 1 f")
}
func Init(e *echo.Echo) {
	panic("unimplemented")
}
// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
  return c.String(http.StatusOK, id)
}


//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}
