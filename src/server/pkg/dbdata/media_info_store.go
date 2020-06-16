package dbdata

import (
	"context"
	"errors"

	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata/conversion"
	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MediaStore struct{}

func (s *MediaStore) Insert(ctx context.Context, conn boil.ContextExecutor, media *pb.Media) error {
	row := conversion.MediaToORM(media)
	return row.Insert(ctx, conn, boil.Infer())
}

func (s *MediaStore) Update(ctx context.Context, conn boil.ContextExecutor, media *pb.Media, fields ...string) error {
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
		mods = append(mods, filter.ModOr(cols.Performer, cols.AlbumPerformer, cols.Album, cols.Track, cols.Genre))
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
	return mods
}

func (s *MediaStore) List(ctx context.Context, conn boil.ContextExecutor, filter MediaFilter, fn func(int64)) ([]*pb.Media, error) {
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

func (s *MediaStore) Select(ctx context.Context, conn boil.ContextExecutor, id string) (*pb.Media, error) {
	row, err := orm.FindMedium(ctx, conn, id)
	if err != nil {
		return nil, err
	}
	return conversion.MediaFromORM(row), nil
}

func (s *MediaStore) Exists(ctx context.Context, conn boil.ContextExecutor, id string) (bool, error) {
	return orm.MediumExists(ctx, conn, id)
}

func (s *MediaStore) ModsFind(media *pb.Media) ([]qm.QueryMod, error) {
	var mods []qm.QueryMod
	var where = orm.MediumWhere
	if media.Location != "" {
		mods = append(mods, where.Location.EQ(media.Location))
	} else if media.Id != "" {
		mods = append(mods, where.ID.EQ(media.Id))
	} else if media.Sha != "" {
		mods = append(mods, where.Sha.EQ(media.Sha))
	} else if media.ShaImage != "" {
		mods = append(mods, where.ShaImage.EQ(media.ShaImage))
	} else {
		return nil, errors.New("could not find any criteria for finding media")
	}
	return mods, nil
}

func (s *MediaStore) FindMedia(ctx context.Context, conn boil.ContextExecutor, media *pb.Media) (*pb.Media, error) {
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

func (s *MediaStore) FindMedias(ctx context.Context, conn boil.ContextExecutor, media *pb.Media, limit int) ([]*pb.Media, error) {
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

func (s *MediaStore) DistinctList(ctx context.Context, conn boil.ContextExecutor, filter MediaFilter) ([]*pb.Media, error) {
	var mods []qm.QueryMod

	if filter.Distinct == "" {
		return nil, errors.New("missing parameter: filter.distinct")
	}

	mods = append(mods, qm.Distinct(filter.Distinct))
	mods = append(mods, s.FilterMods(filter)...)
	mods = append(mods, filter.Mods()...)
	rows, err := orm.Media(mods...).All(ctx, conn)
	if err != nil {
		return nil, err
	}
	return conversion.MediasFromORM(rows...), nil
}

func (s *MediaStore) AlbumListView(ctx context.Context, conn boil.ContextExecutor, filter MediaFilter) ([]*pb.Media, error) {
	var mods []qm.QueryMod
	var cols = orm.ViewAlbumColumns

	if filter.Query != "" {
		mods = append(mods, filter.ModOr(cols.Genre, cols.Performer, cols.Album, cols.Format))
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
