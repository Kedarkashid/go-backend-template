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
func FetchAllUsersService() ([]model.UserRegistrationRequest, error) {

	query := `select id, name, gender, age, email, mobile_number, password, created_at from users;`
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return []model.UserRegistrationRequest{}, err
	}
	defer rows.Close()
	var users []model.UserRegistrationRequest

	for rows.Next() {
		var user model.UserRegistrationRequest
		err := rows.Scan(
			&user.ID,
			&user.FullName,
			&user.Gender,
			&user.Age,
			&user.Email,
			&user.Mobile,
			&user.Password,
			&user.CreatedAt,
		)
		if err != nil {
			fmt.Println("Error scanning user row:", err)
			return []model.UserRegistrationRequest{}, err
		}
		fmt.Println(user)
		users = append(users, user)
	}
	return users, nil
}
