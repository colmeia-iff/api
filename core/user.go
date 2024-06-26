package core

import (
	"context"
	"log"

	"go.mod/db"
	"go.mod/entity"
	"go.mod/middleware"

	"go.mod/rest"
)

type UserManager struct {
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (um *UserManager) User(ctx context.Context, id string) (*entity.UserInfoView, error) {
	userinfo, err := db.ReturnUserById(ctx, id)
	if err != nil {
		return nil, rest.LogError(err)
	}
	return userinfo, nil
}
func (um *UserManager) CreateUser(ctx context.Context, user entity.User) error {
	err := db.Create(ctx, user)
	if err != nil {
		return rest.LogError(err, "um.CreateUser db.Create")
	}
	return nil
}

func (*UserManager) Login(ctx context.Context, email, password string) (string, error) {
	// Implemente a lógica para verificar as credenciais do usuário no banco de dados.
	// Aqui, estou usando uma função fictícia chamada VerifyCredentials como exemplo.
	user, err := db.VerifyCredentials(ctx, email, password)
	if err != nil {
		return "", rest.LogError(err, "um.Login db.VerifyCredentials")
	}
	log.Println("teste agr")
	body := &entity.User{
		ID:    user.ID,
		Email: user.Email,
	}
	token, err := middleware.GenerateToken(body)
	if err != nil {
		return "", rest.LogError(err, "middleware.generatetoken")
	}
	return token, nil
}
