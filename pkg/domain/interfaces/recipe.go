package interfaces

import "github.com/gin-gonic/gin"

type RecipeService interface {
	CreateRecipe(c *gin.Context)
	ListRecipes(c *gin.Context)
	GetRecipe(c *gin.Context)
	UpdateRecipe(c *gin.Context)
	DeleteRecipe(c *gin.Context)
}
