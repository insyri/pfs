package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	var (
		username = os.Getenv("POSTGRES_USERNAME")
		password = os.Getenv("POSTGRES_PASSWORD")
		hostname = "database"
		dbname   = os.Getenv("POSTGRES_NAME")
	)

	dsn := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", hostname, username, password, dbname)

	// Get environment variables
	env := godotenv.Load("./database.env")
	if env != nil {
		log.Fatalf("Error loading .env file")
	}

	time.Sleep(5 * time.Second)

	// Connect to database
	var err error
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	type File struct {
		FileName     string `json:"file_name"`
		FileType     string `json:"file_type"`
		Size         int    `json:"size"`
		Expires      string `json:"expires_at"`
		AutoDelete   bool   `json:"auto_delete"`
		Downloads    int    `json:"downloads"`
		MaxDownloads int    `json:"max_downloads"`
	}

	// type Files struct {
	// 	Files []File `json:"files"`
	// }

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api/:file", func(c echo.Context) error {
		file := c.Param("file")
		return c.String(http.StatusOK, file)
	}).Name = "getFileContent"

	e.GET("/api/:file/info", func(c echo.Context) error {
		file := c.Param("file")
		return c.String(http.StatusOK, file)
	}).Name = "getFileInfo"

	e.POST("/api/:file", func(c echo.Context) error {
		file := c.Param("file")

		f := new(File)
		if err := c.Bind(f); err != nil {
			return err
		}

		sqlStatement := "INSERT INTO " + os.Getenv("DB_NAME") + " (file_name, file_type, size, expires_at, auto_delete, downloads, max_downloads) VALUES ($1, $2, $3, $4, $5, $6, $7)"
		res, err := db.DB().Query(sqlStatement, file, f.FileType, f.Size, f.Expires, f.AutoDelete, f.Downloads, f.MaxDownloads)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, f)
		}
		return c.String(http.StatusOK, "ok")
	}).Name = "postFile"

	e.DELETE("/api/:file", func(c echo.Context) error {
		file := c.Param("file")
		return c.String(http.StatusOK, file)
	}).Name = "deleteFile"

	e.Logger.Fatal(e.Start(":8080"))
}
