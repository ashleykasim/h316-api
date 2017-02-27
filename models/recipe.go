package models

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"recipes-api/db"
)

//TODO: what fields should be stored for projects?
type Recipe struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Created     time.Time     `json:"created"`
	Updated     time.Time     `json:"updated"`
	Description string        `json:"description"`
	Directions  string        `json:"directions"`
	ImageUrl    string        `json:"imageUrl"`
}

func GetAllRecipes(c *gin.Context) {
	recipes := []Recipe{}

	err := db.Db.C("recipes").Find(nil).All(&recipes)
	//TODO: better error handling.
	//mgo treats "not found" as an error condition and having panic()
	//here will cause the service to return a 500
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, recipes)
}

func GetRecipeById(c *gin.Context) {
	recipe := Recipe{}
	id := c.Params.ByName("id")

	recipes := []Recipe{}

	err := db.Db.C("recipes").Find(nil).All(&recipes)

	for _, p := range recipes {
		if p.Id == bson.ObjectIdHex(id) {
			recipe = p
		}
	}

	// err := db.Db.C("recipes").FindId(bson.ObjectIdHex(id)).One(&recipe)
	if err != nil {
		log.Println(err)
	}
	log.Println(recipe.Id)
	if recipe == (Recipe{}) {
		c.JSON(404, gin.H{"error": "recipe not found"})
	} else {
		c.JSON(200, recipe)
	}
}

func CreateRecipe(c *gin.Context) {
	recipe := Recipe{}
	c.BindJSON(&recipe)

	err := db.Db.C("recipes").Insert(&Recipe{Name: recipe.Name,
		Description: recipe.Description,
		Directions:  recipe.Directions,
		ImageUrl:    recipe.ImageUrl,
		Created:     time.Now()})
	if err != nil {
		log.Println(err)
		c.JSON(500, err)
	} else {
		c.JSON(201, "created recipe")
	}
}

func DeleteRecipe(c *gin.Context) {
	id := c.Params.ByName("id")
	err := db.Db.C("recipes").RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		log.Println(err)
		c.JSON(500, err)
	} else {
		c.Status(204)
	}
}

func UpdateRecipe(c *gin.Context) {
	id := c.Params.ByName("id")
	recipe := Recipe{}
	c.BindJSON(&recipe)

	//validate which fields have changed to avoid setting unchanged fields to null
	//TODO: is there a better way to do this?
	orig := Recipe{}
	err1 := db.Db.C("recipes").FindId(bson.ObjectIdHex(id)).One(&orig)
	if err1 != nil {
		log.Println(err1)
		c.JSON(500, err1)
	}
	if recipe.Name == "" {
		recipe.Name = orig.Name
	}
	if recipe.Description == "" {
		recipe.Description = orig.Description
	}
	if recipe.Directions == "" {
		recipe.Directions = orig.Directions
	}
	if recipe.ImageUrl == "" {
		recipe.ImageUrl = orig.ImageUrl
	}
	recipe.Created = orig.Created

	err2 := db.Db.C("recipes").UpdateId(bson.ObjectIdHex(id), &Recipe{Name: recipe.Name, Description: recipe.Description, Directions: recipe.Directions, ImageUrl: recipe.ImageUrl, Created: recipe.Created, Updated: time.Now()})
	if err2 != nil {
		log.Println(err2)
		c.JSON(500, err2)
	} else {
		c.Status(204)
	}
}
