package services

//import (
//	"hotel-management/models"
//	repositories "hotel-management/repositories"
//)
//
//type BookingService interface {
//	CreateBooking(booking *models.Booking) error
//	GetAllBookings() ([]models.Booking, error)
//}
//
//type bookingService struct {
//	BookingRepository repositories.BookingRepository
//}
//
//func NewBookingService(bookingRepo repositories.BookingRepository) BookingService {
//	return &bookingService{
//		BookingRepository: bookingRepo,
//	}
//}
//
//func (bs *bookingService) CreateBooking(booking *models.Booking) error {
//	return bs.BookingRepository.CreateBooking(booking)
//}
//
//func (bs *bookingService) GetAllBookings() ([]models.Booking, error) {
//	return bs.BookingRepository.GetAllBookings()
//}
