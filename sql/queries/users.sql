-- name: CreateUser :one
INSERT INTO "user"(email,username,password,data_limit)
VALUES ($1,$2,$3,$4)
RETURNING *;