package cmd

import (
	"github.com/joshmalbrecht/message-sniffer/internal/topic"
	"github.com/spf13/cobra"
)

var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "Displays messages that are produced on a topic",
	Long:  `Displays messages that are produced on a topic.`,
	Run: func(cmd *cobra.Command, args []string) {
		exchangeName, _ := cmd.Flags().GetString("exchange-name")
		bindingKey, _ := cmd.Flags().GetString("binding-key")
		hostname, _ := cmd.Flags().GetString("hostname")
		port, _ := cmd.Flags().GetInt("port")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		topic.Sniff(exchangeName, bindingKey, hostname, port, username, password)
	},
}

func init() {
	topicCmd.PersistentFlags().StringP("exchange-name", "n", "", "Exchange name to sniff on")
	topicCmd.MarkPersistentFlagRequired("exchange-name")
	topicCmd.PersistentFlags().StringP("binding-key", "b", "#", "Binding key used to filter messages that are published on the defined exchange")

	rootCmd.AddCommand(topicCmd)
}
