package wise

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Transfer struct {
	ID              int         `json:"id"`
	User            int         `json:"user"`
	TargetAccount   int         `json:"targetAccount"`
	SourceAccount   int         `json:"sourceAccount"`
	Quote           int         `json:"quote"`
	QuoteUUID       string      `json:"quoteUuid"`
	Status          string      `json:"status"`
	Reference       string      `json:"reference"`
	Rate            float32     `json:"rate"`
	Created         string      `json:"created"`
	Business        interface{} `json:"business"`
	TransferRequest interface{} `json:"transferRequest"`
	Details         struct {
		Reference string `json:"reference"`
	} `json:"details"`
	HasActiveIssues       bool    `json:"hasActiveIssues"`
	SourceCurrency        string  `json:"sourceCurrency"`
	SourceValue           float64 `json:"sourceValue"`
	TargetCurrency        string  `json:"targetCurrency"`
	TargetValue           float64 `json:"targetValue"`
	CustomerTransactionID string  `json:"customerTransactionId"`
}

type TransferListOptions struct {
	Profile          int    `json:"profile"`
	Status           string `json:"status"`
	SourceCurrency   string `json:"sourceCurrency"`
	TargetCurrency   string `json:"targetCurrency"`
	CreatedDateStart string `json:"createdDateStart"`
	CreatedDateEnd   string `json:"createdDateEnd"`
	Limit            int    `json:"limit"`
	Offset           int    `json:"offset"`
}

func transferListOptionsToQueryParams(opts *TransferListOptions) (map[string]string, error) {
	qp := make(map[string]string)
	raw := map[string]string{
		"profile":          strconv.FormatInt(int64(opts.Profile), 10),
		"status":           opts.Status,
		"sourceCurrency":   opts.SourceCurrency,
		"targetCurrency":   opts.TargetCurrency,
		"createdDateStart": opts.CreatedDateStart,
		"createdDateEnd":   opts.CreatedDateEnd,
		"limit":            strconv.FormatInt(int64(opts.Limit), 10),
		"offset":           strconv.FormatInt(int64(opts.Offset), 10),
	}
	for k, v := range raw {
		if v != "" && v != "0" {
			qp[k] = v
		}
	}
	return qp, nil
}

func PrintTransfersTable(transfers []Transfer, out io.Writer) {
	t := table.NewWriter()
	t.SetStyle(defaultTableStyle)
	t.SetOutputMirror(out)
	t.AppendHeader(table.Row{"Id", "Source Account", "Target Account", "Status", "Reference"})
	for _, transfer := range transfers {
		t.AppendRow(table.Row{transfer.ID, transfer.SourceAccount, transfer.TargetAccount, transfer.Status, transfer.Reference})
	}
	t.Render()
}

func (w WiseClient) ListTransfers(opts *TransferListOptions) ([]Transfer, error) {
	var (
		transfers []Transfer
		err       error
	)

	qp, err := transferListOptionsToQueryParams(opts)
	if err != nil {
		return nil, err
	}

	resp, err := w.cli.R().
		SetQueryParams(qp).
		Get(wiseTransfersPath)
	if err != nil {
		return nil, err
	}

	data := resp.Body()
	err = json.Unmarshal(data, &transfers)
	if err != nil {
		return nil, err
	}

	return transfers, nil
}

func (w WiseClient) GetTransfer(id int) (Transfer, error) {
	var (
		transfer Transfer
		err      error
	)

	resp, err := w.cli.R().Get(fmt.Sprintf("%s/%d", wiseTransfersPath, id))
	if err != nil {
		return transfer, err
	}

	data := resp.Body()
	err = json.Unmarshal(data, &transfer)
	if err != nil {
		return transfer, err
	}

	return transfer, nil
}

func (w WiseClient) GetReceipt(id int) ([]byte, error) {
	resp, err := w.cli.R().Get(fmt.Sprintf("%s/%d/receipt.pdf", wiseTransfersPath, id))
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}
