package brand

import (
	"api/app/config"
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPutBrand(t *testing.T) {
	db := services.DBConnectTest()

	lib.LoadEnvironment(config.Environment)

	app := fiber.New()
	// app.Use(middleware.JwtMiddleware)
	app.Put("/brands/:id", PutBrand)

	initial := model.Brand{
		BrandAPI: model.BrandAPI{
			BrandName: lib.Strptr("Addidas"),
			BrandCode: lib.Strptr("ADD"),
		},
	}

	initial2 := model.Brand{
		BrandAPI: model.BrandAPI{
			BrandName: lib.Strptr("Nike"),
			BrandCode: lib.Strptr("NKK"),
		},
	}

	db.Create(&initial)
	db.Create(&initial2)

	uri := "/brands/" + initial.ID.String()

	payload := `{
		"brand_name": "Nike",
		"brand_code": "NKK"
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
	uri = "/brands/non-existing-id"
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
