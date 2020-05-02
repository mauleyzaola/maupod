// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Medium is an object representing the database table.
type Medium struct {
	ID                    string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Location              string    `boil:"location" json:"location" toml:"location" yaml:"location"`
	FileExtension         string    `boil:"file_extension" json:"file_extension" toml:"file_extension" yaml:"file_extension"`
	Format                string    `boil:"format" json:"format" toml:"format" yaml:"format"`
	FileSize              int64     `boil:"file_size" json:"file_size" toml:"file_size" yaml:"file_size"`
	Duration              float64   `boil:"duration" json:"duration" toml:"duration" yaml:"duration"`
	OverallBitRateMode    string    `boil:"overall_bit_rate_mode" json:"overall_bit_rate_mode" toml:"overall_bit_rate_mode" yaml:"overall_bit_rate_mode"`
	OverallBitRate        int64     `boil:"overall_bit_rate" json:"overall_bit_rate" toml:"overall_bit_rate" yaml:"overall_bit_rate"`
	StreamSize            int64     `boil:"stream_size" json:"stream_size" toml:"stream_size" yaml:"stream_size"`
	Album                 string    `boil:"album" json:"album" toml:"album" yaml:"album"`
	Track                 string    `boil:"track" json:"track" toml:"track" yaml:"track"`
	Title                 string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	TrackPosition         int64     `boil:"track_position" json:"track_position" toml:"track_position" yaml:"track_position"`
	Performer             string    `boil:"performer" json:"performer" toml:"performer" yaml:"performer"`
	Genre                 string    `boil:"genre" json:"genre" toml:"genre" yaml:"genre"`
	RecordedDate          int64     `boil:"recorded_date" json:"recorded_date" toml:"recorded_date" yaml:"recorded_date"`
	FileModifiedDate      time.Time `boil:"file_modified_date" json:"file_modified_date" toml:"file_modified_date" yaml:"file_modified_date"`
	Comment               string    `boil:"comment" json:"comment" toml:"comment" yaml:"comment"`
	Channels              string    `boil:"channels" json:"channels" toml:"channels" yaml:"channels"`
	ChannelPositions      string    `boil:"channel_positions" json:"channel_positions" toml:"channel_positions" yaml:"channel_positions"`
	ChannelLayout         string    `boil:"channel_layout" json:"channel_layout" toml:"channel_layout" yaml:"channel_layout"`
	SamplingRate          int64     `boil:"sampling_rate" json:"sampling_rate" toml:"sampling_rate" yaml:"sampling_rate"`
	SamplingCount         int64     `boil:"sampling_count" json:"sampling_count" toml:"sampling_count" yaml:"sampling_count"`
	BitDepth              int64     `boil:"bit_depth" json:"bit_depth" toml:"bit_depth" yaml:"bit_depth"`
	CompressionMode       string    `boil:"compression_mode" json:"compression_mode" toml:"compression_mode" yaml:"compression_mode"`
	EncodedLibrary        string    `boil:"encoded_library" json:"encoded_library" toml:"encoded_library" yaml:"encoded_library"`
	EncodedLibraryName    string    `boil:"encoded_library_name" json:"encoded_library_name" toml:"encoded_library_name" yaml:"encoded_library_name"`
	EncodedLibraryVersion string    `boil:"encoded_library_version" json:"encoded_library_version" toml:"encoded_library_version" yaml:"encoded_library_version"`
	EncodedLibraryDate    time.Time `boil:"encoded_library_date" json:"encoded_library_date" toml:"encoded_library_date" yaml:"encoded_library_date"`
	BitRateMode           string    `boil:"bit_rate_mode" json:"bit_rate_mode" toml:"bit_rate_mode" yaml:"bit_rate_mode"`
	BitRate               int64     `boil:"bit_rate" json:"bit_rate" toml:"bit_rate" yaml:"bit_rate"`

	R *mediumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mediumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MediumColumns = struct {
	ID                    string
	Location              string
	FileExtension         string
	Format                string
	FileSize              string
	Duration              string
	OverallBitRateMode    string
	OverallBitRate        string
	StreamSize            string
	Album                 string
	Track                 string
	Title                 string
	TrackPosition         string
	Performer             string
	Genre                 string
	RecordedDate          string
	FileModifiedDate      string
	Comment               string
	Channels              string
	ChannelPositions      string
	ChannelLayout         string
	SamplingRate          string
	SamplingCount         string
	BitDepth              string
	CompressionMode       string
	EncodedLibrary        string
	EncodedLibraryName    string
	EncodedLibraryVersion string
	EncodedLibraryDate    string
	BitRateMode           string
	BitRate               string
}{
	ID:                    "id",
	Location:              "location",
	FileExtension:         "file_extension",
	Format:                "format",
	FileSize:              "file_size",
	Duration:              "duration",
	OverallBitRateMode:    "overall_bit_rate_mode",
	OverallBitRate:        "overall_bit_rate",
	StreamSize:            "stream_size",
	Album:                 "album",
	Track:                 "track",
	Title:                 "title",
	TrackPosition:         "track_position",
	Performer:             "performer",
	Genre:                 "genre",
	RecordedDate:          "recorded_date",
	FileModifiedDate:      "file_modified_date",
	Comment:               "comment",
	Channels:              "channels",
	ChannelPositions:      "channel_positions",
	ChannelLayout:         "channel_layout",
	SamplingRate:          "sampling_rate",
	SamplingCount:         "sampling_count",
	BitDepth:              "bit_depth",
	CompressionMode:       "compression_mode",
	EncodedLibrary:        "encoded_library",
	EncodedLibraryName:    "encoded_library_name",
	EncodedLibraryVersion: "encoded_library_version",
	EncodedLibraryDate:    "encoded_library_date",
	BitRateMode:           "bit_rate_mode",
	BitRate:               "bit_rate",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelperfloat64 struct{ field string }

func (w whereHelperfloat64) EQ(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat64) NEQ(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat64) LT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat64) LTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat64) GT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat64) GTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MediumWhere = struct {
	ID                    whereHelperstring
	Location              whereHelperstring
	FileExtension         whereHelperstring
	Format                whereHelperstring
	FileSize              whereHelperint64
	Duration              whereHelperfloat64
	OverallBitRateMode    whereHelperstring
	OverallBitRate        whereHelperint64
	StreamSize            whereHelperint64
	Album                 whereHelperstring
	Track                 whereHelperstring
	Title                 whereHelperstring
	TrackPosition         whereHelperint64
	Performer             whereHelperstring
	Genre                 whereHelperstring
	RecordedDate          whereHelperint64
	FileModifiedDate      whereHelpertime_Time
	Comment               whereHelperstring
	Channels              whereHelperstring
	ChannelPositions      whereHelperstring
	ChannelLayout         whereHelperstring
	SamplingRate          whereHelperint64
	SamplingCount         whereHelperint64
	BitDepth              whereHelperint64
	CompressionMode       whereHelperstring
	EncodedLibrary        whereHelperstring
	EncodedLibraryName    whereHelperstring
	EncodedLibraryVersion whereHelperstring
	EncodedLibraryDate    whereHelpertime_Time
	BitRateMode           whereHelperstring
	BitRate               whereHelperint64
}{
	ID:                    whereHelperstring{field: "\"media\".\"id\""},
	Location:              whereHelperstring{field: "\"media\".\"location\""},
	FileExtension:         whereHelperstring{field: "\"media\".\"file_extension\""},
	Format:                whereHelperstring{field: "\"media\".\"format\""},
	FileSize:              whereHelperint64{field: "\"media\".\"file_size\""},
	Duration:              whereHelperfloat64{field: "\"media\".\"duration\""},
	OverallBitRateMode:    whereHelperstring{field: "\"media\".\"overall_bit_rate_mode\""},
	OverallBitRate:        whereHelperint64{field: "\"media\".\"overall_bit_rate\""},
	StreamSize:            whereHelperint64{field: "\"media\".\"stream_size\""},
	Album:                 whereHelperstring{field: "\"media\".\"album\""},
	Track:                 whereHelperstring{field: "\"media\".\"track\""},
	Title:                 whereHelperstring{field: "\"media\".\"title\""},
	TrackPosition:         whereHelperint64{field: "\"media\".\"track_position\""},
	Performer:             whereHelperstring{field: "\"media\".\"performer\""},
	Genre:                 whereHelperstring{field: "\"media\".\"genre\""},
	RecordedDate:          whereHelperint64{field: "\"media\".\"recorded_date\""},
	FileModifiedDate:      whereHelpertime_Time{field: "\"media\".\"file_modified_date\""},
	Comment:               whereHelperstring{field: "\"media\".\"comment\""},
	Channels:              whereHelperstring{field: "\"media\".\"channels\""},
	ChannelPositions:      whereHelperstring{field: "\"media\".\"channel_positions\""},
	ChannelLayout:         whereHelperstring{field: "\"media\".\"channel_layout\""},
	SamplingRate:          whereHelperint64{field: "\"media\".\"sampling_rate\""},
	SamplingCount:         whereHelperint64{field: "\"media\".\"sampling_count\""},
	BitDepth:              whereHelperint64{field: "\"media\".\"bit_depth\""},
	CompressionMode:       whereHelperstring{field: "\"media\".\"compression_mode\""},
	EncodedLibrary:        whereHelperstring{field: "\"media\".\"encoded_library\""},
	EncodedLibraryName:    whereHelperstring{field: "\"media\".\"encoded_library_name\""},
	EncodedLibraryVersion: whereHelperstring{field: "\"media\".\"encoded_library_version\""},
	EncodedLibraryDate:    whereHelpertime_Time{field: "\"media\".\"encoded_library_date\""},
	BitRateMode:           whereHelperstring{field: "\"media\".\"bit_rate_mode\""},
	BitRate:               whereHelperint64{field: "\"media\".\"bit_rate\""},
}

