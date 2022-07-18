package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/insyri/pfs/backend/fill"
	"github.com/insyri/pfs/backend/routes"
	"github.com/insyri/pfs/backend/structures"
	"github.com/insyri/pfs/backend/util"
	"github.com/insyri/pfs/backend/validation"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"golang.org/x/sys/unix"
)

var (
	Verbose = util.Verbose
	Error   = util.Error
	Panic   = util.Panic
	Info    = util.Info
	Success = util.Success
)

func LoadConfig(f string) (*structures.Config, error) {
	viper.SetConfigName(f)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg structures.Config

	if err := viper.UnmarshalExact(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func main() {
	cfg, err := LoadConfig("pfs.example.toml")
	if err != nil {
		Error.Fatal(err)
	}

	Verbose.Printf("Config out: %+v", cfg)

	// Get environment variables
	// env := godotenv.Load("./database.env")
	// if env != nil {
	// 	Error.Fatal("Error loading .env file")
	// }

	// // Extract environment variables
	// var (
	// 	username = os.Getenv("POSTGRES_USER")
	// 	password = os.Getenv("POSTGRES_PASSWORD")
	// 	hostname = "database"
	// 	dbname   = os.Getenv("POSTGRES_DB")
	// )

	// Connect to database
	connection, err := pgx.Connect(
		context.Background(),
		// fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, hostname, dbname),
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.Database.User, cfg.Database.Pass, "database", cfg.Database.Name),
	)

	if err != nil {
		Error.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer connection.Close(context.Background())
	Success.Println("Connected to database")

	// Create a new Fiber app
	app := fiber.New()
	api := app.Group("/api")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	api.Get("/", routes.Base)

	// Get /api/:hash
	api.Get("/:hash", routes.Hash)

	// Post /api/paste
	api.Post("/upload/paste", func(c *fiber.Ctx) error {
		util.LogAPIConn(c)
		c.Accepts("application/json")
		entry := new(structures.PasteRequest)
		var _ = new(structures.PasteResponse)

		if err := c.BodyParser(entry); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Validate values of entry
		errors := validation.ValidatePaste(entry)
		if errors != nil {
			return c.Status(400).JSON(errors)
		}

		// TODO: check if hash already exists (pass connection to FillPaste)

		ret := fill.FillPaste(entry)
		fmt.Println(ret.Text)

		// Fetch amount of free space left on storage drive
		var stat unix.Statfs_t

		wd, err := os.Getwd()
		if err != nil || wd == "" {
			return err
		}

		unix.Statfs(wd, &stat)

		// Available blocks * size per block = available space in bytes
		free_space := stat.Bavail * uint64(stat.Bsize)

		// Fetch database size
		query := fmt.Sprintf("select pg_database_size( '%s' );", cfg.Database.Name)

		row := connection.QueryRow(context.Background(), query)

		var used_space int64

		if err := row.Scan(&used_space); err != nil {
			Error.Fatal(err)
		}

		// Check if file will exceed remaining free space on storage drive
		if int64(free_space) < used_space+int64(len(entry.Text)) {
			// TODO: Make error log better with more information
			Error.Println("Not enough storage to save text file.")
			return c.SendStatus(507)
		}

		// Insert entry into database
		connection.Exec(context.Background(), fmt.Sprintf("INSERT INTO entries (id, raw_text, created_at) VALUES (DEFAULT, '%s', %d)", ret.Text, ret.Expires_At))

		// Return the entry
		return c.JSON(ret)
	})

	Error.Fatal(app.Listen(":8080"))
}
