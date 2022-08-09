package database

// Import "dumbmerch/models", "dumbmerch/pkg/mysql", "fmt" here ...
import (
	"dumbmerch/models"
	"dumbmerch/pkg/mysql"
	"fmt"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
