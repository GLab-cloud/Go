package main

import (
	"net/http"
	"time"

	"github.com/GLab-cloud/Go/MongoDb/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"github.com/joho/godotenv"
    "os"
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
	println(os.Getenv("MONGO_URI"))
	session, err := mgo.DialWithTimeout(os.Getenv("MONGO_URI"),60 * time.Second)
	if err!=nil{
		panic(err)
	}
	return session
}