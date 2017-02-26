package models

import (
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"recipes-api/db"
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

// func UpdateTitle(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	Title := Title{}
// 	c.BindJSON(&Title)
//
// 	//validate which fields have changed to avoid setting unchanged fields to null
// 	//TODO: is there a better way to do this?
// 	orig := Title{}
// 	err1 := db.Db.C("Titles").FindId(bson.ObjectIdHex(id)).One(&orig)
// 	if err1 != nil {
// 		log.Println(err1)
// 		c.JSON(500, err1)
// 	}
// 	if Title.Name == "" {
// 		Title.Name = orig.Name
// 	}
// 	if Title.Description == "" {
// 		Title.Description = orig.Description
// 	}
// 	if Title.Directions == "" {
// 		Title.Directions = orig.Directions
// 	}
// 	Title.Created = orig.Created
//
// 	err2 := db.Db.C("Titles").UpdateId(bson.ObjectIdHex(id), &Title{Name: Title.Name, Description: Title.Description, Directions: Title.Directions, Created: Title.Created, Updated: time.Now()})
// 	if err2 != nil {
// 		log.Println(err2)
// 		c.JSON(500, err2)
// 	} else {
// 		c.Status(204)
// 	}
// }
