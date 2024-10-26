package repository

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User,error)
}