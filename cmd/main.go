package main

import (
	"log"

	"github.com/cristian-yw/Minitask_W10D2/internal/config"
	"github.com/cristian-yw/Minitask_W10D2/internal/routers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Printf("Error loading .env file: %v", err.Error())
	// 	return
	// }
	// log.Println(os.Getenv("DB_USER"))

	db, err := config.InitDD()
	if err != nil {
		log.Println("Error connecting to database: ", err.Error())
		return
	}
	defer db.Close()

	if err := config.TestDB(db); err != nil {
		log.Println("Error pinging database: ", err.Error())
		return
	}
	log.Println("Database connection successful")

	router := routers.InitRouter(db)
	router.Run("0.0.0.0:8080")
}
