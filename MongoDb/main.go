package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/GLab-cloud/Go/MongoDb/controllers"
	//"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main(){
  r:=httprouter.New()
  uc:=controllers.NewUserController(getSession())
  r.GET("/user/:id",uc.GetUser ) //curl "http://localhost:3000/user/6828191e7c530c8f2a80c278"
  r.POST("/user",uc.CreateUser ) //curl -X POST -H 'Content-Type:application/json' http://localhost:3000/user -d '{"Name":"mario", "Gender":"Male","Age":41}'
  r.DELETE("/user/:id",uc.DeleteUser )
  http.ListenAndServe("localhost:3000",r)

}
func getSession() *mgo.Session{
	//godotenv.Load()
   url:="mongodb+srv://tintin:9zgXJTFzRbAMgp4A@cluster0.5usukl0.mongodb.net/"
	// session,err := mgo.Dial(os.Getenv("MONGO_URI"))
	session,err := mgo.Dial(url)
	// db:=session.DB("")
	// names, err := db.CollectionNames()
	// println(names)
	if err!=nil{
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		//os.Exit(1)
	}
	fmt.Println("connect MONGO Db successful")

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