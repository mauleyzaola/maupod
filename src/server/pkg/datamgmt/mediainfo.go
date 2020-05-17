package datamgmt

import (
	"context"

	"github.com/mauleyzaola/maupod/src/server/pkg/domain"
	"github.com/mauleyzaola/maupod/src/server/pkg/filters"
	"github.com/volatiletech/sqlboiler/boil"
)

type Media interface {
	Insert(ctx context.Context, conn boil.ContextExecutor, media *domain.Media) error
	Update(ctx context.Context, conn boil.ContextExecutor, media *domain.Media, fields []string) error
	Delete(ctx context.Context, conn boil.ContextExecutor, id string) error
	List(ctx context.Context, conn boil.ContextExecutor, filter filters.MediaFilter) ([]*domain.Media, error)
}
