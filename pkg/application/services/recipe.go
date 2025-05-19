package services

import (
	"net/http"

	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/application/storage"
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type RecipesService struct {
	store storage.RecipeStore
}

func NewRecipeHandler(s storage.RecipeStore) *RecipesService {
	service := &RecipesService{
		store: s,
	}

	return service
}

// Define service function signatures
func (h RecipesService) CreateRecipe(c *gin.Context) {
	// Get request body and convert it to models.Recipe
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Create a URL-friendly name
	id := slug.Make(recipe.Name)

	// Add to the store
	h.store.Add(id, recipe)
}

func (h RecipesService) ListRecipes(c *gin.Context) {
	// Call the store to get the list of recipes
	r, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	// Return the list, JSON encoding is implicit
	c.JSON(200, r)
}

func (h RecipesService) GetRecipe(c *gin.Context) {
	// Retrieve the URL parameter
	id := c.Param("id")

	// Get the recipe by ID from the store
	recipe, err := h.store.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Return the recipe, JSON encoding is implicit
	c.JSON(200, recipe)
}

func (h RecipesService) UpdateRecipe(c *gin.Context) {
	// Get request body and convert it to recipes.Recipe
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	// Retrieve URL parameter
	id := c.Param("id")

	// Call the store to update the recipe
	err := h.store.Update(id, recipe)
	if err != nil {
		if err == storage.ErrorNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h RecipesService) DeleteRecipe(c *gin.Context) {
	// Retrieve URL parameter
	id := c.Param("id")

	// Call the store to delete the recipe
	err := h.store.Remove(id)
	if err != nil {
		if err == storage.ErrorNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
