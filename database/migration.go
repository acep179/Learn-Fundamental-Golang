package database

// Import "dumbmerch/models", "dumbmerch/pkg/mysql", "fmt" here ...
import (
	"dumbmerch/models"
	"dumbmerch/pkg/mysql"
	"fmt"
)

//todo Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{})
	//ctt mysql.DB adalah connection untuk ke db nya dan AutoMigrate adalah sebuah method yg telah disediakan oleh gorm yg akan memanggil model user dari pkg models dan nama struck nya adalah user
	//ctt sebagaimana terlihat bahwa utk connection-nya itu kita simpan atau kita buat di dlm folder pkg/mysql dengan nama file-nya adalah mysql.go.

	//. Adapun folder pkg bisa kita isikan terkait pihak ke tiga dlm app kita seperti middleware, jwt, dll. termasuk konfigurasi koneksi mysql.

	//ctt return dari AutoMigrate, yaitu error (err)
	if err != nil {
		//ctt Jika error-nya tidak nil (jika ada isinya, jika ada error). Maka, cetak/tampilkan error-nya dan panic=nya adalah "Migration Failed"
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
