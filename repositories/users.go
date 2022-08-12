package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	// Using "Save" method here ...

	//ctt Terdapat sedikit perbedaan pada Method Update ini di mana pada Raw Query kita membutuhkan dua parameter untuk menjalankan Method UpdateUser ini. Yaitu, terkait models.User-nya alias data yg ingin kita updata dan yang ke dua adalah ID sebagai penunjuk data mana yang akan diupdate.

	//ctt Jika kita menggunakan Method dari GORM-nya, maka kita hanya membutuhkan parameter models.User-nya saja.

	//ctt Kendati demikian, kita tetap harus mengambil ID dari Parameter pada Handler-nya dan membuat pengkondisian utk mengisikan data yang akan di-update-nya. (Silakan lihat pada Handler)
	//ctt Perbedaannya hanya terdapat pada pengiriman parameternya saja. Yaitu, kita sudah tidak perlu lagi mengirimkan parameter ID-nya ke Repository ini.

	err := r.db.Save(&user).Error

	//todo Adapun Method yang digunakan untuk meng-update adalah .Save()

	return user, err
}
