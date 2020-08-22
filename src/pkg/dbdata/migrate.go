package dbdata

import (
	"database/sql"

	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDbFromPath(db *sql.DB, driver, dir string) (int, error) {
	migrations := &migrate.FileMigrationSource{
		Dir: dir,
	}

	// this value needs to match the one in dbconfig.yml
	migrate.SetTable("migrations")
	return migrate.Exec(db, driver, migrations, migrate.Up)
}
