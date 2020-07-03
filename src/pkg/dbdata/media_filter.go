package dbdata

import "strconv"

type MediaFilter struct {
	QueryFilter
	Format          string `schema:"format"`
	Filename        string `schema:"filename"`
	Distinct        string `schema:"distinct"`
	Album           string `schema:"album"`
	Performer       string `schema:"performer"`
	RecordedDate    int    `schema:"recorded_date"`
	Genre           string `schema:"genre"`
	AlbumIdentifier string `schema:"album_identifier"`

	SearchAlbum                string `schema:"search_album"`
	SearchPerformer            string `schema:"search_performer"`
	SearchGenre                string `schema:"search_genre"`
	SearchTrackName            string `schema:"search_track_name"`
	SearchRecordedDateMin      string `schema:"search_recorded_date_min"`
	SearchRecordedDateMax      string `schema:"search_recorded_date_max"`
	SearchRecordedDateMinValue *int64 `schema:"-"`
	SearchRecordedDateMaxValue *int64 `schema:"-"`
}

func (f *MediaFilter) Validate() error {
	if val := f.SearchRecordedDateMin; val != "" {
		value, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		f.SearchRecordedDateMinValue = &value
	}
	if val := f.SearchRecordedDateMax; val != "" {
		value, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		f.SearchRecordedDateMaxValue = &value
	}
	return f.QueryFilter.Validate()
}
