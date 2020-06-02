package data

type MediaFilter struct {
	QueryFilter
	Format       string `schema:"format"`
	Filename     string `schema:"filename"`
	Distinct     string `schema:"distinct"`
	Album        string `schema:"album"`
	Performer    string `schema:"performer"`
	RecordedDate int    `schema:"recorded_date"`
	Genre        string `schema:"genre"`
}
