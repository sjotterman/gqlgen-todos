-- name: GetRestaurant :one
SELECT * from restaurant
where id = $1 LIMIT 1;

-- name: ListRestaurants :many
SELECT * from restaurant
ORDER BY name;

-- name: CreateRestaurant :one
INSERT INTO restaurant (name, description, phone_number)
VALUES ($1, $2, $3)
returning *;

-- name: DeleteRestaurant :exec
DELETE from restaurant
where id = $1;

-- name: UpdateRestaurant :one
UPDATE restaurant
set name= $2,
description = $3,
phone_number = $4
where id = $1
returning *;
