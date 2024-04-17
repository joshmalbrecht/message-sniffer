package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// topicCmd represents the topic command
var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "TODO: Add description",
	Long:  `TODO: Add a longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("topic called")
	},
}

func init() {
	rootCmd.AddCommand(topicCmd)
}
