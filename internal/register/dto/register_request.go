package register

type RegisterRequestBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}

type EmailVerification struct {
	SUBJECT           string
	EMAIL             string
	VERIFICATION_CODE string
}
