package adapter

import (
	"context"

	repository "github.com/thnkrn/go-gin-template/pkg/repository"
)

type mockRepo struct {
}

func NewRepo() repository.Repository {
	return &mockRepo{}
}

func (u *mockRepo) Mock(ctx context.Context) string {
	return "Hello"
}
