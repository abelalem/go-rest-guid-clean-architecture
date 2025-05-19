package rest

import (
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/domain/interfaces"
	"github.com/gin-gonic/gin"
)

type RecipeHandler struct {
	service interfaces.RecipeService
}

func NewRecipesHandler(s interfaces.RecipeService) *RecipeHandler {
	handler := &RecipeHandler{
		service: s,
	}

	return handler
}

// Define handler function signatures
func (h RecipeHandler) CreateRecipe(c *gin.Context) {}

func (h RecipeHandler) ListRecipes(c *gin.Context) {}

func (h RecipeHandler) GetRecipe(c *gin.Context) {}

func (h RecipeHandler) UpdateRecipe(c *gin.Context) {}

func (h RecipeHandler) DeleteRecipe(c *gin.Context) {}
