package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
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

func patchRecipe(c echo.Context, db *sql.DB) error {
	slug := c.Param("slug")
	recipe, err := data.SelectRecipeBySlug(db, slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	recipePayload := make(map[string]interface{})
	if err := json.NewDecoder(c.Request().Body).Decode(&recipePayload); err != nil {
		c.Logger().Error(err)
		return err
	}

	fieldsUpdated := make([]string, 0)
	u := reflect.ValueOf(recipe).Elem()

	for fieldName, rawVal := range recipePayload {
		updateSuccess := false
		field := u.FieldByName(fieldName)
		if !field.CanSet() {
			continue
		}
		v := reflect.ValueOf(rawVal)

		// TODO switch field.Kind(); handle custom JSON field types
		if field.Kind() == v.Kind() {
			field.Set(v)
			updateSuccess = true
		}

		if updateSuccess {
			fieldsUpdated = append(fieldsUpdated, fieldName)
		}
	}

	message := make(map[string]string)
	message["payload"] = fmt.Sprintf("%v", recipePayload)
	message["recipe"] = fmt.Sprintf("%v", recipe)
	message["fieldsUpdated"] = fmt.Sprintf("%v", fieldsUpdated)
	return c.JSONPretty(http.StatusOK, message, " ")
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
	"GET::/recipes":         getAllRecipes,
	"POST::/recipes":        createRecipe,
	"GET::/recipes/:slug":   getRecipeBySlug,
	"PATCH::/recipes/:slug": patchRecipe,
	"PUT::/recipes/:slug":   updateRecipe,
	"DELETE::/recipes/:id":  deleteRecipe,
}
