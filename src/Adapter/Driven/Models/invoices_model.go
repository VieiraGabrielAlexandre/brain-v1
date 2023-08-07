package models

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoices struct {
	gorm.Model
	TokenUser     string  `json:"user_token" validate:"nonzero"`
	Token         string  `json:"token" validate:"nonzero"`
	Value         float32 `json:"value" validate:"nonzero"`
	Description   string  `json:"description" validate:"nonzero"`
	PaymentMethod string  `json:"payment_method" validate:"nonzero"`
}

func (u *Invoices) BeforeCreate(tx *gorm.DB) (err error) {
	token := uuid.New().String()
	u.Token = fmt.Sprintf("%x", token[:])
	return nil
}
