package db

import (
	"fmt"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
	DB      *mgo.Database
)

const (
	MongoServerURL = "mongodb://127.0.0.1"
)

func Connect(r *gin.Engine) {
	session, err := mgo.Dial(MongoServerURL)

	if err != nil {
		fmt.Printf("Can't connect to mongoDB. %v\n", err)
		panic(err.Error())
	}

	session.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", MongoServerURL)

	db := session.DB("admin")

	Session = session
	DB = db
}
