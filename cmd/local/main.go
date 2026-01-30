package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pos-backend/internal/infrastructure"
	"pos-backend/internal/infrastructure/database"
	"pos-backend/internal/utility"
	"strings"
	"time"

	"pos-backend/internal/infrastructure/logs"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	initTimeZone()
	db, err := database.InitDatabase()
	if err != nil {
		logs.Error("⚠️ database unavailable")
		db = nil
	}

	conn := infrastructure.Connections{
		DB: db,
	}

	app := fiber.New(fiber.Config{
		BodyLimit:      225 * 1024 * 1024,
		ReadBufferSize: 12288,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
		},
	})

	app.Use(recover.New())

	if os.Getenv("ENV") == "uat" || os.Getenv("ENV") == "dev" {
		app.Use(logger.New(logger.Config{
			Format:     "${blue}${time} ${yellow}${status} - ${red}${latency} ${cyan}${method} ${path} ${green} ${ip} ${ua} ${reset}\n",
			TimeFormat: "02-Jan-2006 15:04:05",
			TimeZone:   "Asia/Bangkok",
			Output:     os.Stdout,
		}))
	}

	app.Get("/manage", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "up",
			"service": "emergency-medicine-manage",
			"version": getVersionNumber(),
			"env":     os.Getenv("ENV"),
			"db":      conn.DB != nil,
		})
	})

	// router.SetUpRouter(app, conn)

	if err := app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port"))); err != nil {
		logs.Error(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initConfig() {
	viper.AddConfigPath("../../config")
	viper.SetConfigType("yaml")

	switch os.Getenv("ENV") {
	case "":
		fmt.Println("uat")
		viper.SetConfigName("config_uat")
	case "uat-local":
		fmt.Println("uat")
		viper.SetConfigName("config_uat")
	case "prd-local":
		fmt.Println("prd")
		viper.SetConfigName("config_prd")
	default:
		fmt.Println("default")
		viper.SetConfigName("config")
	}

	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func getVersionNumber() string {
	version := "0.0.1"
	inFile, err := os.Open("./Makefile")
	if err != nil {
		log.Println(err.Error() + `: ` + err.Error())
		return version
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		lineVersion := scanner.Text()
		if strings.TrimSpace(lineVersion) != "" {
			listFirstLine := strings.Split(lineVersion, " ")
			version = listFirstLine[len(listFirstLine)-1]
			break
		} else {
			break
		}
	}

	return version
}
