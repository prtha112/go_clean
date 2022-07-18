package RestApi

import (
	"fmt"
	auth "go_clean/RestApi/Auth"
	"go_clean/RestApi/Config"
	middleware "go_clean/RestApi/Middleware"
	"go_clean/RestApi/Models"
	"go_clean/RestApi/Routers"
	"go_clean/RestApi/Seeder"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/pseidemann/finish"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var err error

func Server() {
	var err error
	var env_auth auth.Env
	var env_middleware middleware.Env

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	Config.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	env_auth.Apikey = os.Getenv("API_TOKEN")
	env_auth.Jwtexpiretime, err = strconv.Atoi(os.Getenv("JWT_EXPIRE_TIME"))
	if err != nil {
		env_auth.Jwtexpiretime = 5 // minute
	}
	env_auth.Jwtsignature = os.Getenv("JWT_SIGNATURE")
	env_middleware.Signature = os.Getenv("JWT_SIGNATURE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	timeOut, err := strconv.Atoi(os.Getenv("GRACEFUL_TIMEOUT"))
	if err != nil {
		timeOut = 30 // second
	}

	Config.DB.AutoMigrate(&Models.Book{}, &Models.Category{}) // Create table
	Seeder.Load()                                             // Mockup data to table

	r := Routers.SetupRouter(env_auth, env_middleware)
	// running
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	graceful := &finish.Finisher{Timeout: time.Duration(timeOut) * time.Second}
	graceful.Add(srv)

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	graceful.Wait()
}
