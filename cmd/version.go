package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v0.4.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the current canconv version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {}
