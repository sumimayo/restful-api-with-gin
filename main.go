package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"example/web-service-gin/bookshelf/controllers"
	"reflect"
    "strconv"
	"fmt"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main () {
	router :=gin.Default()
	router.GET("/:id", getUser)
	router.GET("/albums", getAlbum)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}

func getUser (c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.JSON(400, err)
		return
	}
	if id <= 0 {
		c.JSON(400, gin.H{"status": "is should be bigger than 0"})
		return
	}
	ctrl := controllers.NewUser()
	fmt.Printf("---ctrl: %v\n---", ctrl)
	result := ctrl.Get(id)
	fmt.Printf("---result : %v\n----", result)
	if reflect.ValueOf(result).IsNil() {
		c.JSON(400, gin.H{})
		return
	}
	c.JSON(200, result)
}

func getAlbum (c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID (c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
