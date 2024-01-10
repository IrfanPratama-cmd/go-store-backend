package routes

import (
	"api/app/controller"
	"api/app/controller/account"
	"api/app/controller/brand"
	"api/app/controller/cart"
	"api/app/controller/category"
	"api/app/controller/product"

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
	// accountRoute.Use(middleware.TokenValidator())
	accountRoute.Post("/login", account.PostLoginAccount)
	accountRoute.Post("/register", account.PostRegisterAccount)
	accountRoute.Post("/verify-account", account.PostVerifyAccount)
	accountRoute.Post("/send-verify-account-code", account.PostSendVerifyAccountCode)

	// Brand
	brandRoute := api.Group("/brands")
	brandRoute.Use(middleware.JwtMiddleware)
	brandRoute.Get("/", brand.GetBrand)
	brandRoute.Post("/", brand.PostBrand)
	brandRoute.Get("/:id", brand.GetBrandID)
	brandRoute.Put("/:id", brand.PutBrand)
	brandRoute.Delete("/:id", brand.DeleteBrand)

	// Category
	categoryRoute := api.Group("/categories")
	categoryRoute.Use(middleware.JwtMiddleware)
	categoryRoute.Get("/", category.GetCategory)
	categoryRoute.Post("/", category.PostCategory)
	categoryRoute.Get("/:id", category.GetCategoryID)
	categoryRoute.Put("/:id", category.PutCategory)
	categoryRoute.Delete("/:id", category.DeleteCategory)

	// Product
	productRoute := api.Group("/products")
	productRoute.Use(middleware.JwtMiddleware)
	productRoute.Get("/", product.GetProduct)
	productRoute.Post("/", lib.HandleSingleFile, product.PostProduct)
	productRoute.Get("/:id", product.GetProductID)
	productRoute.Put("/:id", lib.HandleSingleFile, product.PutProduct)
	productRoute.Delete("/:id", product.DeleteProduct)

	// Cart
	cartRoute := api.Group("/carts")
	cartRoute.Use(middleware.JwtMiddleware)
	cartRoute.Get("/", cart.GetCart)
	cartRoute.Post("/", cart.PostCart)
	cartRoute.Put("/:id", cart.PutCart)
	cartRoute.Delete("/:id", cart.DeleteCart)

	// // Profile
	// profileRoute := api.Group("/profile")
	// profileRoute.Use(middleware.Oauth2Authentication)
	// profileRoute.Get("/", profile.GetProfile)
	// profileRoute.Put("/", profile.PutProfile)
}
