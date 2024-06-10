package routes

//
//import (
//	"github.com/gin-gonic/gin"
//	"hotel-management/controllers"
//)
//
//func SetupRoutes(router *gin.Engine) {
//	v1 := router.Group("/api/v1")
//	{
//		rooms := v1.Group("/rooms")
//		{
//			rooms.GET("/", controllers.GetRooms)
//			rooms.POST("/", controllers.AddRoom)
//			// Add other room routes...
//		}
//
//		bookings := v1.Group("/bookings")
//		{
//			bookings.GET("/", controllers.GetBookings)
//			bookings.POST("/", controllers.MakeBooking)
//			// Add other booking routes...
//		}
//
//		housekeepings := v1.Group("/housekeepings")
//		{
//			housekeepings.GET("/", controllers.GetHousekeepings)
//			housekeepings.POST("/", controllers.AddHousekeeping)
//			// Add other housekeeping routes...
//		}
//
//		services := v1.Group("/services")
//		{
//			services.GET("/", controllers.GetServices)
//			services.POST("/", controllers.AddService)
//			// Add other service routes...
//		}
//
//		employees := v1.Group("/employees")
//		{
//			employees.GET("/", controllers.GetEmployees)
//			employees.POST("/", controllers.AddEmployee)
//			// Add other employee routes...
//		}
//
//		payments := v1.Group("/payments")
//		{
//			payments.GET("/", controllers.GetPayments)
//			payments.POST("/", controllers.AddPayment)
//			// Add other payment routes...
//		}
//	}
//}
