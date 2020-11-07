package dbdata

import (
	"context"
	"errors"
	"github.com/volatiletech/sqlboiler/queries"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata/conversion"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/protos"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MediaStore struct{}

func NewMediaStore() *MediaStore {
	return &MediaStore{}
}

func (s *MediaStore) Insert(ctx context.Context, conn boil.ContextExecutor, media *protos.Media) error {
	row := conversion.MediaToORM(media)
	return row.Insert(ctx, conn, boil.Infer())
}

func (s *MediaStore) Update(ctx context.Context, conn boil.ContextExecutor, media *protos.Media, fields ...string) error {
	row := conversion.MediaToORM(media)
	_, err := row.Update(ctx, conn, boil.Whitelist(fields...))
	return err
}

func (s *MediaStore) Delete(ctx context.Context, conn boil.ContextExecutor, id string) error {
	row, err := orm.FindMedium(ctx, conn, id)
	if err != nil {
		return err
	}
	_, err = row.Delete(ctx, conn)
	return err
}

func (s *MediaStore) FilterMods(filter MediaFilter) []qm.QueryMod {
	var mods []qm.QueryMod
	var cols = orm.MediumColumns
	if filter.Query != "" {
		mods = append(mods, filter.ModOr(cols.Performer, cols.AlbumPerformer, cols.Album, cols.Track, cols.Genre, cols.Track))
	}
	if filter.Album != "" {
		mods = append(mods, filter.ModAnd(KeyValuePair{
			Key:   cols.Album,
			Value: filter.Album,
		}))
	}
	if filter.Album != "" {
		mods = append(mods, filter.ModAnd(KeyValuePair{
			Key:   cols.Album,
			Value: filter.Album,
		}))
	}
	if filter.Genre != "" {
		mods = append(mods, filter.ModAnd(KeyValuePair{
			Key:   cols.Genre,
			Value: filter.Genre,
		}))
	}
	if filter.Performer != "" {
		mods = append(mods, filter.ModAnd(KeyValuePair{
			Key:   cols.Performer,
			Value: filter.Performer,
		}))
	}
	if filter.AlbumIdentifier != "" {
		mods = append(mods, filter.ModAnd(KeyValuePair{
			Key:   cols.AlbumIdentifier,
			Value: filter.AlbumIdentifier,
		}))
	}
	return mods
}

func (s *MediaStore) List(ctx context.Context, conn boil.ContextExecutor, filter MediaFilter, fn func(int64)) ([]*protos.Media, error) {
	var mods []qm.QueryMod

	mods = append(mods, s.FilterMods(filter)...)
	if fn != nil {
		total, err := orm.Media(mods...).Count(ctx, conn)
		if err != nil {
			return nil, err
		}
		fn(total)
	}
	mods = append(mods, Mods(&filter.QueryFilter)...)
	rows, err := orm.Media(mods...).All(ctx, conn)
	if err != nil {
		return nil, err
	}
	return conversion.MediasFromORM(rows...), nil
}

func (s *MediaStore) Select(ctx context.Context, conn boil.ContextExecutor, id string) (*protos.Media, error) {
	row, err := orm.FindMedium(ctx, conn, id)
	if err != nil {
		return nil, err
	}
	return conversion.MediaFromORM(row), nil
}

func (s *MediaStore) Exists(ctx context.Context, conn boil.ContextExecutor, id string) (bool, error) {
	return orm.MediumExists(ctx, conn, id)
}

func (s *MediaStore) ModsFind(media *protos.Media) ([]qm.QueryMod, error) {
	var mods []qm.QueryMod
	var where = orm.MediumWhere
	if media.Location != "" {
		mods = append(mods, where.Location.EQ(media.Location))
	} else if media.Id != "" {
		mods = append(mods, where.ID.EQ(media.Id))
	} else if media.Sha != "" {
		mods = append(mods, where.Sha.EQ(media.Sha))
	} else if media.AlbumIdentifier != "" {
		mods = append(mods, where.AlbumIdentifier.EQ(media.AlbumIdentifier))
	} else {
		return nil, errors.New("could not find any criteria for finding media")
	}
	return mods, nil
}

func (s *MediaStore) FindMedia(ctx context.Context, conn boil.ContextExecutor, media *protos.Media) (*protos.Media, error) {
	mods, err := s.ModsFind(media)
	if err != nil {
		return nil, err
	}
	row, err := orm.Media(mods...).One(ctx, conn)
	if err != nil {
		return nil, err
	}
	return conversion.MediaFromORM(row), nil
}

func (s *MediaStore) FindMedias(ctx context.Context, conn boil.ContextExecutor, media *protos.Media, limit int) ([]*protos.Media, error) {
	mods, err := s.ModsFind(media)
	if err != nil {
		return nil, err
	}
	if limit != 0 {
		mods = append(mods, qm.Limit(limit))
	}
	rows, err := orm.Media(mods...).All(ctx, conn)
	if err != nil {
		return nil, err
	}
	return conversion.MediasFromORM(rows...), nil
}

func (s *MediaStore) DistinctList(ctx context.Context, conn boil.ContextExecutor, filter MediaFilter) ([]*protos.Media, error) {
	var mods []qm.QueryMod

	if filter.Distinct == "" {
		return nil, errors.New("missing parameter: filter.distinct")
	}

	mods = append(mods, qm.Distinct(filter.Distinct))
	mods = append(mods, s.FilterMods(filter)...)
	mods = append(mods, filter.Mods()...)
	rows, err := orm.ViewAlbums(mods...).All(ctx, conn)
	if err != nil {
		return nil, err
	}
	return conversion.ViewAlbumsToMedia(rows...), nil
}

func (s *MediaStore) AlbumListView(ctx context.Context, conn boil.ContextExecutor, filter MediaFilter) ([]*protos.Media, error) {
	var mods []qm.QueryMod
	var cols = orm.ViewAlbumColumns

	if filter.Query != "" {
		mods = append(mods, filter.ModOr(cols.Genre, cols.Performer, cols.Album, cols.Format))
	}

	if filter.AlbumIdentifier != "" {
		mods = append(mods, filter.ModAnd(KeyValuePair{
			Key:   cols.AlbumIdentifier,
			Value: filter.AlbumIdentifier,
		}))
	}
	if filter.Genre != "" {
		mods = append(mods, filter.ModAnd(KeyValuePair{
			Key:   cols.Genre,
			Value: filter.Genre,
		}))
	}
	if filter.Performer != "" {
		mods = append(mods, filter.ModAnd(KeyValuePair{
			Key:   cols.Performer,
			Value: filter.Performer,
		}))
	}
	if filter.Album != "" {
		mods = append(mods, filter.ModAnd(KeyValuePair{
			Key:   cols.Album,
			Value: filter.Album,
		}))
	}
	if filter.Format != "" {
		mods = append(mods, filter.ModAnd(KeyValuePair{
			Key:   cols.Format,
			Value: filter.Format,
		}))
	}
	mods = append(mods, filter.Mods()...)
	rows, err := orm.ViewAlbums(mods...).All(ctx, conn)
	if err != nil {
		return nil, err
	}
	return conversion.ViewAlbumsToMedia(rows...), nil
}

func (s *MediaStore) PlaylistMods(filter QueryFilter) []qm.QueryMod {
	var mods []qm.QueryMod
	var cols = orm.PlaylistColumns
	mods = filter.Mods()
	if filter.Query != "" {
		mods = append(mods, filter.ModOr(cols.Name))
	}
	return mods
}

// PlayedMediaList will return a list of media which has been played at least once
// based on its related events
func (s *MediaStore) PlayedMediaList(ctx context.Context, conn boil.ContextExecutor) ([]*protos.Media, error) {
	var sqlQuery = `
	select m.*
	from media m
	where exists(
		select null
		from   media_event
		where  m.sha = sha
		and     event = $1)
`
	var medias orm.MediumSlice
	query := queries.Raw(sqlQuery, protos.Message_MESSAGE_EVENT_ON_TRACK_PLAY_COUNT_INCREASE)
	err := query.Bind(ctx, conn, &medias)
	if err != nil {
		return nil, err
	}
	return conversion.MediasFromORM(medias...), nil
}
