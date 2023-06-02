package main

import (
	"github.com/gin-gonic/gin"

	mysql "online-course/pkg/db/mysql"

	admin "online-course/internal/admin/injector"
	discount "online-course/internal/discount/injector"
	forgotPassword "online-course/internal/forgot_password/injector"
	oauth "online-course/internal/oauth/injector"
	product "online-course/internal/product/injector"
	productCategory "online-course/internal/product_category/injector"
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
	admin.InitializedService(db).Route(&r.RouterGroup)
	productCategory.InitializedService(db).Route(&r.RouterGroup)
	product.InitializedService(db).Route(&r.RouterGroup)
	discount.InitializedService(db).Route(&r.RouterGroup)

	r.Run()
}
