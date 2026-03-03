package model

type UserRegistrationRequest struct {
	ID        int    `json:"id"`
	FullName  string `json:"name" `
	Gender    string `json:"gender" `
	Age       int    `json:"age" `
	Email     string `json:"email"`
	Mobile    string `json:"mobileNumber" `
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}
