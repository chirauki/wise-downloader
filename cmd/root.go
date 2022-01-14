/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/tetratelabs/multierror"

	"github.com/chirauki/wise-downloader/pkg/wise"
)

var (
	errorMultipleProfiles = errors.New("multiple profiles found, can't get profile ID")
)

func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "wise-downloader",
		PreRunE: rootPreRun,
	}

	cmd.PersistentFlags().StringVarP(&token, flagToken, "t", "", "The Wise API token [REQUIRED]")
	cmd.MarkFlagRequired(flagToken)

	cmd.AddCommand(newListAccounts())
	cmd.AddCommand(newListTransfers())
	cmd.AddCommand(newTransferReceipt())

	return cmd
}

func rootPreRun(cmd *cobra.Command, args []string) error {
	var multi, err error

	for _, f := range []string{flagToken, flagStartDate, flagEndDate} {
		if !cmd.Flags().Lookup(f).Changed {
			multi = multierror.Append(multi, fmt.Errorf("flag %s is required", f))
		}
	}
	_, err = time.Parse(dateLayout, startDate)
	if err != nil {
		multi = multierror.Append(multi, errors.New("start date format is invalid, it should be YYYY-mm-dd, for instance, 2020-12-31 for December 31st 2020"))
	}
	_, err = time.Parse(dateLayout, endDate)
	if err != nil {
		multi = multierror.Append(multi, errors.New("end date format is invalid, it should be YYYY-mm-dd, for instance, 2020-12-31 for December 31st 2020"))
	}

	return multi
}

func guessProfile(cli *wise.WiseClient) (int, error) {
	var (
		profileId int
		err       error
	)

	profiles, err := cli.GetProfiles()
	if err != nil {
		return profileId, err
	}

	if len(profiles) == 1 {
		// Only if there is 1 profile, we can guess it
		profileId = profiles[0].ID
	} else {
		return profileId, errorMultipleProfiles
	}

	return profileId, err
}
