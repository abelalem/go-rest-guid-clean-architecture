package memory

import (
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/application/storage"
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/domain/models"
)

type MemoryStore struct {
	list map[string]models.Recipe
}

func NewMemoryStore() *MemoryStore {
	list := make(map[string]models.Recipe)

	return &MemoryStore{
		list: list,
	}
}

func (m MemoryStore) Add(name string, recipe models.Recipe) error {
	m.list[name] = recipe

	return nil
}

func (m MemoryStore) List() (map[string]models.Recipe, error) {
	return m.list, nil
}

func (m MemoryStore) Get(name string) (models.Recipe, error) {
	if val, ok := m.list[name]; ok {
		return val, nil
	}

	return models.Recipe{}, storage.ErrorNotFound
}

func (m MemoryStore) Update(name string, recipe models.Recipe) error {
	if _, ok := m.list[name]; ok {
		m.list[name] = recipe
		return nil
	}

	return storage.ErrorNotFound
}

func (m MemoryStore) Remove(name string) error {
	delete(m.list, name)

	return nil
}
