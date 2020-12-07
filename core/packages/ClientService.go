package services

import (
	"onlineBanking/core/database/postgres"
	"onlineBanking/core/models"
	"github.com/jackc/pgx/pgxpool"
	"context"
	"crypto/md5"
	"fmt"
	"log"
)

func AddClient(name string, surname string, login string, password string, age int, gender string, phone string, db *pgxpool.Pool) (err error) {
	status := true
	password = MakeHash(password)
	_, err = db.Exec(context.Background(), postgres.AddClient, name, surname, login, password, age, gender, phone, status)
	if err != nil {
		log.Fatalf("Пользователь недобавлен: %s", err)
		return err
	}
	return nil
}

func MakeHash(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func QueryError(text string) (err error) {
	return fmt.Errorf(text)
}

func GetAllClients(db *pgxpool.Pool) (clients []models.Client, err error) {
	rows, err := db.Query(context.Background(), postgres.GetAllClients)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("1 wrong")
		return nil, err
	}
	fmt.Println(err)
	defer rows.Close()
	//defer func() {
	//	if innerErr := rows.Close(); innerErr != nil {
	//		clients = nil
	//	}
	//}()

	for rows.Next() {
		client := models.Client{}
		err = rows.Scan(
			&client.ID,
			&client.Name,
			&client.Surname,
			&client.Login,
			&client.Password,
			&client.Age,
			&client.Gender,
			&client.Phone,
			&client.Status,
			&client.VerifiedAt)
		fmt.Println(err)
		if err != nil {
			log.Fatalf("2 wrong")
			return nil, err
		}
		clients = append(clients, client)
	}
	if rows.Err() != nil {
		log.Fatalf("3 wrong")
		return nil, rows.Err()
	}
	return clients, nil
}

func Login(login string, password string, db *pgxpool.Pool) (loginPredicate bool, err error) {
	var dbLogin, dbPassword string
	err = db.QueryRow(context.Background(), postgres.LoginSQL, login).Scan(&dbLogin, &dbPassword)

	if err != nil {
		//		fmt.Printf("%s, %e\n", loginSQL, err)
		return false, err
	}
	err = QueryError("Несовпадение пароля")
	if MakeHash(password) != dbPassword {
		//fmt.Println(makeHash(password), " ", dbPassword)
		return true, err
	}
	//fmt.Println(makeHash(password), " ", dbPassword)
	return true, nil
}

func SearchByLogin(login string, db *pgxpool.Pool) (id int64, surname string) {
	err := db.QueryRow(context.Background(), postgres.SearchClientByLogin, login).Scan(&id, &surname)
	if err != nil {
		log.Fatalf("Ошибка в %s", postgres.SearchClientByLogin)
	}
	return id, surname
}
