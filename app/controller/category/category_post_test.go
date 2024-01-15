package category

import (
	"api/app/config"
	"api/app/lib"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPostCategory(t *testing.T) {
	db := services.DBConnectTest()
	lib.LoadEnvironment(config.Environment)

	app := fiber.New()
	// app.Use(middleware.JwtMiddleware)
	app.Post("/categories", PostCategory)

	uri := "/categories"

	payload := `{
		"category_name": "String",
		"category_code": "str"
	}`

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	response, body, err := lib.PostTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")

	// test invalid json format
	response, _, err = lib.PostTest(app, uri, headers, "invalid json format")
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 400, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
