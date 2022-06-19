package author_usecase

import (
	author_repo "victoria-falls/src/repository/authors"

	"go.uber.org/zap"
)

type AuthorUsecase struct {
	authorRepository author_repo.AuthorRepository
	logger           *zap.Logger
}

func AuthorCase(ar author_repo.AuthorRepository, logger *zap.Logger) *AuthorUsecase {
	return &AuthorUsecase{
		authorRepository: ar,
		logger:           logger,
	}
}
