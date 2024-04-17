package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "message-sniffer",
	Short: "TODO: Add short description",
	Long:  `TODO: Add long description`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("username", "u", "guest", "defines username used to connect to a rabbitmq broker")
	rootCmd.PersistentFlags().StringP("password", "p", "guest", "defines password used to connect to a rabbitmq broker")
	rootCmd.MarkFlagsRequiredTogether("username", "password")

	rootCmd.PersistentFlags().StringP("hostname", "H", "localhost", "defines hostname for rabbitmq broker")
	rootCmd.PersistentFlags().IntP("port", "P", 5672, "defines port for rabbitmq broker")
}
