-- name: CreateFeed :one
 Insert into feed (id,created_at,updated_at,name,url,user_id)
 Values ($1,$2,$3,$4,$5,$6) 
 RETURNING *;


-- name: GetFeeds :many
Select * from feed;