package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"golang.org/x/sys/unix"
)

var (
	Verbose = log.New(os.Stderr, color.New(color.FgYellow).Sprint("Verbose")+" ", log.Lmsgprefix|log.Lshortfile)
	Info    = log.New(os.Stderr, color.New(color.FgBlue).Sprint("Info")+" ", log.Lmsgprefix)
	Success = log.New(os.Stderr, color.New(color.FgGreen).Sprint("Success")+" ", log.Lmsgprefix)
	Panic   = log.New(os.Stderr, color.New(color.BgHiRed|color.Bold).Sprint("Panic")+" ", log.Lmsgprefix|log.Llongfile)
	Error   = log.New(os.Stderr, color.New(color.FgRed).Sprint("Error")+" ", log.Lmsgprefix|log.Llongfile)
)

func main() {
	conf, err := LoadConfig()
	if err != nil {
		Error.Fatal(err)
	}

	Info.Printf("Database Name:     \"%s\"", conf.Db_Name)
	Info.Printf("Database Password: \"%s\"", conf.Db_Pass)
	Info.Printf("Database User:     \"%s\"", conf.Db_User)
	Info.Printf("Expiry:            \"%d\"", conf.Expiry)
	Info.Printf("Max Storage:       \"%d\"", conf.Max_Storage)
	Info.Printf("Save Directory:    \"%s\"", conf.Save_Dir)

	// Get environment variables
	env := godotenv.Load("./database.env")
	if env != nil {
		Error.Fatal("Error loading .env file")
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
		query := fmt.Sprintf("select pg_database_size( '%s' );", dbname)

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
