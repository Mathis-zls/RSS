-- name: CreateUser :one
 Insert into users (id,created_at,updated_at,name,api_key)
 Values ($1,$2,$3,$4, 
    encode(sha256(random()::text::bytea),'hex'))
 RETURNING *;



-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;