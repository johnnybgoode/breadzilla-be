package data

import (
	"database/sql"
	"fmt"
	"reflect"

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
		Title       string                  `db:"title"`
		Slug        string                  `db:"slug"`
		Credit      string                  `db:"credit"`
		Image       string                  `db:"image"`
		Portions    types.JSON[Portions]    `db:"portions" json:"portions"`
		Ingredients types.JSON[Ingredients] `db:"ingredients" json:"ingredients"`
		Steps       types.JSON[Steps]       `db:"steps" json:"steps"`
	}

	Recipes []*Recipe
)

const (
	selectAllQuery    = `SELECT * FROM recipes ORDER BY title ASC`
	selectBySlugQuery = `SELECT * FROM recipes WHERE slug = ? ORDER BY title ASC`
	insertQuery       = `INSERT INTO recipes(title, slug, credit, image, portions, ingredients, steps) VALUES (?, ?, ?, ?, ?, ?, ?)`
	updateQuery       = `UPDATE recipes SET title = ?, slug = ?, credit = ?, image = ?, portions = ?, ingredients = ?, steps = ? WHERE id = ?`
	deleteQuery       = `DELETE FROM recipes WHERE id = ?`
	// patchQuery				= `UPDATE recipes SET`
)

func (rs *Recipes) SelectAll(db *sql.DB) error {
	rows, err := db.Query(selectAllQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return fmt.Errorf("SelectAllQueryRecipes %e", err)
		}

		r := new(Recipe)
		err := rows.Scan(&r.ID, &r.Title, &r.Slug, &r.Credit, &r.Image, &r.Portions, &r.Ingredients, &r.Steps)
		if err != nil {
			return fmt.Errorf("SelectAllQueryRecipes %e", err)
		}

		*rs = append(*rs, r)
	}

	return nil
}

func (r *Recipe) SelectBySlug(db *sql.DB, slug string) error {
	row := db.QueryRow(selectBySlugQuery, slug)

	if err := row.Scan(&r.ID, &r.Title, &r.Slug, &r.Credit, &r.Image, &r.Portions, &r.Ingredients, &r.Steps); err != nil {
		return fmt.Errorf("selectRecipeBySlug %e", err)
	}

	return nil
}

func (r *Recipe) Insert(db *sql.DB, recipe *Recipe) error {
	if _, err := db.Exec(insertQuery, recipe.Title, recipe.Slug, recipe.Credit, recipe.Image, recipe.Portions, recipe.Ingredients, recipe.Steps); err != nil {
		return err
	}
	return nil
}

func (r *Recipe) Patch(db *sql.DB, partialRecipe map[string]interface{}) ([]string, error) {
	fieldsUpdated := make([]string, 0)
	u := reflect.ValueOf(r).Elem()
	t := reflect.TypeOf(r).Elem()

	for fieldName, rawVal := range partialRecipe {
		// updateSuccess := false

		sf, ok := t.FieldByName(fieldName)
		field := u.FieldByName(fieldName)
		if !ok || !field.CanSet() {
			continue
		}
		v := reflect.ValueOf(rawVal)

		// TODO switch field.Kind(); handle custom JSON field types
		// if field.Kind() == v.Kind() {
		// field.Set(v)
		// updateSuccess = true
		// }
		if field.Kind() == v.Kind() {
			field.Set(v)
			tag := sf.Tag.Get("db")
			// updateSuccess = true

			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("Updated: %s: %s", fieldName, tag))
		} else {
			nv  := reflect.New(field.Type()).Elem()
			// fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("new val: %v is %v", nv, nv.Interface()))
			field.Set(nv)
			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("Field Interface() %T", field.Interface()))
			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("Field Type() %T", field.Type()))
			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("Field Kind() %T", field.Kind()))
			// fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("Value type %v", v.Kind()))
			// fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("Raw Value type %T", rawVal))

		}

		// else {
			// switch field.Interface().(type) {
			// case Ingredients:
				// nv := new(Ingredients)
				// k := v.Kind()
				// err := json.NewDecoder(rawVal).Decode(&nv)
 			// default:
				// continue
			// }
		// }

		// if updateSuccess {
		// 	tag := sf.Tag.Get("db")
		// 	fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("%s: %s", fieldName, tag))
		// }
	}

	return fieldsUpdated, r.Update(db)

	// query := `UPDATE recipes SET `
	// for _, fieldName := range fieldsUpdated {
	// query += fieldName + " = ?"
	// }

	// message := make(map[string]string)
	// message["payload"] = fmt.Sprintf("%v", partialRecipe)
	// message["recipe"] = fmt.Sprintf("%v", recipe)
	// message["fieldsUpdated"] = fmt.Sprintf("%v", fieldsUpdated)
}

func (r *Recipe) Update(db *sql.DB) error {
	if _, err := db.Exec(updateQuery, r.Title, r.Slug, r.Credit, r.Image, r.Portions, r.Ingredients, r.Steps, r.ID); err != nil {
		return err
	}
	return nil
}

func (r *Recipe) Delete(db *sql.DB, id int) error {
	if _, err := db.Exec(deleteQuery, id); err != nil {
		return err
	}
	return nil
}
