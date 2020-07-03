package information

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInfoFromFile(t *testing.T) {
	raw, err := MediaInfoFromFile("./test_data/sample1.m4a")
	info, err := ParseMediaInfo(raw)
	require.NoError(t, err, "file should be present")
	require.NotNil(t, info, "media info should return an object")

	assert.EqualValues(t, "./test_data/sample1.m4a", info.CompleteName)
	assert.EqualValues(t, "./test_data", info.FolderName)
	assert.EqualValues(t, "m4a", info.FileExtension)
	assert.EqualValues(t, "MPEG-4", info.Format)
	assert.EqualValues(t, "MPEG-4", info.CommercialName)
	assert.EqualValues(t, "audio/mp4", info.InternetMediaType)
	assert.EqualValues(t, 450690, info.FileSize)
	assert.EqualValues(t, 12005, info.Duration)
	assert.EqualValues(t, "CBR", info.OverallBitRateMode)
	assert.EqualValues(t, 300335, info.OverallBitRate)
	assert.EqualValues(t, 3246, info.StreamSize)
	assert.EqualValues(t, "Infinite Pursuit", info.Title)
	assert.EqualValues(t, "Fables", info.Album)
	assert.EqualValues(t, "Jean-Luc Ponty", info.AlbumPerformer)
	assert.EqualValues(t, 0, info.Part)
	assert.EqualValues(t, 1, info.PartTotal)
	assert.EqualValues(t, "Jean-Luc Ponty", info.Performer)
	assert.EqualValues(t, "Jean-Luc Ponty", info.Composer)
	assert.EqualValues(t, "Front: L R", info.ChannelsPosition)
}
