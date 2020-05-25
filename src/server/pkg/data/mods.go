package data

import (
	"github.com/mauleyzaola/maupod/src/server/pkg/filters"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func Mods(f *filters.QueryFilter) []qm.QueryMod {
	var mods []qm.QueryMod
	if f.Sort != "" {
		mods = append(mods, qm.OrderBy(f.Sort+" "+f.Direction))
	}
	if f.Limit != 0 {
		mods = append(mods, qm.Limit(f.Limit))
	}
	if f.Offset != 0 {
		mods = append(mods, qm.Offset(f.Offset))
	}
	return mods
}

func ModOr(f *filters.QueryFilter, fields ...string) qm.QueryMod {
	var submods []qm.QueryMod
	for _, v := range fields {
		submods = append(submods, qm.Or("LOWER("+v+") LIKE ?", LikeQuoted(f.Query)))
	}
	return qm.Expr(submods...)
}
