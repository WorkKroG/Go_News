package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"./models"
)

type NewsController struct {
	Id          bson.ObjectId `bson:"_id"`
	title       string        `bson:"title"`
	body        string        `bson:"body"`
	image       string        `bson:"image"`
	has_my_like bool          `bson:"has_my_like"`
}

func Connect(r *gin.Engine) {
	models.Connect(r)
}

func (t NewsController) Index(c *gin.Context) {
	tm := models.CreateNewsManager()
	c.JSON(200, tm.All())
}

func (t NewsController) Create(c *gin.Context) {
	tm := models.CreateNewsManager()
	var todo models.Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tm.Add(todo))
}

func (show NewsController) Show(c *gin.Context) {
	tm := models.CreateNewsManager()
	id := c.Params.ByName("id")

	c.JSON(200, tm.GetById(id))
}

func (t NewsController) Update(c *gin.Context) {
	tm := models.CreateNewsManager()
	id := c.Params.ByName("id")

	var todo models.Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tm.UpdateById(id, todo))
}

func (t NewsController) Destroy(c *gin.Context) {
	tm := models.CreateNewsManager()
	id := c.Params.ByName("id")

	tm.RemoveById(id)

	c.JSON(200, gin.H{})
}
