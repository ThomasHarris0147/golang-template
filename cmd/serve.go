/*
Copyright © 2025 NAME HERE thomasharris0147@gmail.com
*/
package cmd

import (
	"log"
	"os"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"harris.com/api"
	"harris.com/config"
	"harris.com/database"
	"harris.com/server"
	"harris.com/services"
	"harris.com/utils"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func startServer() {
	app := fx.New(
		fx.Provide(
			database.NewDatabase,
			initFiberApp,
			services.NewUserService,
		),
		fx.Invoke(
			utils.InitLogger,
			loadEnv,
			registerGeneratedHandler,
			startFiberServer,
		),
	)
	defer func(Logger *zap.Logger) {
		err := Logger.Sync()
		if err != nil {
			panic("Failed to sync logger")
		}
	}(utils.Logger)
	app.Run()
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func loadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return utils.LogError(err)
	}

	if value, exists := os.LookupEnv("DB_DSN"); !exists || value == "" {
		if err := os.Setenv("DB_DSN", "local.db"); err != nil {
			return utils.LogError(err)
		}
	}

	return nil
}

func registerGeneratedHandler(
	app *fiber.App,
	services server.Server,
) {
	api.RegisterHandlers(app, services)
}

func startFiberServer(app *fiber.App) {
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start Fiber server: %v", err)
	}
}

func initFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		// Increase read buffer size to prevent "431 Request Header Fields Too Large" errors
		ReadBufferSize: 1 * 1024 * 1024, // 1MB
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.Config("ALLOW_ORIGINS"),
		AllowCredentials: true,
	}))

	registerBodyDecoder()

	// Define health endpoint before validation middleware to exclude it from validation
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	return app
}

func registerBodyDecoder() {
	contentTypes := []string{
		"image/jpeg", "image/jpg", "image/png", "image/gif", "image/webp",

		"application/pdf",
		"application/msword",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document",

		"text/plain", "text/csv",

		"application/zip",

		"application/octet-stream",
	}

	for _, ct := range contentTypes {
		openapi3filter.RegisterBodyDecoder(ct, openapi3filter.FileBodyDecoder)
	}
}
