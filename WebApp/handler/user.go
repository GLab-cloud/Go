package handler

import (
	"github-trend-BE/model"
	"github-trend-BE/model/req"
	"net/http"

	"github.com/labstack/echo"
)

type UserHandler struct {
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	//return c.String(http.StatusOK,"welcome to Sign In")
	return c.JSON(http.StatusOK, echo.Map{
		"userName": "Batman",
		"email":    "tr.gmail.com",
	})
}
func (u *UserHandler) HandleSignUp(c echo.Context) error {
	req := req.ReqSignUp{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	type User struct {
		EmailUser string `json:"email"`
		FullName  string `json:"name"`
		Age       int    `json:"age"`
	}
	user := User{
		EmailUser: "tr.gmail.com",
		FullName:  "Superman",
		Age:       90,
	}
	//return c.String(http.StatusOK,"welcome to Sign Up")
	return c.JSON(http.StatusOK, user)
}
