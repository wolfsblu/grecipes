-- name: GetRecipe :one
SELECT *
FROM recipes
WHERE id = ?
LIMIT 1;

-- name: ListRecipes :many
SELECT *
FROM recipes
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO recipes (name)
VALUES (?)
RETURNING *;

-- name: UpdateRecipe :exec
UPDATE recipes
set name = ?
WHERE id = ?
RETURNING *;

-- name: DeleteRecipe :exec
DELETE
FROM recipes
WHERE id = ?;