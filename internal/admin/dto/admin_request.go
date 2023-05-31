package admin

type AdminRequestBody struct {
	Email     string  `json:"email" binding:"email"`
	Name      string  `json:"name" binding:"required"`
	Password  *string `json:"password" binding:"required"`
	CreatedBy *int64  `json:"created_by"`
	UpdatedBy *int64  `json:"updated_by"`
}
