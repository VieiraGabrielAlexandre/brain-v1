package controllers

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	models "goravel/src/Adapter/Driven/Models"
	invoices_repository "goravel/src/Adapter/Driven/Repositories"
)

type InvoiceController struct {
	//Dependent services
}

func NewInvoiceController() *InvoiceController {
	return &InvoiceController{
		//Inject services
	}
}

func (r *InvoiceController) Show(ctx http.Context) {
	var invoices []models.Invoices
	tokenUser := ctx.Request().Route("token_user")

	err := ctx.Request().Bind(&invoices)

	if err != nil {
		fmt.Println(err)

		ctx.Response().Success().Json(http.Json{
			"data": err,
		})

		return
	}

	result, _ := invoices_repository.FindByTokenUser(tokenUser)

	ctx.Response().Success().Json(http.Json{
		"data": result,
	})
}

func (r *UserController) FindOneInvoice(ctx http.Context) {
	token := ctx.Request().Route("token")

	user, err := invoices_repository.FindOneByToken(token)

	if err != nil {
		ctx.Response().Status(422).Json(http.Json{
			"error": err,
		})
		return
	}

	ctx.Response().Success().Json(http.Json{
		"data": user,
	})
}

func (r *InvoiceController) Create(ctx http.Context) {
	var invoices models.Invoices

	err := ctx.Request().Bind(&invoices)

	if err != nil {
		fmt.Println(err)

		ctx.Response().Success().Json(http.Json{
			"data": err,
		})

		return
	}

	result := invoices_repository.CreateInvoice(invoices)

	ctx.Response().Success().Json(http.Json{
		"data": result,
	})

}
