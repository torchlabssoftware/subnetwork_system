-- name: GetPoolsbyTags :many
SELECT pool.id FROM pool
WHERE pool.tag = ANY($1::text[]);