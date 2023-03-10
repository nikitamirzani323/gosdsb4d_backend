package controllers

import (
	"log"
	"time"

	"bitbucket.org/isbtotogroup/sdsb4d-backend/entities"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/helpers"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/models"
	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const Field_vietnamnight_redis = "LISTVIETNAMNIGHT_SDSB4D"

func Vietnamnighthome(c *fiber.Ctx) error {
	var obj entities.Model_vietnamenight
	var arraobj []entities.Model_vietnamenight
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Field_vietnamnight_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		vietnamnight_id, _ := jsonparser.GetInt(value, "vietnamnight_id")
		vietnamnight_date, _ := jsonparser.GetString(value, "vietnamnight_date")
		vietnamnight_prize1_1300, _ := jsonparser.GetString(value, "vietnamnight_prize1_1300")
		vietnamnight_prize1_1700, _ := jsonparser.GetString(value, "vietnamnight_prize1_1700")
		vietnamnight_prize1_2000, _ := jsonparser.GetString(value, "vietnamnight_prize1_2000")
		vietnamnight_prize1_2200, _ := jsonparser.GetString(value, "vietnamnight_prize1_2200")
		vietnamnight_create, _ := jsonparser.GetString(value, "vietnamnight_create")
		vietnamnight_update, _ := jsonparser.GetString(value, "vietnamnight_update")

		obj.Vietnamnight_id = int(vietnamnight_id)
		obj.Vietnamnight_date = vietnamnight_date
		obj.Vietnamnight_prize1_1300 = vietnamnight_prize1_1300
		obj.Vietnamnight_prize1_1700 = vietnamnight_prize1_1700
		obj.Vietnamnight_prize1_2000 = vietnamnight_prize1_2000
		obj.Vietnamnight_prize1_2200 = vietnamnight_prize1_2200
		obj.Vietnamnight_create = vietnamnight_create
		obj.Vietnamnight_update = vietnamnight_update

		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_vietnamenightHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Field_vietnamnight_redis, result, 60*time.Minute)
		log.Println("VIETNAMNIGHT MYSQL")
		return c.JSON(result)
	} else {
		log.Println("VIETNAMNIGHT CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func VietnamnightSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_vietnamenightsave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	result, err := models.Save_vietnamenightHome(
		client_admin,
		client.Tanggal, client.Sdata, client.Idrecord)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_vietnamnight()
	return c.JSON(result)
}
func VietnamnightGeneratorSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_vietnamenightprizesave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	field := ""
	switch client.Tipe {
	case "prize1_1300":
		field = "prize_1300"
	case "prize1_1700":
		field = "prize_1700"
	case "prize1_2000":
		field = "prize_2000"
	case "prize1_2200":
		field = "prize_2200"
	}

	result, err := models.Save_vietnamenightGenerator(
		client_admin,
		field, client.Prize, client.Sdata, client.Idrecord)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_vietnamnight()
	return c.JSON(result)
}
func VietnamnightGeneratorNumber(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	result, err := models.Save_Generatorvietnamenight(client_admin)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_vietnamnight()
	return c.JSON(result)
}
func _deleteredis_vietnamnight() {
	val_master := helpers.DeleteRedis(Field_vietnamnight_redis)
	log.Printf("Redis Delete BACKEND VIETNAMENIGHT : %d", val_master)

	field_redis_api := "SDSB4D_LISTSDSBNIGHT_API"
	val_api := helpers.DeleteRedis(field_redis_api)
	log.Printf("Redis Delete API LISTSDSBDAY_SDSB4D : %d", val_api)
}
