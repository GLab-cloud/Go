package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func HandleSignIn(c echo.Context) error {
	//return c.String(http.StatusOK,"welcome to Sign In")
	return c.JSON(http.StatusOK, echo.Map{
		"userName": "Batman",
		"email":    "tr.gmail.com",
	})
}
func HandleSignUp(c echo.Context) error {
	type User struct {
		Email    string
		FullName string
		Age      int
	}
	user := User{
		Email:    "tr.gmail.com",
		FullName: "Superman",
		Age:      90,
	}
	//return c.String(http.StatusOK,"welcome to Sign Up")
	return c.JSON(http.StatusOK, user)
}