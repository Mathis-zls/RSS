-- name: CreatePosts :one
 Insert into posts (id,created_at,updated_at,title,description,published_at,url,feed_id)
 Values ($1,$2,$3,$4,$5,$6,$7,$8)
 RETURNING *;

-- name: GetPostForUser :many
Select posts.* from posts 
Join feed_follows On posts.feed_id = feed_follows.id_feed
WHERE feed_follows.id_user=$1
Order by posts.published_at Desc 
Limit $2;