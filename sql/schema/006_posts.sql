-- +goose Up
Create Table posts(
   id UUID PRIMARY KEY,
   created_at TIMESTAMP Not NULL,
   updated_at TIMESTAMP Not NULL,
   title Text Not NULL,
   description Text,
   published_at Timestamp NOT NULL,
   url Text NOT NULL Unique,
   feed_id UUID NOT NULL References feed(id) ON DElete cascade
);
-- +goose Down
Drop table posts;