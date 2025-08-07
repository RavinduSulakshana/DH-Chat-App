package store

import "chat-app-backend/models"

type Store interface {
	AddUser(user *models.User)
	RemoveUser(userID string)
	GetUser(userID string) (*models.User, bool)
	GetAllUsers() []*models.User
	GetUserByUsername(username string) (*models.User, bool)
}
