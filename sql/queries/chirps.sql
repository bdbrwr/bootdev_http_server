-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES(
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: GetChirps :many
SELECT * FROM chirps
WHERE sqlc.narg(author_id)::uuid IS NULL or user_id = sqlc.narg(author_id)
ORDER BY
  CASE WHEN sqlc.narg(sort)::text = 'desc' THEN created_at END DESC,
  created_at ASC;

-- name: GetChirp :one
SELECT * FROM chirps
WHERE id = $1;

-- name: DeleteChirp :exec
DELETE FROM chirps
WHERE id = $1;
