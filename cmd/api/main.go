package main

import (
	"github.com/gin-gonic/gin"

	mysql "online-course.faerul.com/pkg/db/mysql"

	admin "online-course.faerul.com/internal/admin/injector"
	cart "online-course.faerul.com/internal/cart/injector"
	discount "online-course.faerul.com/internal/discount/injector"
	forgotPassword "online-course.faerul.com/internal/forgot_password/injector"
	oauth "online-course.faerul.com/internal/oauth/injector"
	order "online-course.faerul.com/internal/order/injector"
	product "online-course.faerul.com/internal/product/injector"
	productCategory "online-course.faerul.com/internal/product_category/injector"
	register "online-course.faerul.com/internal/register/injector"
	verificationEmail "online-course.faerul.com/internal/verification_email/injector"
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
	cart.InitializedService(db).Route(&r.RouterGroup)
	order.InitializedService(db).Route(&r.RouterGroup)

	r.Run()
}
