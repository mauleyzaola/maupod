package filters

type MediaFilter struct {
	QueryFilter
	Format   string `schema:"format"`
	Filename string `schema:"filename"`
}
