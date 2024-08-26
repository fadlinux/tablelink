package usecase

import (
	rUser "tablelink/internal/repository"
)

type (
	userUsecase struct {
		postgreUserRepo rUser.Repository
	}
)

func NewUserUsecase(postgreUser rUser.Repository) Usecase {
	return &userUsecase{
		postgreUser,
	}
}
