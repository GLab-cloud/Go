package main

import (
	//"fmt"
	"github-trend-BE/db"
	"github-trend-BE/handler"
	"github-trend-BE/repository/repo_implement"
	"github-trend-BE/router"

	"net/http"

	"github.com/labstack/echo"
)

// func helloWorldPage(w http.ResponseWriter, r *http.Request ){
// 	fmt.Fprint(w,"Hello World")
// }

func main() {
	// 	http.HandleFunc("/",helloWorldPage)
	// 	http.ListenAndServe(":8080",nil)

	//connect Db
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "123",
		DbName:   "golang",
	}
	sql.Connect()
	defer sql.Close()

	//start server
	e := echo.New()

    //middleware
	e.Use(middleware.AddTrailingSlash())

	//handler
	userHandler := handler.UserHandler{
		UserRepo: repo_implement.NewUserRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()
	//Init(e)
	e.GET("/", handler.Welcome)

	e.GET("/users/:id", getUser)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":1000"))
	//e.POST("/users", saveUser)

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

// e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
