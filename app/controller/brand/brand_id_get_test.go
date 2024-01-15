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

func TestGetBrandID(t *testing.T) {
	db := services.DBConnectTest()
	lib.LoadEnvironment(config.Environment)

	app := fiber.New()
	// app.Use(middleware.JwtMiddleware)
	app.Get("/brands/:id", GetBrandID)

	initial := model.Brand{
		BrandAPI: model.BrandAPI{
			BrandName: lib.Strptr("Addidas"),
			BrandCode: lib.Strptr("ADD"),
		},
	}
	db.Create(&initial)

	uri := "/brands/" + initial.ID.String()
	response, body, err := lib.GetTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")
	utils.AssertEqual(t, initial.ID.String(), body["id"], "getting response body")

	// test get non existing id
	uri = "/brands/non-existing-id"
	response, _, err = lib.GetTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
