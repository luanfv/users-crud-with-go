package user

import "github.com/google/uuid"

type Memory struct {
	data map[string]memoryUser
}

func NewMemory() *Memory {
    return &Memory{
        data: make(map[string]memoryUser, 0),
    }
}

type MemoryInsertInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

func (m *Memory) Insert(input MemoryInsertInput) (string, error) {
	id := uuid.NewString()
	user := memoryUser{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Biography: input.Biography,
	}
	m.data[id] = user
	return id, nil
}

func (m *Memory) GetAll() ([]User, error) {
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