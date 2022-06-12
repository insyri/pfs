package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	println("API starting...")

	// Get environment variables
	env := godotenv.Load("./database.env")
	if env != nil {
		log.Fatalf("Error loading .env file")
	}

	// Extract environment variables
	var (
		username = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		hostname = "database" // "localhost"
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

	// Structs
	// type File struct {
	// 	FileName     string `json:"file_name"`
	// 	FileType     string `json:"file_type"`
	// 	Size         int    `json:"size"`
	// 	Expires      string `json:"expires_at"`
	// 	AutoDelete   bool   `json:"auto_delete"`
	// 	Downloads    int    `json:"downloads"`
	// 	MaxDownloads int    `json:"max_downloads"`
	// }

	// type Files struct {
	// 	Files []File `json:"files"`
	// }

	app := fiber.New()
	// e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(c.Path())
	})
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	app.Get("/:file", func(c *fiber.Ctx) error {
		return c.SendString(c.Path())
	})
	// e.GET("/api/:file", func(c echo.Context) error {
	// 	file := c.Param("file")
	// 	return c.String(http.StatusOK, file)
	// }).Name = "getFileContent"

	// e.GET("/api/:file/info", func(c echo.Context) error {
	// 	file := c.Param("file")
	// 	return c.String(http.StatusOK, file)
	// }).Name = "getFileInfo"

	// e.POST("/api/:file", func(c echo.Context) error {
	// 	file := c.Param("file")

	// 	f := new(File)
	// 	if err := c.Bind(f); err != nil {
	// 		return err
	// 	}

	// 	sqlStatement := "INSERT INTO " + os.Getenv("DB_NAME") + " (file_name, file_type, size, expires_at, auto_delete, downloads, max_downloads) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	// 	res, err := conn.Query(sqlStatement, file, f.FileType, f.Size, f.Expires, f.AutoDelete, f.Downloads, f.MaxDownloads)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(res)
	// 		return c.JSON(http.StatusCreated, f)
	// 	}
	// 	return c.String(http.StatusOK, "ok")
	// }).Name = "postFile"

	// e.DELETE("/api/:file", func(c echo.Context) error {
	// 	file := c.Param("file")
	// 	return c.String(http.StatusOK, file)
	// }).Name = "deleteFile"

	// e.Logger.Fatal(e.Start(":8080"))
}
