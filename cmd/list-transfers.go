/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/chirauki/wise-downloader/pkg/wise"
)

func newListTransfers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-transfers",
		Short: "Lists the transfers found in Wise",
		Example: `
# List all accounts
wise-downloader --token asdfg list-transfers

# List all accounts in EUR
wise-downloader --token asdfg list-transfers --currency EUR
`,
		SilenceErrors: false,
		SilenceUsage:  false,
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := wise.NewClient(token)
			profile, err := guessProfile(cli)
			if err != nil {
				return err
			}

			transfers, err := cli.ListTransfers(&wise.TransferListOptions{
				Profile:          profile,
				Status:           status,
				SourceCurrency:   sourceCurrency,
				TargetCurrency:   targetCurrency,
				CreatedDateStart: startDate,
				CreatedDateEnd:   endDate,
				Limit:            limit,
				Offset:           offset,
			})
			if err != nil {
				return err
			}

			wise.PrintTransfersTable(transfers, cmd.OutOrStdout())

			return nil
		},
	}

	cmd.Flags().IntVarP(&profile, flagProfile, "p", 0, "profile to filter on")
	cmd.Flags().StringVarP(&status, flagStatus, "s", "", "status to filter on")
	cmd.Flags().StringVar(&sourceCurrency, flagSourceCurrency, "", "source currency to filter on")
	cmd.Flags().StringVar(&targetCurrency, flagTargetCurrency, "", "target currency to filter on")
	cmd.Flags().StringVar(&startDate, flagStartDate, "", "start date")
	cmd.Flags().StringVar(&endDate, flagEndDate, "", "end date")

	cmd.Flags().IntVarP(&limit, flagLimit, "l", 0, "Max number of items to get")
	cmd.Flags().IntVarP(&offset, flagOffset, "o", 0, "Offset of the items to get")

	return cmd
}
