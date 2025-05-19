package models

type Ingredient struct {
	Id   int
	Name string
}

type Recipe struct {
	Id         int
	Name       string
	Ingredient []Ingredient
}
