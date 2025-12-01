package data

import (
	"database/sql"
	"fmt"

	"github.com/johnnybgoode/breadzilla/internal/types"
)

type (
	Portions struct {
		Unit  string `json:"unit"`
		Units string `json:"units"`
		Value int    `json:"value"`
	}

	Ingredient struct {
		Name  string `json:"name"`
		Unit  string `json:"unit"`
		Value int    `json:"value"`
	}
	Ingredients []Ingredient

	Step struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Time        int    `json:"time"`
	}
	Steps []Step

	Recipe struct {
		ID          int64
		Title       string
		Slug        string
		Credit      string
		Image       string
		Portions    types.JSON[Portions]
		Ingredients types.JSON[Ingredients]
		Steps       types.JSON[Steps]
	}
)

const (
	selectAllQuery    = `SELECT * FROM recipes ORDER BY title ASC`
	selectBySlugQuery = `SELECT * FROM recipes WHERE slug = ? ORDER BY title ASC`
	insertQuery       = `INSERT INTO recipes(title, slug, credit, image, portions, ingredients, steps) VALUES (?, ?, ?, ?, ?, ?, ?)`
)

func SelectAllRecipes(db *sql.DB) ([]Recipe, error) {
	rows, err := db.Query(selectAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []Recipe

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("SelectAllQueryRecipes %e", err)
		}

		var r Recipe
		err := rows.Scan(&r.ID, &r.Title, &r.Slug, &r.Credit, &r.Image, &r.Portions, &r.Ingredients, &r.Steps)
		if err != nil {
			return nil, fmt.Errorf("SelectAllQueryRecipes %e", err)
		}

		recipes = append(recipes, r)
	}

	return recipes, nil
}

func SelectRecipeBySlug(db *sql.DB, slug string) (Recipe, error) {
	row := db.QueryRow(selectBySlugQuery, slug)
	var r Recipe

	if err := row.Scan(&r.ID, &r.Title, &r.Slug, &r.Credit, &r.Image, &r.Portions, &r.Ingredients, &r.Steps); err != nil {
		return r, fmt.Errorf("selectRecipeBySlug %e", err)
	}

	return r, nil
}

func InsertRecipe(db *sql.DB, recipe *Recipe) error {
	if _, err := db.Exec(insertQuery, recipe.Title, recipe.Slug, recipe.Credit, recipe.Image, recipe.Portions, recipe.Ingredients, recipe.Steps); err != nil {
		return err
	}
	return nil
}

func MakeRecipe() Recipe {
	return Recipe{
		Title:  "Rolls",
		Slug:   "rolls",
		Credit: "",
		Image:  "",
		Portions: types.JSON[Portions]{
			Val: Portions{
				Unit:  "roll",
				Units: "rolls",
				Value: 12,
			},
		},
		Ingredients: types.JSON[Ingredients]{
			Val: Ingredients{{
				Name:  "flour",
				Unit:  "g",
				Value: 240,
			}},
		},
		Steps: types.JSON[Steps]{
			Val: Steps{{}},
		},
	}
}