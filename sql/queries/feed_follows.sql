-- name: CreateFeedFollow :one
Insert into feed_follows (id,id_feed,id_user,created_at,updated_at)
Values($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetFeedFollows :many
Select * from feed_follows WHERE id_user=$1;

-- name: DeleteFeedFollow :exec
delete from feed_follows WHERE id=$1 and id_user=$2;