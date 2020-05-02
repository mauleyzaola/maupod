package files

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMediainfoParser(t *testing.T) {
	file, err := os.Open("./test_data/mediainfo.json")
	require.NoError(t, err)
	defer func() {
		err = file.Close()
		assert.NoError(t, err)
	}()

	var mediaInfos []MediaInfo
	err = json.NewDecoder(file).Decode(&mediaInfos)
	require.NoError(t, err)
	assert.Len(t, mediaInfos, 100, "song count should match")
}
