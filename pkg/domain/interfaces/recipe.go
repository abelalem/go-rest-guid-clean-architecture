package interfaces

import (
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/domain/models"
)

type RecipeService interface {
	CreateRecipe(id string, recipe models.Recipe) error
	ListRecipes() (map[string]models.Recipe, error)
	GetRecipe(id string) (models.Recipe, error)
	UpdateRecipe(id string, recipe models.Recipe) error
	DeleteRecipe(id string) error
}
