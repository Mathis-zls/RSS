-- +goose Up
Create Table feed_follows(
    id UUID PRIMARY KEY,
    id_user UUID Not NULL References users(id) On delete cascade,
    id_feed UUID NOT NULL References feed (id),
    created_at TIMESTAMP Not NULL,
    updated_at TIMESTAMP Not NULL,
   Unique(id_user,id_feed) 
);



-- +goose Down
drop TABLE feed_follows;