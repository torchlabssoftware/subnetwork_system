-- name: CreateUser :one
INSERT INTO "user"(email,username,password,data_limit)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: InsertUserPool :many
INSERT INTO user_pools(user_id,pool_id)
SELECT $1, UNNEST($2::uuid[])
RETURNING *;

-- name: InsertUserIpwhitelist :many
INSERT INTO user_ip_whitelist(user_id,ip_cidr)
SELECT $1, UNNEST($2::text[])
RETURNING *;

-- name: GetUserbyId :one
SELECT * FROM "user" as u
WHERE u.id = $1;