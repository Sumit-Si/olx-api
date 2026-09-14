package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sumit-si/olx-api/internal/config"
)

func main() {
	// fmt.Println(os.Args)
	if len(os.Args) < 2 {
		log.Fatal("usage: migrate <up | down>")
	}

	cfg := config.MustLoad()

	m, err := migrate.New(
		"file://migrations",
		cfg.DatabaseUrl)
	// "github://mattes:personal-access-token@mattes/migrate_test",
	// "postgres://localhost:5432/database?sslmode=enable")

	if err != nil {
		log.Fatalf("migration.new: %v", err)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
		log.Println("up called")
	case "down":
		// m.Down -> all the way down to first (means all the previous migration vanishes which is very dangerous) so instead we use Steps: -ve values means down migration by how much you define like < 0 and so on and +ve values means up migration by how much you define like 1,2,3 and so on
		if err := m.Steps(-1); err != nil {
			log.Fatal(err)
		}
		log.Printf("down called")
	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}

	fmt.Println("running migration")
}
