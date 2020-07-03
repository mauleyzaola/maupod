package information

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func ParseMediaInfo(r io.Reader) (*MediaInfo, error) {
	if r == nil {
		return nil, errors.New("missing parameter: r")
	}

	toInt := func(v string) int64 {
		val, _ := strconv.ParseInt(v, 10, 64)
		return val
	}

	toFloat := func(v string) float64 {
		val, _ := strconv.ParseFloat(v, 64)
		return val
	}

	infoData := toInfoData(r)
	var mi = &MediaInfo{}
	for key, value := range infoData {
		defaultValue := value[0]
		switch strings.ToLower(key) {
		case strings.ToLower("Album"):
			mi.Album = defaultValue
		case strings.ToLower("Album/Performer"):
			mi.AlbumPerformer = defaultValue
		case strings.ToLower("Count"):
			mi.AudioCount = toInt(defaultValue)
		case strings.ToLower("Audio_Format_List"):
			mi.AudioFormatList = defaultValue
		case strings.ToLower("Bit depth"):
			for _, v := range value {
				if val := toInt(v); val != 0 {
					mi.BitDepth = val
				} else {
					mi.BitDepthString = v
				}
			}
		case strings.ToLower("Bit rate"):
			for _, v := range value {
				if val := toInt(v); val != 0 {
					mi.BitRate = val
					break
				}
			}
		case strings.ToLower("Bit rate mode"):
			mi.BitRateMode = defaultValue
		case strings.ToLower("Channel(s)"):
			mi.Channels = defaultValue
		case strings.ToLower("Channel layout"):
			mi.ChannelsLayout = defaultValue
		case strings.ToLower("Channel positions"):
			mi.ChannelsPosition = defaultValue
		case strings.ToLower("Comment"):
			mi.Comment = defaultValue
		case strings.ToLower("Commercial name"):
			mi.CommercialName = defaultValue
		case strings.ToLower("Complete name"):
			mi.CompleteName = defaultValue
		case strings.ToLower("Compression mode"):
			mi.Compression = defaultValue
		case strings.ToLower("Count of stream of this kind"):
			mi.CountOfAudioStreams = toInt(defaultValue)
		case strings.ToLower("Duration"):
			mi.Duration = toFloat(defaultValue)
		case strings.ToLower("Encoded_Library_Date"):
			mi.EncodedLibraryDate = defaultValue
		case strings.ToLower("Encoded_Library_Name"):
			mi.EncodedLibraryName = defaultValue
		case strings.ToLower("Encoded_Library_Version"):
			mi.EncodedLibraryVersion = defaultValue
		case strings.ToLower("File extension"):
			mi.FileExtension = defaultValue
		case strings.ToLower("File name extension"):
			mi.FileName = defaultValue
		case strings.ToLower("File size"):
			for _, v := range value {
				if val := toInt(v); val != 0 {
					mi.FileSize = val
					break
				}
			}
		case strings.ToLower("Folder name"):
			mi.FolderName = defaultValue
		case strings.ToLower("Format/Info"):
			mi.FormatInfo = defaultValue
		case strings.ToLower("Format"):
			mi.Format = defaultValue
		case strings.ToLower("Format/Url"):
			mi.FormatURL = defaultValue
		case strings.ToLower("Genre"):
			mi.Genre = defaultValue
		case strings.ToLower("Internet media type"):
			mi.InternetMediaType = defaultValue
		case strings.ToLower("Kind of stream"):
			for _, v := range value {
				if strings.ToLower(v) == strings.ToLower("General") {
					continue
				}
				mi.KindOfStream = v
				break
			}
		case strings.ToLower("Overall bit rate"):
			for _, v := range value {
				if val := toInt(v); val != 0 {
					mi.OverallBitRate = val
					break
				}
			}
		case strings.ToLower("Overall bit rate mode"):
			mi.OverallBitRateMode = defaultValue
		case strings.ToLower("Part"):
			mi.Part = toInt(defaultValue)
		case strings.ToLower("Part/Total"):
			mi.PartTotal = toInt(defaultValue)
		case strings.ToLower("Performer"):
			mi.Performer = defaultValue
		case strings.ToLower("Recorded date"):
			mi.RecordedDate = toInt(defaultValue)
		case strings.ToLower("Samples count"):
			for _, v := range value {
				if val := toInt(v); val != 0 {
					mi.SamplesCount = val
					break
				}
			}
		case strings.ToLower("Sampling rate"):
			for _, v := range value {
				if val := toInt(v); val != 0 {
					mi.SamplingRate = val
					break
				}
			}
		case strings.ToLower("Stream identifier"):
			for _, v := range value {
				if val := toInt(v); val != 0 {
					mi.StreamIdentifier = val
					break
				}
			}
		case strings.ToLower("Stream size"):
			for _, v := range value {
				if val := toInt(v); val != 0 {
					mi.StreamSize = val
					break
				}
			}
		case strings.ToLower("Title"):
			mi.Title = defaultValue
		case strings.ToLower("Track name/Position"):
			for _, v := range value {
				if val := toInt(v); val != 0 {
					mi.TrackNamePosition = val
					break
				}
			}
		case strings.ToLower("Track name/Total"):
			for _, v := range value {
				if val := toInt(v); val != 0 {
					mi.TrackNameTotal = val
					break
				}
			}
		case strings.ToLower("Track name"):
			mi.TrackName = defaultValue
		case strings.ToLower("Writing library"):
			mi.WritingLibrary = defaultValue
		case strings.ToLower("Composer"):
			mi.Composer = defaultValue
		}
	}
	return mi, nil
}

type InfoString string

// Split will return a key and a value, based in mediainfo text output format
func (in InfoString) Split() (key, value string) {
	const sep = ":"
	val := string(in)
	firstSep := strings.Index(val, sep)
	if firstSep == -1 {
		// no sep found, return nothing
		return
	}
	key = val[:firstSep]
	key = strings.TrimSpace(key)
	value = val[firstSep+1:]
	value = strings.TrimSpace(value)
	return
}

type InfoData map[string][]string

// toInfoData will parse a mediainfo result grouped on its keys, considering each key should have a value otherwise will be ignored
func toInfoData(r io.Reader) InfoData {
	var res = make(map[string][]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		info := InfoString(scanner.Text())
		k, v := info.Split()
		// ignore missing values
		if v == "" {
			continue
		}
		val := res[k]
		val = append(val, v)
		res[k] = val
	}
	return res
}
