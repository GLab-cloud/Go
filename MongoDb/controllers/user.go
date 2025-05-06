package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)
type UserController struct {
	session *mgo.Session
}
func NewUserController( s *mgo.Session) *UserController{
  return &UserController{s}
}
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){

}
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){

}
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){

}