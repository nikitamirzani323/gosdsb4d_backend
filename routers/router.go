package routers

import (
	"bitbucket.org/isbtotogroup/sdsb4d-backend/controllers"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "sveltemdb/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("/dashboard", monitor.New())

	app.Post("/api/login", controllers.CheckLogin)
	app.Post("/api/valid", middleware.JWTProtected(), controllers.Home)
	app.Post("/api/alladmin", middleware.JWTProtected(), controllers.Adminhome)
	app.Post("/api/detailadmin", middleware.JWTProtected(), controllers.AdminDetail)
	app.Post("/api/saveadmin", middleware.JWTProtected(), controllers.AdminSave)

	app.Post("/api/alladminrule", middleware.JWTProtected(), controllers.Adminrulehome)
	app.Post("/api/saveadminrule", middleware.JWTProtected(), controllers.AdminruleSave)

	app.Post("/api/sdsbday", middleware.JWTProtected(), controllers.Sdsbdayhome)
	app.Post("/api/savesdsbday", middleware.JWTProtected(), controllers.SdsbdaySave)
	app.Post("/api/savegeneratorsdsbday", middleware.JWTProtected(), controllers.SdsbdayGeneratorSave)
	app.Post("/api/generatornumber", middleware.JWTProtected(), controllers.SdsbdayGeneratorNumber)

	app.Post("/api/sdsbnight", middleware.JWTProtected(), controllers.Sdsbnighthome)
	app.Post("/api/savesdsbnight", middleware.JWTProtected(), controllers.SdsbnightSave)
	app.Post("/api/savegeneratorsdsbnight", middleware.JWTProtected(), controllers.SdsbnightGeneratorSave)
	app.Post("/api/generatornumbernight", middleware.JWTProtected(), controllers.SdsbnightGeneratorNumber)

	app.Post("/api/vietnamnight", middleware.JWTProtected(), controllers.Vietnamnighthome)
	app.Post("/api/savevietnamnight", middleware.JWTProtected(), controllers.VietnamnightSave)
	app.Post("/api/savegeneratorvietnamnight", middleware.JWTProtected(), controllers.VietnamnightGeneratorSave)
	app.Post("/api/generatornumbervietnamnight", middleware.JWTProtected(), controllers.VietnamnightGeneratorNumber)

	app.Post("/api/initprediksi", middleware.JWTProtected(), controllers.TokenPrediksi)
	app.Post("/api/listpasaran", middleware.JWTProtected(), controllers.ListPasaran)
	app.Post("/api/prediksi", middleware.JWTProtected(), controllers.PrediksiWajib)

	return app
}
