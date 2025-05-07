package main

import (
	"net/http"

	"os"

	"github.com/GLab-cloud/Go/MongoDb/controllers"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main(){
  r:=httprouter.New()
  uc:=controllers.NewUserController(getSession())
  r.GET("/user/:id",uc.GetUser )
  r.POST("/user",uc.CreateUser )
  r.DELETE("/user/:id",uc.DeleteUser )
  http.ListenAndServe("localhost:9000",r)

}
func getSession() *mgo.Session{
	godotenv.Load()

	session,err := mgo.Dial(os.Getenv("MONGO_URI"))
	println(session.DB(""))
	if err==nil{
	}
	return session
}