package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var grpcSinkCmd = &cobra.Command{
	Use:   "grpc-sink",
	Short: "sink for GRPC message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grpcSink called")
	},
}

func init() {
	rootCmd.AddCommand(grpcSinkCmd)
}
