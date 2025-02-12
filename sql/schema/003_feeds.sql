-- +goose Up
Create TABLE feed (
 id UUID PRIMARY KEY,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 name TEXT NOT NULL,
 url TEXT Unique Not NULL,
 user_id UUID Not NULL References users(id) on delete cascade
);


-- +goose Down
 DROP TABLE feed;