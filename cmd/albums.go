package cmd

import (
	"fmt"
	"os"

	"github.com/redradrat/tape/lib"
	"github.com/redradrat/tape/search"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(albumsCmd)
	albumsCmd.Flags().StringVarP(&albumsCmdArtistString, "artist", "a", "", "Search for specific artists")
	albumsCmd.Flags().StringVarP(&albumsCmdComposerString, "composer", "c", "", "Search for specific composers")
	albumsCmd.Flags().StringVarP(&albumsCmdLabelString, "label", "l", "", "Search for specific label")
	albumsCmd.Flags().StringVarP(&albumsCmdTrackString, "track", "t", "", "Search for a specific track")
	albumsCmd.Flags().IntVar(&albumsCmdResultLimit, "limit", 0, "Limit the result to a specific amount")
}

var (
	albumsCmd = &cobra.Command{
		Use:     "albums (string)",
		Aliases: []string{"album"},
		Short:   "Search for albums",
		Long:    `Search for albums`,
		RunE:    albums,
	}

	albumsCmdArtistString   string
	albumsCmdComposerString string
	albumsCmdLabelString    string
	albumsCmdTrackString    string
	albumsCmdResultLimit    int
)

func albums(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("%s takes exactly 1 or no arguments", cmd.Name())
	}

	var specSearches []search.Search
	if albumsCmdArtistString != "" {
		specSearches = append(specSearches, search.Search{SearchType: search.Artist, SearchString: albumsCmdArtistString})
	}
	if albumsCmdComposerString != "" {
		specSearches = append(specSearches, search.Search{SearchType: search.Composer, SearchString: albumsCmdComposerString})
	}
	if albumsCmdLabelString != "" {
		specSearches = append(specSearches, search.Search{SearchType: search.Label, SearchString: albumsCmdLabelString})
	}
	if albumsCmdTrackString != "" {
		specSearches = append(specSearches, search.Search{SearchType: search.Track, SearchString: albumsCmdTrackString})
	}

	if len(args) == 0 {
		result, err := lib.Albums("", specSearches)
		if err != nil {
			return err
		}
		result.Print(os.Stdout)
	} else {
		result, err := lib.Albums(args[0], specSearches)
		if err != nil {
			return err
		}
		result.Print(os.Stdout)
	}
	return nil
}
