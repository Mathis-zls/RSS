-- name: CreateFeed :one
 Insert into feed (id,created_at,updated_at,name,url,user_id)
 Values ($1,$2,$3,$4,$5,$6) 
 RETURNING *;


-- name: GetFeeds :many
Select * from feed;



-- name: GetNextFeedsToFetch :many
Select * from feed 
Order by last_fetchedAt Asc nulls first
limit $1; 

-- name: MarkFeedFetched :one
Update feed 
set last_fetchedAt = NOW(),
updated_at = NOW()
WHERE id=$1
RETURNING *;