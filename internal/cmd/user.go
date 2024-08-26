package user

import (
	"database/sql"
	"os"
	"tablelink/common/postgre"

	_ "github.com/lib/pq"

	httpSystem "tablelink/internal/delivery/http"

	userRepo "tablelink/internal/repository"
	postgreUserRepository "tablelink/internal/repository/postgre"

	uUser "tablelink/internal/usecase/user"
)

var (
	postgreDB *sql.DB

	postgreUserRepo userRepo.Repository

	userUsecase uUser.Usecase

	HTTPDelivery httpSystem.Delivery
)

func Initialize() {
	dsn := os.Getenv("DATABASE_URL")
	postgreDB = postgre.NewDBConnection("postgres", dsn)

	postgreUserRepo = postgreUserRepository.NewPostgreUserRepo(postgreDB)

	userUsecase = uUser.NewUserUsecase(postgreUserRepo)
	HTTPDelivery = httpSystem.NewUserHTTP(userUsecase)

}
