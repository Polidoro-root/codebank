package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pg"
	"ithub.com/Polidoro-root/codebank/domain"
	"ithub.com/Polidoro-root/codebank/infrastructure/repository"
	usecase "ithub.com/Polidoro-root/codebank/use_case"
)

func main() {
	db := setupDb()

	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Joao"
	cc.ExpirationYear = 2021
	cc.ExpirationMonth = 7
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	repo.CreateCreditCard(*cc)
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)

	useCase := usecase.NewUseCaseTransaction(transactionRepository)

	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s db=%s sslmode=disable", "db", "5432", "postgres", "root", "codebank")

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("error connection to database")
	}

	return db
}
