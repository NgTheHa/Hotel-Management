package repositories

import (
	"hotel-management/models"
	"hotel-management/utils"
)

type RoomRepository interface {
	CreateRoom(room *models.Room) error
	GetAllRooms() ([]models.Room, error)
}

type roomRepository struct{}

func NewRoomRepository() RoomRepository {
	return &roomRepository{}
}

func (rr *roomRepository) CreateRoom(room *models.Room) error {
	return utils.DB.Create(room).Error
}

func (rr *roomRepository) GetAllRooms() ([]models.Room, error) {
	var rooms []models.Room
	err := utils.DB.Find(&rooms).Error
	return rooms, err
}
