package handler

import (
	"github-trend-BE/log"
	"github-trend-BE/model"
	req2 "github-trend-BE/model/req"
	"github-trend-BE/security"

	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
	//bind request body to struct req
	req := req2.ReqSignUp{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	//verify user - validator
	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())

		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	//USE BCRYPT NOT MD5 TO HASH PASSWORD
	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()
	userId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())

		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
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
