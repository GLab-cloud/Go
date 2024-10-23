package router

import (
	"github-trend-BE/handler"
	"os/user"

	"github.com/labstack/echo"
)

type API struct {
	e *echo.Echo
	UserHandler handler.UserHandler
}
func (api *API) SetupRouter (){
//users APIs
api.e.GET("/user/sign-in",api.UserHandler.HandleSignIn)
api.e.GET("/user/sign-up",api.UserHandler.HandleSignUp)


}