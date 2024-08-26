package postgre

import (
	"context"
	mUser "tablelink/internal/model/users"
)

func (ar *postgreUserRepo) FetchAll(ctx context.Context) (data mUser.ResponseDefault, err error) {
	return data, err
	//

}

func (ar *postgreUserRepo) FetchByid(ctx context.Context, paramsInput mUser.RequestLogin) (err error) {
	return err
	//

}
