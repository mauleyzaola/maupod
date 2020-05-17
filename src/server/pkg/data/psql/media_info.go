package psql

import (
	"context"
	"errors"

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

func (s *MediaStore) List(ctx context.Context, conn boil.ContextExecutor, filter filters.MediaFilter) ([]*domain.Media, error) {
	return nil, errors.New("not implemented")
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
