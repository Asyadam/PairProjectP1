package handler

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/Asyadam/PairProjectP1/entity"
	"github.com/Asyadam/PairProjectP1/repository"
)

type AuthHandler struct {
	AuthRepo *repository.AuthRepository
}

func NewAuthHandler(authRepo *repository.AuthRepository) *AuthHandler {
	return &AuthHandler{AuthRepo: authRepo}
}

func (h *AuthHandler) Register(reader *bufio.Reader) {
	var user entity.User
	var profile entity.Profile

	fmt.Print("Input Email: ")
	email, _ := reader.ReadString('\n')
	user.Email = strings.TrimSpace(email)

	fmt.Print("Input Password: ")
	password, _ := reader.ReadString('\n')
	user.Password = strings.TrimSpace(password)

	fmt.Print("Input Full Name: ")
	fullName, _ := reader.ReadString('\n')
	profile.FullName = strings.TrimSpace(fullName)

	fmt.Print("Input Phone: ")
	phone, _ := reader.ReadString('\n')
	profile.Phone = strings.TrimSpace(phone)

	fmt.Print("Input Address: ")
	address, _ := reader.ReadString('\n')
	profile.Address = strings.TrimSpace(address)

	user.Role = "customer"

	err := h.AuthRepo.CreateUserWithProfile(user, profile)

	if err != nil {
		fmt.Println("Register Failed:", err)
		return
	}

	fmt.Println("Register Success!")
}

func (h *AuthHandler) Login(reader *bufio.Reader) (entity.User, bool) {
	var emptyUser entity.User

	fmt.Print("Input Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Input Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

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
