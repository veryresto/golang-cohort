package profile

type ProfileRequestLogoutBody struct {
	Authorization string `header:"authorization" binding:"required"`
}
