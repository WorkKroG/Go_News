package models

import (
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"./db"
)

type Todo struct {
	Id          bson.ObjectId `bson:"_id"`
	title       string        `bson:"title"`
	body        string        `bson:"body"`
	image       string        `bson:"image"`
	has_my_like bool          `bson:"has_my_like"`
}

type NewsManager struct {
	C *mgo.Collection
}

func Connect(r *gin.Engine) {
	db.Connect(r)
}

func (tm *NewsManager) All() []Todo {
	todoList := []Todo{}

	tm.C.Find(nil).All(&todoList)
	return todoList
}

func (tm *NewsManager) Add(todo Todo) Todo {
	tm.C.Insert(todo)

	return todo
}

func (tm *NewsManager) GetById(id string) Todo {
	todo := Todo{}
	tm.C.FindId(bson.ObjectIdHex(id)).One(&todo)

	return todo
}

func (tm *NewsManager) UpdateById(id string, todo Todo) Todo {
	tm.C.UpdateId(bson.ObjectIdHex(id), todo)
	return todo
}

func (tm *NewsManager) RemoveById(id string) {
	tm.C.RemoveId(bson.ObjectIdHex(id))
}

func CreateNewsManager() *NewsManager {
	tm := new(NewsManager)
	tm.C = db.DB.C("News")

	return tm
}
