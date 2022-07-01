package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"golang.org/x/sys/unix"
)

func main() {
	conf, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// use it later
	fmt.Println(conf)

	// Get environment variables
	env := godotenv.Load("./database.env")
	if env != nil {
		log.Fatalf("Error loading .env file")
	}

	// Extract environment variables
	var (
		username = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		hostname = "database"
		dbname   = os.Getenv("POSTGRES_DB")
	)

	// Connect to database
	connection, err := pgx.Connect(
		context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, hostname, dbname),
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer connection.Close(context.Background())
	fmt.Println("Connected to database")

	// Create a new Fiber app
	app := fiber.New()
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ok its up")
	})

	// Get /api/:hash
	api.Get("/:hash", func(c *fiber.Ctx) error {
		hash := c.Params("hash")

		return c.SendString(fmt.Sprintf("Hello, %s!", hash))
	})

	// Post /api/paste
	api.Post("/upload/paste", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		entry := new(PasteRequest)
		var _ = new(PasteResponse)

		if err := c.BodyParser(entry); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Validate values of entry
		errors := ValidatePaste(entry)
		if errors != nil {
			return c.Status(400).JSON(errors)
		}

		// TODO: check if hash already exists (pass connection to FillPaste)

		ret := FillPaste(entry)

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
		query := fmt.Sprintf("select pg_database_size( '%s' );", dbname)

		row := connection.QueryRow(context.Background(), query)

		var used_space int64

		if err := row.Scan(&used_space); err != nil {
			log.Print(err)
		}

		// Check if file will exceed remaining free space on storage drive
		if int64(free_space) < used_space+int64(len(entry.Text)) {
			// TODO: Make error log better with more information
			log.Println("Not enough storage to save text file.")
			return c.SendStatus(507)
		}

		// Insert entry into database

		// Return the entry
		return c.JSON(ret)
	})

	log.Fatal(app.Listen(":8080"))
}
