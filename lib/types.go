package lib

import (
	"reflect"
)

type tapeType interface {
	typeString() string
}

// Album represents an Album object.
type Album struct {
	artist Artist
}

func (album Album) typeString() string {
	return getTypeString(album)
}

// Artist represents an Artist object.
type Artist struct {
}

func (artist Artist) typeString() string {
	return getTypeString(artist)
}

// Composer represents a Composer object.
type Composer struct {
	Artist Artist
}

func (composer Composer) typeString() string {
	return getTypeString(composer)
}

// Label represents a Label object.
type Label struct {
	Artists []Artist
}

func (label Label) typeString() string {
	return getTypeString(label)
}

// Track represents a Track object.
type Track struct {
	Name     string
	Artist   Artist
	Album    Album
	Composer Composer
}

func (track Track) typeString() string {
	return getTypeString(track)
}

// Result represents the result of a search.
type Result struct {
	Values []tapeType
}

func getTypeString(t interface{ typeString() string }) string {
	return reflect.TypeOf(t).String()
}
