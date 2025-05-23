package handler

import (
	"github-trend-BE/banana"
	"github-trend-BE/log"
	"github-trend-BE/middleware"
	"github-trend-BE/model"
	req2 "github-trend-BE/model/req"
	"github-trend-BE/repository"
	"github-trend-BE/security"
	"net/http"

	//validator "github.com/go-playground/validator/v10"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	//return c.String(http.StatusOK,"welcome to Sign In")
	// return c.JSON(http.StatusOK, echo.Map{
	// 	"userName": "Batman",
	// 	"email":    "tr.gmail.com",
	// })
	// req := req2.ReqSignIn{}
	// if err := c.Bind(&req); err != nil {
	// 	log.Error(err.Error())
	// 	return c.JSON(http.StatusBadRequest, model.Response{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    err.Error(),
	// 		Data:       nil,
	// 	})
	// }

	if err := c.Validate(middleware.Req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	//req := req2.ReqSignIn{}
	//c.Bind(&req)
	user, err := u.UserRepo.CheckLogin(c.Request().Context(), middleware.Req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	isTheSame := security.ComparePasswords(user.Password, []byte(middleware.Req.Password))
	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Sign In Failed!",
			Data:       nil,
		})
	}
	//gen token
	token,err:=security.GenToken(user)
	if err!=nil{
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token=token
	user.Password = ""
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Sign in successfull!",
		Data:       user,
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
	// validate := validator.New()

	// if err := validate.Struct(req); err != nil {
	// 	log.Error(err.Error())

	// 	return c.JSON(http.StatusBadRequest, model.Response{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    err.Error(),
	// 		Data:       nil,
	// 	})
	// }

	//verify user
	if err := c.Validate(req); err != nil {
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
	user := model.User{
		UserId:   userId.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		Role:     role,
		Token:    "",
	}

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	// type User struct {
	// 	EmailUser string `json:"email"`
	// 	FullName  string `json:"name"`
	// 	Age       int    `json:"age"`
	// }
	// user := User{
	// 	EmailUser: "tr.gmail.com",
	// 	FullName:  "Superman",
	// 	Age:       90,
	// }
	//return c.String(http.StatusOK,"welcome to Sign Up")
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	//gen token
	token,err:=security.GenToken(user)
	if err!=nil{
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token=token
	user.Password = ""
	//return c.JSON(http.StatusOK, user)
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Sign up user - insert DB successfull!",
		Data:       user,
	})
}
func (u *UserHandler) Profile (c echo.Context) error {
	//return nil
	tokenData:=c.Get("user").(*jwt.Token)
    claims:=tokenData.Claims.(*model.JwtCustomClaims)
	user,err:=u.UserRepo.SelectUserWithId(c.Request().Context(), claims.UserId)
	if err != nil {
		if err == banana.UserNotFound {
			return c.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
    user.Token=tokenData.Raw
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Get user profile from DB successfull!",
		Data:       user,
	})

}
func (u *UserHandler) UpdateProfile(c echo.Context) error {
	req := req2.ReqUpdateUser{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	// validate
	err := c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}
    //get UserId from auth JWT token
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	user := model.User{
		UserId:   claims.UserId,
		FullName: req.FullName,
		Email:    req.Email,
	}

	user, err = u.UserRepo.UpdateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Update user profile to DB successfull!",
		Data:       user,
	})
}