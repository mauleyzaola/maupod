package media

import (
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

func (m *MediaInfo) ToProto() *pb.Media {
	res := &pb.Media{
		Format:             m.Format,
		FileSize:           m.FileSize,
		Duration:           m.Duration,
		OverallBitRate:     m.OverallBitRate,
		OverallBitRateMode: m.OverallBitRateMode,
		StreamSize:         m.StreamSize,
		Album:              m.Album,
		Title:              m.Title,
		Track:              m.TrackName,
		TrackPosition:      m.TrackNamePosition,
		Performer:          m.Performer,
		Genre:              m.Genre,
		RecordedDate:       m.RecordedDate,
		Comment:            m.Comment,
	}

	return res
}

// TODO: include 	TrackNameTotal        int64

/*
   string comment = 21;
   string channels = 22;
   string channel_positions = 23;
   string channel_layout = 24;
   int64 sampling_rate = 25;
   int64 sampling_count = 26;
   int64 bit_depth = 27;
   string compression_mode = 28;
   string encoded_library = 29;
   string encoded_library_name = 30;
   string encoded_library_version = 31;
   string bit_rate_mode = 32;
   int64 bit_rate = 33;

*/

/*
	AlbumPerformer        string
	AudioCount            int64
	AudioFormatList       string
	BitDepth              int64
	BitDepthString        string
	Channels              int64
	ChannelsLayout        string
	ChannelsPosition      string
	CommercialName        string
	CompleteName          string
	Compression           string
	CountOfAudioStreams   int64
	EncodedLibraryDate    string
	EncodedLibraryName    string
	EncodedLibraryVersion string
	FileExtension         string
	FileName              string
	FolderName            string
	FormatInfo            string
	Format                string
	FormatURL             string
	InternetMediaType     string
	KindOfStream          string
	Part                  int64
	PartTotal             int64
	SamplesCount          int64
	SamplingRate          int64
	StreamIdentifier      int64
	WritingLibrary        string
*/
