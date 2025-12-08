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

-- name: GetCountries :many
SELECT * FROM country;

-- name: AddCountry :one
INSERT INTO country(name,code,region_id)
VALUES($1,$2,$3)
RETURNING *;

-- name: DeleteCountry :exec
DELETE FROM country as c
where c.name = $1
RETURNING *;