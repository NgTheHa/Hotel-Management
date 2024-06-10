package routes

import (
	"github.com/gin-gonic/gin"
	"hotel-management/controllers"
	"hotel-management/repositories"
	"hotel-management/services"
)

func SetupRoutes(router *gin.Engine) {
	// Initialize Repositories
	//bookingRepo := repositories.NewBookingRepository()
	//employeeRepo := repositories.NewEmployeeRepository()
	roomRepo := repositories.NewRoomRepository()

	// Initialize Services
	//bookingService := services.NewBookingService(bookingRepo)
	//employeeService := services.NewEmployeeService(employeeRepo)
	roomService := services.NewRoomService(roomRepo)

	// Initialize Controllers
	//bookingController := controllers.NewBookingController(bookingService)
	//employeeController := controllers.NewEmployeeController(employeeService)
	roomController := controllers.NewRoomController(roomService)

	v1 := router.Group("/api/v1")
	{
		rooms := v1.Group("/rooms")
		{
			rooms.GET("/", roomController.GetRooms)
			rooms.POST("/", roomController.AddRoom)
			// Add other room routes...
		}

		//bookings := v1.Group("/bookings")
		//{
		//	bookings.GET("/", bookingController.GetAllBookings)
		//	bookings.POST("/", bookingController.CreateBooking)
		//	// Add other booking routes...
		//}

		//housekeepings := v1.Group("/housekeepings")
		//{
		//	housekeepings.GET("/", controllers.GetHousekeepings)
		//	housekeepings.POST("/", controllers.AddHousekeeping)
		//	// Add other housekeeping routes...
		//}
		//
		//services := v1.Group("/services")
		//{
		//	services.GET("/", controllers.GetServices)
		//	services.POST("/", controllers.AddService)
		//	// Add other services routes...
		//}
		//
		//employees := v1.Group("/employees")
		//{
		//	employees.GET("/", employeeController.GetAllEmployees)
		//	employees.POST("/", employeeController.GetAllEmployees)
		//	// Add other employee routes...
		//}
		//
		//payments := v1.Group("/payments")
		//{
		//	payments.GET("/", controllers.GetPayments)
		//	payments.POST("/", controllers.AddPayment)
		//	// Add other payment routes...
		//}
	}
}
