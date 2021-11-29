package cmd

import (
	"context"
	"fmt"

	"github.com/foxglove/foxglove-cli/foxglove/svc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func executeImport(baseURL, clientID, deviceID, filename, token string) error {
	ctx := context.Background()
	client := svc.NewRemoteFoxgloveClient(baseURL, clientID, token)
	err := svc.Import(ctx, client, deviceID, filename)
	if err != nil {
		return err
	}
	return nil
}

func newImportCommand(baseURL, clientID string) *cobra.Command {
	var deviceID string
	var filename string
	importCmd := &cobra.Command{
		Use:   "import",
		Short: "Import a data file to the foxglove data platform",
		Run: func(cmd *cobra.Command, args []string) {
			err := executeImport(baseURL, clientID, deviceID, filename, viper.GetString("bearer_token"))
			if err != nil {
				fmt.Printf("Import failed: %s\n", err)
			}
		},
	}
	importCmd.PersistentFlags().StringVarP(&deviceID, "device-id", "", "", "device ID")
	importCmd.PersistentFlags().StringVarP(&filename, "filename", "", "", "filename")
	return importCmd
}
