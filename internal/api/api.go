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
		// return err
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

func AddRoutes(s *server.Server) {
	s.Echo().GET("/recipes", s.WithDB(getAllRecipes))
	s.Echo().GET("/recipes/:slug", s.WithDB(getRecipeBySlug))
	s.Echo().POST("/recipes", s.WithDB(createRecipe))
	s.Echo().PUT("/recipes/:slug", s.WithDB(updateRecipe))
	s.Echo().DELETE("/recipes/:id", s.WithDB(deleteRecipe))
}
