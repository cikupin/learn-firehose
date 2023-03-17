package cmd

import (
	"github.com/cikupin/learn-firehose/config"
	"github.com/spf13/cobra"
)

var grpcProducerCmd = &cobra.Command{
	Use:   "grpc-producer",
	Short: "produce GRPC message to Kafka in every 5 seconds",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv()
	},
}

func init() {
	rootCmd.AddCommand(grpcProducerCmd)
}
