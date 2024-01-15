package category

import (
	"api/app/config"
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPutCategory(t *testing.T) {
	db := services.DBConnectTest()

	lib.LoadEnvironment(config.Environment)

	app := fiber.New()
	// app.Use(middleware.JwtMiddleware)
	app.Put("/categories/:id", PutCategory)

	initial := model.Category{
		CategoryAPI: model.CategoryAPI{
			CategoryName: lib.Strptr("Kendaraan"),
			CategoryCode: lib.Strptr("String"),
		},
	}

	initial2 := model.Category{
		CategoryAPI: model.CategoryAPI{
			CategoryName: lib.Strptr("Samsung"),
			CategoryCode: lib.Strptr("SM"),
		},
	}

	db.Create(&initial)
	db.Create(&initial2)

	uri := "/categories/" + initial.ID.String()

	payload := `{
		"category_name": "String",
		"category_code": "str"
	}`

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	response, body, err := lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")

	// test invalid json body
	response, _, err = lib.PutTest(app, uri, headers, "invalid json format")
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 400, response.StatusCode, "getting response code")

	// test update with non existing id
	uri = "/categories/non-existing-id"
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
