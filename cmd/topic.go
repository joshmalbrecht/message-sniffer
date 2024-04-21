package cmd

import (
	"github.com/joshmalbrecht/message-sniffer/internal/topic"
	"github.com/spf13/cobra"
)

var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "TODO: Add description",
	Long:  `TODO: Add a longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		exchangeName, _ := cmd.Flags().GetString("exchange-name")
		routingKey, _ := cmd.Flags().GetString("routing-key")
		hostname, _ := cmd.Flags().GetString("hostname")
		port, _ := cmd.Flags().GetInt("port")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		topic.Sniff(exchangeName, routingKey, hostname, port, username, password)
	},
}

func init() {
	topicCmd.PersistentFlags().StringP("exchange-name", "n", "", "exchange name to sniff on")
	topicCmd.MarkPersistentFlagRequired("exchange-name")
	topicCmd.PersistentFlags().StringP("routing-key", "r", "", "routing key to sniff on")
	topicCmd.MarkPersistentFlagRequired("routing-key")

	rootCmd.AddCommand(topicCmd)
}
