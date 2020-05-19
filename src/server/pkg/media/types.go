package media

// MediaInfo represents the output from `mediainfo` program
type MediaInfo struct {
	KindOfStream          string
	StreamIdentifier      int64
	CountOfAudioStreams   int64
	AudioFormatList       string
	CompleteName          string
	FolderName            string
	FileName              string
	FileExtension         string
	Format                string
	FormatInfo            string
	FormatURL             string
	CommercialName        string
	InternetMediaType     string
	FileSize              int64
	Duration              int64
	OverallBitRate        string
	OverallBitRateMode    string
	StreamSize            int64
	Title                 string
	Album                 string
	AlbumPerformer        string
	Part                  int64
	PartTotal             int64
	TrackName             string
	TrackNamePosition     int64
	TrackNameTotal        int64
	Performer             string
	Gener                 string
	RecordedDate          int64
	AudioCount            int64
	Channels              int64
	ChannelsPosition      string
	ChannelsLayout        string
	SamplingRate          int64
	SamplesCount          int64
	BitDepth              int64
	BitDepthString        string
	Compression           string
	WritingLibrary        string
	EncodedLibraryName    string
	EncodedLibraryVersion string
	EncodedLibraryDate    string
}
