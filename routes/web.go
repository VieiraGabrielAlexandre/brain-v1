package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/http/middleware"

	"goravel/app/http/controllers"
)

func Api() {
	invoicesController := controllers.NewInvoiceController()
	usersController := controllers.NewUserController()

	facades.Route().Get("/invoices", invoicesController.Show)
	facades.Route().Middleware(middleware.AuthToken()).Post("/invoices", invoicesController.Create)

	facades.Route().Post("/users", usersController.Create)
	facades.Route().Get("/users/{token}", usersController.FindOne)

	facades.Route().Fallback(func(ctx http.Context) {
		ctx.Response().String(404, "not found")
	})

}
