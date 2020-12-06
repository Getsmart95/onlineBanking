package main

import (
	"Golang/onlineBank/cli/clientCli/controllers"
	"context"
	"github.com/jackc/pgx/pgxpool"
	"fmt"
	"log"
	"os"

)
const unauthorizedOperations = `Список доступных операций:
1. Авторизация
3. Список банкоматов
q. Выйти из приложения

Введите команду`
func main() {
	db, err := pgxpool.Connect(context.Background(), `postgresql://root@localhost:5432/postgres?sslmode=disable`)
	//db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatalf("Ошибка открытия базы данных %s", err)
	}
	defer db.Close()

	mainAppFunction(db)
}
func mainAppFunction(db *pgxpool.Pool) {
	var cmd string
	for {
		fmt.Println(unauthorizedOperations)
		fmt.Scan(&cmd)
		switch cmd {
		case "1":
			id, err := controllers.Authorize(db)
			if err != nil {
				fmt.Println("Попробуйте еще раз")
				continue
			} else {
				controllers.AuthorizedOperations(id, db)
			}
		case "3":
			err := controllers.GetATMsForClient(db)
			if err != nil {
				fmt.Printf("Ошибка выдачи списка банкоматов %s:", err)
			}
		case "q":
			os.Exit(0)
		default:
			fmt.Println("Введена неверная команда, попробуйте еще раз", unauthorizedOperations)
			continue
		}
	}
}
