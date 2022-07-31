package main

import (
	"log"
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

	router.POST("/books", postBooksHandler)

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

type BookInput struct {
	Title    string
	Price    int
	SubTitle string `json:"sub_title"` //menangkap json yang bernama sub_title, jadi bisa beda penulisan
}

func postBooksHandler(c *gin.Context) {
	var Book BookInput

	err := c.ShouldBindJSON(&Book)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"title":     Book.Title,
		"price":     Book.Price,
		"sub_title": Book.SubTitle,
	})
}
