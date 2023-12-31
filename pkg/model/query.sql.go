// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package model

import (
	"context"
)

const getUser = `-- name: GetUser :one
SELECT
    id, username
FROM
    users
WHERE
    id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Username)
	return i, err
}
