package router

import (
	"github-trend-BE/handler"

	"github.com/labstack/echo"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	//users APIs
	//group
	g:=api.Echo.Group('/user',middleware.AddTrailingSlash())
	g.POST("/sign-in", api.UserHandler.HandleSignIn)
	g.POST("/sign-up", api.UserHandler.HandleSignUp)
}
