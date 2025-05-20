package dto

import (
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/domain/models"
)

type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

func NewRecipe(r models.Recipe) *Recipe {
	ingredients := make([]Ingredient, len(r.Ingredients))

	for i, ri := range r.Ingredients {
		ingredients[i] = Ingredient{
			Name: ri.Name,
		}
	}

	recipe := &Recipe{
		Name:        r.Name,
		Ingredients: ingredients,
	}

	return recipe
}

func (r Recipe) MapToModel() models.Recipe {
	ingredients := make([]models.Ingredient, len(r.Ingredients))

	for i, ri := range r.Ingredients {
		ingredients[i] = models.Ingredient{
			Name: ri.Name,
		}
	}

	recipe := models.Recipe{
		Name:        r.Name,
		Ingredients: ingredients,
	}

	return recipe
}
