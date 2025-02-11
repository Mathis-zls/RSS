-- +goose Up

Alter Table users add column api_key Varchar(64) Unique not null default(
    encode(sha256(random()::text::bytea),'hex')
);

-- +goose Down
Alter Table users drop column api_key;