package media

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func MediaParser(r io.Reader) (*MediaInfo, error) {
	if r == nil {
		return nil, errors.New("missing parameter: r")
	}

	toInt := func(v string) int64 {
		val, _ := strconv.ParseInt(v, 10, 64)
		return val
	}

	infoData := ToInfoData(r)
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

// ToInfoData will parse a mediainfo result grouped on its keys, considering each key should have a value otherwise will be ignored
func ToInfoData(r io.Reader) InfoData {
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
