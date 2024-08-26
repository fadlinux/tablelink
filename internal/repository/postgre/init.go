package postgre

import (
	"database/sql"
	rUser "tablelink/internal/repository"

	_ "github.com/lib/pq"
)

type (
	postgreUserRepo struct {
		conn *sql.DB
	}
)

func NewPostgreUserRepo(con *sql.DB) rUser.Repository {
	repo := &postgreUserRepo{
		conn: con,
	}

	repo.prepareAddUserStmt()
	repo.prepareUpdateUserStmt()
	repo.prepareDeleteUserStmt()
	repo.prepareLoginStmt()
	repo.prepareAlluserStmt()

	return repo
}
