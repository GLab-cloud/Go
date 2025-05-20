package controllers

import (
	"context"
	"encoding/json"

	//"errors"
	"fmt"
	//"log"
	"net/http"

	"github.com/GLab-cloud/Go/MongoDb/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"gopkg.in/mgo.v2"
)
type UserController struct {
	client *mongo.Client
}
func NewUserController(client *mongo.Client) *UserController{
  return &UserController{client}
}
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id:=p.ByName("id")
	fmt.Println("get user param id = "+id)
	u:=models.User{}
    oid,err:=bson.ObjectIDFromHex(id)
	if err!=nil{
		fmt.Println("id is not valid BSON ObjectID")
	}
	err1:=uc.client.Database("mongo-golang").Collection("users").FindOne(context.TODO(),bson.D{{"_id",oid}}).Decode(&u)

	if err1 == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the id %s\n",id)
		return
	}
	uj,err:= json.Marshal(u)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(u)
	fmt.Fprintf(w,"%s\n",uj)// response json data to client

}


func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u:=models.User{}
	//u.Id=bson.NewObjectID()
	uj:=json.NewDecoder(r.Body).Decode(&u)
	doc:= models.User{Id: bson.NewObjectID(), Name: "Na", Gender: "Female",Age: 1}
	result, err := uc.client.Database("mongo-golang").Collection("users").InsertOne(context.TODO(),doc)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w,"%s\n",uj)
}
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id:=p.ByName("id")
	// if !bson.ObjectIDFromHex(id){
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	oid,err:=bson.ObjectIDFromHex(id)
	if err!=nil{
		fmt.Println("id is not valid")
	}
	if deleteResult,err:=uc.client.Database("mongo-golang").Collection("users").DeleteOne(context.TODO(),oid);err!=nil{
		w.WriteHeader(404)
        fmt.Println(deleteResult)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Deleted user ",id,"\n")
}
// package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/GLab-cloud/Go/MongoDb/models"
// 	"github.com/julienschmidt/httprouter"
// 	"gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// )
// type UserController struct {
// 	client *mgo.Session
// }
// func NewUserController( s *mgo.Session) *UserController{
//   return &UserController{s}
// }
// func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
// 	id:=p.ByName("id")
// 	fmt.Println("get user param id = "+id)
// 	if !bson.IsObjectIdHex(id){
// 		w.WriteHeader(http.StatusNotFound)
// 	}
// 	oid:=bson.ObjectIdHex(id)
// 	u:=models.User{}
// 	if err:=uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u);err!=nil{
// 		w.WriteHeader(404)
// 		return
// 	}
// 	uj,err:= json.Marshal(u)
// 	if err!=nil{
// 		fmt.Println(err)
// 	}
// 	w.Header().Set("Content-Type","application/json")
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w,"%s\n",uj)
// }
// func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
// 	u:=models.User{}
// 	json.NewDecoder(r.Body).Decode(&u)
// 	u.Id=bson.NewObjectId()
// 	uc.session.DB("mongo-golang").C("users").Insert(u)
// 	uj,err:= json.Marshal(u)
// 	if err!=nil{
// 		fmt.Println(err)
// 	}
// 	w.Header().Set("Content-Type","application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprintf(w,"%s\n",uj)
// }
// func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
// 	id:=p.ByName("id")
// 	if !bson.IsObjectIdHex(id){
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	oid:=bson.ObjectIdHex(id)
// 	if err:=uc.session.DB("mongo-golang").C("users").RemoveId(oid);err!=nil{
// 		w.WriteHeader(404)

// 	}
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w,"Deleted user ",oid,"\n")
// }