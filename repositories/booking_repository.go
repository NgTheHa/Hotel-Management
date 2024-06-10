package repositories

//
//import (
//	"hotel-management/models"
//	"hotel-management/utils"
//)
//
//type BookingRepository interface {
//	CreateBooking(booking *models.Booking) error
//	GetAllBookings() ([]models.Booking, error)
//}
//
//type bookingRepository struct{}
//
//func NewBookingRepository() BookingRepository {
//	return &bookingRepository{}
//}
//
//func (br *bookingRepository) CreateBooking(booking *models.Booking) error {
//	return utils.DB.Create(booking).Error
//}
//
//func (br *bookingRepository) GetAllBookings() ([]models.Booking, error) {
//	var bookings []models.Booking
//	err := utils.DB.Find(&bookings).Error
//	return bookings, err
//}
