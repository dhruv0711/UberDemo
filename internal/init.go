package uber

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=uber password=dhruv@11")
	if err != nil {
		fmt.Println(err)

		panic("failed to connect database")
	}

	fmt.Println("successfully connected.")
	// var booking service.Booking

	db.AutoMigrate(&Booking{}, &Cabs{})

}
