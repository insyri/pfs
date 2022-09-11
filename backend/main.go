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
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"golang.org/x/sys/unix"
)

var (
	Verbose = util.Verbose
	Error   = util.Error
	// Panic   = util.Panic
	Info    = util.Info
	Success = util.Success
)

var ISER = "{\"Error\": \"Internal Server Error\""

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

	// Connect to database
	pool, err := pgx.Connect(
		context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.Database.User, cfg.Database.Pass, "database", cfg.Database.Name),
	)

	if err != nil {
		Error.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer pool.Close(context.Background())
	Success.Println("Connected to database")

	// Create a new Fiber app
	app := fiber.New()
	api := app.Group("/api")

	app.Use(func(c *fiber.Ctx) error {
		return util.LogAPIConn(c)
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	api.Get("/", routes.Base)

	// Get /api/:hash
	api.Get("/:hash", routes.Hash)

	// Post /api/paste
	api.Post("/upload/paste", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		entry := &structures.PasteRequest{}

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

		Info.Print("Paste: ", ret)

		// Fetch amount of free space left on storage drive
		var stat unix.Statfs_t

		wd, err := os.Getwd()
		if err != nil || wd == "" {
			Info.Print(err)
			return c.Status(500).JSON(ISER)
		}

		if err := unix.Statfs(wd, &stat); err != nil {
			Info.Print(err)
			return c.Status(500).JSON(ISER)
		}

		// Available blocks * size per block = available space in bytes
		free_space := stat.Bavail * uint64(stat.Bsize)

		// Fetch database size
		query := fmt.Sprintf("select pg_database_size( '%s' );", cfg.Database.Name)

		row := pool.QueryRow(context.Background(), query)

		var used_space int64

		if err := row.Scan(&used_space); err != nil {
			Info.Print(err)
			return c.Status(500).JSON(ISER)
		}

		// Check if file will exceed remaining free space on storage drive
		if int64(free_space) < used_space+int64(len(entry.Text)) {
			// TODO: Make error log better with more information
			Error.Println("Not enough storage to save text file.")
			return c.SendStatus(507)
		}

		// Insert entry into database
		var ret_id pgtype.Int8
		if err := pool.QueryRow(context.Background(), fmt.Sprintf(
			"INSERT INTO entries (id, raw_text, expires_at) VALUES (DEFAULT, '%s', %d) RETURNING id",
			ret.Text, ret.Expires_At)).Scan(&ret_id); err != nil {
			Info.Print(err)

			return c.Status(500).JSON(ISER)
		}

		// Return the entry
		return c.JSON(ret_id)
	})

	Error.Fatal(app.Listen(":8080"))
}
