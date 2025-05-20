package services

import (
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/application/storage"
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/domain/models"
)

type RecipesService struct {
	store storage.RecipeStore
}

func NewRecipeService(s storage.RecipeStore) *RecipesService {
	service := &RecipesService{
		store: s,
	}

	return service
}

// Define service function signatures
func (h RecipesService) CreateRecipe(id string, recipe models.Recipe) error {
	return h.store.Add(id, recipe)
}

func (h RecipesService) ListRecipes() (map[string]models.Recipe, error) {
	return h.store.List()
}

func (h RecipesService) GetRecipe(id string) (models.Recipe, error) {
	return h.store.Get(id)
}

func (h RecipesService) UpdateRecipe(id string, recipe models.Recipe) error {
	return h.store.Update(id, recipe)
}

func (h RecipesService) DeleteRecipe(id string) error {
	return h.store.Remove(id)
}
