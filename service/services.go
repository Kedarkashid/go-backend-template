package service

import (
	"fmt"

	"github.com/Kedarkashid/go-backend-template/database"
	"github.com/Kedarkashid/go-backend-template/model"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(req model.UserRegistrationRequest) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}
	fmt.Println("Hashed Password:", string(hashedPassword))

	query := `
		INSERT INTO users 
		(name, gender, age, email, mobile_number, password)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err = database.DB.Exec(
		query,
		req.FullName,
		req.Gender,
		req.Age,
		req.Email,
		req.Mobile,
		hashedPassword,
	)

	if err != nil {
		return err
	}

	return nil
}
