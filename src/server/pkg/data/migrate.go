package data

import (
	"database/sql"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDbFromPath(db *sql.DB, driver, dir string) (int, error) {
	dirPath := filepath.Join(helpers.PathBackend(), dir)
	migrations := &migrate.FileMigrationSource{
		Dir: dirPath,
	}

	// this value needs to match the one in dbconfig.yml
	migrate.SetTable("migrations")
	return migrate.Exec(db, driver, migrations, migrate.Up)
}
