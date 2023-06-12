package main

import (
	"github.com/gin-gonic/gin"

	mysql "online-course/pkg/db/mysql"

	admin "online-course/internal/admin/injector"
	cart "online-course/internal/cart/injector"
	classRoom "online-course/internal/class_room/injector"
	dashboard "online-course/internal/dashboard/injector"
	discount "online-course/internal/discount/injector"
	forgotPassword "online-course/internal/forgot_password/injector"
	oauth "online-course/internal/oauth/injector"
	order "online-course/internal/order/injector"
	product "online-course/internal/product/injector"
	productCategory "online-course/internal/product_category/injector"
	profile "online-course/internal/profile/injector"
	register "online-course/internal/register/injector"
	user "online-course/internal/user/injector"
	verificationEmail "online-course/internal/verification_email/injector"
	webhook "online-course/internal/webhook/injector"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

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
	classRoom.InitializedService(db).Route(&r.RouterGroup)
	webhook.InitializedService(db).Route(&r.RouterGroup)
	user.InitializedService(db).Route(&r.RouterGroup)
	profile.InitializedService(db).Route(&r.RouterGroup)
	dashboard.InitializedService(db).Route(&r.RouterGroup)

	r.Run()
}
