package cmd

import (
	"fmt"
	"os"

	"github.com/redradrat/tape/lib"
	"github.com/redradrat/tape/search"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tracksCmd)
	tracksCmd.Flags().StringVar(&tracksCmdAlbumsString, "album", "", "Search for specific albums")
	tracksCmd.Flags().StringVar(&tracksCmdAlbumsString, "artist", "", "Search for specific albums")
	tracksCmd.Flags().StringVarP(&tracksCmdComposerString, "composer", "c", "", "Search for specific composers")
	tracksCmd.Flags().StringVarP(&tracksCmdLabelString, "label", "l", "", "Search for specific label")
	tracksCmd.Flags().IntVar(&tracksCmdResultLimit, "limit", 0, "Limit the result to a specific amount")
}

var (
	tracksCmd = &cobra.Command{
		Use:     "tracks (string)",
		Aliases: []string{"track"},
		Short:   "Search for tracks",
		Long:    `Search for tracks`,
		RunE:    tracks,
	}

	tracksCmdAlbumsString   string
	tracksCmdComposerString string
	tracksCmdLabelString    string
	tracksCmdTrackString    string
	tracksCmdResultLimit    int
)

func tracks(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("%s takes exactly 1 or no arguments", cmd.Name())
	}

	var specSearches []search.Search
	if tracksCmdAlbumsString != "" {
		specSearches = append(specSearches, search.Search{SearchType: search.Artist, SearchString: tracksCmdAlbumsString})
	}
	if tracksCmdComposerString != "" {
		specSearches = append(specSearches, search.Search{SearchType: search.Composer, SearchString: tracksCmdComposerString})
	}
	if tracksCmdLabelString != "" {
		specSearches = append(specSearches, search.Search{SearchType: search.Label, SearchString: tracksCmdLabelString})
	}
	if tracksCmdTrackString != "" {
		specSearches = append(specSearches, search.Search{SearchType: search.Track, SearchString: tracksCmdLabelString})
	}

	if len(args) == 0 {
		result, err := lib.Tracks("", specSearches)
		if err != nil {
			return err
		}
		result.Print(os.Stdout)
	} else {
		result, err := lib.Tracks(args[0], specSearches)
		if err != nil {
			return err
		}
		result.Print(os.Stdout)
	}
	return nil
}
