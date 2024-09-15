package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"log"
	"tender_api/src/internal/config"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatal(errors.Wrap(err, "init config"))
	}

	m, err := migrate.New("file://./src/cmd/migrate/migrations/postgresql", cfg.PostgresConn)
	if err != nil {
		log.Fatal(err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Applied migration: %d. Dirty: %t\n", version, dirty)
}
