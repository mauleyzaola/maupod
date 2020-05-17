package psql

import (
	"context"

	"github.com/mauleyzaola/maupod/src/server/pkg/data"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/mauleyzaola/maupod/src/server/pkg/data/conversion"
	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/domain"
	"github.com/mauleyzaola/maupod/src/server/pkg/filters"
	"github.com/volatiletech/sqlboiler/boil"
)

type MediaStore struct{}

func (s *MediaStore) Insert(ctx context.Context, conn boil.ContextExecutor, media *domain.Media) error {
	row := conversion.MediaToORM(media)
	return row.Insert(ctx, conn, boil.Infer())
}

func (s *MediaStore) Update(ctx context.Context, conn boil.ContextExecutor, media *domain.Media, fields []string) error {
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

func (s *MediaStore) List(ctx context.Context, conn boil.ContextExecutor, filter filters.MediaFilter, fn func(int64)) ([]*domain.Media, error) {
	var mods []qm.QueryMod
	// TODO: implement the actual filters
	if fn != nil {
		total, err := orm.Media(mods...).Count(ctx, conn)
		if err != nil {
			return nil, err
		}
		fn(total)
	}
	mods = append(mods, data.Mods(&filter.QueryFilter)...)
	rows, err := orm.Media(mods...).All(ctx, conn)
	if err != nil {
		return nil, err
	}
	return conversion.MediasFromORM(rows...), nil
}

func (s *MediaStore) Select(ctx context.Context, conn boil.ContextExecutor, id string) (*domain.Media, error) {
	row, err := orm.FindMedium(ctx, conn, id)
	if err != nil {
		return nil, err
	}
	return conversion.MediaFromORM(row), nil
}

func (s *MediaStore) Exists(ctx context.Context, conn boil.ContextExecutor, id string) (bool, error) {
	return orm.MediumExists(ctx, conn, id)
}
