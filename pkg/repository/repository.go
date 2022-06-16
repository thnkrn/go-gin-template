package repository

import "context"

type Repository interface {
	Mock(ctx context.Context) string
}
