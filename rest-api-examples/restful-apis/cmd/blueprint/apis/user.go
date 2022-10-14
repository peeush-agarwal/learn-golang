package apis

import (
	"log"
	"net/http"
	"strconv"

	"example/restful-apis/cmd/blueprint/daos"
	"example/restful-apis/cmd/blueprint/services"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func GetUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
