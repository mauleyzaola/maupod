package filters

import (
	"errors"
)

type QueryFilter struct {
	Limit     int    `schema:"limit"`
	Offset    int    `schema:"offset"`
	Sort      string `schema:"sort"`
	Direction string `schema:"direction"`
	Query     string `schema:"query"`
}

func (f *QueryFilter) Validate() error {
	// check and assign limits
	if f.Limit < 0 {
		return errors.New("wrong limit format, expected positive number")
	}

	if f.Offset < 0 {
		return errors.New("wrong offset format, expected zero or greater number")
	}
	// consider the case when the caller is not passing the sort values
	if f.Direction != "" {
		if f.Direction != "asc" && f.Direction != "desc" {
			return errors.New("wrong direction value, expected [asc|desc]")
		}
		if f.Sort == "" {
			return errors.New("missing sort parameter")
		}

	}

	// sanity check, so we avoid errors executing invalid queries
	if f.Limit == 0 && f.Offset != 0 {
		return errors.New("if offset value is provided, it requires a limit value too")
	}

	return nil
}
