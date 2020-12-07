package services

import (
	"onlineBanking/core/models"
	"onlineBanking/core/database/postgres"
	"context"
	"github.com/jackc/pgx/pgxpool"
	"fmt"
)

func AddService(Name string, AccountNumber int64, db *pgxpool.Pool) (err error){
	_, err = db.Exec(context.Background(), postgres.AddService, Name, AccountNumber)
	if err != nil {
		fmt.Errorf("Error in %s, err: %e", postgres.AddService, err)
		return err
	}
	return nil
}


func GetAllServices(db *pgxpool.Pool) (Services []models.Service, err error){
	rows, err := db.Query(context.Background(), postgres.GetAllServices)
	if err != nil {
		fmt.Errorf("%s, %e",postgres.GetAllServices, err)
		return nil, err
	}
	//defer func() {
	//	if innerErr := rows.Close(); innerErr != nil {
	//		Services = nil
	//	}
	//}()

	for rows.Next(){
		Service := models.Service{}
		err := rows.Scan(&Service.ID, &Service.Name)
		if err != nil {
			fmt.Errorf("%s, %e",postgres.GetAllServices, err)
			return nil, err
		}
		Services = append(Services, Service)
	}
	if rows.Err() != nil{
		fmt.Errorf("%s, %e",postgres.GetAllServices, rows.Err())
		return nil, rows.Err()
	}
	return Services, nil
}

func CheckServiceHaving(cmd int64, db *pgxpool.Pool) (err error){
	var id int64
	err = db.QueryRow(context.Background(), `select id from services where id = ($1)`, cmd).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
