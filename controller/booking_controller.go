package controllers

import (
	"hotel-management/models"
	"hotel-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	BookingService services.BookingService
}

func NewBookingController(bookingService services.BookingService) *BookingController {
	return &BookingController{
		BookingService: bookingService,
	}
}

func (bc *BookingController) CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.BookingService.CreateBooking(&booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, booking)
}

func (bc *BookingController) GetAllBookings(c *gin.Context) {
	bookings, err := bc.BookingService.GetAllBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}
