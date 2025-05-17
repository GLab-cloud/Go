package main

import (
	"fmt"
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
	fmt.Println("connect MONGO Db successful")
	// db:=session.DB("")
	// names, err := db.CollectionNames()
	// println(names)
	if err!=nil{
		fmt.Println("connect MONGO Db failed!!!")
	}
	return session
}

// type userInfo struct {
// 	username string `bson:"username"`
// 	pass     string    `bson:"pass"`
//   }

//   func connect() (session *mgo.Session) {
// 	  connectURL := "localhost:27017"
// 	  session, err := mgo.Dial(connectURL)
// 	  if err != nil {
// 		  fmt.Printf("Can't connect to mongo, go error %v\n", err)
// 		  os.Exit(1)
// 	  }
// 	  session.SetSafe(&mgo.Safe{})
// 	  return session
//   }

//   func main() {
// 	session := connect()
// 	defer session.Close()

// 	result := userInfo {}
// 	collection := session.DB("test_db").C("userlogin")
// 	err := collection.Find(bson.M{"username": "super_sam"}).One(&result)
// 		  if err != nil {
// 				  log.Println("Not in db")
// 				  return
// 		  }

// 		  fmt.Println("Username:%s", result.pass)
//   }