package handler

import (
	"fmt"

	"github.com/Asyadam/PairProjectP1/entity"
	"github.com/Asyadam/PairProjectP1/repository"
)

type AuthHandler struct {
	AuthRepo *repository.AuthRepository
}

func NewAuthHandler(authRepo *repository.AuthRepository) *AuthHandler {
	return &AuthHandler{AuthRepo: authRepo}
}

func (h *AuthHandler) Register() {

	var email string
	var password string

	fmt.Print("Input Email: ")
	fmt.Scan(&email)

	fmt.Print("Input Password: ")
	fmt.Scan(&password)

	user := entity.User{
		Email:    email,
		Password: password,
		Role:     "customer",
	}

	err := h.AuthRepo.CreateUser(user)

	if err != nil {
		fmt.Println("Register Failed:", err)
		return
	}

	fmt.Println("Register Success!")
}

func (h *AuthHandler) Login() (entity.User, bool) {
	var email string
	var password string
	var emptyUser entity.User

	fmt.Print("Input Email: ")
	fmt.Scan(&email)

	fmt.Print("Input Password: ")
	fmt.Scan(&password)

	user, err := h.AuthRepo.FindUserByEmail(email)

	if err != nil {
		fmt.Println("Email not found")
		return emptyUser, false
	}

	if user.Password != password {
		fmt.Println("Wrong Password")
		return emptyUser, false
	}

	fmt.Println("Login Success!")
	fmt.Println("Welcome,", user.Email)

	return user, true
}
