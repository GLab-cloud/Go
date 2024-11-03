package repository

import (
	"context"
	//"github-trend-BE/log"
	"github-trend-BE/model"
	"github-trend-BE/model/req"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	SelectUserWithId(context context.Context, userId string) (model.User, error)
}
