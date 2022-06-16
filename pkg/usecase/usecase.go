package usecase

import "context"

type Usecase interface {
	Mock(ctx context.Context) string
}
