package main

import (
	"fmt"

	"Ang2Tea/medods-test/common"
	"Ang2Tea/medods-test/internal/adapter"
	"Ang2Tea/medods-test/internal/adapter/db"
	"Ang2Tea/medods-test/internal/adapter/handler"
	"Ang2Tea/medods-test/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

const (
	configFile = "config.yml"
	envFile    = ".env"
)

func main() {
	var configDirectory string
	common.LookupEnv(&configDirectory, common.CONFIG_DIRECTORY, "config")

	config := common.GetConfig(fmt.Sprintf("%s/%s", configDirectory, configFile))

	if config.Mode == common.DebugMode {
		err := godotenv.Load(fmt.Sprintf("%s/%s", configDirectory, envFile))
		common.PanicIfErr(err)
	}

	var secretKey string
	common.LookupEnv(&secretKey, common.JWT_SECRET_KEY)

	gormDB := common.GetDB()

	err := db.Migration(gormDB)
	common.PanicIfErr(err)

	userStorage := db.NewUserStorage(gormDB)
	tokenManager := adapter.NewJWTTokenManager(secretKey)
	authEvents := adapter.NewMockAuthEvent()

	authUsecase := usecase.NewAuthUsecase(userStorage, tokenManager, authEvents)

	handlers := handler.NewFiberRoutes(authUsecase)

	app := fiber.New(fiber.Config{
		AppName: config.AppName,
	})

	handlers.RegisterRoutes(app)

	ipAddress := fmt.Sprintf("%s:%d", config.WebServer.Host, config.WebServer.Port)
	common.PanicIfErr(app.Listen(ipAddress))
}
