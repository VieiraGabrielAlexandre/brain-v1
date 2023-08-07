package invoices_repository

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/facades"
	models "goravel/src/Adapter/Driven/Models"
)

func FindOne(token string) (models.Users, error) {
	user := models.Users{}

	facades.Orm().Query().Where("token = ?", token).First(&user)

	if user.Token == "" {
		return user, errors.New("user not found")
	}

	return user, nil
}

func CreateUser(user models.Users) models.Users {
	err := facades.Orm().Query().Create(&user)

	fmt.Println(err)
	if err != nil {
		return models.Users{}
	}

	return user
}
