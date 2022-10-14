package daos

import (
	"example/restful-apis/cmd/blueprint/config"
	"example/restful-apis/cmd/blueprint/models"
)

// UserDAO persists user data in database
type UserDAO struct{}

// NewUserDAO creates a new UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) Get(id uint) (*models.User, error) {
	var user models.User

	// Using `Gorm`
	err := config.Config.DB.Where("id = ?", id). // Do the query
							First(&user). // Make it scalar
							Error         // retrieve error or null
	return &user, err
}
