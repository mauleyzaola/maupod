package datamgmt

import (
	"context"

	"github.com/mauleyzaola/maupod/src/server/pkg/filters"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/volatiletech/sqlboiler/boil"
)

type Media interface {
	Insert(ctx context.Context, conn boil.ContextExecutor, media *pb.Media) error
	Update(ctx context.Context, conn boil.ContextExecutor, media *pb.Media, fields []string) error
	Delete(ctx context.Context, conn boil.ContextExecutor, id string) error
	List(ctx context.Context, conn boil.ContextExecutor, filter filters.MediaFilter, fn func(int64)) ([]*pb.Media, error)
}
