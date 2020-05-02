package psql

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDb(t *testing.T) {
	db, err := ConnectPostgres(`user=postgres host=127.0.0.1 password=nevermind sslmode=disable`)
	require.NoError(t, err)
	assert.NotNil(t, db)

	const maupodDb = "maupod"
	ok, err := DatabaseExists(db, maupodDb)
	require.NoError(t, err)

	if !ok {
		err = CreateDbIfNotExists(db, maupodDb)
		require.NoError(t, err)
	}
}
