package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/johnnybgoode/breadzilla/internal/data"
	"github.com/johnnybgoode/breadzilla/pkg/server"
	"github.com/labstack/echo/v4"
)

func getAllRecipes(c echo.Context, db *sql.DB) error {
	recipes, err := data.SelectAllRecipes(db)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, recipes)
}

func getRecipeBySlug(c echo.Context, db *sql.DB) error {
	slug := c.Param("slug")
	recipe, err := data.SelectRecipeBySlug(db, slug)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, recipe)
}

func createRecipe(c echo.Context, db *sql.DB) error {
	recipe := new(data.Recipe)
	if err := c.Bind(recipe); err != nil {
		c.Logger().Error(err)
		return err
	}
	return data.InsertRecipe(db, recipe)
}

func updateRecipe(c echo.Context, db *sql.DB) error {
	recipe := new(data.Recipe)
	if err := c.Bind(recipe); err != nil {
		c.Logger().Error(err)
		return err
	}
	return data.UpdateRecipe(db, recipe)
}

func deleteRecipe(c echo.Context, db *sql.DB) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return data.DeleteRecipe(db, idInt)
}

var Routes = server.RouteMap{
	"GET::/recipes":        getAllRecipes,
	"GET::/recipes/:slug":  getRecipeBySlug,
	"POST::/recipes":       createRecipe,
	"PUT::/recipes/:slug":  updateRecipe,
	"DELETE::/recipes/:id": deleteRecipe,
}
