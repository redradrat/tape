package search

const (
	Album searchType = iota
	Artist
	Composer
	Label
	Track
)

type searchType uint

func (sT searchType) String() string {
	names := [...]string{
		"Album",
		"Artist",
		"Composer",
		"Label",
		"Track"}

	return names[sT]
}

// Search represents an issued Search, that can be restricted by a searchType
type Search struct {
	SearchString string
	SearchType   searchType
}
