package media

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func MediaParser(data []byte) (*MediaInfo, error) {
	return nil, errors.New("not implemented")
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
