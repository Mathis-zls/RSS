-- +goose Up
Alter table feed Add column last_fetchedAt Timestamp;

-- +goose Down
Alter table feed Drop last_fetchedAt;