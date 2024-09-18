package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type User struct {
	User_Id  int    `db:"user_id"`
	Username string `db:"username"`
}

type Database struct {
	databaseUrl string
	Dbx         *sqlx.DB
}

func NewDatabase(databaseUrl string) *Database {
	return &Database{
		databaseUrl: databaseUrl,
	}
}

func (db *Database) Connect() error {
	fmt.Printf("Attempting to connect to database at: %s\n", db.databaseUrl)

	Dbx, err := sqlx.Connect("postgres", db.databaseUrl)
	if err != nil {
		return err
	}

	db.Dbx = Dbx

	fmt.Println("Connected to the database")

	return nil
}

func (db *Database) close() error {
	return db.Dbx.Close()
}

func (db *Database) CreateTable(schema string) error {
	err := db.Connect()
	if err != nil {
		return err
	}
	defer db.close()

	db.Dbx.MustExec(schema)

	return nil
}

func (db *Database) LogTables() {
	query := `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';`
	var tables []string
	err := db.Dbx.Select(&tables, query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tables in the database:")
	for _, table := range tables {
		fmt.Println(table)
	}
}

func (db *Database) GetUsers() {
	users := []User{}

	err := db.Dbx.Select(&users, `SELECT * FROM "user"`)
	if err != nil {
		log.Fatalln(err)
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s\n", user.User_Id, user.Username)
	}
}

func (db *Database) CreateUser() error {
	fmt.Println("createUser")
	query := `INSERT INTO "user" (username) VALUES ('alex') RETURNING user_id`
	var userID int
	err := db.Dbx.QueryRow(query).Scan(&userID)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("New user ID: %d\n", userID)
	return nil
}

func (db *Database) DropTable(table string) {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, table)
	_, err := db.Dbx.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Table dropped successfully.")
}
