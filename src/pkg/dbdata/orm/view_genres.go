// Code generated by SQLBoiler 4.1.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ViewGenre is an object representing the database table.
type ViewGenre struct {
	Genre          null.String  `boil:"genre" json:"genre,omitempty" toml:"genre" yaml:"genre,omitempty"`
	PerformerCount null.Float64 `boil:"performer_count" json:"performer_count,omitempty" toml:"performer_count" yaml:"performer_count,omitempty"`
	AlbumCount     null.Float64 `boil:"album_count" json:"album_count,omitempty" toml:"album_count" yaml:"album_count,omitempty"`
	Duration       null.Float64 `boil:"duration" json:"duration,omitempty" toml:"duration" yaml:"duration,omitempty"`
	Total          null.Float64 `boil:"total" json:"total,omitempty" toml:"total" yaml:"total,omitempty"`

	R *viewGenreR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L viewGenreL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ViewGenreColumns = struct {
	Genre          string
	PerformerCount string
	AlbumCount     string
	Duration       string
	Total          string
}{
	Genre:          "genre",
	PerformerCount: "performer_count",
	AlbumCount:     "album_count",
	Duration:       "duration",
	Total:          "total",
}

// Generated where

type whereHelpernull_Float64 struct{ field string }

func (w whereHelpernull_Float64) EQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Float64) NEQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Float64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Float64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Float64) LT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Float64) LTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Float64) GT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Float64) GTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ViewGenreWhere = struct {
	Genre          whereHelpernull_String
	PerformerCount whereHelpernull_Float64
	AlbumCount     whereHelpernull_Float64
	Duration       whereHelpernull_Float64
	Total          whereHelpernull_Float64
}{
	Genre:          whereHelpernull_String{field: "\"view_genres\".\"genre\""},
	PerformerCount: whereHelpernull_Float64{field: "\"view_genres\".\"performer_count\""},
	AlbumCount:     whereHelpernull_Float64{field: "\"view_genres\".\"album_count\""},
	Duration:       whereHelpernull_Float64{field: "\"view_genres\".\"duration\""},
	Total:          whereHelpernull_Float64{field: "\"view_genres\".\"total\""},
}

// ViewGenreRels is where relationship names are stored.
var ViewGenreRels = struct {
}{}

// viewGenreR is where relationships are stored.
type viewGenreR struct {
}

// NewStruct creates a new relationship struct
func (*viewGenreR) NewStruct() *viewGenreR {
	return &viewGenreR{}
}

// viewGenreL is where Load methods for each relationship are stored.
type viewGenreL struct{}

var (
	viewGenreAllColumns            = []string{"genre", "performer_count", "album_count", "duration", "total"}
	viewGenreColumnsWithoutDefault = []string{"genre", "performer_count", "album_count", "duration", "total"}
	viewGenreColumnsWithDefault    = []string{}
	viewGenrePrimaryKeyColumns     = []string{"genre"}
)

type (
	// ViewGenreSlice is an alias for a slice of pointers to ViewGenre.
	// This should generally be used opposed to []ViewGenre.
	ViewGenreSlice []*ViewGenre

	viewGenreQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	viewGenreType                 = reflect.TypeOf(&ViewGenre{})
	viewGenreMapping              = queries.MakeStructMapping(viewGenreType)
	viewGenrePrimaryKeyMapping, _ = queries.BindMapping(viewGenreType, viewGenreMapping, viewGenrePrimaryKeyColumns)
	viewGenreInsertCacheMut       sync.RWMutex
	viewGenreInsertCache          = make(map[string]insertCache)
	viewGenreUpdateCacheMut       sync.RWMutex
	viewGenreUpdateCache          = make(map[string]updateCache)
	viewGenreUpsertCacheMut       sync.RWMutex
	viewGenreUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single viewGenre record from the query.
func (q viewGenreQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ViewGenre, error) {
	o := &ViewGenre{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for view_genres")
	}

	return o, nil
}

// All returns all ViewGenre records from the query.
func (q viewGenreQuery) All(ctx context.Context, exec boil.ContextExecutor) (ViewGenreSlice, error) {
	var o []*ViewGenre

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to ViewGenre slice")
	}

	return o, nil
}

