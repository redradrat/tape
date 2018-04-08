package cmd

import (
	"fmt"
	"os"

	"github.com/redradrat/tape/lib"
	"github.com/redradrat/tape/search"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(artistsCmd)
	artistsCmd.Flags().StringVarP(&artistsCmdAlbumsString, "album", "a", "", "Search for specific albums")
	artistsCmd.Flags().StringVarP(&artistsCmdComposerString, "composer", "c", "", "Search for specific composers")
	artistsCmd.Flags().StringVarP(&artistsCmdLabelString, "label", "l", "", "Search for specific label")
	artistsCmd.Flags().StringVarP(&artistsCmdTrackString, "track", "t", "", "Search for a specific track")
	artistsCmd.Flags().IntVar(&artistsCmdResultLimit, "limit", 0, "Limit the result to a specific amount")
}

var (
	artistsCmd = &cobra.Command{
		Use:     "artists (string)",
		Aliases: []string{"artist"},
		Short:   "Search for artists",
		Long:    `Search for artists`,
		RunE:    artists,
	}

	artistsCmdAlbumsString   string
	artistsCmdComposerString string
	artistsCmdLabelString    string
	artistsCmdTrackString    string
	artistsCmdResultLimit    int
)

func artists(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("%s takes exactly 1 or no arguments", cmd.Name())
	}

	var specSearches []search.Search
	if artistsCmdAlbumsString != "" {
		specSearches = append(specSearches, search.Search{SearchType: search.Artist, SearchString: artistsCmdAlbumsString})
	}
	if artistsCmdComposerString != "" {
		specSearches = append(specSearches, search.Search{SearchType: search.Composer, SearchString: artistsCmdComposerString})
	}
	if artistsCmdLabelString != "" {
		specSearches = append(specSearches, search.Search{SearchType: search.Label, SearchString: artistsCmdLabelString})
	}
	if artistsCmdTrackString != "" {
		specSearches = append(specSearches, search.Search{SearchType: search.Track, SearchString: artistsCmdLabelString})
	}

	if len(args) == 0 {
		result, err := lib.Artists("", specSearches)
		if err != nil {
			return err
		}
		result.Print(os.Stdout)
	} else {
		result, err := lib.Artists(args[0], specSearches)
		if err != nil {
			return err
		}
		result.Print(os.Stdout)
	}
	return nil
}
