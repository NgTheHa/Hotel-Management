package services

import (
	"hotel-management/models"
	"hotel-management/repositories"
)

type RoomService interface {
	CreateRoom(room *models.Room) error
	GetAllRooms() ([]models.Room, error)
}

type roomService struct {
	RoomRepository repositories.RoomRepository
}

func NewRoomService(roomRepo repositories.RoomRepository) RoomService {
	return &roomService{
		RoomRepository: roomRepo,
	}
}

func (rs *roomService) CreateRoom(room *models.Room) error {
	return rs.RoomRepository.CreateRoom(room)
}

func (rs *roomService) GetAllRooms() ([]models.Room, error) {
	return rs.RoomRepository.GetAllRooms()
}
