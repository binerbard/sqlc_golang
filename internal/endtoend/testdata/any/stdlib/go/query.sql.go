// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/lib/pq"
)

const any = `-- name: Any :many
SELECT id
FROM bar
WHERE id = ANY($1::bigint[])
`

func (q *Queries) Any(ctx context.Context, dollar_1 []int64) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, any, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
