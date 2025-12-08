-- name: GetRegions :many
SELECT * FROM region;

-- name: AddRegion :one
INSERT INTO region(name)
VALUES($1)
RETURNING *;

-- name: DeleteRegion :exec
DELETE FROM region as r
where r.name = $1
RETURNING *;