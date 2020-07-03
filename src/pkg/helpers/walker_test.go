package helpers

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWalkFiles(t *testing.T) {
	var mode = os.ModePerm
	// create file structure
	uid := NewUUID()
	assert.NotEmpty(t, uid)

	root := filepath.Join(os.TempDir(), uid)
	err := os.MkdirAll(root, mode)
	require.NoError(t, err)

	defer func() {
		err = os.RemoveAll(root)
		assert.NoError(t, err)
	}()

	beers := filepath.Join(root, "beers")
	err = os.MkdirAll(beers, mode)
	require.NoError(t, err)

	wines := filepath.Join(root, "wines")
	err = os.MkdirAll(wines, mode)
	require.NoError(t, err)

	err = ioutil.WriteFile(filepath.Join(beers, "leffe1.txt"), nil, mode)
	require.NoError(t, err)
	err = ioutil.WriteFile(filepath.Join(beers, "leffe2.txt"), nil, mode)
	require.NoError(t, err)
	err = ioutil.WriteFile(filepath.Join(beers, "leffe3.txt"), nil, mode)
	require.NoError(t, err)

	err = ioutil.WriteFile(filepath.Join(wines, "rioja1.txt"), nil, mode)
	require.NoError(t, err)
	err = ioutil.WriteFile(filepath.Join(wines, "rioja2.txt"), nil, mode)
	require.NoError(t, err)
	err = ioutil.WriteFile(filepath.Join(wines, "rioja3.txt"), nil, mode)
	require.NoError(t, err)

	type myfile struct {
		base     string
		fullpath string
		isdir    bool
	}

	files := make(map[string]*myfile)

	var fn WalkerFunc = func(name string, isDir bool) bool {
		t.Logf("walking file: %s isDir: %v", name, isDir)
		myfile := &myfile{
			base:     filepath.Base(name),
			fullpath: name,
			isdir:    isDir,
		}
		_, ok := files[myfile.fullpath]
		assert.False(t, ok, "we should not walk two times the same file")
		files[myfile.fullpath] = myfile
		return false
	}

	err = WalkFiles(root, fn)
	assert.NoError(t, err)
}
