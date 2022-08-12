package repositories

import (
	"dumbmerch/models"
	// Import time here ...
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	// Declare CreateUser interface here ...
	CreateUser(user models.User) (models.User, error)
	//ctt isi dari parameternya adalah models.user. Karena, isi dari models.user adalah format yg akan kita kirimkan ke dlm database-nya. Sehingga, setiap apapun yg ingin kita masukkan ke dlm database-nya. Maka, harus sesuai dengan sruck-nya tersebut.
	//ctt Hanya saja, utk ID nya tdk akan kita gunakan karena ID akan diisi secar otomatis dengan auto increment pada mysql.
	//ctt Adapun nilai yg akan dikembalikan dari CreateUser adalah models.user juga. Karena, kita ingin tahu data apa sih yg tadi kita masukkan.

	//todo Setelah ini, barulah kita buatkan method-nya di bawah
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

// Create CreateUser method here ...
// Write this code
func (r *repository) CreateUser(user models.User) (models.User, error) {

	err := r.db.Exec("INSERT INTO users(name,email,password,created_at,updated_at) VALUES (?,?,?,?,?)", user.Name, user.Email, user.Password, time.Now(), time.Now()).Error

	//ctt Ketika kita menggunakan RAW Query untuk melakukan insert, maka kita tidak lagi menggunakan r.db.Raw. Akan, tetapi, Raw-nya kita ganti menjadi Exec alias Execute menjadi r.db.Exec.

	//ctt user.Name, user.Email, dst. itu kita dapatkan dari si parameter function ini
	//ctt sedangkan untuk created_at dan updated_at-nya kita gunakan time.Now()

	return user, err

	//ctt setelah itu semua, barulah kita return-na user dan err nya
}

//todo selanjutnya, kita buatkan handler-nya
