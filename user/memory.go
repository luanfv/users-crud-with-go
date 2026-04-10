package user

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type Memory struct {
	mu sync.RWMutex
	data map[string]memoryUser
}

func NewMemory() *Memory {
	return &Memory{
		data: make(map[string]memoryUser),
	}
}

type MemoryUserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

func (m *Memory) Insert(input MemoryUserInput) (User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id := uuid.NewString()
	mu := memoryUser{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Biography: input.Biography,
	}
	m.data[id] = mu
	return User{
		ID:        id,
		FirstName: mu.FirstName,
		LastName:  mu.LastName,
		Biography: mu.Biography,
	}, nil
}

func (m *Memory) FindAll() ([]User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	users := make([]User, 0, len(m.data))
	for id, user := range m.data {
		users = append(users, User{
			ID:        id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Biography: user.Biography,
		})
	}
	return users, nil
}

func (m *Memory) FindById(id string) (User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	u, ok := m.data[id]
	if !ok {
		return User{}, errors.New("User not found")
	}
	return User{
		ID:        id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Biography: u.Biography,
	}, nil
}

func (m *Memory) Update(id string, user User) (User, error) {
	m.mu.Lock()
    defer m.mu.Unlock()

	userData, ok := m.data[id]; if !ok {
		return User{}, errors.New("User not found")
	}
	userData.FirstName = user.FirstName
	userData.LastName = user.LastName
	userData.Biography = user.Biography
	m.data[id] = userData
	return User{
		ID:        id,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Biography: userData.Biography,
	}, nil
}

func (m *Memory) Delete(id string) (User, error) {
	m.mu.Lock()
    defer m.mu.Unlock()

	user, ok := m.data[id]; if !ok {
		return User{}, errors.New("User not found")
	}
	delete(m.data, id)
	return User{
		ID:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Biography: user.Biography,
	}, nil
}