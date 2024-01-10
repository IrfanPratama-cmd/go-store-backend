package account

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"sync"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2"
)

// PostRegister godoc
// @Summary Registration
// @Description Registration
// @Param data body model.RegistrationAPI true "Payload"
// @Security TokenKey
// @Success 200 {object} lib.Response "registered"
// @Failure 401 {object} lib.Response "Unauthorized"
// @Failure 400 {object} lib.Response "Bad Request"
// @Failure 204 {object} lib.Response "No Content"
// @Failure 404 {object} lib.Response "Not Found"
// @Failure 409 {object} lib.Response "Conflict"
// @Failure 500 {object} lib.Response "Internal Server Error"
// @Router /account/register [post]
// @Tags Account
func PostRegisterAccount(c *fiber.Ctx) error {
	api := &model.RegistrationAPI{}
	if err := lib.BodyParser(c, api); err != nil {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB.WithContext(c.UserContext())

	// Generate unique verification code
	verificationCode := lib.RandomNumber(4)

	// Set verification expiration time
	verificationExpiration := time.Now().Add(5 * time.Minute)

	var email model.User
	checkEmail := db.Model(&model.User{}).Where(`email = ?`, api.Email).First(&email)

	if checkEmail.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email is already used",
		})
	}

	// Create user
	user := &model.User{
		UserAPI: model.UserAPI{
			Fullname:               api.Fullname,
			IsOwner:                lib.Boolptr(true),
			VerificationCode:       lib.Strptr(verificationCode),
			VerificationExpiration: lib.DateTimeptr(strfmt.DateTime(verificationExpiration)),
			Email:                  api.Email,
			Password:               lib.Strptr(lib.RandomChars(10)),
		},
	}

	// Define model
	var contact model.Contact
	contact.ContactName = api.Fullname
	contact.Email = api.Email
	db.Create(&contact)

	var passwordHash string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		passwordHash, _ = lib.HashPassword(*api.Password)
	}()

	wg.Wait()
	user.Password = lib.Strptr(passwordHash)
	if err := db.Save(user).Error; err != nil {
		return lib.ErrorInternal(c, err.Error())
	}

	return lib.OK(c)
}
