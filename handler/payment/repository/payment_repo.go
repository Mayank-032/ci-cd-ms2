package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"microservice2/database"
	"microservice2/domain/entity"
)

func FetchPayment(ctx context.Context, email string) (entity.PaymentDetails, error) {
	db := database.DB

	sqlQuery := `
		SELECT amount_paid FROM payments WHERE email = ?
	`

	var amountPaid string
	err := db.QueryRowContext(ctx, sqlQuery, email).Scan(&amountPaid)
	if err != nil {
		log.Println("Error: " + err.Error())
		if err == sql.ErrNoRows {
			return entity.PaymentDetails{}, nil
		}
		return entity.PaymentDetails{}, errors.New("unable_to_make_db_call")
	}
	return entity.PaymentDetails{
		Email: email,
		AmountPaid: amountPaid,
	}, nil
}
