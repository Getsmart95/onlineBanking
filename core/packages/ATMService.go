package services

import (
	"onlineBanking/core/database/postgres"
	"onlineBanking/core/models"
	"context"
	"github.com/jackc/pgx/pgxpool"
	"log"
)

func AddATM(name string, status bool, db *pgxpool.Pool) (err error) {
	_, err = db.Exec(context.Background(), postgres.AddATMs, name, status)
	if err != nil {
		log.Fatalf("Запись недобавлена: %e", err)
		return err
	}
	return nil
}

func GetAllATMs(db *pgxpool.Pool) (ATMs []models.ATM, err error) {
	rows, err := db.Query(context.Background(), postgres.GetAllATMs)
	if err != nil {
		log.Fatalf("1 wrong")
		return nil, err
	}


	defer rows.Close()
	//defer func() {
	//	if innerErr := rows.Close(); innerErr != nil {
	//		ATMs = nil
	//	}
	//}()

	for rows.Next() {
		ATM := models.ATM{}
		err = rows.Scan(&ATM.ID, &ATM.Name, &ATM.Status)
		if err != nil {
			log.Fatalf("2 wrong")
			return nil, err
		}
		ATMs = append(ATMs, ATM)
	}
	if rows.Err() != nil {
		log.Fatalf("3 wrong")
		return nil, rows.Err()
	}
	return ATMs, nil
}
