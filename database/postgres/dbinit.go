package postgres
import (
	"context"
	"github.com/jackc/pgx/pgxpool"
	"log"
)

func DbInit(database *pgxpool.Pool){
	DDLs := []string{clientsTable, clientsAccountsTable, ATMsTable, servicesTable, historiesTable}
	for _, ddl := range DDLs{
		_, err := database.Exec(context.Background(), ddl)
		if err != nill{
			log.Fatalf("Can't init this #{ddl} error is ${err}")
		}
	}
}