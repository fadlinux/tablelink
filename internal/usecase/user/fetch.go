package usecase

import (
	"context"
	"log"
	mUser "tablelink/internal/model/users"
)

// GetAllUser implements Usecase.
func (u *userUsecase) GetAllUser(ctx context.Context) (data mUser.ResponseDefault, err error) {
	data, err = u.postgreUserRepo.FetchAll(ctx)

	if err != nil {
		log.Println("Usecase GetAllUser error :", err)
	}

	return data, err
}

// Login implements Usecase.
func (u *userUsecase) Login(ctx context.Context, paramsInput mUser.RequestLogin) (exist bool, err error) {
	var count int64

	err = u.postgreUserRepo.FetchByid(ctx, paramsInput)
	if err != nil {
		exist = false
		log.Println("Usecase Login error :", err)
	}

	if count > 0 {
		exist = true
	}

	return exist, err

}
