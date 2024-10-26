package repo_implement

type UserRepoImp struct {
	sql *db.Sql

}
func NewUserRepo(sql *db.Sql ) repository.UserRepo{
	return &UserRepoImp{
		sql:sql,
	}

}
func (u *UserRepoImp)SaveUser(context context.Context, user model.User) (model.User,error){

}