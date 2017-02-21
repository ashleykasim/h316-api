package models

import (
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"recipes-api/db"
)

type Menu struct {
	Id                     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Recipe                 string        `json:"title"`
	Date                   string        `json:"start"`
	Url						         string				 `json:"url"`
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

// func GetRecipeById(c *gin.Context) {
// 	recipe := Recipe{}
// 	id := c.Params.ByName("id")
//
// 	recipes := []Recipe{}
//
// 	err := db.Db.C("recipes").Find(nil).All(&recipes)
//
// 	for _, p := range recipes {
// 		if p.Id == bson.ObjectIdHex(id) {
// 			recipe = p
// 		}
// 	}
//
// 	// err := db.Db.C("recipes").FindId(bson.ObjectIdHex(id)).One(&recipe)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println(recipe.Id)
// 	if recipe == (Recipe{}) {
// 		c.JSON(404, gin.H{"error": "recipe not found"})
// 	} else {
// 		c.JSON(200, recipe)
// 	}
// }

func CreateMenu(c *gin.Context) {
  menu := Menu{}
	c.BindJSON(&menu)

	err := db.Db.C("menus").Insert(&Menu{Recipe: menu.Recipe,
		Url:    menu.Url,
		Date:   menu.Date})
	if err != nil {
		log.Println(err)
		c.JSON(500, err)
	} else {
		c.JSON(201, "created menu")
	}
}

// func DeleteRecipe(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	err := db.Db.C("recipes").RemoveId(bson.ObjectIdHex(id))
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(500, err)
// 	} else {
// 		c.Status(204)
// 	}
// }
//
// func UpdateRecipe(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	recipe := Recipe{}
// 	c.BindJSON(&recipe)
//
// 	//validate which fields have changed to avoid setting unchanged fields to null
// 	//TODO: is there a better way to do this?
// 	orig := Recipe{}
// 	err1 := db.Db.C("recipes").FindId(bson.ObjectIdHex(id)).One(&orig)
// 	if err1 != nil {
// 		log.Println(err1)
// 		c.JSON(500, err1)
// 	}
// 	if recipe.Name == "" {
// 		recipe.Name = orig.Name
// 	}
// 	if recipe.Description == "" {
// 		recipe.Description = orig.Description
// 	}
// 	if recipe.Directions == "" {
// 		recipe.Directions = orig.Directions
// 	}
// 	recipe.Created = orig.Created
//
// 	err2 := db.Db.C("recipes").UpdateId(bson.ObjectIdHex(id), &Recipe{Name: recipe.Name, Description: recipe.Description, Directions: recipe.Directions, Created: recipe.Created, Updated: time.Now()})
// 	if err2 != nil {
// 		log.Println(err2)
// 		c.JSON(500, err2)
// 	} else {
// 		c.Status(204)
// 	}
// }
