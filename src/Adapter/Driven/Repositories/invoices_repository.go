package invoices_repository

import (
	"fmt"
	"github.com/goravel/framework/facades"
	"github.com/pkg/errors"
	models "goravel/src/Adapter/Driven/Models"
)

func CreateInvoice(user models.Invoices) models.Invoices {
	err := facades.Orm().Query().Create(&user)

	fmt.Println(err)
	if err != nil {
		return models.Invoices{}
	}

	return user
}

func FindOneInvoice(token string) (models.Invoices, error) {
	invoice := models.Invoices{}

	facades.Orm().Query().Where("token = ?", token).First(&invoice)

	if invoice.Token == "" {
		return invoice, errors.New("invoice not found")
	}

	return invoice, nil
}
