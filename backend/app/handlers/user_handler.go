package handlers

import "github.com/kishansakhiya/wails-demo/backend/app/models"

type UserHandler struct{}

func (u *UserHandler) GetUser() models.User {
	return models.User{
		Name: "Kishan Sakhiya", 
		Role: "Developer",
		Email: "kishansakhiya@gmail.com",
		Phone: "9876543210",
		Address: "123, Main Street, Anytown, USA",
		City: "Anytown",
		State: "CA",
		Zip: "12345",
	}
}
