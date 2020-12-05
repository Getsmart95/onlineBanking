package services

import (
	"Golang/onlineBank_core/database/postgres"
	models "Golang/onlineBank_core/models"
	"context"
	"github.com/jackc/pgx/pgxpool"
	"fmt"
	"log"

)


/////////////////////----------CLIENT---------//////////////////////
func QueryError(text string) (err error){
	return fmt.Errorf(text)
}

func Login(login, password string, db *pgxpool.Pool) (loginPredicate bool, err error){
	var dbLogin, dbPassword string
	err = db.QueryRow(context.Background(), postgres.LoginSQL, login).Scan(&dbLogin, &dbPassword)
	if err != nil {
		//		fmt.Printf("%s, %e\n", loginSQL, err)
		return false, err
	}
	err = QueryError("Несовпадение пароля")
	if models.MakeHash(password) != dbPassword {
		//fmt.Println(makeHash(password), " ", dbPassword)
		return true, err
	}
	//fmt.Println(makeHash(password), " ", dbPassword)
	return true, nil
}

func SearchByLogin(login string, db *pgxpool.Pool) (id int64, surname string){
	err := db.QueryRow(context.Background(), postgres.SearchClientByLogin, login).Scan(&id, &surname)
	if err != nil {
		log.Fatalf("Ошибка в %s", searchClientByLogin)
	}
	return id, surname
}
