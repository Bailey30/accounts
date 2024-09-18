package main

import (
	"log"
	"os"

	"github.com/Bailey30/accounts/pkg/accounts"
	"github.com/Bailey30/accounts/pkg/config"
	"github.com/Bailey30/accounts/pkg/db"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// allow reading of env variables from env file
	envErr := godotenv.Load("../.env")
	if envErr != nil {
		log.Fatal(envErr)
	}
	// fmt.Println("db url", os.Getenv("DATABASE_URL"))
	//
	// args := os.Args
	//
	// if (len(args)) < 2 {
	// 	fmt.Println("no args")
	// }
	//
	// for i, arg := range args[1:] {
	// 	fmt.Printf("Argument %d: %s\n", i+1, arg)
	// }

	database := db.NewDatabase(os.Getenv("DATABASE_URL"))

	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	//     const (
	// 	selectArticle      = `SELECT * FROM articles WHERE id=$1`
	// 	selectManyArticles = `SELECT * FROM articles LIMIT $1 OFFSET $2`
	// 	insertArticle      = `INSERT INTO articles (title, body, created_at, updated_at) VALUES ($1, $2, now(), now()) RETURNING id`
	// 	updateArticle      = `UPDATE articles SET title = $1, body = $2, updated_at = now() WHERE id = $3`
	// )
	// database.CreateTable(`CREATE TABLE IF NOT EXISTS sale (
	// id SERIAL PRIMARY KEY,
	// amount DECIMAL(10, 2),
	//    payment_date DATE,
	// created_at timestamptz not null,
	// updated_at timestamptz not null
	// );`)

	database.LogTables()
	// database.CreateUser()
	database.GetUsers()
	// database.DropTable(`"sale"`)

	args := config.GetArgs()

	config, err := config.NewConfig(args)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Args: %v\n", args)
	// fmt.Printf("Config: %v\n", config)

	accounts.EntityHandler(config, database)

}
