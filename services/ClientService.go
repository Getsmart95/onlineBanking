package services

import (
	"database/sql"
	"github.com/jackc/pgx/pgxpool"
	"log"
)

func AddClient(name, surname, login, password, phoneNumber string, db *pgxpool.Pool) (err error){
	locked := true
	password = makeHash(password)
	_, err = db.Exec(addClientDML, name, surname, login, password, phoneNumber, locked)
	if err != nil {
		log.Fatalf("Пользователь недобавлен: %s", err)
		return err
	}
	return nil
}


func GetAllClients(db *sql.DB) (clients []cmodels.Client, err error){
	rows, err := db.Query(getAllClients)
	if err != nil {
		log.Fatalf("1 wrong")
		return nil, err
	}

	defer func() {
		if innerErr := rows.Close(); innerErr != nil {
			clients = nil
		}
	}()

	for rows.Next(){
		client := cmodels.Client{}
		err = rows.Scan(&client.ID, &client.Name, &client.Surname, &client.NumberPhone, &client.Login, &client.Password, &client.Locked)
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
