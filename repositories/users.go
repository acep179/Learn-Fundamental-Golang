package repositories

//ctt Pada repositories ini, akan kita buatkan terkait proses bagaimana dari connection yg kita punya itu kita tuliskan query-query-nya. Sehingga, kita tidak lagi menuliskan di dlm handler-nya.

// Import "dumbmerch/models", "gorm.io/gorm"
import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

//ctt interface itu seperti kontrak. Sehingga, ketika kita ingin menggunakan repository user, maka kita harus memenuhi kontrak yg ada di interface tsb.
// Declare UserRepository interface here ...
type UserRepository interface {

	//ctt interface ini berisi function-function yg akan kita deklarasikan

	// nama function dan nilai yg dikembalikan
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)

	//.ctt salah satu function-nya adalah FindUsers yg nilai kembalian atau return-nya adalah dalam bentuk modelsUser yg di-slice dan error
	//ctt Maksud dari slice adalah nantinya, function FindUsers akan mengambil seluruh struct dari model user. Kemudian, barulah akan diambil potongan dari model user tsb.
	//ctt Sebaiknya, ketika kita ingin melakukan pemformatan data apa saja yg kita butuhkan, maka kita lakukan di handler
}

//ctt repository struct ini akan berisikan database
// Declare repository struct here ...
type repository struct {
	db *gorm.DB
}

// Create RepositoryUser function here ...
func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

//ctt RepositoryUser ini akan kita panggil pada saat pendeklarasian routes-nya (routes/users.go). Sehingga, ketika kita mengakses endpoint tertentu, maka kita selipkan si connection ke database-nya agar dapat bisa kita gunakan di dlm repository-nya

// Create FindUsers method here ...
func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

	return users, err
}

//ctt function FindUsers di atas kita buat bukan menggunakan function melainkan menggunakan method. Yg artinya kita .... Tujuannya, agar dari si repostitory bisa mengakses si connection db-nya (r.db).
//ctt repository adalah struct yg artinya ia merupakan kerangkanya
//ctt Itulah mengapa ketika kita ingin melakukan FindUsers() kita membutuhkan kerangkanya itu tadi
//ctt Sama seperti penjelasan sebelumnya bahwa ketika kita membuat function, lalu kita jadikan ia sebagai method maka bisa kita tambahkan ... (r *repository) => r yg mempointing repository struct

// Create GetUser method here ...
func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}
