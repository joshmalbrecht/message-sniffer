package cmd

import (
	"github.com/joshmalbrecht/message-sniffer/internal/queue"
	"github.com/spf13/cobra"
)

var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Prints queue messages",
	Long: `Displays all of the messages that exist in the provided queue and continues to display messages 
	as they are created in the queue. The messages in the queue are not acknowledged.`,
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
	queueCmd.PersistentFlags().StringP("name", "n", "", "Name of a queue to display messages for")
	queueCmd.MarkPersistentFlagRequired("name")

	rootCmd.AddCommand(queueCmd)
}
