package cmd

import (
	"fmt"
	"log"

	"github.com/cikupin/learn-firehose/client"
	"github.com/cikupin/learn-firehose/config"
	"github.com/spf13/cobra"
)

var initSchemaCmd = &cobra.Command{
	Use:   "init-schema",
	Short: "initialize schema to stencil",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv()
		initSchema()
	},
}

func init() {
	rootCmd.AddCommand(initSchemaCmd)
}

func initSchema() {
	namespace := "mynamespace"
	schemaName := "flagschema"
	schemaDescriptor := "/usr/bin/payload/schema.desc"

	stencilHost := fmt.Sprintf("%s:%d", config.Cfg.Stencil.Host, config.Cfg.Stencil.Port)
	stencil := client.NewStencilClient(stencilHost)

	responseCode, err := stencil.CreateNamespace(namespace, "for saving flag schema")
	if err != nil && responseCode != 6 {
		log.Fatalf("[ERROR] create stencil namespace: %s\n", err.Error())
	}

	err = stencil.UploadSchema(namespace, schemaName, schemaDescriptor)
	if err != nil {
		log.Fatalf("[ERROR] create schema: %s\n", err.Error())
	}
}
