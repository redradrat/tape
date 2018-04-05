package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Build information
	BuildStamp string
	GitHash    string
	Version    string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Print tape version tag`,
	Run: func(cmd *cobra.Command, args []string) {
		if Version != "" {
			fmt.Println("tape " + Version)
		} else {
			fmt.Println("tape NIGHTLY (git: " + GitHash + ")")
		}
	},
}
