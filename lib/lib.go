package lib

import (
	"io"

	"github.com/redradrat/tape/search"
)

// Artists returns an array of Artists
func Artists(search string, specSearch []search.Search) (Result, error) {
	var list []tapeType
	list = append(list, Artist{})
	list = append(list, Artist{})
	return Result{Values: list}, nil
}

// Albums returns an array of Albums
func Albums(search string, specSearch []search.Search) (Result, error) {
	var list []tapeType
	list = append(list, Album{})
	list = append(list, Album{})
	return Result{Values: list}, nil
}

// Tracks returns an array of Tracks
func Tracks(search string, specSearch []search.Search) (Result, error) {
	var list []tapeType
	list = append(list, Track{})
	list = append(list, Track{})
	return Result{Values: list}, nil
}

// Print prints a Result to specified io.Writer
func (res Result) Print(o io.Writer) error {
	return nil
}
