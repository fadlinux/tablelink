package repository

import (
	"context"
	mUser "tablelink/internal/model/users"
)

type Repository interface {
	FetchAll(ctx context.Context) (data mUser.ResponseDefault, err error)
	//for login
	FetchByid(ctx context.Context, paramsInput mUser.RequestLogin) (err error)

	AddUser(ctx context.Context, paramsInput mUser.RequestLogin) (lastId int, err error)
	UpdateUser(ctx context.Context, name string) (err error)
	DeleteUser(ctx context.Context, userId int) (err error)
}
