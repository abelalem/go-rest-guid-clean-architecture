package models

type Ingredient struct {
	Name string
}

type Recipe struct {
	Name        string
	Ingredients []Ingredient
}
