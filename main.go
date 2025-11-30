package main

import (
	"context"
	"fmt"
	"log"

	"github.com/johnnybgoode/breadzilla/internal/data"
	"github.com/johnnybgoode/breadzilla/internal/types"
	"github.com/johnnybgoode/breadzilla/pkg/database"
)

func main() {
	ctx := context.Background()
	var config database.Config
	config.ProcessFromEnv(ctx)
	db, err := database.Connect(&config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	recipe := data.Recipe{
		Title:  "Rolls",
		Slug:   "rolls",
		Credit: "",
		Image:  "",
		Portions: types.JSON[data.Portions]{
			Val: data.Portions{
				Unit:  "roll",
				Units: "rolls",
				Value: 12,
			},
		},
		Ingredients: types.JSON[data.Ingredients]{
			Val: data.Ingredients{{
				Name:  "flour",
				Unit:  "g",
				Value: 240,
			}},
		},
		Steps: types.JSON[data.Steps]{
			Val: data.Steps{{}},
		},
	}
	data.InsertRecipe(db, &recipe)

	recipes, _ := data.SelectAllRecipes(db)
	fmt.Printf("Recipes %+v", recipes)
}
