package uber

import "fmt"

func listBookingDB(phoneNumber string) ([]Booking, error) {
	fmt.Println(phoneNumber)
	var list []Booking
	db.Where("customer_phone_number = ?", phoneNumber).Find(&list)
	return list, nil
}

func CreateBookingDB(book Booking) (Booking, error) {
	book.IsActive = true
	db.Create(&book)
	var newBooking Booking
	db.Last(&newBooking)
	return newBooking, nil
}

func AddACabDB(cab Cabs) ([]Cabs, error) {
	db.Create(&cab)

	var cabs []Cabs
	db.Find(&cabs)
	return cabs, nil
}

func FindAllCabsDB() ([]Cabs, error) {
	var cabs []Cabs
	db.Find(&cabs)
	return cabs, nil
}
