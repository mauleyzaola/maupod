package data

type MediaFilter struct {
	QueryFilter
	Format   string `schema:"format"`
	Filename string `schema:"filename"`
	Distinct string `schema:"distinct"`
}
