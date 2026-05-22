package handler

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/Asyadam/PairProjectP1/entity"
	"github.com/Asyadam/PairProjectP1/repository"
)

type ProfileHandler struct {
	ProfileRepo *repository.ProfileRepository
}

func NewProfileHandler(profileRepo *repository.ProfileRepository) *ProfileHandler {
	return &ProfileHandler{ProfileRepo: profileRepo}
}

func (h *ProfileHandler) ViewProfile(userID int) {
	profile, err := h.ProfileRepo.GetProfileByUserID(userID)

	if err != nil {
		fmt.Println("Get Profile Failed:", err)
		return
	}

	fmt.Println("==== MY PROFILE ====")
	fmt.Println("Profile ID:", profile.ID)
	fmt.Println("User ID:", profile.UserID)
	fmt.Println("Full Name:", profile.FullName)
	fmt.Println("Phone:", profile.Phone)
	fmt.Println("Address:", profile.Address)
	fmt.Println("====================")
}

func (h *ProfileHandler) UpdateProfile(reader *bufio.Reader, userID int) {
	var profile entity.Profile

	profile.UserID = userID

	fmt.Print("Input New Full Name: ")
	fullName, _ := reader.ReadString('\n')
	profile.FullName = strings.TrimSpace(fullName)

	fmt.Print("Input New Phone: ")
	phone, _ := reader.ReadString('\n')
	profile.Phone = strings.TrimSpace(phone)

	fmt.Print("Input New Address: ")
	address, _ := reader.ReadString('\n')
	profile.Address = strings.TrimSpace(address)

	err := h.ProfileRepo.UpdateProfile(profile)

	if err != nil {
		fmt.Println("Update Profile Failed:", err)
		return
	}

	fmt.Println("Update Profile Success!")
}
