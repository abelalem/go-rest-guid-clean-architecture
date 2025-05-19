package storage

import (
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/domain/models"
)

type RecipeStore interface {
	Add(name string, recipe models.Recipe) error
	Get(name string) (models.Recipe, error)
	List() (map[string]models.Recipe, error)
	Update(name string, recipe models.Recipe) error
	Remove(name string) error
}
