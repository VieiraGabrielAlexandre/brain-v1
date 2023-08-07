package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateUserRequest struct {
	Nome           string `json:"nome" json:"nome"`
	ContaVinculada string `json:"conta_vinculada" json:"conta_vinculada"`
}

func (r *CreateUserRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateUserRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"nome.required":            "Nome é necessário",
		"conta_vinculada.required": "Conta vinculada é necessário",
	}
}

func (r *CreateUserRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateUserRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateUserRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
