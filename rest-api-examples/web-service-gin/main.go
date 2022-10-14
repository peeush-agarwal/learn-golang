package main

import (
	"net/http"

	_ "example/web-service-gin/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	"github.com/swaggo/gin-swagger"        // gin-swagger middleware
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// GetAlbums godoc
// @Summary Retrieves list of album
// @Produce json
// @Success 200 {array} album
// @Router  /albums [get]
func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

// PostAlbums godoc
// @Summary Creates album for the specified request
// @Accept  json
// @Produce json
// @Param   album body     album true "New Album"
// @Success 200   {object} album
// @Router  /albums [post]
func postAlbums(ctx *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID godoc
// @Summary Retrieves album based on given ID
// @Produce json
// @Param   id  path     integer true "Album ID"
// @Success 200 {object} album
// @Failure 404 {object} string
// @Router  /albums/{id} [get]
func getAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, album := range albums {
		if album.ID == id {
			ctx.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// @title          Example Swagger API
// @version        1.0
// @description    Swagger API for Golang Project Web service.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.email abc@xyz.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:5000
// @BasePath /api/v1
func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		albums := v1.Group("/albums")
		{
			albums.GET("", getAlbums)
			albums.GET(":id", getAlbumByID)
			albums.POST("", postAlbums)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":5000")
}
