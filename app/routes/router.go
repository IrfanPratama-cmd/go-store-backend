package routes

import (
	"api/app/controller"
	"api/app/controller/account"
	// "api/app/controller/profile"
	"api/app/lib"
	"api/app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

// Handle all request to route to controller
func Handle(app *fiber.App) {
	app.Use(cors.New())

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			lib.PrintStackTrace(e)
		},
	}))

	api := app.Group(viper.GetString("ENDPOINT"))

	api.Static("/swagger", "docs/swagger.json")
	api.Get("/", controller.GetAPIIndex)
	api.Get("/info.json", controller.GetAPIInfo)
	api.Post("/logs", controller.PostLogs)

	// Account
	accountRoute := api.Group("/account")
	accountRoute.Use(middleware.TokenValidator())
	accountRoute.Post("/login", account.PostLoginAccount)
	accountRoute.Post("/register", account.PostRegisterAccount)
	accountRoute.Post("/verify-account", account.PostVerifyAccount)
	accountRoute.Post("/send-verify-account-code", account.PostSendVerifyAccountCode)

	// // Profile
	// profileRoute := api.Group("/profile")
	// profileRoute.Use(middleware.Oauth2Authentication)
	// profileRoute.Get("/", profile.GetProfile)
	// profileRoute.Put("/", profile.PutProfile)
}
