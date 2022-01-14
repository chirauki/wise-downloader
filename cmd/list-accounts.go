/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/chirauki/wise-downloader/pkg/wise"
)

func newListAccounts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-accounts",
		Short: "Lists the accounts found in Wise",
		Example: `
# List all accounts
wise-downloader --token asdfg list-accounts

# List all accounts in EUR
wise-downloader --token asdfg list-accounts --currency EUR
`,
		SilenceErrors: false,
		SilenceUsage:  false,
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := wise.NewClient(token)
			profile, err := guessProfile(cli)
			if err != nil {
				return err
			}
			accounts, err := cli.ListAccounts(&wise.AccountListOptions{
				ProfileID: profile,
				Currency:  currency,
				Limit:     limit,
				Offset:    offset,
			})
			if err != nil {
				return err
			}

			wise.PrintAccountsTable(accounts, cmd.OutOrStdout())

			return nil
		},
	}

	cmd.Flags().IntVarP(&profile, flagProfile, "p", 0, "Filter accounts by this profile")
	cmd.Flags().StringVarP(&currency, flagCurrency, "c", "", "Filter accounts by this currency")
	cmd.Flags().IntVarP(&limit, flagLimit, "l", 0, "Max number of items to get")
	cmd.Flags().IntVarP(&offset, flagOffset, "o", 0, "Offset of the items to get")

	return cmd
}
