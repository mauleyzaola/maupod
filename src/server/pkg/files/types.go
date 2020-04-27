package files

type MediaInfo struct {
	Media Media `json:"media"`
}

type Media struct {
	Ref   string  `json:"@ref"`
	Track []Track `json:"track"`
}

type Track struct {
	Type                  string `json:"@type"`
	AudioCount            string `json:"AudioCount,omitempty"`
	FileExtension         string `json:"FileExtension,omitempty"`
	Format                string `json:"Format"`
	FileSize              string `json:"FileSize,omitempty"`
	Duration              string `json:"Duration"`
	OverallBitRateMode    string `json:"OverallBitRate_Mode,omitempty"`
	OverallBitRate        string `json:"OverallBitRate,omitempty"`
	StreamSize            string `json:"StreamSize"`
	Title                 string `json:"Title,omitempty"`
	Album                 string `json:"Album,omitempty"`
	PartPosition          string `json:"Part_Position,omitempty"`
	Composer              string `json:"Composer,omitempty"`
	Cover                 string `json:"Cover,omitempty"`
	CoverMime             string `json:"Cover_Mime,omitempty"`
	Lyrics                string `json:"Lyrics,omitempty"`
	Comment               string `json:"Comment,omitempty"`
	Extra                 Extra  `json:"extra,omitempty"`
	FormatVersion         string `json:"Format_Version,omitempty"`
	FormatProfile         string `json:"Format_Profile,omitempty"`
	SamplesPerFrame       string `json:"SamplesPerFrame,omitempty"`
	FrameRate             string `json:"FrameRate,omitempty"`
	FrameCount            string `json:"FrameCount,omitempty"`
	AlbumPerformer        string `json:"Album_Performer,omitempty"`
	Part                  string `json:"Part,omitempty"`
	PartPositionTotal     string `json:"Part_Position_Total,omitempty"`
	Track                 string `json:"Track,omitempty"`
	TrackPosition         string `json:"Track_Position,omitempty"`
	TrackPositionTotal    string `json:"Track_Position_Total,omitempty"`
	Performer             string `json:"Performer,omitempty"`
	Genre                 string `json:"Genre,omitempty"`
	RecordedDate          string `json:"Recorded_Date,omitempty"`
	FileModifiedDate      string `json:"File_Modified_Date,omitempty"`
	FileModifiedDateLocal string `json:"File_Modified_Date_Local,omitempty"`
	BitRateMode           string `json:"BitRate_Mode,omitempty"`
	BitRate               string `json:"BitRate,omitempty"`
	Channels              string `json:"Channels,omitempty"`
	ChannelPositions      string `json:"ChannelPositions,omitempty"`
	ChannelLayout         string `json:"ChannelLayout,omitempty"`
	SamplingRate          string `json:"SamplingRate,omitempty"`
	SamplingCount         string `json:"SamplingCount,omitempty"`
	BitDepth              string `json:"BitDepth,omitempty"`
	EncodedLibrary        string `json:"Encoded_Library,omitempty"`
	EncodedLibraryName    string `json:"Encoded_Library_Name,omitempty"`
	EncodedLibraryVersion string `json:"Encoded_Library_Version,omitempty"`
	EncodedLibraryDate    string `json:"Encoded_Library_Date,omitempty"`
}

type Extra struct {
	ITunPGAP string `json:"iTunPGAP"`
}
