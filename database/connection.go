package database

import (
	"log"
	"os"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"

)

var DB *gorm.DB

func Connect() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}
	fmt.Println("El archivo .env cargo correctamente!")

	// variables de entorno
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dns := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dbUser, dbPassword, dbName, dbSSLMode)

	database, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("No se pudo conectar a la base de datos")
	}

	DB = database

}
