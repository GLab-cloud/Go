package repository
import(
	"context"
	//"github-trend-BE/log"
	"github-trend-BE/model"

)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User,error)
	CheckLogin(
		context context.Context,
		loginReq model.req.ReqSignIn,
	) (model.User, error)
}