package uber

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookingHistory(c *gin.Context) {
	phoneNumber := c.Param("phoneNumber")

	listBooking, err := listBookingDB(phoneNumber)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"list": listBooking,
	})
}

func CreateBooking(c *gin.Context) {

	var book Booking
	err := c.BindJSON(&book)
	if err != nil {
		fmt.Println(err)
	}

	CabsList, err := FindAllCabsDB()
	if err != nil {
		return
	}

	var minDistance float64 = math.MaxFloat64
	var cab Cabs
	for _, v := range CabsList {
		fmt.Println("cabgps", v.CabCurrentGps, "customergps", book.CustomerGps)
		dis, err := calculateDistance(v.CabCurrentGps, book.CustomerGps)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("dis", dis)
		if dis < minDistance {
			minDistance = dis
			cab = v
		}
	}

	book.CabDriverName = cab.CabDriverName
	book.CabDriverPhoneNumber = cab.CabDriverPhoneNumber
	book.CabMakeModel = cab.CabMakeModel
	book.CabRegistrationNumber = cab.CabRegistrationNumber

	newBooking, err := CreateBookingDB(book)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Booking": newBooking,
	})
}

func GetNearByCabs(c *gin.Context) {
	var book Booking
	c.BindJSON(&book)

	CabsList, err := FindAllCabsDB()
	if err != nil {
		return
	}

	var minDistance float64 = math.MaxFloat64
	var cab Cabs
	for _, v := range CabsList {
		dis, err := calculateDistance(v.CabCurrentGps, book.CustomerGps)
		if err != nil {
			fmt.Println(err)
		}
		if dis < minDistance {
			minDistance = dis
			cab = v
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"nearestCab": cab,
	})
}

func AddACab(c *gin.Context) {
	var cab Cabs
	err := c.BindJSON(&cab)
	if err != nil {
		fmt.Println(err)
	}

	CabList, err := AddACabDB(cab)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"CabList": CabList,
	})
}
