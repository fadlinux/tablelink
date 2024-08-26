package usecase

import (
	"context"
	mUser "tablelink/internal/model/users"
)

type Usecase interface {
	GetAllUser(ctx context.Context) (data mUser.ResponseDefault, err error)

	//for login
	Login(ctx context.Context, paramsInput mUser.RequestLogin) (exist bool, err error)

	AddUser(ctx context.Context, paramsInput mUser.ResponseLogin) (lastId int, err error)
	UpdateUser(ctx context.Context, name string) (err error)
	DeleteUser(ctx context.Context, userId int) (err error)
}
