package api

import (
	"database/sql"
	"net/http"

	"github.com/johnnybgoode/breadzilla/internal/data"
	"github.com/johnnybgoode/breadzilla/pkg/server"
	"github.com/labstack/echo/v4"
)

func getAllRecipes(c echo.Context, db *sql.DB) error {
	recipes, err := data.SelectAllRecipes(db)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, recipes)
}

func getRecipeBySlug(c echo.Context, db *sql.DB) error {
	slug := c.Param("slug")
	recipe, err := data.SelectRecipeBySlug(db, slug)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, recipe)
}

func AddRoutes(s *server.Server) {
	s.Echo().GET("/recipes", s.WithDB(getAllRecipes))
	s.Echo().GET("/recipes/:slug", s.WithDB(getRecipeBySlug))
}
