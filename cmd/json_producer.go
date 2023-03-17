package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var jsonProducerCmd = &cobra.Command{
	Use:   "json-producer",
	Short: "produce JSON message to Kafka in every 5 seconds",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jsonProducer called")
	},
}

func init() {
	rootCmd.AddCommand(jsonProducerCmd)
}
