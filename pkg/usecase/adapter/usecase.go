package adapter

import (
	"context"

	repository "github.com/thnkrn/go-gin-template/pkg/repository"
	usecase "github.com/thnkrn/go-gin-template/pkg/usecase"
)

type mockUsecase struct {
	repo repository.Repository
}

func NewUseCase(r repository.Repository) usecase.Usecase {
	return &mockUsecase{repo: r}
}

func (u *mockUsecase) Mock(ctx context.Context) string {
	repo := u.repo.Mock(ctx)
	return repo
}
