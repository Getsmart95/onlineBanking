package main

import (
	"Golang/onlineBanking/cli/controllers"
	"Golang/onlineBanking/core/database/postgres"
	"Golang/onlineBanking/cli/constants"
	"github.com/jackc/pgx/pgxpool"
	"context"
	"os"
	//"strings"
	//"encoding/xml"
	"fmt"
	//"io/ioutil"
	"log"
)

func main() {
	db, err := pgxpool.Connect(context.Background(), `postgresql://root@localhost:5432/postgres?sslmode=disable`)

	if err != nil {
		fmt.Println("Error", err)
	}
	defer db.Close()
	//err = db.Ping()
	if err != nil {
		fmt.Println("Нет подключения к серверу")
	}

	err = postgres.DbInit(db)
	if err != nil {
		log.Fatal("Error: #{err}")
	}
	AutorizeByAdmin(db)
}

func AutorizeByAdmin(db *pgxpool.Pool) {
	var cmd string
	for {
		fmt.Println(constants.AuthorizedOperations)
		fmt.Scan(&cmd)
		switch cmd {
		case "1":
			controllers.AddClientHandler(db)
		case "2":
			controllers.AddAccountHandler(db)
		case "3":
			controllers.AddServiceHandler(db)
		//case "4":
		//	controllers.AddClientsToJsonXmlFiles(db)
		//case "5":
		//	controllers.AddAccountsToJsonXmlFiles(db)
		//case "6":
		//	controllers.AddATMsToJsonXmlFiles(db)
		//case "7":
		//	controllers.AddClientsFromXmlJson(db)
		//case "8":
		//	controllers.AddAccountsFromXmlJson(db)
		//case "9":
		//	controllers.AddAtmFromXmlJson(db)
		case "10":
			controllers.AddATMHandler(db)
		case "q":
			os.Exit(0)
		default:
			fmt.Println("Введенно неверное значение, пробуйте еще раз\n")
			continue
		}
	}
}
