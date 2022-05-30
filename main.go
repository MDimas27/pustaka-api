package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/hello", helloHandler)

	router.GET("/books/:id", booksHandler)	

	router.GET("/query", queryHandler)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"name": "Dimas Surya",
		"bio": "Software Engineer and Remoter Worker",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"title": "Belajar Golang",
		"subtitle": "Belajar Golang bareng-bareng",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{"title": title})
	
}