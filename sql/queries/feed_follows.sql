-- name: CreateFeedFollow :one
Insert into feed_follows (id,id_feed,id_user,created_at,updated_at)
Values($1,$2,$3,$4,$5)
RETURNING *;