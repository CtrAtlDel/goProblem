package testStore

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"

	_ "github.com/lib/pq"
)

type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepository
}
