package usecase

import (
	"context"
	user "tablelink/internal/model/users"
)

// AddUser implements Usecase.
func (u *userUsecase) AddUser(ctx context.Context, paramsInput user.ResponseLogin) (lastId int, err error) {
	panic("unimplemented")
}

// UpdateUser implements Usecase.
func (u *userUsecase) UpdateUser(ctx context.Context, name string) (err error) {
	panic("unimplemented")
}
