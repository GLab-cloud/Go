package repo_implement

import (
	"context"
	"github-trend-BE/db"
	"github-trend-BE/log"
	"github-trend-BE/model"
	"github-trend-BE/repository"
	"time"

	pq "github.com/lib/pq"

	"github-trend-BE/banana"
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
_,err:=u.sql.Db.NamedExecContext(context,statement,user)
if err!=nil{
	log.Error(err.Error())
if err,ok:=err.(*pq.Error);ok{
	//	"23505": "unique_violation",
	if err.Code.Name()=="unique violation"{
		return user,banana.UserConflict
	}
}
return user,banana.SignUpFailed

}
return user,nil

}