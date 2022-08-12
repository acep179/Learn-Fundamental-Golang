package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	// Using "Find" method here ...

	//ctt Karena kita ingin mendapatkan seluruh data, maka kita akan menggunakan Method dari "Retrieving all objects" yang terdapat pada dokumentasi GORM

	//ctt Dalam hal ini, kita akan menggunakan Method Find, lalu kita isikan model apa yang ingin kita cari, dalam hal ini adalah models users, sehingga kita isikan r.db.Find.(&users)

	err := r.db.Find(&users).Error

	//ctt Karena kita akan memanggil connection-nya menggunakan "db", maka kita tidak dapat menuliskan secara langsung db.Find(), kita harus memanggilnya dari repository yang diinisiasikan sebagai r (di sebelah tulisan func). Maka dari itu, kita ketikkan r.db.Find()

	//ctt Lalu, karena hasil dari query-nya langsung sekalian kita simpan ke dalam variabel users yaitu dengan menggunakan pointing (&users). Maka, kita cukup mengetikkan variabel untuk menampung error-nya. Maka dari itu, kita ketikkan menjadi err := r.db.Find(&users).Error

	return users, err
}

//! Berbeda dengan Method FindUsers, untuk Method GetUser kita gunakan Method First() karena Method Find() hanya digunakan untuk menampilkan data yang banyak, sedangkan dalam Method GetUser kita hanya akan menampilkan satu data berdasarkan ID yang dikirim

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	// Using "First" method here ...

	//ctt Sebenarnya, ada beberapa Method yang dapat kita gunakan. Ada Method First, Take, dan Last.
	//ctt Method First akan mengembalikan baris pertama yang kita urutkan berdasarkan primary_key
	//ctt Take mengembalikan satu data yang tdk kita tentukan urutannya berdasarkan apa
	//ctt Method Last akan mengembalikan baris terakhir yang kita urutkan berdasarkan primary_key

	//ctt Jika kita ingin mencari data berdasarkan ID-nya, maka kita dapat menggunakan Method First() dan diikuti dengan si ID-nya
	//ctt Contohnya adalah First(user, 4)

	err := r.db.First(&user, ID).Error

	//ctt Kita juga dapat melihat query yg dijalankan oleh sistem dengan menambahkan Debug()
	//ctt Contoh: r.db.Debug().First(&users,ID).Error
	//ctt Perlu diperhatikan bahwa Debug harus dihapus ketika telah pada tahap production. Sehignga, sebaiknya Debug() harus segera dihapus ketika tidak lagi digunakan.

	return user, err
}
