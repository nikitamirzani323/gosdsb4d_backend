package controllers

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type responsedefault struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}
type responsetoken struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}
type responseerror struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const PATH_API string = "http://139.59.105.180:7073/"

func TokenPrediksi(c *fiber.Ctx) error {
	hostname := c.Hostname()
	log.Println("Hostname: ", hostname)

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsetoken{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"password":        "asdQWE123!@#",
		}).
		Post(PATH_API + "api/loginother")
	if err != nil {
		fmt.Println(err.Error())
	}
	result := resp.Result().(*responsetoken)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": result.Status,
			"token":  result.Token,
			"time":   time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func ListPasaran(c *fiber.Ctx) error {
	type payload_listpasaran struct {
		Master        string `json:"master"`
		Tokenprediksi string `json:"token"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_listpasaran)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	log.Println(token)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(client.Tokenprediksi).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"master":          client.Master,
		}).
		Post(PATH_API + "api/listpasaranwajib")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func PrediksiWajib(c *fiber.Ctx) error {
	type payload_prediksi struct {
		Master         string `json:"master"`
		Tokenprediksi  string `json:"token"`
		Idpasarantogel string `json:"idpasarantogel" `
		Nomorprediksi  string `json:"nomorprediksi" "`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_prediksi)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	log.Println(token)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(client.Tokenprediksi).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"master":          client.Master,
			"idpasarantogel":  client.Idpasarantogel,
			"nomorprediksi":   client.Nomorprediksi,
		}).
		Post(PATH_API + "api/prediksiwajib")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
