/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/chirauki/wise-downloader/pkg/wise"
)

func newTransferReceipt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-receipt",
		Short: "Downloads the PDF receipt for a transfer",
		Example: `
# Get the receipt for a transfer
wise-downloader --token asdfg transfer-receipt --transfer 12345

# List all accounts in EUR
wise-downloader --token asdfg list-transfers --currency EUR
`,
		SilenceErrors: false,
		SilenceUsage:  false,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !cmd.Flags().Lookup(flagOutput).Changed {
				output = fmt.Sprintf("%v.pdf", transfer)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := wise.NewClient(token)

			receiptData, err := cli.GetReceipt(transfer)
			if err != nil {
				return err
			}

			f, err := os.Create(output)
			if err != nil {
				return err
			}
			defer f.Close()
			f.Write(receiptData)
			f.Sync()

			return nil
		},
	}

	cmd.Flags().IntVarP(&transfer, flagTransfer, "x", 0, "Transfer ID to get the receipt for")
	cmd.MarkFlagRequired(flagTransfer)
	cmd.Flags().StringVarP(&output, flagOutput, "o", "", "Output file [default uses <transferid>.pdf]")

	return cmd
}
