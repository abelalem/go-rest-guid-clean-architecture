package main

import (
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/api/rest"
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/application/services"
	"github.com/abelalem/go-rest-guid-clean-architecture/pkg/infrastructure/memory"
)

func main() {

	// Instantiate recipe Handler and provide a data store implementation
	store := memory.NewMemoryStore()

	service := services.NewRecipeService(store)

	rest.NewRecipesHandler(service)
}
