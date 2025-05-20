package dto

import "github.com/abelalem/go-rest-guid-clean-architecture/pkg/domain/models"

type Ingredient struct {
	Name string `json:"name"`
}

func NewIngredient(i models.Ingredient) *Ingredient {
	ingredient := &Ingredient{
		Name: i.Name,
	}

	return ingredient
}

func (i Ingredient) MapToModel() *models.Ingredient {
	ingredient := &models.Ingredient{
		Name: i.Name,
	}

	return ingredient
}
