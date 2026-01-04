package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/johnnybgoode/breadzilla/internal/data"
	"github.com/johnnybgoode/breadzilla/pkg/server"
	"github.com/labstack/echo/v4"
)

func getAllRecipes(c echo.Context, db *sql.DB) error {
	recipes := new(data.Recipes)
	if err := recipes.SelectAll(db); err != nil {
		c.Logger().Error(err)
		return err
	}
	c.Logger().Printf("Recipes %v", recipes)
	return c.JSON(http.StatusOK, recipes)
}

func getRecipeBySlug(c echo.Context, db *sql.DB) error {
	slug := c.Param("slug")
	recipe := new(data.Recipe)

	if err := recipe.SelectBySlug(db, slug); err != nil {
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
	if err := recipe.Insert(db, recipe); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, recipe)
}

func updateRecipe(c echo.Context, db *sql.DB) error {
	recipe := new(data.Recipe)
	if err := c.Bind(recipe); err != nil {
		c.Logger().Error(err)
		return err
	}
	if err := recipe.Update(db); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, recipe)
}

func patchRecipe(c echo.Context, db *sql.DB) error {
	slug := c.Param("slug")
	recipe := new(data.Recipe)

	if err := recipe.SelectBySlug(db, slug); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	// partialRecipe := make(map[string]interface{})
	if err := json.NewDecoder(c.Request().Body).Decode(&recipe); err != nil {
		c.Logger().Error(err)
		return err
	}
	// fieldsUpdated, err := recipe.Patch(db, partialRecipe);
	// if err != nil {
	// 	c.Logger().Error(err)
	// 	return err
	// }

	return c.JSONPretty(http.StatusOK, recipe, " ")
}

func deleteRecipe(c echo.Context, db *sql.DB) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	if err := (new(data.Recipe)).Delete(db, idInt); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

var Routes = server.RouteMap{
	"GET::/recipes":         getAllRecipes,
	"POST::/recipes":        createRecipe,
	"GET::/recipes/:slug":   getRecipeBySlug,
	"PATCH::/recipes/:slug": patchRecipe,
	"PUT::/recipes/:slug":   updateRecipe,
	"DELETE::/recipes/:id":  deleteRecipe,
}
