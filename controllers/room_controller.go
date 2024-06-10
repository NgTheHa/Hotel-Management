package controllers

import (
	"hotel-management/models"
	"hotel-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	RoomService services.RoomService
}

func NewRoomController(roomService services.RoomService) *RoomController {
	return &RoomController{
		RoomService: roomService,
	}
}

func (rc *RoomController) GetRooms(c *gin.Context) {
	rooms, err := rc.RoomService.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rooms)
}

func (rc *RoomController) AddRoom(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := rc.RoomService.CreateRoom(&room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, room)
}
