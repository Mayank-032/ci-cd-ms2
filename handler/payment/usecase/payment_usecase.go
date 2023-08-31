package usecase

import (
	"context"
	"errors"
	"log"
	"microservice2/handler/payment/repository"
)

func ValidatePayment(ctx context.Context, email string) error {
	fetchedPayment, err := repository.FetchPayment(ctx, email)
	if err != nil {
		return errors.New("unable_to_fetch_password")
	}

	if len(fetchedPayment.AmountPaid) == 0 {
		log.Println("Error: no payment initiated")
		return errors.New("no_payment_initiated")
	}
	return errors.New("unable_to_validate_user")
}
