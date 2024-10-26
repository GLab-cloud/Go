package router

import (
	"github-trend-BE/handler"

	"github.com/labstack/echo"
)

type API struct {
	Echo           *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	//users APIs
	api.Echo.GET("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)

}
