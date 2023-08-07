package models

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Nome           string `json:"nome" validate:"nonzero"`
	Token          string `json:"token" validate:"nonzero"`
	ContaVinculada int    `json:"conta_vinculada" validate:"nonzero"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	token := uuid.New().String()
	u.Token = fmt.Sprintf("%x", token[:])
	return nil
}
