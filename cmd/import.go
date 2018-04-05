package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import music ",
	Long:  `Import music into tape library`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}