// MediumRels is where relationship names are stored.
var MediumRels = struct {
}{}

// mediumR is where relationships are stored.
type mediumR struct {
}

// NewStruct creates a new relationship struct
func (*mediumR) NewStruct() *mediumR {
	return &mediumR{}
}

// mediumL is where Load methods for each relationship are stored.
type mediumL struct{}

var (
	mediumAllColumns            = []string{"id", "location", "file_extension", "format", "file_size", "duration", "overall_bit_rate_mode", "overall_bit_rate", "stream_size", "album", "track", "title", "track_position", "performer", "genre", "recorded_date", "file_modified_date", "comment", "channels", "channel_positions", "channel_layout", "sampling_rate", "sampling_count", "bit_depth", "compression_mode", "encoded_library", "encoded_library_name", "encoded_library_version", "encoded_library_date", "bit_rate_mode", "bit_rate"}
	mediumColumnsWithoutDefault = []string{"id", "location", "file_extension", "format", "file_size", "duration", "overall_bit_rate_mode", "overall_bit_rate", "stream_size", "album", "track", "title", "track_position", "performer", "genre", "recorded_date", "file_modified_date", "comment", "channels", "channel_positions", "channel_layout", "sampling_rate", "sampling_count", "bit_depth", "compression_mode", "encoded_library", "encoded_library_name", "encoded_library_version", "encoded_library_date", "bit_rate_mode", "bit_rate"}
	mediumColumnsWithDefault    = []string{}
	mediumPrimaryKeyColumns     = []string{"id"}
)

