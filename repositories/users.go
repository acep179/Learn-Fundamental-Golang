package repositories

// Import "dumbmerch/models", "gorm.io/gorm"
import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

// Declare UserRepository interface here ...
type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
}

// Declare repository struct here ...
type repository struct {
	db *gorm.DB
}

// Create RepositoryUser function here ...
func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

// Create FindUsers method here ...
func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

	return users, err
}

// Create GetUser method here ...
func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}
