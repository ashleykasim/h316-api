package models

import (
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"h316-api/db"
)

type Menu struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title  string        `json:"title"`
	Date   string        `json:"start"`
	Url    string        `json:"url"`
	Meal   string        `json:"meal"`
	Recipe Recipe        `json:"recipe"`
}

func GetAllMenus(c *gin.Context) {
	menus := []Menu{}

	err := db.Db.C("menus").Find(nil).All(&menus)
	//TODO: better error handling.
	//mgo treats "not found" as an error condition and having panic()
	//here will cause the service to return a 500
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, menus)
}

func GetMenusByDate(c *gin.Context) {
	menus := []Menu{}
	d := c.Params.ByName("date")

	err := db.Db.C("menus").Find(bson.M{"date": d}).All(&menus)

	if err != nil {
		log.Println(err)
	}
	c.JSON(200, menus)
}

func CreateMenu(c *gin.Context) {
	menu := Menu{}
	c.BindJSON(&menu)

	err := db.Db.C("menus").Insert(&Menu{Title: menu.Title,
		Url:    menu.Url,
		Meal:   menu.Meal,
		Date:   menu.Date,
		Recipe: menu.Recipe})
	if err != nil {
		log.Println(err)
		c.JSON(500, err)
	} else {
		c.JSON(201, "created menu")
	}
}

func DeleteMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	err := db.Db.C("menus").RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		log.Println(err)
		c.JSON(500, err)
	} else {
		c.Status(204)
	}
}
