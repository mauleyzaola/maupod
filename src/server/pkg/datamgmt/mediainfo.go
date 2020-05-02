package datamgmt

import (
	"context"

	"github.com/mauleyzaola/maupod/src/server/pkg/filters"

	"github.com/mauleyzaola/maupod/src/server/pkg/media"
	"github.com/volatiletech/sqlboiler/boil"
)

type MediaInfo interface {
	Insert(ctx context.Context, conn boil.ContextExecutor, media *media.MediaInfo) error
	Update(ctx context.Context, conn boil.ContextExecutor, media *media.MediaInfo, fields []string) error
	Delete(ctx context.Context, conn boil.ContextExecutor, id string) error
	List(ctx context.Context, conn boil.ContextExecutor, filter filters.MediaFilter) error
}
