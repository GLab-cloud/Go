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
e.GET("/users/:id", getUser)
//e.PUT("/users/:id", updateUser)
//e.DELETE("/users/:id", deleteUser)
e.GET("/show", show)
e.Logger.Fatal(e.Start(":1323"))
//e.POST("/users", saveUser)

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
