package main

import (
	service "github.com/dhruv0711/UberDemo/internal"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/getBookingHistory/:phoneNumber", service.GetBookingHistory)
	r.POST("/createBooking", service.CreateBooking)
	r.GET("/getNearByCabs/:customerGps", service.GetNearByCabs)
	r.POST("/addACab", service.AddACab)
	r.Run(":8888")
}
