package repositories

import (
	"dumbmerch/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	// Declare UpdateUser interface here ...
	UpdateUser(user models.User, ID int) (models.User, error)
	//ctt jika kita berbicara terkait update, tentu saja kita membutuhkan dua parameter sekaligus. Adapun parameter yg dibutuhkan adalah ID dari data yang ingin kita ubah isinya dan isi dari data tersebut agar kita dapat menampilkan data yang telah ada sebelumnya, sebelum kita update.

	//todo sekarang, kita buatkan Method UpdateUser-nya di bawah
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Exec("INSERT INTO users(name,email,password,created_at,updated_at) VALUES (?,?,?,?,?)", user.Name, user.Email, user.Password, time.Now(), time.Now()).Error

	return user, err
}

//todo Write UpdateUser method here ...
//todo Write this code
func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
	err := r.db.Raw("UPDATE users SET name=?, email=?, password=? WHERE id=?", user.Name, user.Email, user.Password, ID).Scan(&user).Error

	return user, err
}
