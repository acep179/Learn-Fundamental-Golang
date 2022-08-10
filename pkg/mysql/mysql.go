package mysql

// Import "fmt", "gorm.io/driver/mysql", "gorm.io/gorm" here ...
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Declare DB varible from *gorm.DB here ...
var DB *gorm.DB

//todo yg pertama kali kita lakukan adalah menyiapkan sebuah variabel dengan nama DB yg akan mem-pointing ke dlm gorm.DB yg merupakan bawaan dari gorm.
//todo Variabel DB tersebut, nantinya akan kita isikan dengan DatabaseInit di bawah

// Connection Database
func DatabaseInit() {
	var err error
	//ctt dsn := "{USER}:{PASSWORD}@tcp({HOST}:{PORT})/{DATABASE}?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:@tcp(localhost:3306)/dumbmerch?charset=utf8mb4&parseTime=True&loc=Local"
	//ctt dsn adalah destination yg berisikan destinasi database kita. Dns tsb diisikan dengan user, password, dst. Jika pada akun user kita tidak menggunakan password, maka kita boleh mengosongkan passwordnya.

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}
