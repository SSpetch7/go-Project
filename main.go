package main

import (
	"fmt"
	"go-project/handler"
	"go-project/repository"
	"go-project/service"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initTimeZome()
	initConfig()

	db := initDatabase()

	userRepository := repository.NewUserRepositoryDB(db) // DB
	// customerRepositoryMock := repository.NewCustomerRepositoryMock() // Mock

	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	authService := service.NewAuthService()
	authHandler := handler.NewAuthHandler(authService)

	// _ = userRepository
	// _ = customerRepositoryMock

	app := fiber.New()
	app.Get("/users", userHandler.GetUsers)
	app.Post("/users", userHandler.RegisterUser)
	app.Post("/login", userHandler.Login)
	app.Get("/auth/:token", authHandler.VerifyToken)
	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// เป็นการ set ค่าเฉพาะ เช่น เปลี่ยนจาก app.port = 8000 เป็น 5000 จะได้ว่า APP_PORT=5000 ใช้กับ dockerfiler
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZome() {
	ict, err := time.LoadLocation("Asia/Bangkok")

	if err != nil {
		panic(err)
	}

	time.Local = ict

}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"))

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	return db

}
