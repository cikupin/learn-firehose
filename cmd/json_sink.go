/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// jsonSinkCmd represents the jsonSink command
var jsonSinkCmd = &cobra.Command{
	Use:   "json-sink",
	Short: "sink for json message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jsonSink called")
	},
}

func init() {
	rootCmd.AddCommand(jsonSinkCmd)
}
