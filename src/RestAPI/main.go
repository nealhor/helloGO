package main

import (
	"io/ioutil"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"github.com/night-codes/mgo-wrapper"
	)
// http://localhost:8080/
func HomePage (c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HELLO WORLD",
	})
}

func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}

func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}

func main() {
	fmt.Println("Hello")

	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryStrings) // /query?name=Dirx&age=24
	r.GET("/path/:name/:age", PathParameters) // /path/dirx/24/
	r.Run()
}