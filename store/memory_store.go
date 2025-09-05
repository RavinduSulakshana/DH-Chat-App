package store

import (
	"chat-app-backend/models"
	"os/user"
	"sync"

	"golang.org/x/tools/go/analysis/passes/defers"
)

type MemoryStore struct {
	users map[string]*models.User
	mutex sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users: make(map[string]*models.User),
	}
}

func (s *MemoryStore) AddUser(user *models.User) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.users[user.ID] = user
}

func (s *MemoryStore) RemoveUser(userID string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.users, userID)

}

func (s *MemoryStore) GetUser(userID string) (*models.User, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	user, exits := s.users[userID]
	return user, exits
}

func (s *MemoryStore) GetAllUsers() []*models.User {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	users := make([]*models.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

func (s *MemoryStore) GetUserByUsername(username string) (*models.User, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, user := range s.users {
		if user.Username == username {
			return user, true
		}
	}
	return nil, false
}
