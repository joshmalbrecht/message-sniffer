package cmd

import (
	"github.com/joshmalbrecht/message-sniffer/internal/queue"
	"github.com/spf13/cobra"
)

var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "TODO: Add description for queue",
	Long:  `TODO: Add longer description for queue`,
	Run: func(cmd *cobra.Command, args []string) {
		queueName, _ := cmd.Flags().GetString("name")
		hostname, _ := cmd.Flags().GetString("hostname")
		port, _ := cmd.Flags().GetInt("port")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		queue.Sniff(queueName, hostname, port, username, password)
	},
}

func init() {
	queueCmd.PersistentFlags().StringP("name", "n", "", "queue name to sniff on")
	queueCmd.MarkPersistentFlagRequired("name")

	rootCmd.AddCommand(queueCmd)
}
