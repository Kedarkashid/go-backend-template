package model

type UserRegistrationRequest struct {
	FullName string `json:"name" `
	Gender   string `json:"gender" `
	Age      int    `json:"age" `
	Email    string `json:"email"`
	Mobile   string `json:"mobileNumber" `
	Password string `json:"password"`
}
