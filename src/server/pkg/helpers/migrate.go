package helpers

import (
	"database/sql"
	"path/filepath"

	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDbFromPath(db *sql.DB, driver, dir string) (int, error) {
	dirPath := filepath.Join(PathBackend(), dir)
	migrations := &migrate.FileMigrationSource{
		Dir: dirPath,
	}

	// this value needs to match the one in dbconfig.yml
	migrate.SetTable("migrations")
	return migrate.Exec(db, driver, migrations, migrate.Up)
}
