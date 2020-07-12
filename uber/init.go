package uber

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open("postgres", "host=172.16.1.1 port=5432 user=postgres dbname=uber password=password123 sslmode=disable")
	if err != nil {
		fmt.Println(err)

		panic("failed to connect database")
	}

	fmt.Println("successfully connected.")
	// var booking service.Booking

	db.AutoMigrate(&Booking{}, &Cabs{})

}
