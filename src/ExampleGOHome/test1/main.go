package main

import (
	"log"
	"text/template"
	"time"

	"github.com/night-codes/mgo-ai"

	"github.com/gin-gonic/gin"
	ttpl "github.com/night-codes/go-gin-ttpl"
	mongo "github.com/night-codes/mgo-wrapper"
	"gopkg.in/mgo.v2"
	"gopkg.in/night-codes/types.v1"
)

type (
	obj map[string]interface{}
	arr []interface{}

	UsersStruct struct {
		collection *mgo.Collection
	}

	UserStruct struct {
		ID       uint64 `form:"id" json:"id" bson:"_id"`
		Name     string `form:"name" json:"name" bson:"name"`
		LastName string `form:"lastName" json:"lastName" bson:"lastName"`
	}
)

var (
	dbName = "myDB"
	port   = "8888"
	users  = UsersStruct{collection: mongo.DB(dbName).C("users")}
)

func init() {
	// init mongo autoincrement package
	ai.Connect(mongo.DB(dbName).C("ai"))
}

func main() {
	r := gin.Default()
	//var r *gin.Engine

	ttpl.Use(r, "templates/*", template.FuncMap{
		"FuncX": FuncX,
	})

	r.Static("files", "files")

	r.GET("/", Page)
	r.POST("/ajax/:module/:method", ajax)

	log.Println("Application started at :" + port)
	r.Run(":" + port)
}

func FuncX(x, y string) string {
	return x + "+" + y
}

func Page(c *gin.Context) {
	usersArr := []UserStruct{}

	if err := users.collection.Find(obj{}).All(&usersArr); err != nil {
		log.Panicln("Error Page:", err)
	}

	c.HTML(200, "index.html", obj{
		"users": usersArr,
	})
}

// chooise module
func ajax(c *gin.Context) {
	c.Header("Expires", time.Now().String())
	c.Header("Cache-Control", "no-cache")
	switch c.Param("module") {
	case "users":
		users.Ajax(c)
	default:
		c.String(400, "Module not found!")
	}
}

func (u *UsersStruct) Ajax(c *gin.Context) {
	switch c.Param("method") {
	case "add":
		u.Add(c)
	case "edit":
		u.Edit(c)
	case "getAll":
		u.GetAll(c)
	case "remove":
		u.Remove(c)

	default:
		c.String(400, "Method not found in module \"Leads\"!")
	}
}

func (u *UsersStruct) Add(c *gin.Context) {
	user := UserStruct{}
	if err := c.Bind(&user); err != nil {
		log.Printf("Error users.Add: %s", err)
	}

	user.ID = ai.Next("users")

	if err := u.collection.Insert(user); err != nil {
		log.Printf("Error(DB) users.Add: %s", err)
	}

	c.JSON(200, obj{"data": obj{"user": user}})
}

func (u *UsersStruct) Edit(c *gin.Context) {
	c.HTML(200, "index.html", obj{"Name": "Vasya", "LastName": "Ivanov"})
}

func (u *UsersStruct) GetAll(c *gin.Context) {
	c.HTML(200, "index.html", obj{"Name": "Vasya", "LastName": "Ivanov"})
}

func (u *UsersStruct) Remove(c *gin.Context) {
	id := types.Uint64(c.PostForm("id"))

	if err := u.collection.RemoveId(id); err != nil {
		log.Println("Error 213:", err)
	}

	c.JSON(200, obj{"data": obj{"id": id}})
}
