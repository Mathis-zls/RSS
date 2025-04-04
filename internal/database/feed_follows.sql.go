// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
Insert into feed_follows (id,id_feed,id_user,created_at,updated_at)
Values($1,$2,$3,$4,$5)
RETURNING id, id_user, id_feed, created_at, updated_at
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	IDFeed    uuid.UUID
	IDUser    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.IDFeed,
		arg.IDUser,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.IDUser,
		&i.IDFeed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
delete from feed_follows WHERE id=$1 and id_user=$2
`

type DeleteFeedFollowParams struct {
	ID     uuid.UUID
	IDUser uuid.UUID
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, arg.ID, arg.IDUser)
	return err
}

const getFeedFollows = `-- name: GetFeedFollows :many
Select id, id_user, id_feed, created_at, updated_at from feed_follows WHERE id_user=$1
`

func (q *Queries) GetFeedFollows(ctx context.Context, idUser uuid.UUID) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollows, idUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.IDUser,
			&i.IDFeed,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