// Count returns the count of all ViewGenre records in the query.
func (q viewGenreQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count view_genres rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q viewGenreQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if view_genres exists")
	}

	return count > 0, nil
}

// ViewGenres retrieves all the records using an executor.
func ViewGenres(mods ...qm.QueryMod) viewGenreQuery {
	mods = append(mods, qm.From("\"view_genres\""))
	return viewGenreQuery{NewQuery(mods...)}
}

// FindViewGenre retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindViewGenre(ctx context.Context, exec boil.ContextExecutor, genre null.String, selectCols ...string) (*ViewGenre, error) {
	viewGenreObj := &ViewGenre{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"view_genres\" where \"genre\"=$1", sel,
	)

	q := queries.Raw(query, genre)

	err := q.Bind(ctx, exec, viewGenreObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from view_genres")
	}

	return viewGenreObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ViewGenre) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no view_genres provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(viewGenreColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	viewGenreInsertCacheMut.RLock()
	cache, cached := viewGenreInsertCache[key]
	viewGenreInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			viewGenreAllColumns,
			viewGenreColumnsWithDefault,
			viewGenreColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(viewGenreType, viewGenreMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(viewGenreType, viewGenreMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"view_genres\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"view_genres\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "orm: unable to insert into view_genres")
	}

	if !cached {
		viewGenreInsertCacheMut.Lock()
		viewGenreInsertCache[key] = cache
		viewGenreInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the ViewGenre.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ViewGenre) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	viewGenreUpdateCacheMut.RLock()
	cache, cached := viewGenreUpdateCache[key]
	viewGenreUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			viewGenreAllColumns,
			viewGenrePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update view_genres, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"view_genres\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, viewGenrePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(viewGenreType, viewGenreMapping, append(wl, viewGenrePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "orm: unable to update view_genres row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for view_genres")
	}

	if !cached {
		viewGenreUpdateCacheMut.Lock()
		viewGenreUpdateCache[key] = cache
		viewGenreUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q viewGenreQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for view_genres")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for view_genres")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ViewGenreSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), viewGenrePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"view_genres\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, viewGenrePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in viewGenre slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all viewGenre")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ViewGenre) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no view_genres provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(viewGenreColumnsWithDefault, o)

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

	viewGenreUpsertCacheMut.RLock()
	cache, cached := viewGenreUpsertCache[key]
	viewGenreUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			viewGenreAllColumns,
			viewGenreColumnsWithDefault,
			viewGenreColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			viewGenreAllColumns,
			viewGenrePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("orm: unable to upsert view_genres, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(viewGenrePrimaryKeyColumns))
			copy(conflict, viewGenrePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"view_genres\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(viewGenreType, viewGenreMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(viewGenreType, viewGenreMapping, ret)
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
		return errors.Wrap(err, "orm: unable to upsert view_genres")
	}

	if !cached {
		viewGenreUpsertCacheMut.Lock()
		viewGenreUpsertCache[key] = cache
		viewGenreUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single ViewGenre record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ViewGenre) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no ViewGenre provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), viewGenrePrimaryKeyMapping)
	sql := "DELETE FROM \"view_genres\" WHERE \"genre\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from view_genres")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for view_genres")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q viewGenreQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no viewGenreQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from view_genres")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for view_genres")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ViewGenreSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), viewGenrePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"view_genres\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, viewGenrePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from viewGenre slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for view_genres")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ViewGenre) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindViewGenre(ctx, exec, o.Genre)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ViewGenreSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ViewGenreSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), viewGenrePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"view_genres\".* FROM \"view_genres\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, viewGenrePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in ViewGenreSlice")
	}

	*o = slice

	return nil
}

// ViewGenreExists checks if the ViewGenre row exists.
func ViewGenreExists(ctx context.Context, exec boil.ContextExecutor, genre null.String) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"view_genres\" where \"genre\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, genre)
	}
	row := exec.QueryRowContext(ctx, sql, genre)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if view_genres exists")
	}

	return exists, nil
}
