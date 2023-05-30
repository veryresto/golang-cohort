package main

import (
	"github.com/gin-gonic/gin"

	mysql "online-course/pkg/db/mysql"

	forgotPassword "online-course/internal/forgot_password/injector"
	oauth "online-course/internal/oauth/injector"
	register "online-course/internal/register/injector"
	verificationEmail "online-course/internal/verification_email/injector"
)

func main() {
	r := gin.Default()

	db := mysql.DB()

	register.InitializedService(db).Route(&r.RouterGroup)
	oauth.InitializedService(db).Route(&r.RouterGroup)
	verificationEmail.InitializedService(db).Route(&r.RouterGroup)
	forgotPassword.InitializedService(db).Route(&r.RouterGroup)

	r.Run()
}
