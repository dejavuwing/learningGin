package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/someGet", someMethod)
	r.POST("/somePost", someMethod)
	r.PUT("/somePut", someMethod)
	r.DELETE("/someDelete", someMethod)
	r.PATCH("/somePatch", someMethod)
	r.HEAD("/someHead", someMethod)
	r.OPTIONS("/someOptins", someMethod)

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		message := name + " is very handsome!"
		c.JSON(200, gin.H{"message": message})
	})

	r.GET("/user/:name/age/:old", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("old")
		message := name + " is " + age + " years old."
		c.JSON(200, gin.H{"message": message})
	})

	r.GET("/colour/:colour/*fruits", func(c *gin.Context) {
		color := c.Param("colour")
		fruits := c.Param("fruits")
		fruitArray := strings.Split(fruits, "/")
		// remove the first element in fruit slice.
		fruitArray = append(fruitArray[:0], fruitArray[1:]...)
		c.JSON(200, gin.H{"color": color, "fruits": fruitArray})
	})

	r.GET("/welcome", welcomeQueryString)
	r.POST("/form_post", formPost)
	r.POST("/form_post_with_querystring", formPostWithQueryString)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func someMethod(c *gin.Context) {
	httpMethod := c.Request.Method
	c.JSON(200, gin.H{"status": "good", "sending": httpMethod})
}

func welcomeQueryString(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.JSON(200, gin.H{"firstname": firstname, "lastname": lastname})
}

func formPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	headerType := c.GetHeader("Content-Type")

	c.JSON(200, gin.H{
		"status":              "posted",
		"message":             message,
		"nick":                nick,
		"header-content-type": headerType,
	})
}

func formPostWithQueryString(c *gin.Context) {
	id := c.Query("id")
	strPage := c.DefaultQuery("page", "0")
	intPage, _ := strconv.Atoi(strPage)

	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	headerType := c.GetHeader("Content-Type")

	c.JSON(200, gin.H{
		"status":              "posted",
		"message":             message,
		"nick":                nick,
		"header-content-type": headerType,
		"id":                  id,
		"page":                intPage,
	})
}
