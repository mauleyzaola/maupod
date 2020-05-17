package domain

import "time"

type Media struct {
	// uuid
	ID string `json:"id"`

	// the sha 256 to uniquely identify a file (dup cleanup)
	Sha string `json:"sha"`

	// the full path to the file on the file store
	Location string `json:"location,omitempty"`

	// date the media file was scanned
	LastScan time.Time `json:"last_scan"`

	// file system modified date
	ModifiedDate time.Time `json:"modified_date"`

	// media info stuff
	FileExtension         string    `json:"file_extension,omitempty"`
	Format                string    `json:"format,omitempty"`
	FileSize              int64     `json:"file_size"`
	Duration              float64   `json:"duration"`
	OverallBitRateMode    string    `json:"overall_bit_rate_mode,omitempty"`
	OverallBitRate        int64     `json:"overall_bit_rate"`
	StreamSize            int64     `json:"stream_size"`
	Album                 string    `json:"album"`
	Title                 string    `json:"title"`
	Track                 string    `json:"track"`
	TrackPosition         int64     `json:"track_position"`
	Performer             string    `json:"performer,omitempty"`
	Genre                 string    `json:"genre,omitempty"`
	RecordedDate          int64     `json:"recorded_date"`
	FileModifiedDate      time.Time `json:"file_modified_date"`
	Comment               string    `json:"comment,omitempty"`
	Channels              string    `json:"channels,omitempty"`
	ChannelPositions      string    `json:"channel_positions,omitempty"`
	ChannelLayout         string    `json:"channel_layout,omitempty"`
	SamplingRate          int64     `json:"sampling_rate,omitempty"`
	SamplingCount         int64     `json:"sampling_count,omitempty"`
	BitDepth              int64     `json:"bit_depth"`
	CompressionMode       string    `json:"compression_mode,omitempty"`
	EncodedLibrary        string    `json:"encoded_library,omitempty"`
	EncodedLibraryName    string    `json:"encoded_library_name,omitempty"`
	EncodedLibraryVersion string    `json:"encoded_library_version,omitempty"`
	BitRateMode           string    `json:"bit_rate_mode"`
	BitRate               int64     `json:"bit_rate"`
}
