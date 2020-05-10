package filters

type MediaFilter struct {
	QueryFilter
	Format string `schema:"format"`
}
