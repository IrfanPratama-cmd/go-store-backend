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

func TestDeleteCategory(t *testing.T) {
	db := services.DBConnectTest()
	lib.LoadEnvironment(config.Environment)

	app := fiber.New()
	// app.Use(middleware.JwtMiddleware)
	app.Delete("/categories/:id", DeleteCategory)

	initial := model.Category{
		CategoryAPI: model.CategoryAPI{
			CategoryName: lib.Strptr("Kendaraan"),
			CategoryCode: lib.Strptr("String"),
		},
	}

	db.Create(&initial)

	uri := "/categories/" + initial.ID.String()
	response, _, err := lib.DeleteTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")

	// test delete with non existing id
	uri = "/catalog/categories/non-existing-id"
	response, _, err = lib.DeleteTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
