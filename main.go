package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Gempur Bayu Aji",
			"bio":  "A Software Engineer",
		})
	})

	router.GET("/hello", helloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)

	//custom port :8888
	router.Run()
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":  "Hello Guys!",
		"number": 2,
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

// localhost:8080/query?title=[title]&name=[name]
func queryHandler(c *gin.Context) {
	title := c.Query("title")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{"title": title, "name": name})
}
