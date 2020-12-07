package controllers

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"log"
	"onlineBanking/cli/constants"
	"onlineBanking/core/packages"
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
			err := PayServiceHandler(id, db)
			if err != nil {
				log.Fatal("Uliya")
			}
		case "q":
			return
		}
	}
}

func ChooseAccountById(id int64, db *pgxpool.Pool) (err error) {
	AccountNumber, err := ChooseAccount(id, db)
	fmt.Println("Введите номер карты")
	var TransferCardNumber string
	fmt.Scan(&TransferCardNumber)
	fmt.Println("Введите сумму перевода")
	var Amount int64
	fmt.Scan(&Amount)
	fmt.Println("Введите сообщение получателю")
	var Message string
	fmt.Scan(&Message)
	var TransferAccountNumber int64

	err = db.QueryRow(context.Background(), `select account_number from accounts where card_number = ($1)`, TransferCardNumber).Scan(&TransferAccountNumber)
	if err != nil {
		fmt.Println("Карта не существует")
		return nil
	}
	err = TransferToAccount(AccountNumber, TransferCardNumber, Amount, Message, TransferAccountNumber, db)
	if err != nil {
		fmt.Println("Невозможно перевести деньги на этот счет")
	}
	return nil
}

////////////////////////
func TransferToAccount(AccountNumber int64, TransferCardNumber string, Amount int64, Message string, TransferAccountNumber int64, db *pgxpool.Pool) (err error) {

	var ServiceId int64
	ServiceId = 1
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
	_, err = db.Exec(context.Background(), `UPDATE accounts set balance = balance - ($1) 
								                 where account_Number = ($2)`, Amount, AccountNumber)
	if err != nil {
		return err
	}

	_, err = db.Exec(context.Background(), `UPDATE accounts set balance = balance + ($1) 
                                                 where card_number = ($2)`, Amount, TransferCardNumber)

	if err != nil {
		return err
	}
	fmt.Println(AccountNumber,TransferCardNumber,Amount, Message,ServiceId)
	_, err = db.Exec(context.Background(), `insert into histories(sender_account_number, recipient_account_number, money, message, service_id)
											values( $1, $2, $3, $4, $5 )`, AccountNumber, TransferAccountNumber, Amount, Message, ServiceId)
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
func PayServiceHandler(id int64, db *pgxpool.Pool) (err error) {
	fmt.Println("Выберите счет:")
	accounts, err := SearchAccountByIdHandler(id, db)
	if err != nil {
		return err
	}

	for {
		var cmd int64
		fmt.Scan(&cmd)
		switch int64(len(accounts)) >= cmd && cmd > 0 {
		case true:
			ChooseToService(accounts[cmd], db)
			return nil
		case false:
			fmt.Println("Введите заново в пределах количество ваших счетов")
		}
	}
	return nil
}

func GetAllServicesHandler(db *pgxpool.Pool) (err error) {
services, err := services.GetAllServices(db)
if err != nil {
	fmt.Errorf("Get all services didn't work %e", err)
	return nil
}

for _, service := range services {
	fmt.Println(service.ID, service.Name, service.AccountNumber)
}
return nil
}

func ChooseToService(AccountNumber int64, db *pgxpool.Pool) (err error) {
	fmt.Println("Выберите услугу: ")
	err = GetAllServicesHandler(db)
	if err != nil {
		fmt.Errorf("GetServiceHandler %e", err)
		return err
	}
	for {
		var cmd int64
		fmt.Scan(&cmd)
		err := services.CheckServiceHaving(cmd, db)
		if err != nil {
			fmt.Println("Такой услуги нет, попробуйте еще раз")
			continue
		} else {
			fmt.Println("Введите сумму оплаты: ")
			var Ammount int64
			fmt.Scan(&Ammount)
			err := Transfer(AccountNumber, Ammount, cmd, db)
			if err != nil {
				fmt.Println("Перевод невозможен")
			}
		}
		return nil
	}
}

func Transfer(accountNumber int64, Ammount int64, ServiceID int64, db *pgxpool.Pool) (err error) {
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
	var Message string
	Message = "Оплата услуги"

	//var AccountBalance int64
	//err = db.QueryRow(context.Background(), `select balance from accounts where accountNumber = ($1)`, accountNumber).Scan(&AccountBalance)
	//if err != nil {
	//	return err
	//}
	var ServiceAccountNumber int64
	err = db.QueryRow(context.Background(), `select account_number from services where id = ($1)`, ServiceID).Scan(&ServiceAccountNumber)
	if err != nil {
		return err
	}

	_, err = db.Exec(context.Background(), `UPDATE accounts set balance = balance - ($1) where account_number = ($2)`, Ammount, accountNumber)
	if err != nil {
		return err
	}
	_, err = db.Exec(context.Background(), `UPDATE accounts set balance = balance + ($1) where account_number = ($2)`, Ammount, ServiceAccountNumber)
	if err != nil {
		return err
	}
	_, err = db.Exec(context.Background(), `insert into histories(sender_account_number, recipient_account_number, money, message, service_id)
											values( $1, $2, $3, $4, $5 )`, accountNumber, ServiceAccountNumber, Ammount, Message, ServiceID)
	if err != nil {
		return err
	}
	//_, err = db.Exec(context.Background(), `UPDATE accounts set balance = balance - ? where accountNumber = ?`, ServicePrice, accountNumber)
	//if err != nil {
	//	return err
	//}

	return nil
}
