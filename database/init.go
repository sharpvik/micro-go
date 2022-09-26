package database

import (
	"embed"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // the DB driver

	"github.com/sharpvik/micro-go/configs"
	"github.com/sharpvik/micro-go/migrations"
)

func MustInit(config *configs.Database) (db *sqlx.DB) {
	details := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name)
	db = mustConnect(details)
	up(db)
	return
}

func mustConnect(details string) (db *sqlx.DB) {
	log.Print("connecting to the database ...")

	var err error
	for i := 0; i < 10; i++ {
		if db, err = sqlx.Connect("postgres", details); err == nil {
			return
		}
		time.Sleep(time.Second)
	}

	if err != nil {
		log.Fatalln("failed to connect to the database:", err)
	}
	return
}

func up(db *sqlx.DB) {
	files, err := migrations.Up.ReadDir(".")
	if err != nil {
		log.Fatalln("failed to list up migrations:", err)
		return
	}
	log.Print("applying migrations ...")
	for _, file := range files {
		if err := readAndApply(db, migrations.Up, file.Name()); err != nil {
			log.Fatal(err)
		}
	}
}

func readAndApply(conn *sqlx.DB, fs embed.FS, path string) (err error) {
	log.Print(path)
	migration, err := fs.ReadFile(path)
	if err != nil {
		return
	}
	_, err = conn.Exec(string(migration))
	return
}
