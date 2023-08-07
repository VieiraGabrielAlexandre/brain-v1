package controllers

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	models "goravel/src/Adapter/Driven/Models"
	users_repository "goravel/src/Adapter/Driven/Repositories"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) FindOne(ctx http.Context) {
	token := ctx.Request().Route("token")

	user, err := users_repository.FindOne(token)

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

func (r *UserController) Create(ctx http.Context) {
	var users models.Users

	err := ctx.Request().Bind(&users)

	if err != nil {
		fmt.Println(err)

		ctx.Response().Success().Json(http.Json{
			"data": err,
		})

		return
	}

	result := users_repository.CreateUser(users)

	ctx.Response().Success().Json(http.Json{
		"data": result,
	})
}
