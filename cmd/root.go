package cmd

import (
	"fmt"
	"os"

	"github.com/joshmalbrecht/message-sniffer/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "message-sniffer",
	Short: "Tool used for displaying published messages on a RabbitMQ broker",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("username", "u", "guest", "Username used to connect to a rabbitmq broker")
	rootCmd.PersistentFlags().StringP("password", "p", "guest", "Password used to connect to a rabbitmq broker")
	rootCmd.MarkFlagsRequiredTogether("username", "password")

	rootCmd.PersistentFlags().StringP("hostname", "H", "localhost", "Hostname used to connect to a rabbitmq broker")
	rootCmd.PersistentFlags().StringP("virtual-host", "V", "", "Virtual host (vhost) used to connect to a rabbitmq broker")
	rootCmd.PersistentFlags().IntP("port", "P", 5672, "Port used to connect to a rabbitmq broker")

	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of message-sniffer",
	Long:  `Print the version number of message-sniffer`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.Intance().GetVersion())
	},
}