type (
	// MediumSlice is an alias for a slice of pointers to Medium.
	// This should generally be used opposed to []Medium.
	MediumSlice []*Medium

	mediumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mediumType                 = reflect.TypeOf(&Medium{})
	mediumMapping              = queries.MakeStructMapping(mediumType)
	mediumPrimaryKeyMapping, _ = queries.BindMapping(mediumType, mediumMapping, mediumPrimaryKeyColumns)
	mediumInsertCacheMut       sync.RWMutex
	mediumInsertCache          = make(map[string]insertCache)
	mediumUpdateCacheMut       sync.RWMutex
	mediumUpdateCache          = make(map[string]updateCache)
	mediumUpsertCacheMut       sync.RWMutex
	mediumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single medium record from the query.
func (q mediumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Medium, error) {
	o := &Medium{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for media")
	}

	return o, nil
}

// All returns all Medium records from the query.
func (q mediumQuery) All(ctx context.Context, exec boil.ContextExecutor) (MediumSlice, error) {
	var o []*Medium

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to Medium slice")
	}

	return o, nil
}

// Count returns the count of all Medium records in the query.
func (q mediumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count media rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mediumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if media exists")
	}

	return count > 0, nil
}

// Media retrieves all the records using an executor.
func Media(mods ...qm.QueryMod) mediumQuery {
	mods = append(mods, qm.From("\"media\""))
	return mediumQuery{NewQuery(mods...)}
}

// FindMedium retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMedium(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Medium, error) {
	mediumObj := &Medium{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"media\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, mediumObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from media")
	}

	return mediumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Medium) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no media provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(mediumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mediumInsertCacheMut.RLock()
	cache, cached := mediumInsertCache[key]
	mediumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mediumAllColumns,
			mediumColumnsWithDefault,
			mediumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mediumType, mediumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mediumType, mediumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"media\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"media\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "orm: unable to insert into media")
	}

	if !cached {
		mediumInsertCacheMut.Lock()
		mediumInsertCache[key] = cache
		mediumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Medium.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Medium) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	mediumUpdateCacheMut.RLock()
	cache, cached := mediumUpdateCache[key]
	mediumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mediumAllColumns,
			mediumPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update media, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"media\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, mediumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mediumType, mediumMapping, append(wl, mediumPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update media row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for media")
	}

	if !cached {
		mediumUpdateCacheMut.Lock()
		mediumUpdateCache[key] = cache
		mediumUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q mediumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for media")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for media")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MediumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("orm: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mediumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"media\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, mediumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in medium slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all medium")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Medium) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no media provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(mediumColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	mediumUpsertCacheMut.RLock()
	cache, cached := mediumUpsertCache[key]
	mediumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mediumAllColumns,
			mediumColumnsWithDefault,
			mediumColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			mediumAllColumns,
			mediumPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("orm: unable to upsert media, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(mediumPrimaryKeyColumns))
			copy(conflict, mediumPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"media\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(mediumType, mediumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mediumType, mediumMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "orm: unable to upsert media")
	}

	if !cached {
		mediumUpsertCacheMut.Lock()
		mediumUpsertCache[key] = cache
		mediumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Medium record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Medium) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no Medium provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mediumPrimaryKeyMapping)
	sql := "DELETE FROM \"media\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from media")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for media")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mediumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no mediumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from media")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for media")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MediumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mediumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"media\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mediumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from medium slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for media")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Medium) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMedium(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MediumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MediumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mediumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"media\".* FROM \"media\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mediumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in MediumSlice")
	}

	*o = slice

	return nil
}

// MediumExists checks if the Medium row exists.
func MediumExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"media\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if media exists")
	}

	return exists, nil
}
