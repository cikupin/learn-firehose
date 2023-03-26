package cmd

import (
	"fmt"
	"log"

	"github.com/cikupin/learn-firehose/config"
	stencil "github.com/odpf/stencil/clients/go"
	"github.com/spf13/cobra"
)

var getDescriptorCmd = &cobra.Command{
	Use: "get-descriptor",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv()
		getDescriptor()
	},
}

func init() {
	rootCmd.AddCommand(getDescriptorCmd)
}

func getDescriptor() {
	namespace := "mynamespace"
	schemaName := "flagschema"

	url := fmt.Sprintf("%s:%d/v1beta1/namespaces/%s/schemas/%s", config.Cfg.Stencil.Host, config.Cfg.Stencil.Port, namespace, schemaName)
	client, err := stencil.NewClient([]string{url}, stencil.Options{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = client.GetDescriptor("payload.Flag")
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("descriptor payload.Flag found")

	_, err = client.GetDescriptor("payload.FlagReply")
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("descriptor payload.FlagReply found")
}
