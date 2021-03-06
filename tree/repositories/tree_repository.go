package repositories

import (
	"database/sql"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"tree/entities"
)

var dbHost string
var dbPort int
var dbPassword string
var dbName string
var dbUser string
var db *pg.DB

func initDatabaseClient(){
	db = pg.Connect(&pg.Options{
		User: dbUser,
		Password: dbPassword,
		Database: dbName,
	})
}

func InitTestDatabase() {
	dbHost = os.Getenv("DB_HOST")
	dbPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	dbPassword = os.Getenv("DB_PASSWORD")
	dbUser = os.Getenv("DB_USER")
	dbName = os.Getenv("DB_NAME")
	if dbName == "" { // not launched with docker
		dbHost = "localhost"
		dbPort = 5432
		dbPassword = "supergarden"
		dbUser = "garden"
		dbName = "garden_test"
	}
	initDatabaseClient()

	// reset test database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("DROP DATABASE IF EXISTS " + dbName)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		log.Fatal(err)
	}

	createTables()
	populateTestDatabase()
}

func createTables(){
	models := []interface{}{ // easily add persistant models here
		(*entities.Tree)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func InitDatabase() {
	dbHost = os.Getenv("DB_HOST")
	dbPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	dbUser = os.Getenv("DB_USER")
	initDatabaseClient()
	createTables()
	populateDatabase()
}

func populateTestDatabase() {
	_, err := db.Model(&entities.Tree{
		Name: "Oak",
		Id: 1,
	}).Insert()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Model(&entities.Tree{
		Name: "Baobab",
		Id: 2,
	}).Insert()
	if err != nil {
		log.Fatal(err)
	}
}

func populateDatabase() {
	_, err := db.Model(&entities.Tree{
		Name: "Mahogany",
		Id: 1,
	}).Insert()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Model(&entities.Tree{
		Name: "Thuja ",
		Id: 2,
	}).Insert()
	if err != nil {
		log.Fatal(err)
	}
}

func FindAll() []entities.Tree {
	var Trees []entities.Tree
	err := db.Model(&Trees).
		Select()
	if err != nil {
		log.Fatal(err)
	}
	return Trees
}
