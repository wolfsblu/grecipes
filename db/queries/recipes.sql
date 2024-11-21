-- name: GetRecipe :one
SELECT *
FROM recipes
WHERE id = ?
LIMIT 1;

-- name: ListRecipes :many
SELECT *
FROM recipes
WHERE created_by = ?
ORDER BY name;

-- name: CreateRecipe :one
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