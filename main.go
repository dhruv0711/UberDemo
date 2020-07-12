package main

import (
	service "github.com/dhruv0711/UberDemo/uber"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/getBookingHistory/:phoneNumber", service.GetBookingHistory)
	r.POST("/createBooking", service.CreateBooking)
	r.POST("/getNearByCabs", service.GetNearByCabs)
	r.POST("/addACab", service.AddACab)
	r.Run(":8888")
}
