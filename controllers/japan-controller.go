package controllers

import (
	"fmt"
	"time"

	"bitbucket.org/isbtotogroup/sdsb4d-backend/entities"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/helpers"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/models"
	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const Field_japan = "LISTJAPAN_SDSB4D"

func Japanhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_japan)
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
	Fieldredis := Field_japan + "_day"
	if client.Japan_tipe == "night" {
		Fieldredis = Field_japan + "_night"
	}

	var obj entities.Model_japan
	var arraobj []entities.Model_japan
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldredis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		japan_id, _ := jsonparser.GetInt(value, "japan_id")
		japan_date, _ := jsonparser.GetString(value, "japan_date")
		japan_prize1, _ := jsonparser.GetString(value, "japan_prize1")
		japan_prize2, _ := jsonparser.GetString(value, "japan_prize2")
		japan_prize3, _ := jsonparser.GetString(value, "japan_prize3")
		japan_create, _ := jsonparser.GetString(value, "japan_create")
		japan_update, _ := jsonparser.GetString(value, "japan_update")

		obj.Japan_id = int(japan_id)
		obj.Japan_date = japan_date
		obj.Japan_prize1 = japan_prize1
		obj.Japan_prize2 = japan_prize2
		obj.Japan_prize3 = japan_prize3
		obj.Japan_create = japan_create
		obj.Japan_update = japan_update

		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_japanHome(client.Japan_tipe)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldredis, result, 60*time.Minute)
		fmt.Println("JAPAN MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("JAPAN CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func JapanSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_japansave)
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

	result, err := models.Save_japanHome(
		client_admin,
		client.Tanggal, client.Tipe, client.Sdata, client.Idrecord)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_japan(client.Tipe)
	return c.JSON(result)
}
func JapanGeneratorSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_japanprizesave)
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
	case "prize1":
		if client.Tipejapan == "day" {
			field = "prize1_japanday"
		} else {
			field = "prize1_japannight"
		}
	case "prize2":
		if client.Tipejapan == "day" {
			field = "prize2_japanday"
		} else {
			field = "prize2_japannight"
		}

	case "prize3":
		if client.Tipejapan == "day" {
			field = "prize3_japanday"
		} else {
			field = "prize3_japannight"
		}
	}
	//admin, field, prize, tipe, sData string, idrecord int
	result, err := models.Save_japanGenerator(
		client_admin,
		field, client.Prize, client.Tipejapan, client.Sdata, client.Idrecord)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_japan(client.Tipejapan)
	return c.JSON(result)
}
func JapanGeneratorNumber(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_japan)
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

	result, err := models.Save_Generatorjapan(client_admin, client.Japan_tipe)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_japan(client.Japan_tipe)
	return c.JSON(result)
}
func _deleteredis_japan(tipe string) {
	txt := "_day"
	if tipe == "night" {
		txt = "_night"
	}
	val_master := helpers.DeleteRedis(Field_japan + txt)
	fmt.Printf("Redis Delete BACKEND JAPAN : %d", val_master)

	field_redis_api := "SDSB4D_LISTSDSBNIGHT_API"
	val_api := helpers.DeleteRedis(field_redis_api)
	fmt.Printf("Redis Delete API LISTSDSBDAY_SDSB4D : %d", val_api)
}
