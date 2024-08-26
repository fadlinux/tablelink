package postgre

import (
	"context"
	"log"
	mUser "tablelink/internal/model/users"
)

func (ar *postgreUserRepo) AddUser(ctx context.Context, paramsInput mUser.RequestLogin) (lastId int, err error) {
	err = addUserStmt.QueryRow(paramsInput.Email, paramsInput.Password).Scan(&lastId)
	if err != nil {
		log.Println("Repository AddCustomer error", err)
		return
	}
	return lastId, err

}

func (ar *postgreUserRepo) UpdateUser(ctx context.Context, name string) (err error) {
	return err
}
