package sqlite

import (
	"database/sql"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

//go:embed migrations/*.sql
var embedMigration embed.FS

// Migrate applies embedded migration files to the provided database instance.
// After applying, a *_db_ver file is generated containing the already applied migrations.
//
// The implementation is not robust to say the least, but a package
// that would work with atlas-generated files has not been found.
// It is better to migrate manually using tools such as atlas
// (the *_db_ver file should also be updated manually so that the app does not try to reapply migrations).
func Migrate(db *sql.DB, name string) error {
	entries, err := embedMigration.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("failed to read migration directory: %w", err)
	}

	appliedMigrations, appliedMigrationsFile := getAppliedMigrations(name)
	defer appliedMigrationsFile.Close()

	for i, entry := range entries {
		if i < len(appliedMigrations) && appliedMigrations[i] == entry.Name() {
			continue
		}

		content, err := fs.ReadFile(embedMigration, "migrations/"+entry.Name())
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", entry.Name(), err)
		}

		queries := strings.Split(string(content), ";")
		for _, query := range queries {
			if strings.TrimSpace(query) == "" {
				continue
			}
			if _, err := db.Exec(query); err != nil {
				return fmt.Errorf("failed to execute migration query in file %s: %w", entry.Name(), err)
			}
		}

		_, err = appliedMigrationsFile.Write([]byte(entry.Name() + ";"))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Executed migration file: %s", entry.Name())
	}

	return nil
}

func getAppliedMigrations(name string) ([]string, os.File) {
	appliedMigrationsFile, err := os.OpenFile(name+"_db_ver", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	appliedMigrationsRaw, err := io.ReadAll(appliedMigrationsFile)
	if err != nil {
		log.Fatal(err)
	}

	var appliedMigrations []string
	if len(appliedMigrationsRaw) > 0 {
		appliedMigrations = strings.Split(string(appliedMigrationsRaw), ";")
	}

	return appliedMigrations, *appliedMigrationsFile
}
