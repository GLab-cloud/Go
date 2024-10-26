package repo_implement

import (
	"context"
	//"github-trend-BE/log"
	"github-trend-BE/db"
	"github-trend-BE/model"
	"github-trend-BE/repository"
)
type UserRepoImp struct {
	sql *db.Sql

}
func NewUserRepo(sql *db.Sql ) repository.UserRepo{
	return &UserRepoImp{
		sql:sql,
	}

}
func (u *UserRepoImp)SaveUser(context context.Context, user model.User) (model.User,error){
statement := `
INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
`
user.CreatedAt = time.Now()
user.UpdatedAt = time.Now()
}