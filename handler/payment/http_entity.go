package payment

import (
	"errors"
)

type VerifyPaymentReq struct {
	Email string `json:"email"`
}

func (req VerifyPaymentReq) Validate() error {
	if len(req.Email) == 0 {
		return errors.New("invalid email")
	}
	return nil
}
