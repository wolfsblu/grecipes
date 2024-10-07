// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO recipes (name)
VALUES (?)
RETURNING id, name
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, name)
	var i Recipe
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteRecipe = `-- name: DeleteRecipe :exec
DELETE
FROM recipes
WHERE id = ?
`

func (q *Queries) DeleteRecipe(ctx context.Context, id interface{}) error {
	_, err := q.db.ExecContext(ctx, deleteRecipe, id)
	return err
}

const getRecipe = `-- name: GetRecipe :one
SELECT id, name
FROM recipes
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetRecipe(ctx context.Context, id interface{}) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, getRecipe, id)
	var i Recipe
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listRecipes = `-- name: ListRecipes :many
SELECT id, name
FROM recipes
ORDER BY name
`

func (q *Queries) ListRecipes(ctx context.Context) ([]Recipe, error) {
	rows, err := q.db.QueryContext(ctx, listRecipes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRecipe = `-- name: UpdateRecipe :exec
UPDATE recipes
set name = ?
WHERE id = ?
RETURNING id, name
`

type UpdateRecipeParams struct {
	Name string
	ID   interface{}
}

func (q *Queries) UpdateRecipe(ctx context.Context, arg UpdateRecipeParams) error {
	_, err := q.db.ExecContext(ctx, updateRecipe, arg.Name, arg.ID)
	return err
}