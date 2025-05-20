package rest

import (
	"net/http"

	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/api/dto"
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/application/storage"
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/domain/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type RecipeHandler struct {
	service interfaces.RecipeService
}

func NewRecipesHandler(s interfaces.RecipeService) {
	handler := &RecipeHandler{
		service: s,
	}

	// Create Gin router
	router := gin.Default()

	// Register Routes
	router.GET("/", handler.HomePage)
	router.POST("/recipes", handler.CreateRecipe)
	router.GET("/recipes", handler.ListRecipes)
	router.GET("/recipes/:id", handler.GetRecipe)
	router.PUT("/recipes/:id", handler.UpdateRecipe)
	router.DELETE("/recipes/:id", handler.DeleteRecipe)

	// Start the server
	router.Run()
}

// Define handler function signatures
func (h RecipeHandler) HomePage(c *gin.Context) {
	c.String(http.StatusOK, "This is recipe's home page")
}

func (h RecipeHandler) CreateRecipe(c *gin.Context) {
	// Get request body and convert it to dto.Recipe
	var recipe dto.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := slug.Make(recipe.Name)

	h.service.CreateRecipe(id, recipe.MapToModel())
}

func (h RecipeHandler) ListRecipes(c *gin.Context) {
	// Call the store to get the list of recipes\
	r, err := h.service.ListRecipes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	// Return the list, JSON encoding is implicit
	c.JSON(200, r)
}

func (h RecipeHandler) GetRecipe(c *gin.Context) {
	// Retrieve the URL parameter
	id := c.Param("id")

	// Get the recipe by ID from the store
	recipe, err := h.service.GetRecipe(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Return the recipe, JSON encoding is implicit
	c.JSON(200, recipe)
}

func (h RecipeHandler) UpdateRecipe(c *gin.Context) {
	// Get request body and convert it to recipes.Recipe
	var recipe dto.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	// Retrieve URL parameter
	id := c.Param("id")

	// Call the store to Update the recipe
	err := h.service.UpdateRecipe(id, recipe.MapToModel())
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

func (h RecipeHandler) DeleteRecipe(c *gin.Context) {
	// Retrieve URL parameter
	id := c.Param("id")

	// Call the store to delete the recipe
	err := h.service.DeleteRecipe(id)
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
