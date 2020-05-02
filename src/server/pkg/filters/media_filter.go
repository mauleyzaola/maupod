package filters

type MediaFilter struct {
	QueryFilter
	ID     string `schema:"id"`
	Format string `schema:"format"`
}
