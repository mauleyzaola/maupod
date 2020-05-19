package media

// MediaInfo represents the output from `mediainfo` program
type MediaInfo struct {
	AlbumPerformer        string
	Album                 string
	AudioCount            int64
	AudioFormatList       string
	BitDepth              int64
	BitDepthString        string
	Channels              int64
	ChannelsLayout        string
	ChannelsPosition      string
	Comment               string
	CommercialName        string
	CompleteName          string
	Compression           string
	CountOfAudioStreams   int64
	Duration              float64
	EncodedLibraryDate    string
	EncodedLibraryName    string
	EncodedLibraryVersion string
	FileExtension         string
	FileName              string
	FileSize              int64
	FolderName            string
	FormatInfo            string
	Format                string
	FormatURL             string
	Genre                 string
	InternetMediaType     string
	KindOfStream          string
	OverallBitRate        int64
	OverallBitRateMode    string
	Part                  int64
	PartTotal             int64
	Performer             string
	RecordedDate          int64
	SamplesCount          int64
	SamplingRate          int64
	StreamIdentifier      int64
	StreamSize            int64
	Title                 string
	TrackNamePosition     int64
	TrackName             string
	TrackNameTotal        int64
	WritingLibrary        string
}
