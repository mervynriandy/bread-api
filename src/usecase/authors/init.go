package author_usecase

import (
	author_repo "bread-api/src/repository/authors"
)

type AuthorUsecase struct {
	authorRepository author_repo.AuthorRepository
}

func AuthorCase(ar author_repo.AuthorRepository) *AuthorUsecase {
	return &AuthorUsecase{
		authorRepository: ar,
	}
}
