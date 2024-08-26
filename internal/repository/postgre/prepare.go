package postgre

import (
	"database/sql"
	"log"
)

var (
	addUserStmt    *sql.Stmt
	updateUserStmt *sql.Stmt
	deleteUserStmt *sql.Stmt
	loginUserStmt  *sql.Stmt
	getAllUserStmt *sql.Stmt
)

func (ar *postgreUserRepo) prepareAddUserStmt() {
	var err error
	addUserStmt, err = ar.conn.Prepare("INSERT INTO users(role_id,email,password, name, last_access, created_at) VALUES($1,$2,$3,$4,$5,now()) RETURNING user_id")
	if err != nil {
		log.Fatal("[Postgre: prepareAddUserStmt] Prepare statement fail :", err)
	}
}

func (ar *postgreUserRepo) prepareUpdateUserStmt() {
	var err error
	updateUserStmt, err = ar.conn.Prepare("UPDATE users SET name = $1 WHERE user_id = $4 ")
	if err != nil {
		log.Fatal("[Postgre: prepareUpdateUserStmt] Prepare statement fail :", err)
	}
}

func (ar *postgreUserRepo) prepareDeleteUserStmt() {
	var err error
	deleteUserStmt, err = ar.conn.Prepare("DELETE FROM users WHERE user_id = $1 ")
	if err != nil {
		log.Fatal("[Postgre: prepareDeleteUserStmt] Prepare statement fail :", err)
	}
}

func (ar *postgreUserRepo) prepareLoginStmt() {
	var err error
	loginUserStmt, err = ar.conn.Prepare("SELECT COUNT(user_id) AS total FROM users WHERE user_id = $1 ")
	if err != nil {
		log.Fatal("[Postgre: prepareLoginStmt] Prepare statement fail :", err)
	}
}

func (ar *postgreUserRepo) prepareAlluserStmt() {
	var err error
	getAllUserStmt, err = ar.conn.Prepare("SELECT * FROM users")
	if err != nil {
		log.Fatal("[Postgre: prepareAlluserStmt] Prepare statement fail :", err)
	}
}
