package services

import (
	"onlineBanking/core/database/postgres"
	"onlineBanking/core/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"log"
)

func AddAccount(clientId int64, balance int64, status bool, cardNumber string, db *pgxpool.Pool) (err error) {
	var count int
	//Number = 1_000_000_000_000_000
	err = db.QueryRow(context.Background(), `select count(*) from accounts`).Scan(&count)
	if err != nil {
		fmt.Errorf("cant %e", err)
		return err
	}
	var accountNumber int64
	var lastAccountNumber int64
	accountNumber = 2
	if count != 0 {
		err := db.QueryRow(context.Background(), `select max(account_number) from accounts`).Scan(&lastAccountNumber)
		if err != nil {
			fmt.Errorf("cant find last AccountWithUserName Number %e", err)
			return err
		}

		accountNumber = lastAccountNumber + 1
	}
	_, err = db.Exec(context.Background(), postgres.AddAccount, clientId, accountNumber, balance, status, cardNumber)
	fmt.Println(err)
	if err != nil {
		fmt.Errorf("cant insert %e", err)
		return err
	}

	fmt.Println("Success")
	return nil
}

func SearchAccountById(id int64, db *pgxpool.Pool) (Accounts []models.AccountForUser, err error) {
	var account models.AccountForUser

	rows, err := db.Query(context.Background(), postgres.SearchAccountByID, id)
	fmt.Println(err)
	if err != nil {
		fmt.Errorf("Активных аккаунтов нет %e\n", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(
			&account.ID,
			&account.ClientId,
			&account.AccountNumber,
			&account.Balance,
			&account.Status,
			&account.CardNumber)

		Accounts = append(Accounts, account)
	}
	return Accounts, nil
}

func GetAllAccounts(db *pgxpool.Pool) (accounts []models.AccountWithUserName, err error) {
	rows, err := db.Query(context.Background(), postgres.GetAllAccounts)
	if err != nil {
		log.Fatalf("1 wrong (Accc)")
		return nil, err
	}

	// defer func() {
	// 	if innerErr := rows.Close(); innerErr != nil {
	// 		accounts = nil
	// 	}
	// }()

	for rows.Next() {
		account := models.AccountWithUserName{}
		err = rows.Scan(
			&account.Account.ID,
			&account.Account.ClientId,
			&account.Account.AccountNumber,
			&account.Account.Balance,
			&account.Account.Status,
			&account.Client.ID,
			&account.Client.Name,
			&account.Client.Surname,
			&account.Client.Login,
			&account.Client.Password,
			&account.Client.Phone,
			&account.Client.Status,
			&account.Client.VerifiedAt)
		fmt.Println(err)
		if err != nil {
			log.Fatalf("2 wrong (Accc)")
			return nil, err
		}
		accounts = append(accounts, account)

	}
	if rows.Err() != nil {
		log.Fatalf("3 wrong (Accc)")
		return nil, rows.Err()
	}
	return accounts, nil
}
