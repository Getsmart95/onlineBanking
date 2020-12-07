package controllers

import (
	"onlineBanking/cli/constants"
	"onlineBanking/core/packages"
	"context"
	"github.com/jackc/pgx/pgxpool"
	"fmt"
)

func Authorize(db *pgxpool.Pool) (id int64, err error) {
	fmt.Println("Введите логин:")
	var login string
	fmt.Scan(&login)
	fmt.Println("Введите пароль:")
	var password string
	fmt.Scan(&password)

	predicate, err := services.Login(login, password, db)
	if predicate == false {
		fmt.Println("Введен неправильный логин")
		return 0, err
	}

	if predicate == true && err != nil {
		fmt.Println("Введен неправильный пароль")
		return 0, err
	}
	id, surname := services.SearchByLogin(login, db)

	fmt.Printf("Добро пожаловать мистер %s\n", surname)
	return id, nil
}

func GetATMsForClient(db *pgxpool.Pool) (err error) {
	ms, err := services.GetAllATMs(db)
	if err != nil {
		return err
	}
	i := 0
	for _, value := range ms {
		i++
		fmt.Println(value)
	}
	if i == 0 {
		fmt.Println("Список банкоматов пуст")
	}
	return nil
}

//go install ./...
func SearchAccountByIdHandler(id int64, db *pgxpool.Pool) (accounts map[int64]int64, err error) {
	list, err := services.SearchAccountById(id, db)
	accounts = map[int64]int64{}
	if err != nil {
		fmt.Errorf("cant : %e", err)
		return nil, err
	}

	fmt.Println("Список ваших счетов:")
	for index, account := range list {
		index64 := int64(1 + index)
		accounts[index64] = account.AccountNumber
		fmt.Println(index+1, ".", account.ClientId, account.AccountNumber, account.Balance)
	}
	return accounts, nil
}

func AuthorizedOperations(id int64, db *pgxpool.Pool) {
	var cmd string
	for {
		fmt.Println(constants.AuthorizedTextOperations)
		fmt.Scan(&cmd)
		switch cmd {
		case "1":
			SearchAccountByIdHandler(id, db)
		case "2":
			ChooseAccountById(id, db)
		case "3":
			//err := PayServiceHandler(id, db)
			//if err != nil {
			//	log.Fatal("Uliya")
			//}
		case "q":
			return
		}
	}
}

func ChooseAccountById(id int64, db *pgxpool.Pool)(err error){
	AccountNumber, err := ChooseAccount(id, db)
	fmt.Println("Введите номер карты")
	var TransferCardNumber string
	fmt.Scan(&TransferCardNumber)
	fmt.Println("Введите сумму перевода")
	var amount int64
	fmt.Scan(&amount)
	err = TransferToAccount(AccountNumber, TransferCardNumber, amount, db)
	if err != nil {
		fmt.Println("Невозможно перевести деньги на этот счет")
	}
	return nil
}
////////////////////////
func TransferToAccount(AccountNumber int64, TransferCardNumber string, Amount int64, db *pgxpool.Pool) (err error) {
	//tx, err := db.Begin()
	//if err != nil {
	//	return err
	//}
	//defer func() {
	//	if err != nil {
	//		_ = tx.Rollback()
	//		return
	//	}
	//	err = tx.Commit()
	//}()
	_, err = db.Exec(context.Background(), `UPDATE accounts set balance = balance - ($1) where account_Number = ($2)`, Amount, AccountNumber)
	if err != nil {
		return err
	}

	_, err = db.Exec(context.Background(), `UPDATE accounts set balance = balance + ($1) where card_number = ($2)`, Amount, TransferCardNumber)
	if err != nil {
		return err
	}
	fmt.Println("Перевод денег успешно выполнено!")
	return nil
}

///////////////////////
func ChooseAccount(id int64, db *pgxpool.Pool) (AccountNumber int64, err error) {
	fmt.Println("Выберите счет:")
	accounts, err := SearchAccountByIdHandler(id, db)
	if err != nil {
		return -1, err
	}
	//	fmt.Println(accounts)

	for {
		var cmd int64
		fmt.Scan(&cmd)
		switch int64(len(accounts)) >= cmd && cmd > 0 {
		case true:
			return accounts[cmd], nil
		case false:
			fmt.Println("Введите заново в пределах количество ваших счетов")
		}
	}
	return -1, nil
}

///////////////////////
//
//func PayServiceHandler(id int64, db *pgxpool.Pool) (err error) {
//	fmt.Println("Выберите счет:")
//	accounts, err := SearchAccountByIdHandler(id, db)
//	if err != nil {
//		return err
//	}
//
//	for {
//		var cmd int64
//		fmt.Scan(&cmd)
//		switch int64(len(accounts)) >= cmd && cmd > 0 {
//		case true:
//			ChooseToService(accounts[cmd], db)
//			return nil
//		case false:
//			fmt.Println("Введите заново в пределах количество ваших счетов")
//		}
//	}
//	return nil
//}

//func GetAllServicesHandler(db *pgxpool.Pool) (err error) {
	//services, err := postgres.GetAllServices(db)
	//if err != nil {
	//	fmt.Errorf("Get all services didn't work %e", err)
	//	return nil
	//}
	//
	//for _, service := range services {
	//	fmt.Println(service.ID, service.Name, service.Price)
	//}
	//return nil
//}
//
//func ChooseToService(AccountNumber int64, db *pgxpool.Pool) (err error) {
//	fmt.Println("Выберите услугу: ")
//	err = GetAllServicesHandler(db)
//	if err != nil {
//		fmt.Errorf("GetServiceHandler %e", err)
//		return err
//	}
//	for {
//		var cmd int64
//		fmt.Scan(&cmd)
//		err := services.CheckServiceHaving(cmd, db)
//		if err != nil {
//			fmt.Println("Такой услуги нет, попробуйте еще раз")
//			continue
//		} else {
//			err := Transfer(AccountNumber, cmd, db)
//			if err != nil {
//				fmt.Println("Перевод невозможен")
//			}
//		}
//		return nil
//	}
//}
//
//func Transfer(accountNumber, ServiceID int64, db *pgxpool.Pool) (err error) {
//	tx, err := db.Begin()
//	if err != nil {
//		return err
//	}
//	defer func() {
//		if err != nil {
//			_ = tx.Rollback()
//			return
//		}
//		err = tx.Commit()
//	}()
//	var AccountBalance int64
//	err = tx.QueryRow(`select balance from accounts where accountNumber = ?`, accountNumber).Scan(&AccountBalance)
//	if err != nil {
//		return err
//	}
//	var ServicePrice int64
//	err = tx.QueryRow(`select price from services where id = ?`, ServiceID).Scan(&ServicePrice)
//	if err != nil {
//		return err
//	}
//
//	_, err = tx.Exec(`UPDATE services set balance = balance + price where id = ?`, ServiceID)
//	if err != nil {
//		return err
//	}
//	_, err = tx.Exec(`UPDATE accounts set balance = balance - ? where accountNumber = ?`, ServicePrice, accountNumber)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
